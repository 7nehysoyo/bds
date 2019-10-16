package ripple

import (
	"errors"
	"github.com/jdcloud-bds/bds/common/json"
	"github.com/jdcloud-bds/bds/common/jsonrpc"
	"github.com/jdcloud-bds/bds/common/log"
	"strconv"
	"strings"
)

type rpcHandler struct {
	client *jsonrpc.Client
}

func newRPCHandler(c *jsonrpc.Client) (*rpcHandler, error) {
	h := new(rpcHandler)
	h.client = c
	return h, nil
}

type CompleteLedgers struct {
	startLedger int64
	endLedger   int64
}

func (h *rpcHandler) GetCompleteLedgers() (map[int]*CompleteLedgers, error) {
	totalCompleteLedgers := make(map[int]*CompleteLedgers, 0)

	res, err := h.client.CallXRP("server_info")
	if err != nil {
		return nil, err
	}
	data := string(res)
	clStr := json.Get(data, "result.info.complete_ledgers").String()
	//clStr demo:"47025320,47025422-47025425,47025527-47025528,47025629-47025661,47025763,47025864-47025877"
	log.Info("splitter ripple: get completed ledgers: %s\n\n", clStr)
	if strings.ToLower(clStr) == "empty" {
		return nil, nil
	}

	clList := strings.Split(clStr, ",")
	for i, v := range clList {
		cl := new(CompleteLedgers)
		index := strings.Index(v, "-")
		if index > 0 {
			cl.startLedger, _ = strconv.ParseInt(v[:index], 10, 64)
			cl.endLedger, _ = strconv.ParseInt(v[index+1:], 10, 64)
		} else {
			cl.startLedger, _ = strconv.ParseInt(v, 10, 64)
			cl.endLedger, _ = strconv.ParseInt(v, 10, 64)
		}
		totalCompleteLedgers[i] = cl
	}
	return totalCompleteLedgers, nil
}

func (h *rpcHandler) SendBatchLedger(start, end int64) error {
	defer stats.Add(MetricRPCCall, 1)
	params := make(map[string]int64, 0)
	params[ParamStartLedgerIndex] = start
	params[ParamEndLedgerIndex] = end

	result, err := h.client.CallXRP("send_batch_ledger", params)
	log.Debug("splitter ripple: send batch ledger result: %s", string(result))
	if err != nil {
		return err
	}
	status := json.Get(string(result), "result.status").String()
	errorMessage := json.Get(string(result), "result.error_message").String()
	if status == "error" {
		log.Error("splitter ripple: send batch ledger result error : %s", errorMessage)
		log.DetailDebug("splitter ripple: send batch ledger result error : %s", errorMessage)
		return errors.New("result status is error")
	}
	return nil
}


