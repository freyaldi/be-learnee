package dto

type CheckoutRequest struct {
	CartId      []int   `json:"cart_id" validate:"required"`
	VoucherCode *string `json:"voucher_code"`
}

type InvoiceResponse struct {
	Id              int     `json:"id"`
	Name            string  `json:"name"`
	VoucherCode     string  `json:"voucher_code"`
	BenefitDiscount float32 `json:"benefit_discount"`
	VoucherDiscount float32 `json:"voucher_discount"`
	Status          string  `json:"status"`
	Price           float64 `json:"price"`
	Discount        float64 `json:"discount"`
	Cost            float64 `json:"cost"`
}
