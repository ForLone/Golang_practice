package main

import (
	"context"
	"fmt"
	"study/feature_postgres/simple_connection"
	"study/feature_postgres/simple_sql"

	// "github.com/k0kubun/pp"
	// "time"
)



func main() {
	ctx := context.Background()
	conn, err := simple_connection.CreateConnection(ctx)
	if err != nil{
		panic(err)
	}
	if err := simple_sql.CreateTable(ctx, conn); err != nil{
		panic(err)
	}
	
	// tasks, err := simple_sql.SelectRows(ctx, conn)
	// if err!=nil{
	// 	panic(err)
	// }

	// for _, task := range tasks{
	// 	if task.ID == 3{
	// 		task.Title = "Покормить кошку"
	// 		task.Description = "Отсыпать кошке 30гр корма"
	// 		task.Completed = true
	// 		now := time.Now()
	// 		task.Completed_at = &now
	// 		if err := simple_sql.UpdateTask(ctx, conn, task); err != nil{
	// 			panic(err)
	// 		}
	// 		break
 	// 	}
		
	// }
	if err := simple_sql.AlterTable(ctx, conn); err != nil{
		panic(err)
	}

	fmt.Println("success")
}
