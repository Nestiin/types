package types

import "github.com/jinzhu/gorm"

type Order struct {
	gorm.Model
	MerchantID           int    `json:"merchantId" gorm:"type:int;column:merchant_id"`
	UserID               int    `json:"userId"  gorm:"type:int;column:user_id"`
	PTID                 int    `json:"productTypeId" gorm:"type:int;column:product_type_id"`
	PCID                 int    `json:"productCategoryId" gorm:"type:int;column:product_category_id"`
	ProductID            int    `json:"productId" gorm:"type:int;column:product_id"`
	UOMID                int    `json:"uomId" gorm:"type:int;column:uom_id"`
	OrderCode            string `json:"-" gorm:"type:varchar(50);column:order_code"`
	UNITPRICE            int    `json:"unitPrice" gorm:"type:int;column:unit_price"`
	Quantity             int    `json:"quantity" gorm:"type:int;column:quantity"`
	DiscountAmount       int    `json:"discountAmount" gorm:"type:int;column:discount_amount"`
	OrderCost            int    `json:"orderCost" gorm:"type:int;column:order_cost"`
	DeliveryCost         int    `json:"deliveryCost" gorm:"type:int;column:delivery_cost"`
	TotalCost            int    `json:"totalCost" gorm:"type:int;column:total_cost"`
	PaymentMode          string `json:"paymentMode" gorm:"type:varchar(250);column:payment_mode"`
	PaymentRefID         int    `json:"paymentRefId" gorm:"type:int;column:payment_ref_id"`
	BillingCustomerName  string `json:"billingCustomerName" gorm:"type:varchar(50);column:billing_customer_name"`
	BillingAddress       string `json:"billingAddress" gorm:"type:varchar(50);column:billing_address"`
	BillingCity          string `json:"billingCity" gorm:"type:varchar(50);column:billing_city"`
	BillingPincode       string `json:"billingPincode" gorm:"type:varchar(50);column:billing_pincode"`
	BillingState         string `json:"billingState" gorm:"type:varchar(50);column:billing_state"`
	BillingCountry       string `json:"billingCountry" gorm:"type:varchar(50);column:billing_country"`
	BillingEmail         string `json:"billingEmail" gorm:"type:varchar(50);column:billing_email"`
	BillingPhone         string `json:"billingPhone" gorm:"type:varchar(50);column:billing_phone"`
	ShippingISBilling    bool   `json:"shippingIsBilling"  gorm:"type:bool;column:shipping_is_billing"`
	ShippingCustomerName string `json:"shippingCustomerName" gorm:"type:varchar(50);column:shipping_customer_name"`
	ShippingAddress      string `json:"shippingAddress" gorm:"type:varchar(50);column:shipping_address"`
	ShippingCity         string `json:"shippingCity" gorm:"type:varchar(50);column:shipping_city"`
	ShippingPincode      string `json:"shippingPincode" gorm:"type:varchar(50);column:shipping_pincode"`
	ShippingState        string `json:"shippingState" gorm:"type:varchar(50);column:shipping_state"`
	ShippingCountry      string `json:"shippingCountry" gorm:"type:varchar(50);column:shipping_country"`
	ShippingEmail        string `json:"shippingEmail" gorm:"type:varchar(50);column:shipping_email"`
	ShippingPhone        string `json:"shippingPhone" gorm:"type:varchar(50);column:shipping_phone"`
	Status               string `json:"status" gorm:"type:varchar(100);column:status"`
	CancelReason         string `json:"cancelReason" gorm:"type:varchar(5000);column:cancel_reason"`
	IsShipmentDone       bool   `json:"isShipmentDone" gorm:"type:varchar(5000);column:is_shipment_done"`
}

func (Order) TableName() string {
	return "order_info"
}