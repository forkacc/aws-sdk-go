// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package cloudwatchrumiface provides an interface to enable mocking the CloudWatch RUM service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package cloudwatchrumiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/cloudwatchrum"
)

// CloudWatchRUMAPI provides an interface to enable mocking the
// cloudwatchrum.CloudWatchRUM service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// CloudWatch RUM.
//	func myFunc(svc cloudwatchrumiface.CloudWatchRUMAPI) bool {
//	    // Make svc.CreateAppMonitor request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := cloudwatchrum.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockCloudWatchRUMClient struct {
//	    cloudwatchrumiface.CloudWatchRUMAPI
//	}
//	func (m *mockCloudWatchRUMClient) CreateAppMonitor(input *cloudwatchrum.CreateAppMonitorInput) (*cloudwatchrum.CreateAppMonitorOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockCloudWatchRUMClient{}
//
//	    myfunc(mockSvc)
//
//	    // Verify myFunc's functionality
//	}
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type CloudWatchRUMAPI interface {
	CreateAppMonitor(*cloudwatchrum.CreateAppMonitorInput) (*cloudwatchrum.CreateAppMonitorOutput, error)
	CreateAppMonitorWithContext(aws.Context, *cloudwatchrum.CreateAppMonitorInput, ...request.Option) (*cloudwatchrum.CreateAppMonitorOutput, error)
	CreateAppMonitorRequest(*cloudwatchrum.CreateAppMonitorInput) (*request.Request, *cloudwatchrum.CreateAppMonitorOutput)

	DeleteAppMonitor(*cloudwatchrum.DeleteAppMonitorInput) (*cloudwatchrum.DeleteAppMonitorOutput, error)
	DeleteAppMonitorWithContext(aws.Context, *cloudwatchrum.DeleteAppMonitorInput, ...request.Option) (*cloudwatchrum.DeleteAppMonitorOutput, error)
	DeleteAppMonitorRequest(*cloudwatchrum.DeleteAppMonitorInput) (*request.Request, *cloudwatchrum.DeleteAppMonitorOutput)

	GetAppMonitor(*cloudwatchrum.GetAppMonitorInput) (*cloudwatchrum.GetAppMonitorOutput, error)
	GetAppMonitorWithContext(aws.Context, *cloudwatchrum.GetAppMonitorInput, ...request.Option) (*cloudwatchrum.GetAppMonitorOutput, error)
	GetAppMonitorRequest(*cloudwatchrum.GetAppMonitorInput) (*request.Request, *cloudwatchrum.GetAppMonitorOutput)

	GetAppMonitorData(*cloudwatchrum.GetAppMonitorDataInput) (*cloudwatchrum.GetAppMonitorDataOutput, error)
	GetAppMonitorDataWithContext(aws.Context, *cloudwatchrum.GetAppMonitorDataInput, ...request.Option) (*cloudwatchrum.GetAppMonitorDataOutput, error)
	GetAppMonitorDataRequest(*cloudwatchrum.GetAppMonitorDataInput) (*request.Request, *cloudwatchrum.GetAppMonitorDataOutput)

	GetAppMonitorDataPages(*cloudwatchrum.GetAppMonitorDataInput, func(*cloudwatchrum.GetAppMonitorDataOutput, bool) bool) error
	GetAppMonitorDataPagesWithContext(aws.Context, *cloudwatchrum.GetAppMonitorDataInput, func(*cloudwatchrum.GetAppMonitorDataOutput, bool) bool, ...request.Option) error

	ListAppMonitors(*cloudwatchrum.ListAppMonitorsInput) (*cloudwatchrum.ListAppMonitorsOutput, error)
	ListAppMonitorsWithContext(aws.Context, *cloudwatchrum.ListAppMonitorsInput, ...request.Option) (*cloudwatchrum.ListAppMonitorsOutput, error)
	ListAppMonitorsRequest(*cloudwatchrum.ListAppMonitorsInput) (*request.Request, *cloudwatchrum.ListAppMonitorsOutput)

	ListAppMonitorsPages(*cloudwatchrum.ListAppMonitorsInput, func(*cloudwatchrum.ListAppMonitorsOutput, bool) bool) error
	ListAppMonitorsPagesWithContext(aws.Context, *cloudwatchrum.ListAppMonitorsInput, func(*cloudwatchrum.ListAppMonitorsOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*cloudwatchrum.ListTagsForResourceInput) (*cloudwatchrum.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *cloudwatchrum.ListTagsForResourceInput, ...request.Option) (*cloudwatchrum.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*cloudwatchrum.ListTagsForResourceInput) (*request.Request, *cloudwatchrum.ListTagsForResourceOutput)

	PutRumEvents(*cloudwatchrum.PutRumEventsInput) (*cloudwatchrum.PutRumEventsOutput, error)
	PutRumEventsWithContext(aws.Context, *cloudwatchrum.PutRumEventsInput, ...request.Option) (*cloudwatchrum.PutRumEventsOutput, error)
	PutRumEventsRequest(*cloudwatchrum.PutRumEventsInput) (*request.Request, *cloudwatchrum.PutRumEventsOutput)

	TagResource(*cloudwatchrum.TagResourceInput) (*cloudwatchrum.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *cloudwatchrum.TagResourceInput, ...request.Option) (*cloudwatchrum.TagResourceOutput, error)
	TagResourceRequest(*cloudwatchrum.TagResourceInput) (*request.Request, *cloudwatchrum.TagResourceOutput)

	UntagResource(*cloudwatchrum.UntagResourceInput) (*cloudwatchrum.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *cloudwatchrum.UntagResourceInput, ...request.Option) (*cloudwatchrum.UntagResourceOutput, error)
	UntagResourceRequest(*cloudwatchrum.UntagResourceInput) (*request.Request, *cloudwatchrum.UntagResourceOutput)

	UpdateAppMonitor(*cloudwatchrum.UpdateAppMonitorInput) (*cloudwatchrum.UpdateAppMonitorOutput, error)
	UpdateAppMonitorWithContext(aws.Context, *cloudwatchrum.UpdateAppMonitorInput, ...request.Option) (*cloudwatchrum.UpdateAppMonitorOutput, error)
	UpdateAppMonitorRequest(*cloudwatchrum.UpdateAppMonitorInput) (*request.Request, *cloudwatchrum.UpdateAppMonitorOutput)
}

var _ CloudWatchRUMAPI = (*cloudwatchrum.CloudWatchRUM)(nil)
