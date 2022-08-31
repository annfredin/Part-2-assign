package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"money-transfer/lib"
	"money-transfer/models"

	"github.com/gin-gonic/gin"
)

//Ping ...
var Ping = func(c *gin.Context) {
	c.String(200, "pong")
}

//Reset...
var Reset = func(c *gin.Context) {
	
	err:= updateBalance("user1.txt", "10000")
	if err != nil {
		log.Println("ERROR: Writing user1=>", err.Error())
		c.JSON(http.StatusOK, "0")
		return
	}

	err= updateBalance("user2.txt", "10000")
	if err != nil {
		log.Println("ERROR: Writing user2=>", err.Error())
		c.JSON(http.StatusOK, "0")
		return
	}

	c.JSON(http.StatusOK, "1")
}


func updateBalance(fileName, data string) error{
	err:= lib.WriteFile(fileName, data)
	return err
}

// Transfer ...
var Transfer = func(c *gin.Context) {
	var requestData models.UserTransferRequest
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, "0")
		return
	}

	//getting bal
	data, err:= lib.ReadFile(fmt.Sprintf("user%d.txt",requestData.ToUserID))
	if err != nil {
		log.Println("ERROR: Reading user1=>", err.Error())
		c.JSON(http.StatusOK, "0")
		return
	}
	toUserBalance := float64(0)
	if d, err := strconv.ParseFloat(data, 64); err == nil {
		toUserBalance = d
	}

	data, err= lib.ReadFile(fmt.Sprintf("user%d.txt",requestData.FromUserID))
	if err != nil {
		log.Println("ERROR: Reading user1=>", err.Error())
		c.JSON(http.StatusOK, "0")
		return
	}
	if d, err := strconv.ParseFloat(data, 64); err == nil {
		
		err:= updateBalance(fmt.Sprintf("user%d.txt",requestData.ToUserID), fmt.Sprintf("%v", toUserBalance + requestData.Amount))
		if err != nil {
			log.Println("ERROR: Writing =>", fmt.Sprintf("user%d.txt Error",requestData.ToUserID), err.Error())
			c.JSON(http.StatusOK, "0")
			return
		}

		err= updateBalance(fmt.Sprintf("user%d.txt",requestData.FromUserID), fmt.Sprintf("%v", d - requestData.Amount))
		if err != nil {
			log.Println("ERROR: Writing =>", fmt.Sprintf("user%d.txt Error",requestData.FromUserID), err.Error())
			c.JSON(http.StatusOK, "0")
			return
		}

	}

	
	c.JSON(http.StatusOK, "1")
}

//GetBalance ...
var GetBalance = func(c *gin.Context) {
	var respList []models.UserBalanceResp
	data, err:= lib.ReadFile("user1.txt")
	if err != nil {
		log.Println("ERROR: Reading user1=>", err.Error())
		c.JSON(http.StatusOK, "0")
		return
	}
	if d, err := strconv.ParseFloat(data, 64); err == nil {
		respList =append(respList, models.UserBalanceResp{UserID: 1, Balance: d})
	}
	

	data, err= lib.ReadFile("user2.txt")
	if err != nil {
		log.Println("ERROR: Reading user2=>", err.Error())
		c.JSON(http.StatusOK, "0")
		return
	}
	if d, err := strconv.ParseFloat(data, 64); err == nil {
		respList =append(respList, models.UserBalanceResp{UserID: 2, Balance: d})
	}
	
	c.JSON(http.StatusOK, respList)
}
