
package handlers

import (
    "net/http"

    "github.com/Butternut01/ecommerce-platform/api-gateway/proto/inventory"
    "github.com/Butternut01/ecommerce-platform/api-gateway/proto/order"
    "github.com/Butternut01/ecommerce-platform/api-gateway/proto/user"
    "github.com/gin-gonic/gin"
)


type GatewayHandler struct {
    inventoryClient inventory.InventoryServiceClient
    orderClient     order.OrderServiceClient
    userClient      user.UserServiceClient
}

func NewGatewayHandler(
    inventoryClient inventory.InventoryServiceClient,
    orderClient order.OrderServiceClient,
    userClient user.UserServiceClient,
) *GatewayHandler {
    return &GatewayHandler{
        inventoryClient: inventoryClient,
        orderClient:    orderClient,
        userClient:     userClient,
    }
}

func (h *GatewayHandler) GetInventory(c *gin.Context) {
    // Logic to call inventory service and return response
}

func (h *GatewayHandler) CreateOrder(c *gin.Context) {
    // Logic to call order service and return response
}

func (h *GatewayHandler) RegisterUser(c *gin.Context) {
    // Logic to call user service and return response
}



func (h *GatewayHandler) HandleInventoryCreate(c *gin.Context) {
    // Implement logic for creating inventory
    c.JSON(http.StatusOK, gin.H{"message": "Inventory created"})
}

func (h *GatewayHandler) HandleInventoryGet(c *gin.Context) {
    // Implement logic for getting inventory
    c.JSON(http.StatusOK, gin.H{"message": "Inventory retrieved"})
}

func (h *GatewayHandler) HandleOrderCreate(c *gin.Context) {
    // Implement logic for creating order
    c.JSON(http.StatusOK, gin.H{"message": "Order created"})
}

func (h *GatewayHandler) HandleOrderGet(c *gin.Context) {
    // Implement logic for getting order
    c.JSON(http.StatusOK, gin.H{"message": "Order retrieved"})
}

func (h *GatewayHandler) HandleUserRegister(c *gin.Context) {
    // Implement logic for user registration
    c.JSON(http.StatusOK, gin.H{"message": "User registered"})
}

func (h *GatewayHandler) HandleUserLogin(c *gin.Context) {
    // Implement logic for user login
    c.JSON(http.StatusOK, gin.H{"message": "User logged in"})
}

func (h *GatewayHandler) CreateProduct(c *gin.Context) {
    // Implement the logic for creating a product
    c.JSON(200, gin.H{"message": "CreateProduct method is not yet implemented"})
}
func (h *GatewayHandler) GetProduct(c *gin.Context) {
    productID := c.Param("id")
    if productID == "" {
        c.JSON(400, gin.H{"error": "Product ID is required"})
        return
    }

    // Call the Inventory service to get the product
    product, err := h.inventoryClient.GetProduct(c, &inventory.GetProductRequest{Id: productID})
    if err != nil {
        c.JSON(500, gin.H{"error": "Failed to retrieve product"})
        return
    }

    c.JSON(200, product)
}
func (h *GatewayHandler) GetOrder(c *gin.Context) {
    orderID := c.Param("id")
    res, err := h.orderClient.GetOrder(c, &order.GetOrderRequest{OrderId: orderID})
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, res)
}
func (h *GatewayHandler) LoginUser(c *gin.Context) {
    // Implement the logic for user login here
    c.JSON(200, gin.H{"message": "LoginUser method is not yet implemented"})
}