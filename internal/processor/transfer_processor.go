package processor

import (
	"encoding/json"
	"net/http"

	"github.com/golang/glog"

	"jianghai-hu/wallet-service/internal/service"
	"jianghai-hu/wallet-service/utils"
)

func TransferProcessor(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req struct {
		FromUserID int `json:"from_user_id"`
		ToUserID   int `json:"to_user_id"`
		Amount     int `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// process
	err := service.NewOrderService().Transfer(ctx, req.FromUserID, req.ToUserID, req.Amount)
	//

	code, msg := utils.ResolveError(err)
	resp := &Response{
		ErrorCode: code,
		ErrorMsg:  msg,
		Data:      nil,
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	if err = json.NewEncoder(w).Encode(resp); err != nil {
		glog.ErrorContextf(ctx, "write resp back to response err:%v", err)
	}
}