package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	bookHotel(ctx)

}

func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Requisição expirada")
		return
	case <-time.After(time.Second * 3):
		fmt.Println("Requisição finalizada")
	}
}
