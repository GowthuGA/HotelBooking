package models

import (
    "time"
)


type Booking struct {
    ID         uint      `json:"id" gorm:"primaryKey"`                
    CustomerID uint      `json:"customer_id" gorm:"not null"`        
    RoomID     uint      `json:"room_id" gorm:"not null"`             
    HotelID    uint      `json:"hotel_id" gorm:"not null"`           
    CheckIn    time.Time `json:"check_in" gorm:"not null"`           
    CheckOut   time.Time `json:"check_out" gorm:"not null"`           
    BookingDate time.Time `json:"booking_date" gorm:"not null"`        
    TotalPrice float64   `json:"total_price" gorm:"type:decimal(10,2);not null"` 
    Status     string    `json:"status" gorm:"type:enum('Pending', 'Confirmed', 'Cancelled');default:'Pending'"` // Booking status
}
