package main

import (
	"machine"
)

var (
	uart = machine.UART0
	tx   = machine.UART_TX_PIN
	rx   = machine.UART_RX_PIN
	led  = machine.D8
)

func main() {
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	uart.Configure(machine.UARTConfig{TX: tx, RX: rx})

	for {
		if uart.Buffered() > 0 {
			data, _ := uart.ReadByte()
			char := string(data)

			if char == "A" {
				led.High()
			}
			if char == "B" {
				led.Low()
			}
		}
	}
}
