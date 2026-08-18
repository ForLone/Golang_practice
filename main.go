package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

var money = 1000
var bank = 0
var mtx = sync.Mutex{}
var paymentHistory = make([]payment, 0)

type payment struct{
	Description string `json:"description"`
	USD int `json:"usd"`
	FullName string `json:"fullname"`
	Address string `json:"address"`
	Time time.Time `json:"date"`
}

type HttpResponse struct{
	Money int `json:"money"`
	Payments []payment `json:"paymentHistory"`
}

func (p payment) Println(){
	fmt.Println("Description:", p.Description)
	fmt.Println("Price:", p.USD)
	fmt.Println("Fullname:", p.FullName)
	fmt.Println("Address:", p.Address)
}

func (p payment) Validate() error{
	if p.Description == ""{
		return http.ErrBodyNotAllowed
	}
	if p.USD == 0{
		return http.ErrBodyNotAllowed
	}
	return nil
}

func payHandler(w http.ResponseWriter, r *http.Request){
	var p payment
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil{
		fmt.Println("Failed to read body: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	p.Time = time.Now()
	
	if money < p.USD{
		str := "Недостаточно денег на балансе"
		fmt.Println(str)
		w.Write([]byte(str, ))
		return
	}
	mtx.Lock()
	money -= p.USD
	mtx.Unlock() 
	if err := p.Validate(); err != nil{
		fmt.Println("Пользователь не заполнил необходимые поля:", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	paymentHistory = append(paymentHistory, p)
	
	httpResponse := HttpResponse{
		Money: money,
		Payments: paymentHistory,
	}

	l, err := json.Marshal(httpResponse)
	if err != nil{
		fmt.Println("Не удалось конвертировать в json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	
	if _, err := w.Write(l); err != nil{
		fmt.Println("Что то пошло не так")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	
}

func main(){
	http.HandleFunc("/pay", payHandler)
	
	if err := http.ListenAndServe(":9091", nil); err != nil{
		fmt.Println("Произошла ошибка", err.Error())
	}
}