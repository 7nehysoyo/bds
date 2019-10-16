package ripple

import (
	"github.com/jdcloud-bds/bds/common/json"
	"github.com/jdcloud-bds/bds/common/log"
	"github.com/jdcloud-bds/bds/common/math"
	model "github.com/jdcloud-bds/bds/service/model/ripple"
	"time"
)

func parseTransactionCommonFields(commonFields *model.TransactionCommonFields, input string, index, timestamp int64) {
	commonFields.TransactionType = json.Get(input, "TransactionType").String()
	commonFields.ID = 0
	commonFields.Account = json.Get(input, "Account").String()
	commonFields.Hash = json.Get(input, "hash").String()
	commonFields.LedgerIndex = index
	commonFields.Timestamp = timestamp
	commonFields.Fee = json.Get(input, "Fee").Int()
	commonFields.Sequence = json.Get(input, "Sequence").Int()
	commonFields.AccountTxnID = json.Get(input, "AccountTxnID").String()
	commonFields.Flags = json.Get(input, "Flags").Int()
	commonFields.LastLedgerSequence = json.Get(input, "LastLedgerSequence").Int()
	commonFields.Memos = json.Get(input, "Memos").String()
	commonFields.Signers = json.Get(input, "Signers").String()
	commonFields.SourceTag = json.Get(input, "SourceTag").Int()
	commonFields.SigningPubKey = json.Get(input, "SigningPubKey").String()
	commonFields.TxnSignature = json.Get(input, "TxnSignature").String()

	commonFields.TransactionResult = json.Get(input, "metaData.TransactionResult").String()
	commonFields.TransactionIndex = int(json.Get(input, "metaData.TransactionIndex").Int())
	commonFields.AffectedNodes = json.Get(input, "metaData.AffectedNodes").String()
	commonFields.DeliveredAmount = json.Get(input, "metaData.DeliveredAmount").String()
}

