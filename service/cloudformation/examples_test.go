// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudformation_test

import (
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
)

var _ time.Duration
var _ strings.Reader
var _ aws.Config

func parseTime(layout, value string) *time.Time {
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return &t
}

// To create a generated template
// This example creates a generated template with a resources file.
func ExampleCloudFormation_CreateGeneratedTemplate_shared00() {
	svc := cloudformation.New(session.New())
	input := &cloudformation.CreateGeneratedTemplateInput{
		GeneratedTemplateName: aws.String("JazzyTemplate"),
		Resources: []*cloudformation.ResourceDefinition{
			{
				ResourceIdentifier: map[string]*string{
					"BucketName": aws.String("jazz-bucket"),
				},
				ResourceType: aws.String("AWS::S3::Bucket"),
			},
			{
				ResourceIdentifier: map[string]*string{
					"DhcpOptionsId": aws.String("random-id123"),
				},
				ResourceType: aws.String("AWS::EC2::DHCPOptions"),
			},
		},
	}

	result, err := svc.CreateGeneratedTemplate(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case cloudformation.ErrCodeAlreadyExistsException:
				fmt.Println(cloudformation.ErrCodeAlreadyExistsException, aerr.Error())
			case cloudformation.ErrCodeLimitExceededException:
				fmt.Println(cloudformation.ErrCodeLimitExceededException, aerr.Error())
			case cloudformation.ErrCodeConcurrentResourcesLimitExceededException:
				fmt.Println(cloudformation.ErrCodeConcurrentResourcesLimitExceededException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To delete a generated template
// This example deletes a generated template
func ExampleCloudFormation_DeleteGeneratedTemplate_shared00() {
	svc := cloudformation.New(session.New())
	input := &cloudformation.DeleteGeneratedTemplateInput{
		GeneratedTemplateName: aws.String("JazzyTemplate"),
	}

	result, err := svc.DeleteGeneratedTemplate(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case cloudformation.ErrCodeGeneratedTemplateNotFoundException:
				fmt.Println(cloudformation.ErrCodeGeneratedTemplateNotFoundException, aerr.Error())
			case cloudformation.ErrCodeConcurrentResourcesLimitExceededException:
				fmt.Println(cloudformation.ErrCodeConcurrentResourcesLimitExceededException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To describe a generated template
// This example describes a generated template
func ExampleCloudFormation_DescribeGeneratedTemplate_shared00() {
	svc := cloudformation.New(session.New())
	input := &cloudformation.DescribeGeneratedTemplateInput{
		GeneratedTemplateName: aws.String("JazzyTemplate"),
	}

	result, err := svc.DescribeGeneratedTemplate(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case cloudformation.ErrCodeGeneratedTemplateNotFoundException:
				fmt.Println(cloudformation.ErrCodeGeneratedTemplateNotFoundException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To describe a selected resource scan
// This example describes a selected resource scan
func ExampleCloudFormation_DescribeResourceScan_shared00() {
	svc := cloudformation.New(session.New())
	input := &cloudformation.DescribeResourceScanInput{
		ResourceScanId: aws.String("arn:aws:cloudformation:us-east-1:123456789012:resourceScan/c19304f6-c4f1-4ff8-8e1f-35162e41d7e1"),
	}

	result, err := svc.DescribeResourceScan(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case cloudformation.ErrCodeResourceScanNotFoundException:
				fmt.Println(cloudformation.ErrCodeResourceScanNotFoundException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To get a generated template in JSON format
// This example gets a generated template ins JSON format.
func ExampleCloudFormation_GetGeneratedTemplate_shared00() {
	svc := cloudformation.New(session.New())
	input := &cloudformation.GetGeneratedTemplateInput{
		GeneratedTemplateName: aws.String("JazzyTemplate"),
	}

	result, err := svc.GetGeneratedTemplate(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case cloudformation.ErrCodeGeneratedTemplateNotFoundException:
				fmt.Println(cloudformation.ErrCodeGeneratedTemplateNotFoundException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To get a generated template in YAML format
// This example gets a generated template in YAML format.
func ExampleCloudFormation_GetGeneratedTemplate_shared01() {
	svc := cloudformation.New(session.New())
	input := &cloudformation.GetGeneratedTemplateInput{
		Format:                aws.String("YAML"),
		GeneratedTemplateName: aws.String("JazzyTemplate"),
	}

	result, err := svc.GetGeneratedTemplate(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case cloudformation.ErrCodeGeneratedTemplateNotFoundException:
				fmt.Println(cloudformation.ErrCodeGeneratedTemplateNotFoundException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To list generated templates
// This example lists the generated templates.
func ExampleCloudFormation_ListGeneratedTemplates_shared00() {
	svc := cloudformation.New(session.New())
	input := &cloudformation.ListGeneratedTemplatesInput{}

	result, err := svc.ListGeneratedTemplates(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To list resource scan related resources
// This example lists the resources related to the passed in resources
func ExampleCloudFormation_ListResourceScanRelatedResources_shared00() {
	svc := cloudformation.New(session.New())
	input := &cloudformation.ListResourceScanRelatedResourcesInput{
		ResourceScanId: aws.String("arn:aws:cloudformation:us-east-1:123456789012:resourceScan/c19304f6-c4f1-4ff8-8e1f-35162e41d7e1"),
		Resources: []*cloudformation.ScannedResourceIdentifier{
			{
				ResourceIdentifier: map[string]*string{
					"BucketName": aws.String("jazz-bucket"),
				},
				ResourceType: aws.String("AWS::S3::Bucket"),
			},
			{
				ResourceIdentifier: map[string]*string{
					"DhcpOptionsId": aws.String("random-id123"),
				},
				ResourceType: aws.String("AWS::EC2::DHCPOptions"),
			},
		},
	}

	result, err := svc.ListResourceScanRelatedResources(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case cloudformation.ErrCodeResourceScanNotFoundException:
				fmt.Println(cloudformation.ErrCodeResourceScanNotFoundException, aerr.Error())
			case cloudformation.ErrCodeResourceScanInProgressException:
				fmt.Println(cloudformation.ErrCodeResourceScanInProgressException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To list the resources in your resource scan
// This example lists the resources in your resource scan
func ExampleCloudFormation_ListResourceScanResources_shared00() {
	svc := cloudformation.New(session.New())
	input := &cloudformation.ListResourceScanResourcesInput{
		ResourceScanId: aws.String("arn:aws:cloudformation:us-east-1:123456789012:resourceScan/c19304f6-c4f1-4ff8-8e1f-35162e41d7e1"),
	}

	result, err := svc.ListResourceScanResources(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case cloudformation.ErrCodeResourceScanNotFoundException:
				fmt.Println(cloudformation.ErrCodeResourceScanNotFoundException, aerr.Error())
			case cloudformation.ErrCodeResourceScanInProgressException:
				fmt.Println(cloudformation.ErrCodeResourceScanInProgressException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To list the resources in your resource scan for specific resource type
// This example lists the resources in your resource scan filtering only the resources
// that start with the passed in prefix
func ExampleCloudFormation_ListResourceScanResources_shared01() {
	svc := cloudformation.New(session.New())
	input := &cloudformation.ListResourceScanResourcesInput{
		ResourceScanId:     aws.String("arn:aws:cloudformation:us-east-1:123456789012:resourceScan/c19304f6-c4f1-4ff8-8e1f-35162e41d7e1"),
		ResourceTypePrefix: aws.String("AWS::S3"),
	}

	result, err := svc.ListResourceScanResources(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case cloudformation.ErrCodeResourceScanNotFoundException:
				fmt.Println(cloudformation.ErrCodeResourceScanNotFoundException, aerr.Error())
			case cloudformation.ErrCodeResourceScanInProgressException:
				fmt.Println(cloudformation.ErrCodeResourceScanInProgressException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// Listing Resource Scans
// This example shows how to list resource scans
func ExampleCloudFormation_ListResourceScans_shared00() {
	svc := cloudformation.New(session.New())
	input := &cloudformation.ListResourceScansInput{}

	result, err := svc.ListResourceScans(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To start a resource scan
// This example shows how to start a new resource scan
func ExampleCloudFormation_StartResourceScan_shared00() {
	svc := cloudformation.New(session.New())
	input := &cloudformation.StartResourceScanInput{}

	result, err := svc.StartResourceScan(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case cloudformation.ErrCodeResourceScanInProgressException:
				fmt.Println(cloudformation.ErrCodeResourceScanInProgressException, aerr.Error())
			case cloudformation.ErrCodeResourceScanLimitExceededException:
				fmt.Println(cloudformation.ErrCodeResourceScanLimitExceededException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To update a generated template's name
// This example updates a generated template with a new name.
func ExampleCloudFormation_UpdateGeneratedTemplate_shared00() {
	svc := cloudformation.New(session.New())
	input := &cloudformation.UpdateGeneratedTemplateInput{
		GeneratedTemplateName:    aws.String("JazzyTemplate"),
		NewGeneratedTemplateName: aws.String("JazzierTemplate"),
	}

	result, err := svc.UpdateGeneratedTemplate(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case cloudformation.ErrCodeAlreadyExistsException:
				fmt.Println(cloudformation.ErrCodeAlreadyExistsException, aerr.Error())
			case cloudformation.ErrCodeGeneratedTemplateNotFoundException:
				fmt.Println(cloudformation.ErrCodeGeneratedTemplateNotFoundException, aerr.Error())
			case cloudformation.ErrCodeLimitExceededException:
				fmt.Println(cloudformation.ErrCodeLimitExceededException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To remove resources from a generated template
// This example removes resources from a generated template
func ExampleCloudFormation_UpdateGeneratedTemplate_shared01() {
	svc := cloudformation.New(session.New())
	input := &cloudformation.UpdateGeneratedTemplateInput{
		GeneratedTemplateName: aws.String("JazzyTemplate"),
		RemoveResources: []*string{
			aws.String("LogicalResourceId1"),
			aws.String("LogicalResourceId2"),
		},
	}

	result, err := svc.UpdateGeneratedTemplate(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case cloudformation.ErrCodeAlreadyExistsException:
				fmt.Println(cloudformation.ErrCodeAlreadyExistsException, aerr.Error())
			case cloudformation.ErrCodeGeneratedTemplateNotFoundException:
				fmt.Println(cloudformation.ErrCodeGeneratedTemplateNotFoundException, aerr.Error())
			case cloudformation.ErrCodeLimitExceededException:
				fmt.Println(cloudformation.ErrCodeLimitExceededException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To add resources to a generated template
// This example adds resources to a generated template
func ExampleCloudFormation_UpdateGeneratedTemplate_shared02() {
	svc := cloudformation.New(session.New())
	input := &cloudformation.UpdateGeneratedTemplateInput{
		AddResources: []*cloudformation.ResourceDefinition{
			{
				ResourceIdentifier: map[string]*string{
					"BucketName": aws.String("jazz-bucket"),
				},
				ResourceType: aws.String("AWS::S3::Bucket"),
			},
			{
				ResourceIdentifier: map[string]*string{
					"DhcpOptionsId": aws.String("random-id123"),
				},
				ResourceType: aws.String("AWS::EC2::DHCPOptions"),
			},
		},
		GeneratedTemplateName: aws.String("JazzyTemplate"),
	}

	result, err := svc.UpdateGeneratedTemplate(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case cloudformation.ErrCodeAlreadyExistsException:
				fmt.Println(cloudformation.ErrCodeAlreadyExistsException, aerr.Error())
			case cloudformation.ErrCodeGeneratedTemplateNotFoundException:
				fmt.Println(cloudformation.ErrCodeGeneratedTemplateNotFoundException, aerr.Error())
			case cloudformation.ErrCodeLimitExceededException:
				fmt.Println(cloudformation.ErrCodeLimitExceededException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}
