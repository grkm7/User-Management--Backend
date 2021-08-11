package product_api

import (
	"config"
	"encoding/json"
	"net/http"
	"strconv"
)

func FindAll(response http.ResponseWriter, request *http.Request) {
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		products, err2 := productModel.FindAll()
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err.Error())
		} else {
			respondWithJson(response, http.StatusOK, products)

		}
	}
}

func SearchID(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		products, err2 := productModel.SearchID(id)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err.Error())
		} else {
			respondWithJson(response, http.StatusOK, products)

		}
	}
}

func Delete(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	sid := vars["id"]
	id, _ := strconv.ParseInt(sid, 10, 64)
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		RowsAffected, err2 := productModel.Delete(id)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err.Error())
		} else {
			respondWithJson(response, http.StatusOK, map[string]int64{
				"RowsAffected": RowsAffected,
			})

		}
	}
}

func Create(response http.ResponseWriter, request *http.Request) {
	var product entities.Product
	err := json.NewDecoder(request.Body).Decode(&product)
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		err2 := productModel.Create(&product)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err.Error())
		} else {
			respondWithJson(response, http.StatusOK, product)

		}
	}
}

func Update(response http.ResponseWriter, request *http.Request) {
	var product entities.Product
	err := json.NewDecoder(request.Body).Decode(&product)
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		_, err2 := productModel.Update(&product)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err.Error())
		} else {
			respondWithJson(response, http.StatusOK, product)

		}
	}
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
