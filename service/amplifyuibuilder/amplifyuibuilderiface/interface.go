// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package amplifyuibuilderiface provides an interface to enable mocking the AWS Amplify UI Builder service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package amplifyuibuilderiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/amplifyuibuilder"
)

// AmplifyUIBuilderAPI provides an interface to enable mocking the
// amplifyuibuilder.AmplifyUIBuilder service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// AWS Amplify UI Builder.
//	func myFunc(svc amplifyuibuilderiface.AmplifyUIBuilderAPI) bool {
//	    // Make svc.CreateComponent request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := amplifyuibuilder.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockAmplifyUIBuilderClient struct {
//	    amplifyuibuilderiface.AmplifyUIBuilderAPI
//	}
//	func (m *mockAmplifyUIBuilderClient) CreateComponent(input *amplifyuibuilder.CreateComponentInput) (*amplifyuibuilder.CreateComponentOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockAmplifyUIBuilderClient{}
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
type AmplifyUIBuilderAPI interface {
	CreateComponent(*amplifyuibuilder.CreateComponentInput) (*amplifyuibuilder.CreateComponentOutput, error)
	CreateComponentWithContext(aws.Context, *amplifyuibuilder.CreateComponentInput, ...request.Option) (*amplifyuibuilder.CreateComponentOutput, error)
	CreateComponentRequest(*amplifyuibuilder.CreateComponentInput) (*request.Request, *amplifyuibuilder.CreateComponentOutput)

	CreateTheme(*amplifyuibuilder.CreateThemeInput) (*amplifyuibuilder.CreateThemeOutput, error)
	CreateThemeWithContext(aws.Context, *amplifyuibuilder.CreateThemeInput, ...request.Option) (*amplifyuibuilder.CreateThemeOutput, error)
	CreateThemeRequest(*amplifyuibuilder.CreateThemeInput) (*request.Request, *amplifyuibuilder.CreateThemeOutput)

	DeleteComponent(*amplifyuibuilder.DeleteComponentInput) (*amplifyuibuilder.DeleteComponentOutput, error)
	DeleteComponentWithContext(aws.Context, *amplifyuibuilder.DeleteComponentInput, ...request.Option) (*amplifyuibuilder.DeleteComponentOutput, error)
	DeleteComponentRequest(*amplifyuibuilder.DeleteComponentInput) (*request.Request, *amplifyuibuilder.DeleteComponentOutput)

	DeleteTheme(*amplifyuibuilder.DeleteThemeInput) (*amplifyuibuilder.DeleteThemeOutput, error)
	DeleteThemeWithContext(aws.Context, *amplifyuibuilder.DeleteThemeInput, ...request.Option) (*amplifyuibuilder.DeleteThemeOutput, error)
	DeleteThemeRequest(*amplifyuibuilder.DeleteThemeInput) (*request.Request, *amplifyuibuilder.DeleteThemeOutput)

	ExchangeCodeForToken(*amplifyuibuilder.ExchangeCodeForTokenInput) (*amplifyuibuilder.ExchangeCodeForTokenOutput, error)
	ExchangeCodeForTokenWithContext(aws.Context, *amplifyuibuilder.ExchangeCodeForTokenInput, ...request.Option) (*amplifyuibuilder.ExchangeCodeForTokenOutput, error)
	ExchangeCodeForTokenRequest(*amplifyuibuilder.ExchangeCodeForTokenInput) (*request.Request, *amplifyuibuilder.ExchangeCodeForTokenOutput)

	ExportComponents(*amplifyuibuilder.ExportComponentsInput) (*amplifyuibuilder.ExportComponentsOutput, error)
	ExportComponentsWithContext(aws.Context, *amplifyuibuilder.ExportComponentsInput, ...request.Option) (*amplifyuibuilder.ExportComponentsOutput, error)
	ExportComponentsRequest(*amplifyuibuilder.ExportComponentsInput) (*request.Request, *amplifyuibuilder.ExportComponentsOutput)

	ExportComponentsPages(*amplifyuibuilder.ExportComponentsInput, func(*amplifyuibuilder.ExportComponentsOutput, bool) bool) error
	ExportComponentsPagesWithContext(aws.Context, *amplifyuibuilder.ExportComponentsInput, func(*amplifyuibuilder.ExportComponentsOutput, bool) bool, ...request.Option) error

	ExportThemes(*amplifyuibuilder.ExportThemesInput) (*amplifyuibuilder.ExportThemesOutput, error)
	ExportThemesWithContext(aws.Context, *amplifyuibuilder.ExportThemesInput, ...request.Option) (*amplifyuibuilder.ExportThemesOutput, error)
	ExportThemesRequest(*amplifyuibuilder.ExportThemesInput) (*request.Request, *amplifyuibuilder.ExportThemesOutput)

	ExportThemesPages(*amplifyuibuilder.ExportThemesInput, func(*amplifyuibuilder.ExportThemesOutput, bool) bool) error
	ExportThemesPagesWithContext(aws.Context, *amplifyuibuilder.ExportThemesInput, func(*amplifyuibuilder.ExportThemesOutput, bool) bool, ...request.Option) error

	GetComponent(*amplifyuibuilder.GetComponentInput) (*amplifyuibuilder.GetComponentOutput, error)
	GetComponentWithContext(aws.Context, *amplifyuibuilder.GetComponentInput, ...request.Option) (*amplifyuibuilder.GetComponentOutput, error)
	GetComponentRequest(*amplifyuibuilder.GetComponentInput) (*request.Request, *amplifyuibuilder.GetComponentOutput)

	GetTheme(*amplifyuibuilder.GetThemeInput) (*amplifyuibuilder.GetThemeOutput, error)
	GetThemeWithContext(aws.Context, *amplifyuibuilder.GetThemeInput, ...request.Option) (*amplifyuibuilder.GetThemeOutput, error)
	GetThemeRequest(*amplifyuibuilder.GetThemeInput) (*request.Request, *amplifyuibuilder.GetThemeOutput)

	ListComponents(*amplifyuibuilder.ListComponentsInput) (*amplifyuibuilder.ListComponentsOutput, error)
	ListComponentsWithContext(aws.Context, *amplifyuibuilder.ListComponentsInput, ...request.Option) (*amplifyuibuilder.ListComponentsOutput, error)
	ListComponentsRequest(*amplifyuibuilder.ListComponentsInput) (*request.Request, *amplifyuibuilder.ListComponentsOutput)

	ListComponentsPages(*amplifyuibuilder.ListComponentsInput, func(*amplifyuibuilder.ListComponentsOutput, bool) bool) error
	ListComponentsPagesWithContext(aws.Context, *amplifyuibuilder.ListComponentsInput, func(*amplifyuibuilder.ListComponentsOutput, bool) bool, ...request.Option) error

	ListThemes(*amplifyuibuilder.ListThemesInput) (*amplifyuibuilder.ListThemesOutput, error)
	ListThemesWithContext(aws.Context, *amplifyuibuilder.ListThemesInput, ...request.Option) (*amplifyuibuilder.ListThemesOutput, error)
	ListThemesRequest(*amplifyuibuilder.ListThemesInput) (*request.Request, *amplifyuibuilder.ListThemesOutput)

	ListThemesPages(*amplifyuibuilder.ListThemesInput, func(*amplifyuibuilder.ListThemesOutput, bool) bool) error
	ListThemesPagesWithContext(aws.Context, *amplifyuibuilder.ListThemesInput, func(*amplifyuibuilder.ListThemesOutput, bool) bool, ...request.Option) error

	RefreshToken(*amplifyuibuilder.RefreshTokenInput) (*amplifyuibuilder.RefreshTokenOutput, error)
	RefreshTokenWithContext(aws.Context, *amplifyuibuilder.RefreshTokenInput, ...request.Option) (*amplifyuibuilder.RefreshTokenOutput, error)
	RefreshTokenRequest(*amplifyuibuilder.RefreshTokenInput) (*request.Request, *amplifyuibuilder.RefreshTokenOutput)

	UpdateComponent(*amplifyuibuilder.UpdateComponentInput) (*amplifyuibuilder.UpdateComponentOutput, error)
	UpdateComponentWithContext(aws.Context, *amplifyuibuilder.UpdateComponentInput, ...request.Option) (*amplifyuibuilder.UpdateComponentOutput, error)
	UpdateComponentRequest(*amplifyuibuilder.UpdateComponentInput) (*request.Request, *amplifyuibuilder.UpdateComponentOutput)

	UpdateTheme(*amplifyuibuilder.UpdateThemeInput) (*amplifyuibuilder.UpdateThemeOutput, error)
	UpdateThemeWithContext(aws.Context, *amplifyuibuilder.UpdateThemeInput, ...request.Option) (*amplifyuibuilder.UpdateThemeOutput, error)
	UpdateThemeRequest(*amplifyuibuilder.UpdateThemeInput) (*request.Request, *amplifyuibuilder.UpdateThemeOutput)
}

var _ AmplifyUIBuilderAPI = (*amplifyuibuilder.AmplifyUIBuilder)(nil)
