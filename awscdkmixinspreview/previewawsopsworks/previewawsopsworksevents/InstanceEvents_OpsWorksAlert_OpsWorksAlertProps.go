package previewawsopsworksevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for Instance aws.opsworks@OpsWorksAlert event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   opsWorksAlertProps := &OpsWorksAlertProps{
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
//   	InstanceId: []*string{
//   		jsii.String("instanceId"),
//   	},
//   	Message: []*string{
//   		jsii.String("message"),
//   	},
//   	StackId: []*string{
//   		jsii.String("stackId"),
//   	},
//   	Type: []*string{
//   		jsii.String("type"),
//   	},
//   }
//
// Experimental.
type InstanceEvents_OpsWorksAlert_OpsWorksAlertProps struct {
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// instance-id property.
	//
	// Specify an array of string values to match this event if the actual value of instance-id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Filter with the Instance reference.
	//
	// Experimental.
	InstanceId *[]*string `field:"optional" json:"instanceId" yaml:"instanceId"`
	// message property.
	//
	// Specify an array of string values to match this event if the actual value of message is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Message *[]*string `field:"optional" json:"message" yaml:"message"`
	// stack-id property.
	//
	// Specify an array of string values to match this event if the actual value of stack-id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StackId *[]*string `field:"optional" json:"stackId" yaml:"stackId"`
	// type property.
	//
	// Specify an array of string values to match this event if the actual value of type is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Type *[]*string `field:"optional" json:"type" yaml:"type"`
}

