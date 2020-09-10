package alidns

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeGtmInstanceStatus invokes the alidns.DescribeGtmInstanceStatus API synchronously
// api document: https://help.aliyun.com/api/alidns/describegtminstancestatus.html
func (client *Client) DescribeGtmInstanceStatus(request *DescribeGtmInstanceStatusRequest) (response *DescribeGtmInstanceStatusResponse, err error) {
	response = CreateDescribeGtmInstanceStatusResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeGtmInstanceStatusWithChan invokes the alidns.DescribeGtmInstanceStatus API asynchronously
// api document: https://help.aliyun.com/api/alidns/describegtminstancestatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeGtmInstanceStatusWithChan(request *DescribeGtmInstanceStatusRequest) (<-chan *DescribeGtmInstanceStatusResponse, <-chan error) {
	responseChan := make(chan *DescribeGtmInstanceStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeGtmInstanceStatus(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeGtmInstanceStatusWithCallback invokes the alidns.DescribeGtmInstanceStatus API asynchronously
// api document: https://help.aliyun.com/api/alidns/describegtminstancestatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeGtmInstanceStatusWithCallback(request *DescribeGtmInstanceStatusRequest, callback func(response *DescribeGtmInstanceStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeGtmInstanceStatusResponse
		var err error
		defer close(result)
		response, err = client.DescribeGtmInstanceStatus(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeGtmInstanceStatusRequest is the request struct for api DescribeGtmInstanceStatus
type DescribeGtmInstanceStatusRequest struct {
	*requests.RpcRequest
	InstanceId   string `position:"Query" name:"InstanceId"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

// DescribeGtmInstanceStatusResponse is the response struct for api DescribeGtmInstanceStatus
type DescribeGtmInstanceStatusResponse struct {
	*responses.BaseResponse
	RequestId                   string `json:"RequestId" xml:"RequestId"`
	AddrNotAvailableNum         int    `json:"AddrNotAvailableNum" xml:"AddrNotAvailableNum"`
	AddrPoolNotAvailableNum     int    `json:"AddrPoolNotAvailableNum" xml:"AddrPoolNotAvailableNum"`
	SwitchToFailoverStrategyNum int    `json:"SwitchToFailoverStrategyNum" xml:"SwitchToFailoverStrategyNum"`
	StrategyNotAvailableNum     int    `json:"StrategyNotAvailableNum" xml:"StrategyNotAvailableNum"`
	Status                      string `json:"Status" xml:"Status"`
	StatusReason                string `json:"StatusReason" xml:"StatusReason"`
}

// CreateDescribeGtmInstanceStatusRequest creates a request to invoke DescribeGtmInstanceStatus API
func CreateDescribeGtmInstanceStatusRequest() (request *DescribeGtmInstanceStatusRequest) {
	request = &DescribeGtmInstanceStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "DescribeGtmInstanceStatus", "alidns", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeGtmInstanceStatusResponse creates a response to parse from DescribeGtmInstanceStatus response
func CreateDescribeGtmInstanceStatusResponse() (response *DescribeGtmInstanceStatusResponse) {
	response = &DescribeGtmInstanceStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
