package services

import (
	"context"
	"errors"
	"quoation-backend/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type QuoteServiceImpl struct {
	quotecollection *mongo.Collection
	ctx   context.Context

}
func NewQuoteService(quotecollection *mongo.Collection,ctx   context.Context)QuoteService{
	return &QuoteServiceImpl{
		quotecollection: quotecollection,
	ctx :  ctx,
	}
}
func (q *QuoteServiceImpl) CreateQuote(quote *models.Quote) error{
	_, err := q.quotecollection.InsertOne(q.ctx,quote)
  return err
}

func (q *QuoteServiceImpl) GetQuote(id *string) (*models.Quote,error){
	var quotename *models.Quote
	query := bson.D{bson.E{Key:"id",Value:id}}
	err := q.quotecollection.FindOne(q.ctx,query).Decode(&quotename)
	return quotename,err
  }

  func (q *QuoteServiceImpl) GetAll() ([]*models.Quote,error) {
	var quotes []*models.Quote
	cursor,err := q.quotecollection.Find(q.ctx,bson.D{{}})
	if err!=nil{
		return nil,err
	}
	defer cursor.Close(q.ctx)
    for cursor.Next(q.ctx){
		var quote models.Quote
		err := cursor.Decode(&quote)
		if err != nil{
			return nil,err
		}
		quotes = append(quotes,&quote)
	}
	if err:= cursor.Err(); err!=nil {
		return nil,err
	}
	// cursor.Close(q.ctx)
	if len(quotes) == 0{
		return nil,errors.New("No quotes found")
	}
  return quotes,nil
	// return nil,nil
  }

  func (q *QuoteServiceImpl) UpdateQuote(quote *models.Quote) error{
	filter := bson.D{bson.E{Key:"id",Value:quote.ID}}
	update := bson.D{bson.E{Key:"$set",Value:bson.D{bson.E{Key:"id",Value:quote.ID},bson.E{Key:"name",Value:quote.Name},bson.E{Key:"expiryDate",Value:quote.ExpiryDate},bson.E{Key:"totalAmount",Value:quote.TotalAmount},bson.E{Key:"files",Value:quote.Files},bson.E{Key:"tables",Value:quote.Tables}}}}
	result,_ :=q.quotecollection.UpdateOne(q.ctx,filter,update)
	if result.MatchedCount !=1 {
		 return errors.New("No matched document found for updating")
		
	}
	return nil
  }
  func (q *QuoteServiceImpl) DeleteQuote(quoteID *string) error {
	filter := bson.D{bson.E{Key:"id",Value:quoteID}}
	result,_:=q.quotecollection.DeleteOne(q.ctx,filter)
	if result.DeletedCount !=1 {
		return errors.New("No matched document found for updating")
	   
   }
	return nil
  }







