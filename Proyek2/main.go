package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
    "sync"
)

func main() {
    defer fmt.Println("Program selesai")
    restaurantOrder := &RestaurantOrder{}

    // Menambahkan beberapa item ke menu
    restaurantOrder.Menu = []MenuItem{
        {Name: "Nasi Goreng", Price: 25000.00},
        {Name: "Mie Goreng", Price: 22000.00},
        {Name: "Ayam Bakar", Price: 30000.00},
    }

    // Menampilkan menu yang tersedia
    clearScreen()
    printHeader("    SELAMAT DATANG DI RESTORAN KAMI.")
    fmt.Println("Menu yang tersedia:")
    for _, item := range restaurantOrder.Menu {
        fmt.Printf("- %s: Rp.%.2f\n", item.Name, item.Price)
    }
    printSeparator()

    // Menerima pesanan dari pengguna
    scanner := bufio.NewScanner(os.Stdin)

    for {
        fmt.Print("Masukkan perintah (add/edit/delete/selesai): ")
        scanner.Scan()
        command := strings.TrimSpace(strings.ToLower(scanner.Text()))

        if command == "selesai" {
            break
        }

        switch command {
        case "add":
            fmt.Print("Masukkan nama item: ")
            scanner.Scan()
            itemName := strings.TrimSpace(scanner.Text())

            fmt.Print("Masukkan jumlah: ")
            scanner.Scan()
            orderQuantity := scanner.Text()

            safeInput(func() {
                quantity, err := strconv.Atoi(orderQuantity)
                if err != nil {
                    panic("Jumlah tidak valid.")
                }

                var foundItem *MenuItem
                for i := range restaurantOrder.Menu {
                    if strings.EqualFold(strings.ReplaceAll(restaurantOrder.Menu[i].Name, " ", ""), strings.ReplaceAll(itemName, " ", "")) {
                        foundItem = &restaurantOrder.Menu[i]
                        foundItem.Quantity += quantity
                        fmt.Printf("Jumlah item %s berhasil diupdate menjadi %d\n", foundItem.Name, foundItem.Quantity)
                        break
                    }
                }

                if foundItem == nil {
                    fmt.Println("Item tidak ditemukan. Pastikan nama yang dimasukkan sesuai dengan yang ada di menu.")
                }
            })

        case "edit":
            fmt.Print("Masukkan nama item yang ingin diubah jumlahnya: ")
            scanner.Scan()
            itemName := strings.TrimSpace(scanner.Text())

            fmt.Print("Masukkan jumlah baru: ")
            scanner.Scan()
            newQuantityInput := scanner.Text()

            safeInput(func() {
                newQuantity, err := strconv.Atoi(newQuantityInput)
                if err != nil {
                    panic("Jumlah tidak valid.")
                }

                var foundItem *MenuItem
                for i := range restaurantOrder.Menu {
                    if strings.EqualFold(strings.ReplaceAll(restaurantOrder.Menu[i].Name, " ", ""), strings.ReplaceAll(itemName, " ", "")) {
                        foundItem = &restaurantOrder.Menu[i]
                        if newQuantity >= 0 {
                            foundItem.Quantity = newQuantity
                            fmt.Printf("Jumlah item %s berhasil diubah menjadi %d\n", foundItem.Name, foundItem.Quantity)
                        } else {
                            fmt.Println("Jumlah harus bernilai positif atau nol.")
                        }
                        break
                    }
                }

                if foundItem == nil {
                    fmt.Println("Item tidak ditemukan. Pastikan nama yang dimasukkan sesuai dengan yang ada di pesanan.")
                }
            })

        case "delete":
            fmt.Print("Masukkan nama item yang ingin dihapus: ")
            scanner.Scan()
            itemName := strings.TrimSpace(scanner.Text())

            restaurantOrder.DeleteItem(itemName)

        default:
            fmt.Println("Perintah tidak dikenal. Silakan masukkan 'add', 'edit', 'delete', atau 'selesai'.")
        }

        // Menampilkan pesanan sementara tanpa membersihkan layar
        fmt.Println("\nPesanan Sementara Anda:")
        for _, item := range restaurantOrder.Menu {
            if item.Quantity > 0 {
                fmt.Printf("- %s (x%d)\n", item.Name, item.Quantity)
            }
        }
        printSeparator()
    }

    // Menampilkan pesanan terakhir
    clearScreen()
    printHeader("          RINCIAN PESANAN ANDA")
    for _, item := range restaurantOrder.Menu {
        if item.Quantity > 0 {
            fmt.Printf("- %s (x%d)\n", item.Name, item.Quantity)
        }
    }
    printSeparator()
    totalPrice := 0.0
    for _, item := range restaurantOrder.Menu {
        if item.Quantity > 0 {
            totalPrice += float64(item.Quantity) * item.Price
        }
    }
    fmt.Printf("Total Harga: Rp%.2f\n", totalPrice)
    printSeparator()

    // Menggunakan goroutine untuk memproses pesanan
    ch := make(chan string, len(restaurantOrder.Menu))
    var wg sync.WaitGroup

    for _, item := range restaurantOrder.Menu {
        if item.Quantity > 0 {
            wg.Add(1)
            go processOrder(item, ch, &wg)
        }
    }

    go func() {
        wg.Wait()
        close(ch)
    }()

    // Menerima hasil dari goroutine
    for msg := range ch {
        fmt.Println(msg)
    }
    printSeparator()

    // Pembayaran
    fmt.Print("Masukkan jumlah yang dibayar: ")
    scanner.Scan()
    paymentInput := scanner.Text()

    safeInput(func() {
        payment, err := validateNumericInput(paymentInput)
        if err != nil {
            panic("Input pembayaran tidak valid")
        }

        if payment < totalPrice {
            fmt.Println("Jumlah yang dibayar kurang.")
        } else {
            change := payment - totalPrice
            fmt.Printf("Jumlah yang dibayar valid. Kembalian: Rp%.2f\n", change)
        }
    })
    printSeparator()

    fmt.Println("Memproses pesanan di goroutine lain...")
}
