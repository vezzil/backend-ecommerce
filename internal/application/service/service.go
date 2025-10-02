package service

var (
	IUserService = &userService{}
	IProductService = &productService{}
	ICategoryService = &categoryService{}
	IOrderService = &orderService{}
	IPaymentService = &paymentService{}
	ICartService = &cartService{}
)