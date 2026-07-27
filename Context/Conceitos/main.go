package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Reserva cancelada, tempo expirado")
		return
	case <-time.After(3 * time.Second):
		fmt.Println("Hotel reservado com sucesso")
		return
	}
}
