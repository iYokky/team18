package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/team18/app/ent"
	"github.com/team18/app/ent/checkout"
)

// CheckoutController defines the struct for the checkout controller
type CheckoutController struct {
	client *ent.Client
	router gin.IRouter
}

type CheckOut struct {
	CheckoutDate    string
	StatussID       int
	CounterstaffsID int
	CheckinsID      int
}

// CreateCheckout handles POST requests for adding checkout entities
// @Summary Create checkout
// @Description Create checkout
// @ID create-checkout
// @Accept   json
// @Produce  json
// @Param checkout body ent.Checkout true "Checkout entity"
// @Success 200 {object} ent.Checkout
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /checkouts [post]
func (ctl *CheckoutController) CreateCheckout(c *gin.Context) {
	obj := CheckOut{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "checkout binding failed",
		})
		return
	}

	t1 := time.Now()
	t2 := t1.Format("2006-01-02T15:04:05Z07:00")
	time, err := time.Parse(time.RFC3339, t2)
	u, err := ctl.client.Checkout.
		Create().
		SetCheckoutDate(time).
		SetStatussID(obj.StatussID).
		SetCounterstaffsID(obj.CounterstaffsID).
		SetCheckinsID(obj.CheckinsID).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, u)
}

// ListCheckout handles request to get a list of checkout entities
// @Summary List checkout entities
// @Description list checkout entities
// @ID list-checkout
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Checkout
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /checkouts [get]
func (ctl *CheckoutController) ListCheckout(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	checkouts, err := ctl.client.Checkout.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, checkouts)
}

// GetCheckout handles GET requests to retrieve a checkout entity
// @Summary Get a checkout entity by ID
// @Description get checkout by ID
// @ID get-checkout
// @Produce  json
// @Param id path int true "Checkout ID"
// @Success 200 {object} ent.Checkout
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /checkouts/{id} [get]
func (ctl *CheckoutController) GetCheckout(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.Checkout.
		Query().
		Where(checkout.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// DeleteCheckout handles DELETE requests to delete a checkout entity
// @Summary Delete a checkout entity by ID
// @Description get checkout by ID
// @ID delete-checkout
// @Produce  json
// @Param id path int true "Checkout ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /checkouts/{id} [delete]
func (ctl *CheckoutController) DeleteCheckout(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Checkout.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// UpdateCheckout handles PUT requests to update a checkout entity
// @Summary Update a checkout entity by ID
// @Description update checkout by ID
// @ID update-checkout
// @Accept   json
// @Produce  json
// @Param id path int true "Checkout ID"
// @Param checkout body ent.Checkout true "Checkout entity"
// @Success 200 {object} ent.Checkout
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /checkouts/{id} [put]
func (ctl *CheckoutController) UpdateCheckout(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Checkout{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "checkout binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.Checkout.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, u)
}

// NewCheckoutController creates and registers handles for the checkout controller
func NewCheckoutController(router gin.IRouter, client *ent.Client) *CheckoutController {
	uc := &CheckoutController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitCheckoutController registers routes to the main engine
func (ctl *CheckoutController) register() {
	checkouts := ctl.router.Group("/checkouts")

	checkouts.GET("", ctl.ListCheckout)

	// CRUD
	checkouts.POST("", ctl.CreateCheckout)
	checkouts.GET(":id", ctl.GetCheckout)
	checkouts.PUT(":id", ctl.UpdateCheckout)
	checkouts.DELETE(":id", ctl.DeleteCheckout)
}
