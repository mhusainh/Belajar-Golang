package main

import (
    "fmt"
    "strings"
)

// Struct untuk representasi item di menu
type MenuItem struct {
    Name     string
    Price    float64
    Quantity int
}

// Struct untuk menampung menu restoran
type RestaurantOrder struct {
    Menu []MenuItem
}

// Interface untuk mengelola menu
type Order interface {
    AddItem(item MenuItem)
    EditItem(name string, field string, newValue interface{})
    DeleteItem(name string)
}

// Method untuk menambah item ke menu
func (ro *RestaurantOrder) AddItem(item MenuItem) {
    ro.Menu = append(ro.Menu, item)
    fmt.Println("Item berhasil ditambahkan:", item.Name)
}

// Method untuk mengedit item di menu
func (ro *RestaurantOrder) EditItem(name string, field string, newValue interface{}) {
    for i, item := range ro.Menu {
        if strings.EqualFold(strings.ReplaceAll(item.Name, " ", ""), strings.ReplaceAll(name, " ", "")) {
            switch field {
            case "price":
                if value, ok := newValue.(float64); ok {
                    ro.Menu[i].Price = value
                    fmt.Printf("Harga item %s berhasil diupdate menjadi Rp%.2f\n", name, value)
                } else {
                    fmt.Println("Kesalahan: Harga harus berupa angka desimal (float64).")
                }
            case "quantity":
                if value, ok := newValue.(int); ok {
                    ro.Menu[i].Quantity = value
                    fmt.Printf("Jumlah item %s berhasil diupdate menjadi %d\n", name, value)
                } else {
                    fmt.Println("Kesalahan: Jumlah harus berupa angka bulat (int).")
                }
            default:
                fmt.Println("Field yang diminta tidak ditemukan.")
            }
            return
        }
    }
    fmt.Println("Item tidak ditemukan:", name)
}

// Method untuk menghapus item dari menu
func (ro *RestaurantOrder) DeleteItem(name string) {
    for i, item := range ro.Menu {
        if strings.EqualFold(strings.ReplaceAll(item.Name, " ", ""), strings.ReplaceAll(name, " ", "")) {
            ro.Menu = append(ro.Menu[:i], ro.Menu[i+1:]...)
            fmt.Println("Item berhasil dihapus:", name)
            return
        }
    }
    fmt.Println("Item tidak ditemukan:", name)
}
