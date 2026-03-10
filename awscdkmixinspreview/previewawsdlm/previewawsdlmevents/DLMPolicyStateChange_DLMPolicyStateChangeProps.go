package previewawsdlmevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.dlm@DLMPolicyStateChange event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dLMPolicyStateChangeProps := &DLMPolicyStateChangeProps{
//   	Cause: []*string{
//   		jsii.String("cause"),
//   	},
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
//   	PolicyId: []*string{
//   		jsii.String("policyId"),
//   	},
//   	State: []*string{
//   		jsii.String("state"),
//   	},
//   }
//
// Experimental.
type DLMPolicyStateChange_DLMPolicyStateChangeProps struct {
	// cause property.
	//
	// Specify an array of string values to match this event if the actual value of cause is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Cause *[]*string `field:"optional" json:"cause" yaml:"cause"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// policy_id property.
	//
	// Specify an array of string values to match this event if the actual value of policy_id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PolicyId *[]*string `field:"optional" json:"policyId" yaml:"policyId"`
	// state property.
	//
	// Specify an array of string values to match this event if the actual value of state is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	State *[]*string `field:"optional" json:"state" yaml:"state"`
}

