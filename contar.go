package contar

import (
	"fmt"
	"time"
)

// Contar123 exibe mensagem no console contando de 1 até 3
func Contar123() {
	for i := 1; i < 3; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}
