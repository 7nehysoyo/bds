package ripple

import (
	"github.com/jdcloud-bds/bds/common/httputils"
	"github.com/jdcloud-bds/bds/common/jsonrpc"
	"testing"
)

func TestSendBatchLedger(t *testing.T) {
	httpClient := httputils.NewRestClientWithAuthentication(nil)
	remoteHandler, err := newRPCHandler(jsonrpc.New(httpClient, "http://116.196.114.8:5555"))
	if err != nil {
		t.Fatal(err)
	}
	err = remoteHandler.SendBatchLedger(50588462, 50588462)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("over")
}

func TestGetCompleteLedgers(t *testing.T) {
	httpClient := httputils.NewRestClientWithAuthentication(nil)
	remoteHandler, err := newRPCHandler(jsonrpc.New(httpClient, "http://116.196.114.8:5555"))
	if err != nil {
		t.Fatal(err)
	}
	result, err := remoteHandler.GetCompleteLedgers()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("complete ledgers result : %v", result)
}
