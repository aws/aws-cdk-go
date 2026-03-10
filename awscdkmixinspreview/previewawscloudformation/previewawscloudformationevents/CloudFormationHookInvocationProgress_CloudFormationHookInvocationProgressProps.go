package previewawscloudformationevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.cloudformation@CloudFormationHookInvocationProgress event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cloudFormationHookInvocationProgressProps := &CloudFormationHookInvocationProgressProps{
//   	EventMetadata: &AWSEventMetadataProps{
//   		Region: []*string{
//   			jsii.String("region"),
//   		},
//   		Resources: []*string{
//   			jsii.String("resources"),
//   		},
//   		Version: []*string{
//   			jsii.String("version"),
//   		},
//   	},
//   	HookDetail: &HookDetail{
//   		Description: []*string{
//   			jsii.String("description"),
//   		},
//   		DocumentationUrl: []*string{
//   			jsii.String("documentationUrl"),
//   		},
//   		FailureMode: []*string{
//   			jsii.String("failureMode"),
//   		},
//   		HookTypeArn: []*string{
//   			jsii.String("hookTypeArn"),
//   		},
//   		HookTypeName: []*string{
//   			jsii.String("hookTypeName"),
//   		},
//   		HookVersion: []*string{
//   			jsii.String("hookVersion"),
//   		},
//   		SourceUrl: []*string{
//   			jsii.String("sourceUrl"),
//   		},
//   	},
//   	Result: &Result{
//   		Data: []*string{
//   			jsii.String("data"),
//   		},
//   	},
//   	Status: []*string{
//   		jsii.String("status"),
//   	},
//   	StatusReason: []*string{
//   		jsii.String("statusReason"),
//   	},
//   	TargetDetail: &TargetDetail{
//   		TargetAction: []*string{
//   			jsii.String("targetAction"),
//   		},
//   		TargetId: []*string{
//   			jsii.String("targetId"),
//   		},
//   		TargetInvocationPoint: []*string{
//   			jsii.String("targetInvocationPoint"),
//   		},
//   		TargetName: []*string{
//   			jsii.String("targetName"),
//   		},
//   		TargetType: []*string{
//   			jsii.String("targetType"),
//   		},
//   	},
//   }
//
// Experimental.
type CloudFormationHookInvocationProgress_CloudFormationHookInvocationProgressProps struct {
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// hook-detail property.
	//
	// Specify an array of string values to match this event if the actual value of hook-detail is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	HookDetail *CloudFormationHookInvocationProgress_HookDetail `field:"optional" json:"hookDetail" yaml:"hookDetail"`
	// result property.
	//
	// Specify an array of string values to match this event if the actual value of result is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Result *CloudFormationHookInvocationProgress_Result `field:"optional" json:"result" yaml:"result"`
	// status property.
	//
	// Specify an array of string values to match this event if the actual value of status is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Status *[]*string `field:"optional" json:"status" yaml:"status"`
	// status-reason property.
	//
	// Specify an array of string values to match this event if the actual value of status-reason is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StatusReason *[]*string `field:"optional" json:"statusReason" yaml:"statusReason"`
	// target-detail property.
	//
	// Specify an array of string values to match this event if the actual value of target-detail is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TargetDetail *CloudFormationHookInvocationProgress_TargetDetail `field:"optional" json:"targetDetail" yaml:"targetDetail"`
}

