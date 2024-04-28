package file

import (
	"devtask/internal/model"
	"time"
)

type OrderDTO struct {
	IdOrder     int64             // ID of an order
	IdClient    int64             // ID of a client
	IsGiven     bool              // Flag whether an order was given to the client or not
	IsReturned  bool              // Flag whether an order was returned to the courier or not
	IsRefunded  bool              // Flag whether an order was refunded to the pickup point or not
	ExpireDate  time.Time         // The expiration date of an order in pickup point. The client can get an order UNTIL this date.
	WeightGrams int64             // The weight of the order in grams
	PriceKopeck int64             // Price of the order in kopeck depends on the type of the package
	PackageType model.PackageType // Type of the package (0-packet, 1-box, 2-tape)
}
