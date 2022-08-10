// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package ssooidc provides the client and types for making API
// requests to AWS SSO OIDC.
//
// Amazon Web Services Single Sign On OpenID Connect (OIDC) is a web service
// that enables a client (such as Amazon Web Services CLI or a native application)
// to register with Amazon Web Services SSO. The service also enables the client
// to fetch the user’s access token upon successful authentication and authorization
// with Amazon Web Services SSO.
//
// Although Amazon Web Services Single Sign-On was renamed, the sso and identitystore
// API namespaces will continue to retain their original name for backward compatibility
// purposes. For more information, see Amazon Web Services SSO rename (https://docs.aws.amazon.com/singlesignon/latest/userguide/what-is.html#renamed).
//
// # Considerations for Using This Guide
//
// Before you begin using this guide, we recommend that you first review the
// following important information about how the Amazon Web Services SSO OIDC
// service works.
//
//   - The Amazon Web Services SSO OIDC service currently implements only the
//     portions of the OAuth 2.0 Device Authorization Grant standard (https://tools.ietf.org/html/rfc8628
//     (https://tools.ietf.org/html/rfc8628)) that are necessary to enable single
//     sign-on authentication with the AWS CLI. Support for other OIDC flows
//     frequently needed for native applications, such as Authorization Code
//     Flow (+ PKCE), will be addressed in future releases.
//
//   - The service emits only OIDC access tokens, such that obtaining a new
//     token (For example, token refresh) requires explicit user re-authentication.
//
//   - The access tokens provided by this service grant access to all AWS account
//     entitlements assigned to an Amazon Web Services SSO user, not just a particular
//     application.
//
//   - The documentation in this guide does not describe the mechanism to convert
//     the access token into AWS Auth (“sigv4”) credentials for use with
//     IAM-protected AWS service endpoints. For more information, see GetRoleCredentials
//     (https://docs.aws.amazon.com/singlesignon/latest/PortalAPIReference/API_GetRoleCredentials.html)
//     in the Amazon Web Services SSO Portal API Reference Guide.
//
// For general information about Amazon Web Services SSO, see What is Amazon
// Web Services SSO? (https://docs.aws.amazon.com/singlesignon/latest/userguide/what-is.html)
// in the Amazon Web Services SSO User Guide.
//
// See https://docs.aws.amazon.com/goto/WebAPI/sso-oidc-2019-06-10 for more information on this service.
//
// See ssooidc package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/ssooidc/
//
// # Using the Client
//
// To contact AWS SSO OIDC with the SDK use the New function to create
// a new service client. With that client you can make API requests to the service.
// These clients are safe to use concurrently.
//
// See the SDK's documentation for more information on how to use the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws.Config documentation for more information on configuring SDK clients.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the AWS SSO OIDC client SSOOIDC for more
// information on creating client for this service.
// https://docs.aws.amazon.com/sdk-for-go/api/service/ssooidc/#New
package ssooidc
