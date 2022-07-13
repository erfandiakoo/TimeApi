package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/erfandiakoo/TimeApi/models"
	"github.com/gofiber/fiber/v2"
)

func GetTime(c *fiber.Ctx) error {
	dt := time.Now()
	return c.JSON(&models.ResponseModel{
		Data:    map[string]interface{}{"result": dt},
		Message: "Ok",
		Code:    200,
	})
}

func GetTimeUnix(c *fiber.Ctx) error {
	dt := time.Now().UnixNano()
	return c.JSON(&models.ResponseModel{
		Data:    map[string]interface{}{"result": dt},
		Message: "Ok",
		Code:    200,
	})
}

func GetTimeUtc(c *fiber.Ctx) error {
	dt := time.Now().UTC()
	return c.JSON(&models.ResponseModel{
		Data:    map[string]interface{}{"result": dt},
		Message: "Ok",
		Code:    200,
	})
}

func GetTimeIP(c *fiber.Ctx) error {
	timezone := GetIPZone(string(c.IP()))

	loc, _ := time.LoadLocation(timezone)
	dt := time.Now().In(loc)

	return c.JSON(&models.ResponseModel{
		Data:    map[string]interface{}{"result": dt},
		Message: "Ok",
		Code:    200,
	})
}

func GetTimeTimeZone(c *fiber.Ctx) error {
	timezone := c.Params("area")

	fmt.Println(timezone)

	//loc, _ := time.LoadLocation(timezone)
	dt := time.Now()

	return c.JSON(&models.ResponseModel{
		Data:    map[string]interface{}{"result": dt},
		Message: "Ok",
		Code:    200,
	})

}

func GetIPZone(ip string) string {
	resp, err := http.Get("http://ip-api.com/json/" + ip + "?fields=256")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
	return string(body)
}
