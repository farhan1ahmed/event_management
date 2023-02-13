package dataservice

import (
	"context"
	"event_ticket_service/models"
	"log"
	"time"
)

func (q *Queries)AddToCartInDB(ctx context.Context, reqPayload models.AddToCart) error{
	cartItem := models.Cart{
		UserID: reqPayload.UserID,
		TicketTypeID: reqPayload.TicketTypeID,
		Quantity: reqPayload.Quantity,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	result := q.db.Model(&cartItem).Create(&cartItem)
	if result.Error != nil{
		log.Println(result.Error)
		return result.Error
	}
	return nil
}


func (q *Queries)UpdateCartItemInDB(ctx context.Context, reqPayload models.UpdateCartItem) error{
	cartItem := models.Cart{
		ID: reqPayload.CartItemID,
		Quantity: reqPayload.Quantity,
	}
	result := q.db.Model(&cartItem).Updates(&cartItem)
	if result.Error != nil{
		log.Println(result.Error)
		return result.Error
	}
	return nil
}


func (q *Queries)DeleteCartItemFromDB(ctx context.Context, reqPayload models.DeleteCartItem) error{
	cartItem := models.Cart{}
	result := q.db.Model(&cartItem).Delete(&cartItem, reqPayload.CartItemID)
	if result.Error != nil{
		log.Println(result.Error)
		return result.Error
	}
	return nil
}

func (q *Queries)GetCartItemsFromDB(ctx context.Context) ([]models.Cart, error){
	var cartItems []models.Cart
	result := q.db.Find(&cartItems)
	if result.Error != nil{
		log.Println(result.Error)
		return nil, result.Error
	}
	return cartItems, nil
}