package utils_test

import (
	"testing"
	. "../utils"
)

var isMatchTest = []struct {
	data string
	reg string
} {
	{"123hello321", "hello\\d+"},
	{"hello world", "\\shero"},
}

func TestIsMatch(t *testing.T) {
	for _, v := range isMatchTest {
		t.Logf("%b\n", IsMatch(v.data, v.reg))
	}
}

var content = `{"SERVICE_RESPONSE_RESULT":{"ORG":"242","accountBlockCode1":"","accountBlockCode2":"","accountIndex":-2019816923,"accountMemo2":"","accountNo":"2998009863535879","accountNoIndex":0,"accountSign":"1","accountType":"","accountUSBlockCode1":"","accountUSBlockCode2":"","accountsOfBill":[],"acctDesc":"","acctLogoDesc":"","acountType":0,"adjustIntegral":0,"affinityUnitCode":"","available":false,"availableLimit":"","billDate":"","billTime":"","cardBlockCode":"","cardNo":"","cardSet":[],"cashLimit":"","chgIntegralTotal":0,"cpAdjustIntegral":0,"cpNewIntegralTotal":0,"currentPoint":0,"date1":"","date2":"","dateOpen":"","ddAcctNbr":"","ddAcctNbrF":"","ddBankId":"","ddBankIdF":"","ddPmt":"","ddPmtF":"","ddStatus":"","ddStatusF":"","defaultBillDate":"","dollarAmountToPayOffTotal":0,"dollarAmtTotal":0,"dollarMinAmountToPayOff2Total":0,"dollarMinAmountToPayOffTotal":0,"dollarMinRemianAmountPayOffTotal":0,"dollarRemianAmountPayOffTotal":0,"emailAddr":"","entitySubscribeFlag":false,"firstSetDate":"","flag":"","foreignORG":"","haveTwoOldAccount":"","initialLimit":0,"isDoubleCurr":false,"isShowTOACard":"","lastIntegral":0,"limit":"","localORG":"","logo":"","maskAccountNo":"2998********5879","maskMasterCardNo":"","maskOldaccountNo":"","monthlistYYYYMM":[],"monthsOfBill":[],"newIntegralTotal":0,"nextDate":"","nextaccountNo":"","oldaccountNo":"","oldaccountNoIndex":0,"partyNo":"","payOffDate":"","payoffDetail":[{"accrualAmount":0,"adjustAmount":0,"amountToPayOff":0,"amountToPayOff2":0,"availableLimit":"","billAmount":0,"cashLimit":"","creditLimit":"","creditLmt":"","currencyType":"242","currendPage":"1","limit":"","minAmountToPayOff":0,"minAmountToPayOff2":0,"newPayBackMoney":"0.00","pageSize":"4","payBackMoney":"0.0","payOffDate":"","payRecords":[{"accountNo":"","cardNo":"8995","consumeAmount":"100.00","consumeArea":"银联收单消费","consumeCurType":"RMB","currencyType":"242","ntxnDescTxt":"银联收单消费","settleDate":"0403","stxnDescTxt":"银联收单消费","txnAmount":"100.00","txnDate":"20160402","txnDateView":"","txnDescTxt":"银联收单消费","txnMonth":""},{"accountNo":"","cardNo":"8995","consumeAmount":"359.00","consumeArea":"银联收单消费","consumeCurType":"RMB","currencyType":"242","ntxnDescTxt":"银联收单消费","settleDate":"0403","stxnDescTxt":"银联收单消费","txnAmount":"359.00","txnDate":"20160402","txnDateView":"","txnDescTxt":"银联收单消费","txnMonth":""},{"accountNo":"","cardNo":"8995","consumeAmount":"49.90","consumeArea":"支付宝（快捷支付）","consumeCurType":"RMB","currencyType":"242","ntxnDescTxt":"支付宝（快捷支付）","settleDate":"0403","stxnDescTxt":"支付宝（快捷支付）","txnAmount":"49.90","txnDate":"20160402","txnDateView":"","txnDescTxt":"支付宝（快捷支付）","txnMonth":""},{"accountNo":"","cardNo":"","consumeAmount":"20.00","consumeArea":"滞纳金","consumeCurType":"RMB","currencyType":"242","ntxnDescTxt":"滞纳金","settleDate":"0405","stxnDescTxt":"滞纳金","txnAmount":"20.00","txnDate":"20160405","txnDateView":"","txnDescTxt":"滞纳金","txnMonth":""}],"preAmountPaidOff":0,"preBillAmount":0,"remaindAmount":0,"remaindAmountValue":0,"summaryRecords":[],"totalConsumeAmt":"528.9","totalConsumeAmt2":"0","totalPage":"","totalRecNum":"4"}],"platinumFlg":"","postBillFlag":false,"preDate":"","prePayOffDate":"","preaccountNo":"","rmbAmountToPayOffTotal":0,"rmbAmtTotal":0,"rmbMinAmountToPayOff2Total":0,"rmbMinAmountToPayOffTotal":0,"rmbMinRemianAmountPayOffTotal":0,"rmbRemianAmountPayOffTotal":0,"settleDate":"","status":"","subscribeFlag":false,"thisBillIsNull":"","totalIntegral":0,"userAmt1":"","userAmt2":"","userCode1":"","yearIntegral":0,"ztype":false},"ret_code":"000","payoffDetailDto":{"accountNo":"2998009863535879","accountNoIndex":-2019816923,"acountType":0,"billType":0,"currentPage":1,"currentPayRecords":[{"accountNo":"","cardNo":"8995","consumeAmount":"100.00","consumeArea":"银联收单消费","consumeCurType":"RMB","currencyType":"242","ntxnDescTxt":"银联收单消费","settleDate":"0403","stxnDescTxt":"银联收单消费","txnAmount":"100.00","txnDate":"20160402","txnDateView":"","txnDescTxt":"银联收单消费","txnMonth":""},{"accountNo":"","cardNo":"8995","consumeAmount":"359.00","consumeArea":"银联收单消费","consumeCurType":"RMB","currencyType":"242","ntxnDescTxt":"银联收单消费","settleDate":"0403","stxnDescTxt":"银联收单消费","txnAmount":"359.00","txnDate":"20160402","txnDateView":"","txnDescTxt":"银联收单消费","txnMonth":""},{"accountNo":"","cardNo":"8995","consumeAmount":"49.90","consumeArea":"支付宝（快捷支付）","consumeCurType":"RMB","currencyType":"242","ntxnDescTxt":"支付宝（快捷支付）","settleDate":"0403","stxnDescTxt":"支付宝（快捷支付）","txnAmount":"49.90","txnDate":"20160402","txnDateView":"","txnDescTxt":"支付宝（快捷支付）","txnMonth":""},{"accountNo":"","cardNo":"","consumeAmount":"20.00","consumeArea":"滞纳金","consumeCurType":"RMB","currencyType":"242","ntxnDescTxt":"滞纳金","settleDate":"0405","stxnDescTxt":"滞纳金","txnAmount":"20.00","txnDate":"20160405","txnDateView":"","txnDescTxt":"滞纳金","txnMonth":""}],"maskCardNo":"2998********5879","nextTenPageIndex":10,"oldaccountNo":"","pageCount":1,"pagesShowAccount":10,"pagesize":100,"payRecords":[{"accountNo":"","cardNo":"8995","consumeAmount":"100.00","consumeArea":"银联收单消费","consumeCurType":"RMB","currencyType":"242","ntxnDescTxt":"银联收单消费","settleDate":"0403","stxnDescTxt":"银联收单消费","txnAmount":"100.00","txnDate":"20160402","txnDateView":"","txnDescTxt":"银联收单消费","txnMonth":""},{"accountNo":"","cardNo":"8995","consumeAmount":"359.00","consumeArea":"银联收单消费","consumeCurType":"RMB","currencyType":"242","ntxnDescTxt":"银联收单消费","settleDate":"0403","stxnDescTxt":"银联收单消费","txnAmount":"359.00","txnDate":"20160402","txnDateView":"","txnDescTxt":"银联收单消费","txnMonth":""},{"accountNo":"","cardNo":"8995","consumeAmount":"49.90","consumeArea":"支付宝（快捷支付）","consumeCurType":"RMB","currencyType":"242","ntxnDescTxt":"支付宝（快捷支付）","settleDate":"0403","stxnDescTxt":"支付宝（快捷支付）","txnAmount":"49.90","txnDate":"20160402","txnDateView":"","txnDescTxt":"支付宝（快捷支付）","txnMonth":""},{"accountNo":"","cardNo":"","consumeAmount":"20.00","consumeArea":"滞纳金","consumeCurType":"RMB","currencyType":"242","ntxnDescTxt":"滞纳金","settleDate":"0405","stxnDescTxt":"滞纳金","txnAmount":"20.00","txnDate":"20160405","txnDateView":"","txnDescTxt":"滞纳金","txnMonth":""}],"preTenPageIndex":1,"queryMonth":"","recordAccount":4},"paybackmoney":"0.0"}`

var matchReg = []struct {
	Reg string
	Type string
} {
	{`accountMemo2.*?accountNo\":\"(\d+)\"`, "M"},
	{`accountNo\":\"(\d+)\"`, "MS"},
	{`consumeAmount":"(.*?)".*?consumeArea":"(.*?)"`, "MM"},
}

func TestMatchData(t *testing.T) {
	t.Logf("%s\n", MatchData(content, matchReg[0].Reg))
}

func TestMatchMutilData(t *testing.T) {
	t.Logf("%s\n", MatchMutilData(content, matchReg[2].Reg))
}

func TestMatchSingleLine(t *testing.T) {
	t.Log(MatchSingleLine(content, matchReg[1].Reg))
}

func TestMatchMutilLine(t *testing.T) {
	t.Log(MatchMutilLine(content, matchReg[2].Reg))
}