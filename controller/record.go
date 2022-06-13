package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/rusdiy/amezink_task/config/utils"
	"github.com/rusdiy/amezink_task/service"
)

type RecordController struct {
	RecordService service.IRecordService
}

func (rc RecordController) GetMarks(w http.ResponseWriter, r *http.Request) {
	minCount, err := strconv.Atoi(r.FormValue("minCount"))
	if err != nil {
		utils.SendError(w, err, http.StatusBadRequest)
		return
	}
	maxCount, err := strconv.Atoi(r.FormValue("maxCount"))
	if err != nil {
		utils.SendError(w, err, http.StatusBadRequest)
		return
	}
	startDate, err := time.Parse("2006-01-02", r.FormValue("startDate"))
	if err != nil {
		utils.SendError(w, err, http.StatusBadRequest)
		return
	}
	endDate, err := time.Parse("2006-01-02", r.FormValue("endDate"))
	if err != nil {
		utils.SendError(w, err, http.StatusBadRequest)
		return
	}

	data, err := rc.RecordService.GetMarks(startDate.Format("2006-01-02"), endDate.Format("2006-01-02"), minCount, maxCount)
	if err != nil {
		utils.SendError(w, err, http.StatusBadRequest)
		return
	}

	resp := utils.JSONResponse{
		Message: "Success",
		Code:    0,
		Records: data,
	}
	resp.SendResponse(w, http.StatusOK)
}
