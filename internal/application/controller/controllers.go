package controller

// Controllers holds all the controller instances
var (
	// User related
	UserCtrl  = &UserController{}

	// Product related
	ProductCtrl  = &ProductController{}
	CategoryCtrl = &CategoryController{}

	// Order related
	OrderCtrl   = &OrderController{}
	PaymentCtrl = &PaymentController{}

	// Cart related
	CartCtrl = &CartController{}
)
