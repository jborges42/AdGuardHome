package home

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/AdguardTeam/AdGuardHome/querylog"
)

func handleQueryLogEnable(w http.ResponseWriter, r *http.Request) {
	config.DNS.QueryLogEnabled = true
	httpUpdateConfigReloadDNSReturnOK(w, r)
}

func handleQueryLogDisable(w http.ResponseWriter, r *http.Request) {
	config.DNS.QueryLogEnabled = false
	httpUpdateConfigReloadDNSReturnOK(w, r)
}

func handleQueryLog(w http.ResponseWriter, r *http.Request) {
	data := config.queryLog.GetData()

	jsonVal, err := json.Marshal(data)
	if err != nil {
		httpError(w, http.StatusInternalServerError, "Couldn't marshal data into json: %s", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonVal)
	if err != nil {
		httpError(w, http.StatusInternalServerError, "Unable to write response json: %s", err)
	}
}

type qlogConfig struct {
	Interval uint `json:"interval"`
}

// Get configuration
func handleQueryLogInfo(w http.ResponseWriter, r *http.Request) {
	resp := qlogConfig{}
	resp.Interval = config.DNS.QueryLogInterval

	jsonVal, err := json.Marshal(resp)
	if err != nil {
		httpError(w, http.StatusInternalServerError, "json encode: %s", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonVal)
	if err != nil {
		httpError(w, http.StatusInternalServerError, "http write: %s", err)
	}
}

// Set configuration
func handleQueryLogConfig(w http.ResponseWriter, r *http.Request) {

	reqData := qlogConfig{}
	err := json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		httpError(w, http.StatusBadRequest, "json decode: %s", err)
		return
	}

	if !checkStatsInterval(reqData.Interval) {
		httpError(w, http.StatusBadRequest, "Unsupported interval")
		return
	}

	config.DNS.QueryLogInterval = reqData.Interval
	conf := querylog.Config{
		Interval: time.Duration(config.DNS.QueryLogInterval) * 24 * time.Hour,
	}
	config.queryLog.Configure(conf)

	returnOK(w)
}

// RegisterQueryLogHandlers - register handlers
func RegisterQueryLogHandlers() {
	httpRegister(http.MethodGet, "/control/querylog", handleQueryLog)
	httpRegister(http.MethodGet, "/control/querylog_info", handleQueryLogInfo)
	httpRegister(http.MethodPost, "/control/querylog_config", handleQueryLogConfig)
	httpRegister(http.MethodPost, "/control/querylog_enable", handleQueryLogEnable)
	httpRegister(http.MethodPost, "/control/querylog_disable", handleQueryLogDisable)
}
