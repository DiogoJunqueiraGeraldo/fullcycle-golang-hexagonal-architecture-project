package handler

import (
	"encoding/json"
	"net/http"

	"github.com/DiogoJunqueiraGeraldo/fullcycle-golang-hexagonal-architecture-project/adapters/dto"
	"github.com/DiogoJunqueiraGeraldo/fullcycle-golang-hexagonal-architecture-project/application"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func MakeProductHandler(r *mux.Router, n *negroni.Negroni, service application.ProductServiceInterface) {
	r.Handle("/product/{id}", n.With(negroni.Wrap(getProduct(service)))).Methods("GET", "OPTIONS")
	r.Handle("/product", n.With(negroni.Wrap(createProduct(service)))).Methods("POST", "OPTIONS")
	r.Handle("/product/{id}", n.With(negroni.Wrap(changeProduct(service)))).Methods("PUT", "OPTIONS")
}

func changeProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		id := vars["id"]

		var updateProductStatusDTO dto.UpdateProductStatusDTO
		err := json.NewDecoder(r.Body).Decode(&updateProductStatusDTO)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(jsonError(err.Error()))
			return
		}

		var product application.ProductInterface
		product, err = service.Get(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write(jsonError(err.Error()))
			return
		}

		if updateProductStatusDTO.Status == "enabled" {
			product, err = service.Enable(product)
		} else if updateProductStatusDTO.Status == "disabled" {
			product, err = service.Disable(product)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(jsonError("Invalid status change operation"))
			return
		}

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}

		err = json.NewEncoder(w).Encode(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}
	})
}

func createProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var createProductDTO dto.CreateProductDTO
		err := json.NewDecoder(r.Body).Decode(&createProductDTO)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(jsonError(err.Error()))
			return
		}

		product, err := createProductDTO.Bind()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}

		product, err = service.Create(product.GetName(), product.GetPrice())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}

		err = json.NewEncoder(w).Encode(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}
	})
}

func getProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		id := vars["id"]

		product, err := service.Get(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write(jsonError(err.Error()))
			return
		}

		err = json.NewEncoder(w).Encode(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}
	})
}
