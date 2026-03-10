package previewawsemrevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.emr@EMRClusterStateChange event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eMRClusterStateChangeProps := &EMRClusterStateChangeProps{
//   	ClusterId: []*string{
//   		jsii.String("clusterId"),
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
//   	Message: []*string{
//   		jsii.String("message"),
//   	},
//   	Name: []*string{
//   		jsii.String("name"),
//   	},
//   	Severity: []*string{
//   		jsii.String("severity"),
//   	},
//   	State: []*string{
//   		jsii.String("state"),
//   	},
//   	StateChangeReason: []*string{
//   		jsii.String("stateChangeReason"),
//   	},
//   }
//
// Experimental.
type EMRClusterStateChange_EMRClusterStateChangeProps struct {
	// clusterId property.
	//
	// Specify an array of string values to match this event if the actual value of clusterId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ClusterId *[]*string `field:"optional" json:"clusterId" yaml:"clusterId"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// message property.
	//
	// Specify an array of string values to match this event if the actual value of message is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Message *[]*string `field:"optional" json:"message" yaml:"message"`
	// name property.
	//
	// Specify an array of string values to match this event if the actual value of name is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Name *[]*string `field:"optional" json:"name" yaml:"name"`
	// severity property.
	//
	// Specify an array of string values to match this event if the actual value of severity is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Severity *[]*string `field:"optional" json:"severity" yaml:"severity"`
	// state property.
	//
	// Specify an array of string values to match this event if the actual value of state is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	State *[]*string `field:"optional" json:"state" yaml:"state"`
	// stateChangeReason property.
	//
	// Specify an array of string values to match this event if the actual value of stateChangeReason is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StateChangeReason *[]*string `field:"optional" json:"stateChangeReason" yaml:"stateChangeReason"`
}

