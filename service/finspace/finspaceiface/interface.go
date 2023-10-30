// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package finspaceiface provides an interface to enable mocking the FinSpace User Environment Management service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package finspaceiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/finspace"
)

// FinspaceAPI provides an interface to enable mocking the
// finspace.Finspace service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// FinSpace User Environment Management service.
//	func myFunc(svc finspaceiface.FinspaceAPI) bool {
//	    // Make svc.CreateEnvironment request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := finspace.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockFinspaceClient struct {
//	    finspaceiface.FinspaceAPI
//	}
//	func (m *mockFinspaceClient) CreateEnvironment(input *finspace.CreateEnvironmentInput) (*finspace.CreateEnvironmentOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockFinspaceClient{}
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
type FinspaceAPI interface {
	CreateEnvironment(*finspace.CreateEnvironmentInput) (*finspace.CreateEnvironmentOutput, error)
	CreateEnvironmentWithContext(aws.Context, *finspace.CreateEnvironmentInput, ...request.Option) (*finspace.CreateEnvironmentOutput, error)
	CreateEnvironmentRequest(*finspace.CreateEnvironmentInput) (*request.Request, *finspace.CreateEnvironmentOutput)

	CreateKxChangeset(*finspace.CreateKxChangesetInput) (*finspace.CreateKxChangesetOutput, error)
	CreateKxChangesetWithContext(aws.Context, *finspace.CreateKxChangesetInput, ...request.Option) (*finspace.CreateKxChangesetOutput, error)
	CreateKxChangesetRequest(*finspace.CreateKxChangesetInput) (*request.Request, *finspace.CreateKxChangesetOutput)

	CreateKxCluster(*finspace.CreateKxClusterInput) (*finspace.CreateKxClusterOutput, error)
	CreateKxClusterWithContext(aws.Context, *finspace.CreateKxClusterInput, ...request.Option) (*finspace.CreateKxClusterOutput, error)
	CreateKxClusterRequest(*finspace.CreateKxClusterInput) (*request.Request, *finspace.CreateKxClusterOutput)

	CreateKxDatabase(*finspace.CreateKxDatabaseInput) (*finspace.CreateKxDatabaseOutput, error)
	CreateKxDatabaseWithContext(aws.Context, *finspace.CreateKxDatabaseInput, ...request.Option) (*finspace.CreateKxDatabaseOutput, error)
	CreateKxDatabaseRequest(*finspace.CreateKxDatabaseInput) (*request.Request, *finspace.CreateKxDatabaseOutput)

	CreateKxEnvironment(*finspace.CreateKxEnvironmentInput) (*finspace.CreateKxEnvironmentOutput, error)
	CreateKxEnvironmentWithContext(aws.Context, *finspace.CreateKxEnvironmentInput, ...request.Option) (*finspace.CreateKxEnvironmentOutput, error)
	CreateKxEnvironmentRequest(*finspace.CreateKxEnvironmentInput) (*request.Request, *finspace.CreateKxEnvironmentOutput)

	CreateKxUser(*finspace.CreateKxUserInput) (*finspace.CreateKxUserOutput, error)
	CreateKxUserWithContext(aws.Context, *finspace.CreateKxUserInput, ...request.Option) (*finspace.CreateKxUserOutput, error)
	CreateKxUserRequest(*finspace.CreateKxUserInput) (*request.Request, *finspace.CreateKxUserOutput)

	DeleteEnvironment(*finspace.DeleteEnvironmentInput) (*finspace.DeleteEnvironmentOutput, error)
	DeleteEnvironmentWithContext(aws.Context, *finspace.DeleteEnvironmentInput, ...request.Option) (*finspace.DeleteEnvironmentOutput, error)
	DeleteEnvironmentRequest(*finspace.DeleteEnvironmentInput) (*request.Request, *finspace.DeleteEnvironmentOutput)

	DeleteKxCluster(*finspace.DeleteKxClusterInput) (*finspace.DeleteKxClusterOutput, error)
	DeleteKxClusterWithContext(aws.Context, *finspace.DeleteKxClusterInput, ...request.Option) (*finspace.DeleteKxClusterOutput, error)
	DeleteKxClusterRequest(*finspace.DeleteKxClusterInput) (*request.Request, *finspace.DeleteKxClusterOutput)

	DeleteKxDatabase(*finspace.DeleteKxDatabaseInput) (*finspace.DeleteKxDatabaseOutput, error)
	DeleteKxDatabaseWithContext(aws.Context, *finspace.DeleteKxDatabaseInput, ...request.Option) (*finspace.DeleteKxDatabaseOutput, error)
	DeleteKxDatabaseRequest(*finspace.DeleteKxDatabaseInput) (*request.Request, *finspace.DeleteKxDatabaseOutput)

	DeleteKxEnvironment(*finspace.DeleteKxEnvironmentInput) (*finspace.DeleteKxEnvironmentOutput, error)
	DeleteKxEnvironmentWithContext(aws.Context, *finspace.DeleteKxEnvironmentInput, ...request.Option) (*finspace.DeleteKxEnvironmentOutput, error)
	DeleteKxEnvironmentRequest(*finspace.DeleteKxEnvironmentInput) (*request.Request, *finspace.DeleteKxEnvironmentOutput)

	DeleteKxUser(*finspace.DeleteKxUserInput) (*finspace.DeleteKxUserOutput, error)
	DeleteKxUserWithContext(aws.Context, *finspace.DeleteKxUserInput, ...request.Option) (*finspace.DeleteKxUserOutput, error)
	DeleteKxUserRequest(*finspace.DeleteKxUserInput) (*request.Request, *finspace.DeleteKxUserOutput)

	GetEnvironment(*finspace.GetEnvironmentInput) (*finspace.GetEnvironmentOutput, error)
	GetEnvironmentWithContext(aws.Context, *finspace.GetEnvironmentInput, ...request.Option) (*finspace.GetEnvironmentOutput, error)
	GetEnvironmentRequest(*finspace.GetEnvironmentInput) (*request.Request, *finspace.GetEnvironmentOutput)

	GetKxChangeset(*finspace.GetKxChangesetInput) (*finspace.GetKxChangesetOutput, error)
	GetKxChangesetWithContext(aws.Context, *finspace.GetKxChangesetInput, ...request.Option) (*finspace.GetKxChangesetOutput, error)
	GetKxChangesetRequest(*finspace.GetKxChangesetInput) (*request.Request, *finspace.GetKxChangesetOutput)

	GetKxCluster(*finspace.GetKxClusterInput) (*finspace.GetKxClusterOutput, error)
	GetKxClusterWithContext(aws.Context, *finspace.GetKxClusterInput, ...request.Option) (*finspace.GetKxClusterOutput, error)
	GetKxClusterRequest(*finspace.GetKxClusterInput) (*request.Request, *finspace.GetKxClusterOutput)

	GetKxConnectionString(*finspace.GetKxConnectionStringInput) (*finspace.GetKxConnectionStringOutput, error)
	GetKxConnectionStringWithContext(aws.Context, *finspace.GetKxConnectionStringInput, ...request.Option) (*finspace.GetKxConnectionStringOutput, error)
	GetKxConnectionStringRequest(*finspace.GetKxConnectionStringInput) (*request.Request, *finspace.GetKxConnectionStringOutput)

	GetKxDatabase(*finspace.GetKxDatabaseInput) (*finspace.GetKxDatabaseOutput, error)
	GetKxDatabaseWithContext(aws.Context, *finspace.GetKxDatabaseInput, ...request.Option) (*finspace.GetKxDatabaseOutput, error)
	GetKxDatabaseRequest(*finspace.GetKxDatabaseInput) (*request.Request, *finspace.GetKxDatabaseOutput)

	GetKxEnvironment(*finspace.GetKxEnvironmentInput) (*finspace.GetKxEnvironmentOutput, error)
	GetKxEnvironmentWithContext(aws.Context, *finspace.GetKxEnvironmentInput, ...request.Option) (*finspace.GetKxEnvironmentOutput, error)
	GetKxEnvironmentRequest(*finspace.GetKxEnvironmentInput) (*request.Request, *finspace.GetKxEnvironmentOutput)

	GetKxUser(*finspace.GetKxUserInput) (*finspace.GetKxUserOutput, error)
	GetKxUserWithContext(aws.Context, *finspace.GetKxUserInput, ...request.Option) (*finspace.GetKxUserOutput, error)
	GetKxUserRequest(*finspace.GetKxUserInput) (*request.Request, *finspace.GetKxUserOutput)

	ListEnvironments(*finspace.ListEnvironmentsInput) (*finspace.ListEnvironmentsOutput, error)
	ListEnvironmentsWithContext(aws.Context, *finspace.ListEnvironmentsInput, ...request.Option) (*finspace.ListEnvironmentsOutput, error)
	ListEnvironmentsRequest(*finspace.ListEnvironmentsInput) (*request.Request, *finspace.ListEnvironmentsOutput)

	ListKxChangesets(*finspace.ListKxChangesetsInput) (*finspace.ListKxChangesetsOutput, error)
	ListKxChangesetsWithContext(aws.Context, *finspace.ListKxChangesetsInput, ...request.Option) (*finspace.ListKxChangesetsOutput, error)
	ListKxChangesetsRequest(*finspace.ListKxChangesetsInput) (*request.Request, *finspace.ListKxChangesetsOutput)

	ListKxChangesetsPages(*finspace.ListKxChangesetsInput, func(*finspace.ListKxChangesetsOutput, bool) bool) error
	ListKxChangesetsPagesWithContext(aws.Context, *finspace.ListKxChangesetsInput, func(*finspace.ListKxChangesetsOutput, bool) bool, ...request.Option) error

	ListKxClusterNodes(*finspace.ListKxClusterNodesInput) (*finspace.ListKxClusterNodesOutput, error)
	ListKxClusterNodesWithContext(aws.Context, *finspace.ListKxClusterNodesInput, ...request.Option) (*finspace.ListKxClusterNodesOutput, error)
	ListKxClusterNodesRequest(*finspace.ListKxClusterNodesInput) (*request.Request, *finspace.ListKxClusterNodesOutput)

	ListKxClusterNodesPages(*finspace.ListKxClusterNodesInput, func(*finspace.ListKxClusterNodesOutput, bool) bool) error
	ListKxClusterNodesPagesWithContext(aws.Context, *finspace.ListKxClusterNodesInput, func(*finspace.ListKxClusterNodesOutput, bool) bool, ...request.Option) error

	ListKxClusters(*finspace.ListKxClustersInput) (*finspace.ListKxClustersOutput, error)
	ListKxClustersWithContext(aws.Context, *finspace.ListKxClustersInput, ...request.Option) (*finspace.ListKxClustersOutput, error)
	ListKxClustersRequest(*finspace.ListKxClustersInput) (*request.Request, *finspace.ListKxClustersOutput)

	ListKxDatabases(*finspace.ListKxDatabasesInput) (*finspace.ListKxDatabasesOutput, error)
	ListKxDatabasesWithContext(aws.Context, *finspace.ListKxDatabasesInput, ...request.Option) (*finspace.ListKxDatabasesOutput, error)
	ListKxDatabasesRequest(*finspace.ListKxDatabasesInput) (*request.Request, *finspace.ListKxDatabasesOutput)

	ListKxDatabasesPages(*finspace.ListKxDatabasesInput, func(*finspace.ListKxDatabasesOutput, bool) bool) error
	ListKxDatabasesPagesWithContext(aws.Context, *finspace.ListKxDatabasesInput, func(*finspace.ListKxDatabasesOutput, bool) bool, ...request.Option) error

	ListKxEnvironments(*finspace.ListKxEnvironmentsInput) (*finspace.ListKxEnvironmentsOutput, error)
	ListKxEnvironmentsWithContext(aws.Context, *finspace.ListKxEnvironmentsInput, ...request.Option) (*finspace.ListKxEnvironmentsOutput, error)
	ListKxEnvironmentsRequest(*finspace.ListKxEnvironmentsInput) (*request.Request, *finspace.ListKxEnvironmentsOutput)

	ListKxEnvironmentsPages(*finspace.ListKxEnvironmentsInput, func(*finspace.ListKxEnvironmentsOutput, bool) bool) error
	ListKxEnvironmentsPagesWithContext(aws.Context, *finspace.ListKxEnvironmentsInput, func(*finspace.ListKxEnvironmentsOutput, bool) bool, ...request.Option) error

	ListKxUsers(*finspace.ListKxUsersInput) (*finspace.ListKxUsersOutput, error)
	ListKxUsersWithContext(aws.Context, *finspace.ListKxUsersInput, ...request.Option) (*finspace.ListKxUsersOutput, error)
	ListKxUsersRequest(*finspace.ListKxUsersInput) (*request.Request, *finspace.ListKxUsersOutput)

	ListTagsForResource(*finspace.ListTagsForResourceInput) (*finspace.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *finspace.ListTagsForResourceInput, ...request.Option) (*finspace.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*finspace.ListTagsForResourceInput) (*request.Request, *finspace.ListTagsForResourceOutput)

	TagResource(*finspace.TagResourceInput) (*finspace.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *finspace.TagResourceInput, ...request.Option) (*finspace.TagResourceOutput, error)
	TagResourceRequest(*finspace.TagResourceInput) (*request.Request, *finspace.TagResourceOutput)

	UntagResource(*finspace.UntagResourceInput) (*finspace.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *finspace.UntagResourceInput, ...request.Option) (*finspace.UntagResourceOutput, error)
	UntagResourceRequest(*finspace.UntagResourceInput) (*request.Request, *finspace.UntagResourceOutput)

	UpdateEnvironment(*finspace.UpdateEnvironmentInput) (*finspace.UpdateEnvironmentOutput, error)
	UpdateEnvironmentWithContext(aws.Context, *finspace.UpdateEnvironmentInput, ...request.Option) (*finspace.UpdateEnvironmentOutput, error)
	UpdateEnvironmentRequest(*finspace.UpdateEnvironmentInput) (*request.Request, *finspace.UpdateEnvironmentOutput)

	UpdateKxClusterCodeConfiguration(*finspace.UpdateKxClusterCodeConfigurationInput) (*finspace.UpdateKxClusterCodeConfigurationOutput, error)
	UpdateKxClusterCodeConfigurationWithContext(aws.Context, *finspace.UpdateKxClusterCodeConfigurationInput, ...request.Option) (*finspace.UpdateKxClusterCodeConfigurationOutput, error)
	UpdateKxClusterCodeConfigurationRequest(*finspace.UpdateKxClusterCodeConfigurationInput) (*request.Request, *finspace.UpdateKxClusterCodeConfigurationOutput)

	UpdateKxClusterDatabases(*finspace.UpdateKxClusterDatabasesInput) (*finspace.UpdateKxClusterDatabasesOutput, error)
	UpdateKxClusterDatabasesWithContext(aws.Context, *finspace.UpdateKxClusterDatabasesInput, ...request.Option) (*finspace.UpdateKxClusterDatabasesOutput, error)
	UpdateKxClusterDatabasesRequest(*finspace.UpdateKxClusterDatabasesInput) (*request.Request, *finspace.UpdateKxClusterDatabasesOutput)

	UpdateKxDatabase(*finspace.UpdateKxDatabaseInput) (*finspace.UpdateKxDatabaseOutput, error)
	UpdateKxDatabaseWithContext(aws.Context, *finspace.UpdateKxDatabaseInput, ...request.Option) (*finspace.UpdateKxDatabaseOutput, error)
	UpdateKxDatabaseRequest(*finspace.UpdateKxDatabaseInput) (*request.Request, *finspace.UpdateKxDatabaseOutput)

	UpdateKxEnvironment(*finspace.UpdateKxEnvironmentInput) (*finspace.UpdateKxEnvironmentOutput, error)
	UpdateKxEnvironmentWithContext(aws.Context, *finspace.UpdateKxEnvironmentInput, ...request.Option) (*finspace.UpdateKxEnvironmentOutput, error)
	UpdateKxEnvironmentRequest(*finspace.UpdateKxEnvironmentInput) (*request.Request, *finspace.UpdateKxEnvironmentOutput)

	UpdateKxEnvironmentNetwork(*finspace.UpdateKxEnvironmentNetworkInput) (*finspace.UpdateKxEnvironmentNetworkOutput, error)
	UpdateKxEnvironmentNetworkWithContext(aws.Context, *finspace.UpdateKxEnvironmentNetworkInput, ...request.Option) (*finspace.UpdateKxEnvironmentNetworkOutput, error)
	UpdateKxEnvironmentNetworkRequest(*finspace.UpdateKxEnvironmentNetworkInput) (*request.Request, *finspace.UpdateKxEnvironmentNetworkOutput)

	UpdateKxUser(*finspace.UpdateKxUserInput) (*finspace.UpdateKxUserOutput, error)
	UpdateKxUserWithContext(aws.Context, *finspace.UpdateKxUserInput, ...request.Option) (*finspace.UpdateKxUserOutput, error)
	UpdateKxUserRequest(*finspace.UpdateKxUserInput) (*request.Request, *finspace.UpdateKxUserOutput)
}

var _ FinspaceAPI = (*finspace.Finspace)(nil)
