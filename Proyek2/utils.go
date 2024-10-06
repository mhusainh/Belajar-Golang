package main

import (
    "fmt"
    "os"
    "os/exec"
    "runtime"
    "strings"
    "regexp"
    "strconv"
    "encoding/base64"
)

// Fungsi untuk membersihkan layar terminal
func clearScreen() {
    var cmd *exec.Cmd
    switch runtime.GOOS {
    case "windows":
        cmd = exec.Command("cmd", "/c", "cls")
    default:
        cmd = exec.Command("clear")
    }
    cmd.Stdout = os.Stdout
    cmd.Run()
}

// Fungsi untuk mencetak garis pemisah
func printSeparator() {
    fmt.Println(strings.Repeat("=", 40))
}

// Fungsi untuk mencetak header dengan garis pemisah
func printHeader(title string) {
    printSeparator()
    fmt.Printf("%s\n", title)
    printSeparator()
}

// Fungsi validasi input untuk angka
func validateNumericInput(input string) (float64, error) {
    input = strings.TrimSpace(input) // Menghapus spasi
    re := regexp.MustCompile(`^\d+(\.\d+)?$`) // Mengizinkan angka dan angka desimal
    if !re.MatchString(input) {
        return 0, fmt.Errorf("input tidak valid, harus berupa angka")
    }
    parsedValue, err := strconv.ParseFloat(input, 64)
    if err != nil {
        return 0, fmt.Errorf("input tidak dapat dikonversi menjadi angka")
    }
    return parsedValue, nil
}

// Fungsi untuk encoding informasi pesanan ke dalam base64
func encodeOrderDetails(details string) string {
    return base64.StdEncoding.EncodeToString([]byte(details))
}

// Fungsi untuk menangani error menggunakan recover
func safeInput(handler func()) {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from error:", r)
        }
    }()
    handler()
}
