package cms

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

// TaskConfigModify invokes the cms.TaskConfigModify API synchronously
// api document: https://help.aliyun.com/api/cms/taskconfigmodify.html
func (client *Client) TaskConfigModify(request *TaskConfigModifyRequest) (response *TaskConfigModifyResponse, err error) {
	response = CreateTaskConfigModifyResponse()
	err = client.DoAction(request, response)
	return
}

// TaskConfigModifyWithChan invokes the cms.TaskConfigModify API asynchronously
// api document: https://help.aliyun.com/api/cms/taskconfigmodify.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) TaskConfigModifyWithChan(request *TaskConfigModifyRequest) (<-chan *TaskConfigModifyResponse, <-chan error) {
	responseChan := make(chan *TaskConfigModifyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.TaskConfigModify(request)
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

// TaskConfigModifyWithCallback invokes the cms.TaskConfigModify API asynchronously
// api document: https://help.aliyun.com/api/cms/taskconfigmodify.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) TaskConfigModifyWithCallback(request *TaskConfigModifyRequest, callback func(response *TaskConfigModifyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *TaskConfigModifyResponse
		var err error
		defer close(result)
		response, err = client.TaskConfigModify(request)
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

// TaskConfigModifyRequest is the request struct for api TaskConfigModify
type TaskConfigModifyRequest struct {
	*requests.RpcRequest
	GroupId      requests.Integer `position:"Query" name:"GroupId"`
	TaskType     string           `position:"Query" name:"TaskType"`
	Id           requests.Integer `position:"Query" name:"Id"`
	TaskName     string           `position:"Query" name:"TaskName"`
	TaskScope    string           `position:"Query" name:"TaskScope"`
	GroupName    string           `position:"Query" name:"GroupName"`
	JsonData     string           `position:"Query" name:"JsonData"`
	InstanceList *[]string        `position:"Query" name:"InstanceList"  type:"Repeated"`
	AlertConfig  string           `position:"Query" name:"AlertConfig"`
}

// TaskConfigModifyResponse is the response struct for api TaskConfigModify
type TaskConfigModifyResponse struct {
	*responses.BaseResponse
	ErrorCode    int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool   `json:"Success" xml:"Success"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
}

// CreateTaskConfigModifyRequest creates a request to invoke TaskConfigModify API
func CreateTaskConfigModifyRequest() (request *TaskConfigModifyRequest) {
	request = &TaskConfigModifyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2018-03-08", "TaskConfigModify", "cms", "openAPI")
	return
}

// CreateTaskConfigModifyResponse creates a response to parse from TaskConfigModify response
func CreateTaskConfigModifyResponse() (response *TaskConfigModifyResponse) {
	response = &TaskConfigModifyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
