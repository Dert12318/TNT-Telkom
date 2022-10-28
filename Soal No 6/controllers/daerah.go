package controllers

import (
	"CartApp/models"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"gorm.io/gorm"
)

type DaerahController struct {
	db *gorm.DB
}

func NewDaerahController(r *gin.RouterGroup, db *gorm.DB) {
	handler := &DaerahController{db}

	r.GET("/daerah", handler.InsertDaerah)
	r.GET("/kotakabupaten", handler.GetKotaKab)
}

func (d DaerahController) InsertDaerah(c *gin.Context) {
	var res models.DaerahApi

	err := json.Unmarshal([]byte(string(Data)), &res)
	if err != nil {
		fmt.Println(err)
		return
	}

	tx := d.db.Begin()
	for _, v := range res.Provinsi {
		var temp models.Provinsi
		temp.Nama = v.Nama
		temp.ID = v.ID
		err := tx.Debug().Create(&temp).Error
		if err != nil {
			tx.Rollback()
			fmt.Println(err)
			c.JSON(400, "bad request")
			return
		}

		var temp2 models.KotaKabupaten
		client := resty.New()
		client.SetDebug(true)
		client.SetContentLength(true)

		idRes := strconv.Itoa(v.ID)
		url := fmt.Sprintf("https://dev.farizdotid.com/api/daerahindonesia/kota?id_provinsi=%s", idRes)
		resp, err := client.R().
			EnableTrace().
			SetResult(&temp2).
			Get(url)

		f := resp.Body()
		response := string(f)

		if errs := json.Unmarshal([]byte(string(response)), &temp2); errs != nil {
			fmt.Println(errs.Error())
			c.JSON(400, "bad request")
			return
		}

		for _, detail := range temp2.KotaKabupaten {
			errs := tx.Debug().Create(&models.KotaKab{
				ID:         detail.ID,
				IdProvinsi: detail.IdProvinsi,
				Nama:       strings.ToUpper(detail.Nama),
			}).Error

			if errs != nil {
				fmt.Println(errs.Error())
				c.JSON(400, "bad request")
			}
		}
	}

	tx.Commit()

	c.JSON(200, "OK")
}

func (d DaerahController) GetKotaKab(c *gin.Context) {
	kotaKab := c.Query("nama_kota_kab")
	var res []models.KotaKab
	err := d.db.Debug().Where("nama LIKE ?", "%"+strings.ToUpper(kotaKab)+"%").Find(&res).Error
	if err != nil {
		fmt.Println(err)
		c.JSON(400, "bad request")
		return
	}

	c.JSON(200, res)
}

var Data = `{
"provinsi": [
{
"id": 11,
"nama": "Aceh"
},
{
"id": 12,
"nama": "Sumatera Utara"
},
{
"id": 13,
"nama": "Sumatera Barat"
},
{
"id": 14,
"nama": "Riau"
},
{
"id": 15,
"nama": "Jambi"
},
{
"id": 16,
"nama": "Sumatera Selatan"
},
{
"id": 17,
"nama": "Bengkulu"
},
{
"id": 18,
"nama": "Lampung"
},
{
"id": 19,
"nama": "Kepulauan Bangka Belitung"
},
{
"id": 21,
"nama": "Kepulauan Riau"
},
{
"id": 31,
"nama": "Dki Jakarta"
},
{
"id": 32,
"nama": "Jawa Barat"
},
{
"id": 33,
"nama": "Jawa Tengah"
},
{
"id": 34,
"nama": "Di Yogyakarta"
},
{
"id": 35,
"nama": "Jawa Timur"
},
{
"id": 36,
"nama": "Banten"
},
{
"id": 51,
"nama": "Bali"
},
{
"id": 52,
"nama": "Nusa Tenggara Barat"
},
{
"id": 53,
"nama": "Nusa Tenggara Timur"
},
{
"id": 61,
"nama": "Kalimantan Barat"
},
{
"id": 62,
"nama": "Kalimantan Tengah"
},
{
"id": 63,
"nama": "Kalimantan Selatan"
},
{
"id": 64,
"nama": "Kalimantan Timur"
},
{
"id": 65,
"nama": "Kalimantan Utara"
},
{
"id": 71,
"nama": "Sulawesi Utara"
},
{
"id": 72,
"nama": "Sulawesi Tengah"
},
{
"id": 73,
"nama": "Sulawesi Selatan"
},
{
"id": 74,
"nama": "Sulawesi Tenggara"
},
{
"id": 75,
"nama": "Gorontalo"
},
{
"id": 76,
"nama": "Sulawesi Barat"
},
{
"id": 81,
"nama": "Maluku"
},
{
"id": 82,
"nama": "Maluku Utara"
},
{
"id": 91,
"nama": "Papua Barat"
},
{
"id": 94,
"nama": "Papua"
}
]
}`