func ParseLedger(data string) (*XRPLedgerData, error) {
	startTime := time.Now()
	var err error

	b := new(XRPLedgerData)
	b.Ledger = new(model.Ledger)
	b.AccountSets = make([]*model.AccountSet, 0)
	b.DepositPreauths = make([]*model.DepositPreauth, 0)
	b.EscrowCancels = make([]*model.EscrowCancel, 0)
	b.EscrowCreates = make([]*model.EscrowCreate, 0)
	b.EscrowFinishes = make([]*model.EscrowFinish, 0)
	b.OfferCancels = make([]*model.OfferCancel, 0)
	b.OfferCreates = make([]*model.OfferCreate, 0)
	b.Payments = make([]*model.Payment, 0)
	b.PaymentChannelClaims = make([]*model.PaymentChannelClaim, 0)
	b.PaymentChannelCreates = make([]*model.PaymentChannelCreate, 0)
	b.PaymentChannelFunds = make([]*model.PaymentChannelFund, 0)
	b.SetRegularKeys = make([]*model.SetRegularKey, 0)
	b.SignerListSets = make([]*model.SignerListSet, 0)
	b.TrustSets = make([]*model.TrustSet, 0)

	//Ledger
	b.Ledger.Accepted = int(json.Get(data, "accepted").Int())
	b.Ledger.AccountHash = json.Get(data, "account_hash").String()
	b.Ledger.CloseFlags = int(json.Get(data, "close_flags").Int())
	b.Ledger.CloseTime = json.Get(data, "close_time").Int()
	b.Ledger.Timestamp = json.Get(data, "close_time").Int() + 946656000
	b.Ledger.CloseTimeHuman = json.Get(data, "close_time_human").String()
	b.Ledger.CloseTimeResolution = int(json.Get(data, "close_time_resolution").Int())
	b.Ledger.Closed = int(json.Get(data, "closed").Int())
	b.Ledger.Hash = json.Get(data, "hash").String()
	b.Ledger.LedgerHash = json.Get(data, "ledger_hash").String()
	b.Ledger.LedgerIndex = json.Get(data, "ledger_index").Int()
	b.Ledger.ParentCloseTime = json.Get(data, "parent_close_time").Int()
	b.Ledger.ParentHash = json.Get(data, "parent_hash").String()
	b.Ledger.SeqNum = json.Get(data, "seqNum").Int()
	totalCoins := json.Get(data, "total_coins").String()
	b.Ledger.TotalCoins, err = parseBigInt(totalCoins)
	if err != nil {
		log.Error("splitter ripple: Ledger %d TotalCoins '%s' parse error", b.Ledger.LedgerIndex, totalCoins)
		return nil, err
	}
	b.Ledger.TransactionHash = json.Get(data, "transaction_hash").String()

	transactionList := json.Get(data, "transactions").Array()
	for _, transaction := range transactionList {

		transactionType := json.Get(transaction.String(), "TransactionType").String()

		switch transactionType {
		case AccountSet:
			accountSet := new(model.AccountSet)
			parseTransactionCommonFields(&accountSet.TransactionCommonFields, transaction.String(), b.Ledger.LedgerIndex, b.Ledger.Timestamp)
			accountSet.ClearFlag = json.Get(transaction.String(), "ClearFlag").Int()
			accountSet.Domain = json.Get(transaction.String(), "Domain").String()
			accountSet.EmailHash = json.Get(transaction.String(), "EmailHash").String()
			accountSet.MessageKey = json.Get(transaction.String(), "MessageKey").String()
			accountSet.SetFlag = json.Get(transaction.String(), "SetFlag").Int()
			accountSet.TransferRate = json.Get(transaction.String(), "TransferRate").Int()

			b.AccountSets = append(b.AccountSets, accountSet)
			log.Debug("splitter ripple: %s transaction %s .", transactionType, accountSet.Hash)

		case DepositPreauth:
			depositPreauth := new(model.DepositPreauth)
			parseTransactionCommonFields(&depositPreauth.TransactionCommonFields, transaction.String(), b.Ledger.LedgerIndex, b.Ledger.Timestamp)
			depositPreauth.Authorize = json.Get(transaction.String(), "Authorize").String()
			depositPreauth.UnAuthorize = json.Get(transaction.String(), "Unauthorize").String()

			b.DepositPreauths = append(b.DepositPreauths, depositPreauth)
			log.Debug("splitter ripple: %s transaction %s .", transactionType, depositPreauth.Hash)

		case EscrowCancel:
			escrowCancel := new(model.EscrowCancel)
			parseTransactionCommonFields(&escrowCancel.TransactionCommonFields, transaction.String(), b.Ledger.LedgerIndex, b.Ledger.Timestamp)
			escrowCancel.Owner = json.Get(transaction.String(), "Owner").String()
			escrowCancel.OfferSequence = json.Get(transaction.String(), "OfferSequence").Int()

			b.EscrowCancels = append(b.EscrowCancels, escrowCancel)
			log.Debug("splitter ripple: %s transaction %s .", transactionType, escrowCancel.Hash)

		case EscrowCreate:
			escrowCreate := new(model.EscrowCreate)
			parseTransactionCommonFields(&escrowCreate.TransactionCommonFields, transaction.String(), b.Ledger.LedgerIndex, b.Ledger.Timestamp)
			escrowCreate.Amount = json.Get(transaction.String(), "Amount").String()
			escrowCreate.Destination = json.Get(transaction.String(), "Destination").String()
			escrowCreate.CancelAfter = json.Get(transaction.String(), "CancelAfter").Int()
			escrowCreate.FinishAfter = json.Get(transaction.String(), "FinishAfter").Int()
			escrowCreate.Condition = json.Get(transaction.String(), "Condition").String()
			escrowCreate.DestinationTag = json.Get(transaction.String(), "DestinationTag").Int()

			b.EscrowCreates = append(b.EscrowCreates, escrowCreate)
			log.Debug("splitter ripple: %s transaction %s .", transactionType, escrowCreate.Hash)

		case EscrowFinish:
			escrowFinish := new(model.EscrowFinish)
			parseTransactionCommonFields(&escrowFinish.TransactionCommonFields, transaction.String(), b.Ledger.LedgerIndex, b.Ledger.Timestamp)
			escrowFinish.Owner = json.Get(transaction.String(), "Owner").String()
			escrowFinish.OfferSequence = json.Get(transaction.String(), "OfferSequence").Int()
			escrowFinish.Condition = json.Get(transaction.String(), "Condition").String()
			escrowFinish.Fulfillment = json.Get(transaction.String(), "Fulfillment").String()

			b.EscrowFinishes = append(b.EscrowFinishes, escrowFinish)
			log.Debug("splitter ripple: %s transaction %s .", transactionType, escrowFinish.Hash)

		case OfferCancel:
			offerCancel := new(model.OfferCancel)
			parseTransactionCommonFields(&offerCancel.TransactionCommonFields, transaction.String(), b.Ledger.LedgerIndex, b.Ledger.Timestamp)
			offerCancel.OfferSequence = json.Get(transaction.String(), "OfferSequence").Int()

			b.OfferCancels = append(b.OfferCancels, offerCancel)
			log.Debug("splitter ripple: %s transaction %s .", transactionType, offerCancel.Hash)

		case OfferCreate:
			offerCreate := new(model.OfferCreate)
			parseTransactionCommonFields(&offerCreate.TransactionCommonFields, transaction.String(), b.Ledger.LedgerIndex, b.Ledger.Timestamp)
			offerCreate.Expiration = json.Get(transaction.String(), "Expiration").Int()
			offerCreate.OfferSequence = json.Get(transaction.String(), "OfferSequence").Int()
			offerCreate.TakerGets = json.Get(transaction.String(), "TakerGets").String()
			offerCreate.TakerPays = json.Get(transaction.String(), "TakerPays").String()

			b.OfferCreates = append(b.OfferCreates, offerCreate)
			log.Debug("splitter ripple: %s transaction %s .", transactionType, offerCreate.Hash)

		case Payment:
			payment := new(model.Payment)
			parseTransactionCommonFields(&payment.TransactionCommonFields, transaction.String(), b.Ledger.LedgerIndex, b.Ledger.Timestamp)
			payment.Amount = json.Get(transaction.String(), "Amount").String()
			payment.Destination = json.Get(transaction.String(), "Destination").String()
			payment.DestinationTag = json.Get(transaction.String(), "DestinationTag").Int()
			payment.InvoiceID = json.Get(transaction.String(), "InvoiceID").String()
			payment.Paths = json.Get(transaction.String(), "Paths").String()
			payment.SendMax = json.Get(transaction.String(), "SendMax").String()
			payment.DeliverMin = json.Get(transaction.String(), "DeliverMin").String()

			b.Payments = append(b.Payments, payment)
			log.Debug("splitter ripple: %s transaction %s .", transactionType, payment.Hash)

		case PaymentChannelClaim:
			paymentChannelClaim := new(model.PaymentChannelClaim)
			parseTransactionCommonFields(&paymentChannelClaim.TransactionCommonFields, transaction.String(), b.Ledger.LedgerIndex, b.Ledger.Timestamp)
			paymentChannelClaim.Amount = json.Get(transaction.String(), "Amount").String()
			paymentChannelClaim.Channel = json.Get(transaction.String(), "Channel").String()
			paymentChannelClaim.Balance = json.Get(transaction.String(), "Balance").String()
			paymentChannelClaim.PublicKey = json.Get(transaction.String(), "PublicKey").String()
			paymentChannelClaim.Signature = json.Get(transaction.String(), "Signature").String()

			b.PaymentChannelClaims = append(b.PaymentChannelClaims, paymentChannelClaim)
			log.Debug("splitter ripple: %s transaction %s .", transactionType, paymentChannelClaim.Hash)

		case PaymentChannelCreate:
			paymentChannelCreate := new(model.PaymentChannelCreate)
			parseTransactionCommonFields(&paymentChannelCreate.TransactionCommonFields, transaction.String(), b.Ledger.LedgerIndex, b.Ledger.Timestamp)
			paymentChannelCreate.Amount = json.Get(transaction.String(), "Amount").String()
			paymentChannelCreate.Destination = json.Get(transaction.String(), "Destination").String()
			paymentChannelCreate.DestinationTag = json.Get(transaction.String(), "DestinationTag").Int()
			paymentChannelCreate.SettleDelay = json.Get(transaction.String(), "SettleDelay").Int()
			paymentChannelCreate.PublicKey = json.Get(transaction.String(), "PublicKey").String()
			paymentChannelCreate.CancelAfter = json.Get(transaction.String(), "CancelAfter").Int()

			b.PaymentChannelCreates = append(b.PaymentChannelCreates, paymentChannelCreate)
			log.Debug("splitter ripple: %s transaction %s .", transactionType, paymentChannelCreate.Hash)

		case PaymentChannelFund:
			paymentChannelFund := new(model.PaymentChannelFund)
			parseTransactionCommonFields(&paymentChannelFund.TransactionCommonFields, transaction.String(), b.Ledger.LedgerIndex, b.Ledger.Timestamp)
			paymentChannelFund.Amount = json.Get(transaction.String(), "Amount").String()
			paymentChannelFund.Channel = json.Get(transaction.String(), "Channel").String()
			paymentChannelFund.Expiration = json.Get(transaction.String(), "Expiration").Int()

			b.PaymentChannelFunds = append(b.PaymentChannelFunds, paymentChannelFund)
			log.Debug("splitter ripple: %s transaction %s .", transactionType, paymentChannelFund.Hash)

		case SetRegularKey:
			setRegularKey := new(model.SetRegularKey)
			parseTransactionCommonFields(&setRegularKey.TransactionCommonFields, transaction.String(), b.Ledger.LedgerIndex, b.Ledger.Timestamp)
			setRegularKey.RegularKey = json.Get(transaction.String(), "RegularKey").String()

			b.SetRegularKeys = append(b.SetRegularKeys, setRegularKey)
			log.Debug("splitter ripple: %s transaction %s .", transactionType, setRegularKey.Hash)

		case SignerListSet:
			signerListSet := new(model.SignerListSet)
			parseTransactionCommonFields(&signerListSet.TransactionCommonFields, transaction.String(), b.Ledger.LedgerIndex, b.Ledger.Timestamp)
			signerListSet.SignerQuorum = json.Get(transaction.String(), "SignerQuorum").Int()
			signerListSet.SignerEntries = json.Get(transaction.String(), "SignerEntries").String()

			b.SignerListSets = append(b.SignerListSets, signerListSet)
			log.Debug("splitter ripple: %s transaction %s .", transactionType, signerListSet.Hash)

		case TrustSet:
			trustSet := new(model.TrustSet)
			parseTransactionCommonFields(&trustSet.TransactionCommonFields, transaction.String(), b.Ledger.LedgerIndex, b.Ledger.Timestamp)
			trustSet.LimitAmount = json.Get(transaction.String(), "LimitAmount").String()
			trustSet.QualityIn = json.Get(transaction.String(), "QualityIn").Int()
			trustSet.QualityOut = json.Get(transaction.String(), "QualityOut").Int()

			b.TrustSets = append(b.TrustSets, trustSet)
			log.Debug("splitter ripple: %s transaction %s .", transactionType, trustSet.Hash)
		}
	}

	b.Ledger.TransactionLength = len(transactionList)

	elaspedTime := time.Now().Sub(startTime)
	log.Debug("splitter ripple: parse ledger %d, txs %d, elasped time %s", b.Ledger.LedgerIndex, b.Ledger.TransactionLength, elaspedTime.String())

	return b, nil
}

func parseBigInt(s string) (math.HexOrDecimal256, error) {
	var n math.HexOrDecimal256
	if s == "0x" {
		s = "0x0"
	}

	v, ok := math.ParseBig256(s)
	if !ok {
		n = math.HexOrDecimal256(*defaultBigNumber)
	} else {
		if v.Cmp(maxBigNumber) >= 0 {
			n = math.HexOrDecimal256(*defaultBigNumber)
		} else {
			n = math.HexOrDecimal256(*v)
		}
	}
	return n, nil
}
