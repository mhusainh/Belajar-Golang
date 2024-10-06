package main

import (
    "fmt"
    "sync"
    "time"
)

// Fungsi untuk memproses pesanan
func processOrder(order MenuItem, ch chan<- string, wg *sync.WaitGroup) {
    defer wg.Done()
    select {
    case <-time.After(5 * time.Second): // Timeout untuk pemrosesan pesanan
        ch <- fmt.Sprintf("Order processing timed out: %s", order.Name)
    default:
        time.Sleep(2 * time.Second) // Simulasi pemrosesan pesanan
        encodedDetails := encodeOrderDetails(fmt.Sprintf("Order: %s, Jumlah: %d", order.Name, order.Quantity))
        ch <- fmt.Sprintf("Pesanan diproses: %s (encoded: %s)", order.Name, encodedDetails)
    }
}
