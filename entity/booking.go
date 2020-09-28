package entity

type BookingHTTP struct {
	SenderName  string `json:"sender_name"`
	SenderPhone string `json:"sender_phone"`
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
	Weight      int    `json:"weight"`
	Service     string `json:"service"`
}

type BookingDB struct {
	BookingID   int64  `db:"booking_id"`
	Resi        string `db:"resi"`
	Price       int64  `db:"price"`
	SenderName  string `db:"sender_name"`
	SenderPhone string `db:"sender_phone"`
	Origin      string `db:"origin"`
	Destination string `db:"destination"`
	Weight      int    `db:"weight"`
	Service     string `db:"service"`
}
