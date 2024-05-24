package wenxin

import (
	"bufio"
	"bytes"
	"errors"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/constant"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/dao"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/api/baidu"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
)

func ChatCompletion(params *baidu.WenXinChatCompletionParams, callback baidu.WenXinChatCompletionCallbackFunc) (err error) {
	configData, err := baidu.Instance().GetConfig(constant.ConfigBaiduFeatureWenXin)
	if err != nil {
		glog.Line(true).Println(params, err)
		return err
	}
	if configData == nil {
		return errors.New("获取文心一言配置失败")
	}
	// 调用接口次数增加
	if _, err = dao.ConfigBaidu.Data(g.Map{
		"call_num": gdb.Raw("call_num+1"),
	}).Where("id=?", configData.Id).Update(); err != nil {
		glog.Line(true).Println(params, err)
		return err
	}

	// 请求数据
	requestData := &baidu.WenXinChatCompletionRequestParams{
		Messages:        params.Messages,
		MaxOutputTokens: 1000,
		Stream:          true,
	}
	requestDataJson, err := gjson.Encode(requestData)
	if err != nil {
		glog.Line(true).Println(requestData, err.Error())
		return err
	}

	httpClient := ghttp.NewClient()
	httpClient.Timeout(0)

	resp, err := httpClient.Post(baidu.WenXinERNIE4_0_8K+"?access_token="+configData.AccessToken, requestDataJson)
	if err != nil {
		glog.Line(true).Println(baidu.WenXinERNIE4_0_8K+"?access_token="+configData.AccessToken, requestDataJson, err.Error())
		return err
	}
	defer resp.Body.Close()
	defer resp.Close()
	glog.Line(true).Debug("请求完成")
	reader := bufio.NewReader(resp.Body)
	emptyMessagesCount := 0
	hasErrorPrefix := false
	errorData := ""
	for {
		lineContent, err := reader.ReadBytes('\n')
		if err != nil {
			// 如果遇到错误标识，则不在以流的形式获取数据直接拿到拿到错误内容报错
			if hasErrorPrefix {
				errDecode, errErr := gjson.Decode(errorData)
				if errErr == nil {
					responseErr := &baidu.WenXinChatCompletionResponseError{}
					errErr = gconv.Scan(errDecode, responseErr)
					if errErr == nil && responseErr.Error != nil {
						return errors.New(responseErr.Error.Message)
					} else {
						glog.Line(true).Println(baidu.WenXinERNIE4_0_8K+"?access_token="+configData.AccessToken, requestData, errorData)
					}
				} else {
					glog.Line(true).Println(baidu.WenXinERNIE4_0_8K+"?access_token="+configData.AccessToken, requestData, errorData)
				}
			}
			glog.Line(true).Println(baidu.WenXinERNIE4_0_8K+"?access_token="+configData.AccessToken, requestData, err.Error(), resp.ReadAllString())
			return err
		}
		lineContent = bytes.TrimSpace(lineContent)
		// 有错误标识
		errorPrefix := []byte(`"error":`)
		if bytes.HasPrefix(lineContent, errorPrefix) {
			hasErrorPrefix = true
			errorData += "{"
		}

		if hasErrorPrefix {
			errorData += gconv.String(lineContent)
			continue
		}

		headerData := []byte("data: ")
		if !bytes.HasPrefix(lineContent, headerData) {
			emptyMessagesCount++
			if emptyMessagesCount > 300 {
				return errors.New("连续空消息过多")
			}
			continue
		}
		// 只要不是连续的空消息就把空消息数量归零
		emptyMessagesCount = 0
		lineContent = bytes.TrimPrefix(lineContent, headerData)
		if gconv.String(lineContent) == "[DONE]" {
			break
		}
		lineContentStr := gconv.String(lineContent)
		//glog.Line().Debug(lineContentStr)
		// 解析返回内容
		decodeContent, err := gjson.Decode(lineContentStr)
		if err != nil {
			glog.Line(true).Println(baidu.WenXinERNIE4_0_8K+"?access_token="+configData.AccessToken, requestData, lineContentStr, err.Error())
			return err
		}
		linObj := &baidu.WenXinChatCompletionResponse{}
		err = gconv.Scan(decodeContent, linObj)
		if err != nil {
			glog.Line(true).Println(baidu.WenXinERNIE4_0_8K+"?access_token="+configData.AccessToken, requestData, decodeContent, err.Error())
			return err
		}
		err = callback(gconv.String(lineContent), linObj)
		if err != nil {
			glog.Line(true).Println(baidu.WenXinERNIE4_0_8K+"?access_token="+configData.AccessToken, requestData, gconv.String(lineContent), err.Error())
			return err
		}
	}

	return nil
}
