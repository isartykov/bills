package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/isartykov/user"
)

// @Summary Create bill
// @Security ApiKeyAuth
// @Tags bills
// @Description create bill
// @ID create-bill
// @Accept json
// @Produce json
// @Param input body user.Bill true "bill info"
// @Success 200 {integer} 1
// @Failure 400,404 {object} Error
// @Failure 500 {object} Error
// @Failure default {object} Error
// @Router /api/bills [post]

func (h *Handler) createBill(w http.ResponseWriter, r *http.Request, userID int) {
	var input user.Bill

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Bill.Create(userID, input)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id": id,
	})
}

type getAllBillsResponse struct {
	Data []user.Bill `json:"data"`
}

// @Summary GetAll bill
// @Security ApiKeyAuth
// @Tags bills
// @Description get all bills
// @ID get-bill
// @Accept json
// @Produce json
// @Success 200 {integer} 1
// @Failure 400,404 {object} Error
// @Failure 500 {object} Error
// @Failure default {object} Error
// @Router /api/bills [get]

func (h *Handler) getAllBills(w http.ResponseWriter, r *http.Request, userID int) {
	bills, err := h.services.Bill.GetAll(userID)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(getAllBillsResponse{
		Data: bills,
	})
}

func (h *Handler) getBill(w http.ResponseWriter, r *http.Request, userID int) {
	param := mux.Vars(r)

	id, err := strconv.Atoi(param["id"])
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, "invalid id param")
		return
	}

	bill, err := h.services.Bill.GetByID(userID, id)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(bill)
}

func (h *Handler) updateBill(w http.ResponseWriter, r *http.Request, userID int) {
	param := mux.Vars(r)

	id, err := strconv.Atoi(param["id"])
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, "invalid id param")
		return
	}

	var input user.UpdateBillInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Bill.Update(userID, id, input)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(statusResponse{
		Status: "ok",
	})
}

func (h *Handler) deleteBill(w http.ResponseWriter, r *http.Request, userID int) {
	param := mux.Vars(r)

	id, err := strconv.Atoi(param["id"])
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.Bill.Delete(userID, id)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(statusResponse{
		Status: "ok",
	})
}

func (h *Handler) random(w http.ResponseWriter, r *http.Request, userID int) {
	bill, err := h.services.Bill.GetRandomBill(userID)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(bill)
}

func (h *Handler) calculate(w http.ResponseWriter, r *http.Request, userID int) {
	var input user.Times

	json.NewDecoder(r.Body).Decode(&input)

	bills, err := h.services.Bill.GetAllCalculate(userID, input)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(getAllBillsResponse{
		Data: bills,
	})
}
