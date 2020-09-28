package repository

const (
	queryInsertBooking = `
	INSERT INTO booking(
		resi,
		price,  
		sender_name,
		sender_phone, 
		origin,
		destination,
		weight,
		service     
	) VALUES (
		:resi,
		:price,
		:sender_name,
		:sender_phone,
		:origin,
		:destination,
		:weight,
		:service
	);`
)
