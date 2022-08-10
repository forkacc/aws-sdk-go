// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package panoramaiface provides an interface to enable mocking the AWS Panorama service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package panoramaiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/panorama"
)

// PanoramaAPI provides an interface to enable mocking the
// panorama.Panorama service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// AWS Panorama.
//	func myFunc(svc panoramaiface.PanoramaAPI) bool {
//	    // Make svc.CreateApplicationInstance request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := panorama.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockPanoramaClient struct {
//	    panoramaiface.PanoramaAPI
//	}
//	func (m *mockPanoramaClient) CreateApplicationInstance(input *panorama.CreateApplicationInstanceInput) (*panorama.CreateApplicationInstanceOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockPanoramaClient{}
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
type PanoramaAPI interface {
	CreateApplicationInstance(*panorama.CreateApplicationInstanceInput) (*panorama.CreateApplicationInstanceOutput, error)
	CreateApplicationInstanceWithContext(aws.Context, *panorama.CreateApplicationInstanceInput, ...request.Option) (*panorama.CreateApplicationInstanceOutput, error)
	CreateApplicationInstanceRequest(*panorama.CreateApplicationInstanceInput) (*request.Request, *panorama.CreateApplicationInstanceOutput)

	CreateJobForDevices(*panorama.CreateJobForDevicesInput) (*panorama.CreateJobForDevicesOutput, error)
	CreateJobForDevicesWithContext(aws.Context, *panorama.CreateJobForDevicesInput, ...request.Option) (*panorama.CreateJobForDevicesOutput, error)
	CreateJobForDevicesRequest(*panorama.CreateJobForDevicesInput) (*request.Request, *panorama.CreateJobForDevicesOutput)

	CreateNodeFromTemplateJob(*panorama.CreateNodeFromTemplateJobInput) (*panorama.CreateNodeFromTemplateJobOutput, error)
	CreateNodeFromTemplateJobWithContext(aws.Context, *panorama.CreateNodeFromTemplateJobInput, ...request.Option) (*panorama.CreateNodeFromTemplateJobOutput, error)
	CreateNodeFromTemplateJobRequest(*panorama.CreateNodeFromTemplateJobInput) (*request.Request, *panorama.CreateNodeFromTemplateJobOutput)

	CreatePackage(*panorama.CreatePackageInput) (*panorama.CreatePackageOutput, error)
	CreatePackageWithContext(aws.Context, *panorama.CreatePackageInput, ...request.Option) (*panorama.CreatePackageOutput, error)
	CreatePackageRequest(*panorama.CreatePackageInput) (*request.Request, *panorama.CreatePackageOutput)

	CreatePackageImportJob(*panorama.CreatePackageImportJobInput) (*panorama.CreatePackageImportJobOutput, error)
	CreatePackageImportJobWithContext(aws.Context, *panorama.CreatePackageImportJobInput, ...request.Option) (*panorama.CreatePackageImportJobOutput, error)
	CreatePackageImportJobRequest(*panorama.CreatePackageImportJobInput) (*request.Request, *panorama.CreatePackageImportJobOutput)

	DeleteDevice(*panorama.DeleteDeviceInput) (*panorama.DeleteDeviceOutput, error)
	DeleteDeviceWithContext(aws.Context, *panorama.DeleteDeviceInput, ...request.Option) (*panorama.DeleteDeviceOutput, error)
	DeleteDeviceRequest(*panorama.DeleteDeviceInput) (*request.Request, *panorama.DeleteDeviceOutput)

	DeletePackage(*panorama.DeletePackageInput) (*panorama.DeletePackageOutput, error)
	DeletePackageWithContext(aws.Context, *panorama.DeletePackageInput, ...request.Option) (*panorama.DeletePackageOutput, error)
	DeletePackageRequest(*panorama.DeletePackageInput) (*request.Request, *panorama.DeletePackageOutput)

	DeregisterPackageVersion(*panorama.DeregisterPackageVersionInput) (*panorama.DeregisterPackageVersionOutput, error)
	DeregisterPackageVersionWithContext(aws.Context, *panorama.DeregisterPackageVersionInput, ...request.Option) (*panorama.DeregisterPackageVersionOutput, error)
	DeregisterPackageVersionRequest(*panorama.DeregisterPackageVersionInput) (*request.Request, *panorama.DeregisterPackageVersionOutput)

	DescribeApplicationInstance(*panorama.DescribeApplicationInstanceInput) (*panorama.DescribeApplicationInstanceOutput, error)
	DescribeApplicationInstanceWithContext(aws.Context, *panorama.DescribeApplicationInstanceInput, ...request.Option) (*panorama.DescribeApplicationInstanceOutput, error)
	DescribeApplicationInstanceRequest(*panorama.DescribeApplicationInstanceInput) (*request.Request, *panorama.DescribeApplicationInstanceOutput)

	DescribeApplicationInstanceDetails(*panorama.DescribeApplicationInstanceDetailsInput) (*panorama.DescribeApplicationInstanceDetailsOutput, error)
	DescribeApplicationInstanceDetailsWithContext(aws.Context, *panorama.DescribeApplicationInstanceDetailsInput, ...request.Option) (*panorama.DescribeApplicationInstanceDetailsOutput, error)
	DescribeApplicationInstanceDetailsRequest(*panorama.DescribeApplicationInstanceDetailsInput) (*request.Request, *panorama.DescribeApplicationInstanceDetailsOutput)

	DescribeDevice(*panorama.DescribeDeviceInput) (*panorama.DescribeDeviceOutput, error)
	DescribeDeviceWithContext(aws.Context, *panorama.DescribeDeviceInput, ...request.Option) (*panorama.DescribeDeviceOutput, error)
	DescribeDeviceRequest(*panorama.DescribeDeviceInput) (*request.Request, *panorama.DescribeDeviceOutput)

	DescribeDeviceJob(*panorama.DescribeDeviceJobInput) (*panorama.DescribeDeviceJobOutput, error)
	DescribeDeviceJobWithContext(aws.Context, *panorama.DescribeDeviceJobInput, ...request.Option) (*panorama.DescribeDeviceJobOutput, error)
	DescribeDeviceJobRequest(*panorama.DescribeDeviceJobInput) (*request.Request, *panorama.DescribeDeviceJobOutput)

	DescribeNode(*panorama.DescribeNodeInput) (*panorama.DescribeNodeOutput, error)
	DescribeNodeWithContext(aws.Context, *panorama.DescribeNodeInput, ...request.Option) (*panorama.DescribeNodeOutput, error)
	DescribeNodeRequest(*panorama.DescribeNodeInput) (*request.Request, *panorama.DescribeNodeOutput)

	DescribeNodeFromTemplateJob(*panorama.DescribeNodeFromTemplateJobInput) (*panorama.DescribeNodeFromTemplateJobOutput, error)
	DescribeNodeFromTemplateJobWithContext(aws.Context, *panorama.DescribeNodeFromTemplateJobInput, ...request.Option) (*panorama.DescribeNodeFromTemplateJobOutput, error)
	DescribeNodeFromTemplateJobRequest(*panorama.DescribeNodeFromTemplateJobInput) (*request.Request, *panorama.DescribeNodeFromTemplateJobOutput)

	DescribePackage(*panorama.DescribePackageInput) (*panorama.DescribePackageOutput, error)
	DescribePackageWithContext(aws.Context, *panorama.DescribePackageInput, ...request.Option) (*panorama.DescribePackageOutput, error)
	DescribePackageRequest(*panorama.DescribePackageInput) (*request.Request, *panorama.DescribePackageOutput)

	DescribePackageImportJob(*panorama.DescribePackageImportJobInput) (*panorama.DescribePackageImportJobOutput, error)
	DescribePackageImportJobWithContext(aws.Context, *panorama.DescribePackageImportJobInput, ...request.Option) (*panorama.DescribePackageImportJobOutput, error)
	DescribePackageImportJobRequest(*panorama.DescribePackageImportJobInput) (*request.Request, *panorama.DescribePackageImportJobOutput)

	DescribePackageVersion(*panorama.DescribePackageVersionInput) (*panorama.DescribePackageVersionOutput, error)
	DescribePackageVersionWithContext(aws.Context, *panorama.DescribePackageVersionInput, ...request.Option) (*panorama.DescribePackageVersionOutput, error)
	DescribePackageVersionRequest(*panorama.DescribePackageVersionInput) (*request.Request, *panorama.DescribePackageVersionOutput)

	ListApplicationInstanceDependencies(*panorama.ListApplicationInstanceDependenciesInput) (*panorama.ListApplicationInstanceDependenciesOutput, error)
	ListApplicationInstanceDependenciesWithContext(aws.Context, *panorama.ListApplicationInstanceDependenciesInput, ...request.Option) (*panorama.ListApplicationInstanceDependenciesOutput, error)
	ListApplicationInstanceDependenciesRequest(*panorama.ListApplicationInstanceDependenciesInput) (*request.Request, *panorama.ListApplicationInstanceDependenciesOutput)

	ListApplicationInstanceDependenciesPages(*panorama.ListApplicationInstanceDependenciesInput, func(*panorama.ListApplicationInstanceDependenciesOutput, bool) bool) error
	ListApplicationInstanceDependenciesPagesWithContext(aws.Context, *panorama.ListApplicationInstanceDependenciesInput, func(*panorama.ListApplicationInstanceDependenciesOutput, bool) bool, ...request.Option) error

	ListApplicationInstanceNodeInstances(*panorama.ListApplicationInstanceNodeInstancesInput) (*panorama.ListApplicationInstanceNodeInstancesOutput, error)
	ListApplicationInstanceNodeInstancesWithContext(aws.Context, *panorama.ListApplicationInstanceNodeInstancesInput, ...request.Option) (*panorama.ListApplicationInstanceNodeInstancesOutput, error)
	ListApplicationInstanceNodeInstancesRequest(*panorama.ListApplicationInstanceNodeInstancesInput) (*request.Request, *panorama.ListApplicationInstanceNodeInstancesOutput)

	ListApplicationInstanceNodeInstancesPages(*panorama.ListApplicationInstanceNodeInstancesInput, func(*panorama.ListApplicationInstanceNodeInstancesOutput, bool) bool) error
	ListApplicationInstanceNodeInstancesPagesWithContext(aws.Context, *panorama.ListApplicationInstanceNodeInstancesInput, func(*panorama.ListApplicationInstanceNodeInstancesOutput, bool) bool, ...request.Option) error

	ListApplicationInstances(*panorama.ListApplicationInstancesInput) (*panorama.ListApplicationInstancesOutput, error)
	ListApplicationInstancesWithContext(aws.Context, *panorama.ListApplicationInstancesInput, ...request.Option) (*panorama.ListApplicationInstancesOutput, error)
	ListApplicationInstancesRequest(*panorama.ListApplicationInstancesInput) (*request.Request, *panorama.ListApplicationInstancesOutput)

	ListApplicationInstancesPages(*panorama.ListApplicationInstancesInput, func(*panorama.ListApplicationInstancesOutput, bool) bool) error
	ListApplicationInstancesPagesWithContext(aws.Context, *panorama.ListApplicationInstancesInput, func(*panorama.ListApplicationInstancesOutput, bool) bool, ...request.Option) error

	ListDevices(*panorama.ListDevicesInput) (*panorama.ListDevicesOutput, error)
	ListDevicesWithContext(aws.Context, *panorama.ListDevicesInput, ...request.Option) (*panorama.ListDevicesOutput, error)
	ListDevicesRequest(*panorama.ListDevicesInput) (*request.Request, *panorama.ListDevicesOutput)

	ListDevicesPages(*panorama.ListDevicesInput, func(*panorama.ListDevicesOutput, bool) bool) error
	ListDevicesPagesWithContext(aws.Context, *panorama.ListDevicesInput, func(*panorama.ListDevicesOutput, bool) bool, ...request.Option) error

	ListDevicesJobs(*panorama.ListDevicesJobsInput) (*panorama.ListDevicesJobsOutput, error)
	ListDevicesJobsWithContext(aws.Context, *panorama.ListDevicesJobsInput, ...request.Option) (*panorama.ListDevicesJobsOutput, error)
	ListDevicesJobsRequest(*panorama.ListDevicesJobsInput) (*request.Request, *panorama.ListDevicesJobsOutput)

	ListDevicesJobsPages(*panorama.ListDevicesJobsInput, func(*panorama.ListDevicesJobsOutput, bool) bool) error
	ListDevicesJobsPagesWithContext(aws.Context, *panorama.ListDevicesJobsInput, func(*panorama.ListDevicesJobsOutput, bool) bool, ...request.Option) error

	ListNodeFromTemplateJobs(*panorama.ListNodeFromTemplateJobsInput) (*panorama.ListNodeFromTemplateJobsOutput, error)
	ListNodeFromTemplateJobsWithContext(aws.Context, *panorama.ListNodeFromTemplateJobsInput, ...request.Option) (*panorama.ListNodeFromTemplateJobsOutput, error)
	ListNodeFromTemplateJobsRequest(*panorama.ListNodeFromTemplateJobsInput) (*request.Request, *panorama.ListNodeFromTemplateJobsOutput)

	ListNodeFromTemplateJobsPages(*panorama.ListNodeFromTemplateJobsInput, func(*panorama.ListNodeFromTemplateJobsOutput, bool) bool) error
	ListNodeFromTemplateJobsPagesWithContext(aws.Context, *panorama.ListNodeFromTemplateJobsInput, func(*panorama.ListNodeFromTemplateJobsOutput, bool) bool, ...request.Option) error

	ListNodes(*panorama.ListNodesInput) (*panorama.ListNodesOutput, error)
	ListNodesWithContext(aws.Context, *panorama.ListNodesInput, ...request.Option) (*panorama.ListNodesOutput, error)
	ListNodesRequest(*panorama.ListNodesInput) (*request.Request, *panorama.ListNodesOutput)

	ListNodesPages(*panorama.ListNodesInput, func(*panorama.ListNodesOutput, bool) bool) error
	ListNodesPagesWithContext(aws.Context, *panorama.ListNodesInput, func(*panorama.ListNodesOutput, bool) bool, ...request.Option) error

	ListPackageImportJobs(*panorama.ListPackageImportJobsInput) (*panorama.ListPackageImportJobsOutput, error)
	ListPackageImportJobsWithContext(aws.Context, *panorama.ListPackageImportJobsInput, ...request.Option) (*panorama.ListPackageImportJobsOutput, error)
	ListPackageImportJobsRequest(*panorama.ListPackageImportJobsInput) (*request.Request, *panorama.ListPackageImportJobsOutput)

	ListPackageImportJobsPages(*panorama.ListPackageImportJobsInput, func(*panorama.ListPackageImportJobsOutput, bool) bool) error
	ListPackageImportJobsPagesWithContext(aws.Context, *panorama.ListPackageImportJobsInput, func(*panorama.ListPackageImportJobsOutput, bool) bool, ...request.Option) error

	ListPackages(*panorama.ListPackagesInput) (*panorama.ListPackagesOutput, error)
	ListPackagesWithContext(aws.Context, *panorama.ListPackagesInput, ...request.Option) (*panorama.ListPackagesOutput, error)
	ListPackagesRequest(*panorama.ListPackagesInput) (*request.Request, *panorama.ListPackagesOutput)

	ListPackagesPages(*panorama.ListPackagesInput, func(*panorama.ListPackagesOutput, bool) bool) error
	ListPackagesPagesWithContext(aws.Context, *panorama.ListPackagesInput, func(*panorama.ListPackagesOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*panorama.ListTagsForResourceInput) (*panorama.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *panorama.ListTagsForResourceInput, ...request.Option) (*panorama.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*panorama.ListTagsForResourceInput) (*request.Request, *panorama.ListTagsForResourceOutput)

	ProvisionDevice(*panorama.ProvisionDeviceInput) (*panorama.ProvisionDeviceOutput, error)
	ProvisionDeviceWithContext(aws.Context, *panorama.ProvisionDeviceInput, ...request.Option) (*panorama.ProvisionDeviceOutput, error)
	ProvisionDeviceRequest(*panorama.ProvisionDeviceInput) (*request.Request, *panorama.ProvisionDeviceOutput)

	RegisterPackageVersion(*panorama.RegisterPackageVersionInput) (*panorama.RegisterPackageVersionOutput, error)
	RegisterPackageVersionWithContext(aws.Context, *panorama.RegisterPackageVersionInput, ...request.Option) (*panorama.RegisterPackageVersionOutput, error)
	RegisterPackageVersionRequest(*panorama.RegisterPackageVersionInput) (*request.Request, *panorama.RegisterPackageVersionOutput)

	RemoveApplicationInstance(*panorama.RemoveApplicationInstanceInput) (*panorama.RemoveApplicationInstanceOutput, error)
	RemoveApplicationInstanceWithContext(aws.Context, *panorama.RemoveApplicationInstanceInput, ...request.Option) (*panorama.RemoveApplicationInstanceOutput, error)
	RemoveApplicationInstanceRequest(*panorama.RemoveApplicationInstanceInput) (*request.Request, *panorama.RemoveApplicationInstanceOutput)

	TagResource(*panorama.TagResourceInput) (*panorama.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *panorama.TagResourceInput, ...request.Option) (*panorama.TagResourceOutput, error)
	TagResourceRequest(*panorama.TagResourceInput) (*request.Request, *panorama.TagResourceOutput)

	UntagResource(*panorama.UntagResourceInput) (*panorama.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *panorama.UntagResourceInput, ...request.Option) (*panorama.UntagResourceOutput, error)
	UntagResourceRequest(*panorama.UntagResourceInput) (*request.Request, *panorama.UntagResourceOutput)

	UpdateDeviceMetadata(*panorama.UpdateDeviceMetadataInput) (*panorama.UpdateDeviceMetadataOutput, error)
	UpdateDeviceMetadataWithContext(aws.Context, *panorama.UpdateDeviceMetadataInput, ...request.Option) (*panorama.UpdateDeviceMetadataOutput, error)
	UpdateDeviceMetadataRequest(*panorama.UpdateDeviceMetadataInput) (*request.Request, *panorama.UpdateDeviceMetadataOutput)
}

var _ PanoramaAPI = (*panorama.Panorama)(nil)
