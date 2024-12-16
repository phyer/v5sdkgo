package v5sdkgo

import (
	"context"
	"fmt"
	. "github.com/phyer/v5sdkgo/rest"
	. "github.com/phyer/v5sdkgo/ws"

	"log"
	"time"
)

/*
rest API请求
更多示例请查看 rest/rest_test.go
*/
func REST() {
	// 设置您的APIKey
	apikey := APIKeyInfo{
		ApiKey:     "eca767d4-fda5-4a1b-bb28-49ae18093307",
		SecKey:     "8CA3628A556ED137977DB298D37BC7F3",
		PassPhrase: "Op3Druaron",
	}

	// 第三个参数代表是否为模拟环境，更多信息查看接口说明
	cli := NewRESTClient("https://www.okex.win", &apikey, false)
	rsp, err := cli.Get(context.Background(), "/api/v5/account/balance", nil)
	if err != nil {
		return
	}

	fmt.Println("Response:")
	fmt.Println("\thttp code: ", rsp.Code)
	fmt.Println("\t总耗时: ", rsp.TotalUsedTime)
	fmt.Println("\t请求耗时: ", rsp.ReqUsedTime)
	fmt.Println("\t返回消息: ", rsp.Body)
	fmt.Println("\terrCode: ", rsp.V5Response.Code)
	fmt.Println("\terrMsg: ", rsp.V5Response.Msg)
	fmt.Println("\tdata: ", rsp.V5Response.Data)

}

// 订阅私有频道
func wsPriv() {
	ep := "wss://ws.okex.com:8443/ws/v5/private?brokerId=9999"

	// 填写您自己的APIKey信息
	apikey := "xxxx"
	secretKey := "xxxxx"
	passphrase := "xxxxx"

	// 创建ws客户端
	r, err := NewWsClient(ep)
	if err != nil {
		log.Println(err)
		return
	}

	// 设置连接超时
	r.SetDailTimeout(time.Second * 2)
	err = r.Start()
	if err != nil {
		log.Println(err)
		return
	}
	defer r.Stop()
	var res bool

	res, _, err = r.Login(apikey, secretKey, passphrase)
	if res {
		fmt.Println("登录成功！")
	} else {
		fmt.Println("登录失败！", err)
		return
	}

	// 订阅账户频道
	var args []map[string]string
	arg := make(map[string]string)
	arg["ccy"] = "BTC"
	args = append(args, arg)

	start := time.Now()
	res, _, err = r.PrivAccout(OP_SUBSCRIBE, args)
	if res {
		usedTime := time.Since(start)
		fmt.Println("订阅成功！耗时:", usedTime.String())
	} else {
		fmt.Println("订阅失败！", err)
	}

	time.Sleep(100 * time.Second)
	start = time.Now()
	res, _, err = r.PrivAccout(OP_UNSUBSCRIBE, args)
	if res {
		usedTime := time.Since(start)
		fmt.Println("取消订阅成功！", usedTime.String())
	} else {
		fmt.Println("取消订阅失败！", err)
	}

}

// 订阅公共频道
func wsPub() {
	ep := "wss://wsaws.okex.com:8443/ws/v5/public?brokerId=9999"

	// 创建ws客户端
	r, err := NewWsClient(ep)
	if err != nil {
		log.Println(err)
		return
	}

	// 设置连接超时
	r.SetDailTimeout(time.Second * 2)
	err = r.Start()
	if err != nil {
		log.Println(err)
		return
	}
	defer r.Stop()
	// 订阅产品频道
	// 在这里初始化instrument列表
	var args []map[string]string
	arg := make(map[string]string)
	arg["instType"] = FUTURES
	//arg["instType"] = OPTION
	args = append(args, arg)

	start := time.Now()
	//订阅
	res, _, err := r.PubInstruemnts(OP_SUBSCRIBE, args)
	fmt.Println("args:", args)
	if res {
		usedTime := time.Since(start)
		fmt.Println("订阅成功！", usedTime.String())
	} else {
		fmt.Println("订阅失败！", err)
	}

	// 在这里 args1 初始化tickerList的列表
	var args1 []map[string]string
	arg1 := make(map[string]string)
	arg1["instId"] = "ETH-USDT"
	//arg["instType"] = OPTION
	args1 = append(args1, arg1)
	//------------------------------------------------------
	start1 := time.Now()
	res, _, err = r.PubTickers(OP_SUBSCRIBE, args1)
	fmt.Println("args:", args)
	if res {
		usedTime := time.Since(start1)
		fmt.Println("订阅成功！", usedTime.String())
	} else {
		fmt.Println("订阅失败！", err)
	}
	time.Sleep(300 * time.Second)
	//
	// start = time.Now()
	// res, _, err = r.PubInstruemnts(OP_UNSUBSCRIBE, args)
	// if res {
	// usedTime := time.Since(start)
	// fmt.Println("取消订阅成功！", usedTime.String())
	// } else {
	// fmt.Println("取消订阅失败！", err)
	// }
}

// websocket交易
func wsJrpc() {
	ep := "wss://ws.okex.com:8443/ws/v5/private?brokerId=9999"

	// 填写您自己的APIKey信息
	apikey := "xxxx"
	secretKey := "xxxxx"
	passphrase := "xxxxx"

	var res bool
	var req_id string

	// 创建ws客户端
	r, err := NewWsClient(ep)
	if err != nil {
		log.Println(err)
		return
	}

	// 设置连接超时
	r.SetDailTimeout(time.Second * 2)
	err = r.Start()
	if err != nil {
		log.Println(err)
		return
	}

	defer r.Stop()

	res, _, err = r.Login(apikey, secretKey, passphrase)
	if res {
		fmt.Println("登录成功！")
	} else {
		fmt.Println("登录失败！", err)
		return
	}

	start := time.Now()
	param := map[string]interface{}{}
	param["instId"] = "BTC-USDT"
	param["tdMode"] = "cash"
	param["side"] = "buy"
	param["ordType"] = "market"
	param["sz"] = "200"
	req_id = "00001"

	res, _, err = r.PlaceOrder(req_id, param)
	if res {
		usedTime := time.Since(start)
		fmt.Println("下单成功！", usedTime.String())
	} else {
		usedTime := time.Since(start)
		fmt.Println("下单失败！", usedTime.String(), err)
	}
}

func main() {
	// 公共订阅
	wsPub()

	// 私有订阅
	// wsPriv()

	// websocket交易
	// wsJrpc()

	// rest请求
	// REST()
}
