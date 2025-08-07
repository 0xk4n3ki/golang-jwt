package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/0xk4n3ki/golang-jwt/database"
	helper "github.com/0xk4n3ki/golang-jwt/helpers"
	"github.com/0xk4n3ki/golang-jwt/models"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func HashPassword()

func VerifyPassword()

func Signup() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user models.User
		if err := ctx.BindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return
		}

		validateErr := validate.Struct(user)
		if validateErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error":validateErr.Error()})
			return
		}

		var c, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		
		count, err := userCollection.CountDocuments(c, bson.M{"email":user.Email})
		defer cancel()
		if err != nil {
			log.Panic(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error":"error occured while checking for the email"})
		}

		count, err = userCollection.CountDocuments(c, bson.M{"phone":user.Phone})
		defer cancel()
		if err != nil {
			log.Panic(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error":"error occured while checking for the phone"})
		}

		if count > 0 {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error":"this email or phone number already exists"})
		}
	}
}

func Login()

func GetUsers()

func GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.Param("user_id")

		if err := helper.MatchUserTypeToUid(ctx, userId); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var c, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User

		err := userCollection.FindOne(c, bson.M{"user_id":userId}).Decode(&user)
		defer cancel()

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, user)
	}
}
