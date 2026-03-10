package previewawsemrevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.emr@EMRConfigurationError event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eMRConfigurationErrorProps := &EMRConfigurationErrorProps{
//   	ClusterId: []*string{
//   		jsii.String("clusterId"),
//   	},
//   	Description: []*string{
//   		jsii.String("description"),
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
//   	Severity: []*string{
//   		jsii.String("severity"),
//   	},
//   }
//
// Experimental.
type EMRConfigurationError_EMRConfigurationErrorProps struct {
	// clusterId property.
	//
	// Specify an array of string values to match this event if the actual value of clusterId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ClusterId *[]*string `field:"optional" json:"clusterId" yaml:"clusterId"`
	// description property.
	//
	// Specify an array of string values to match this event if the actual value of description is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Description *[]*string `field:"optional" json:"description" yaml:"description"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// severity property.
	//
	// Specify an array of string values to match this event if the actual value of severity is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Severity *[]*string `field:"optional" json:"severity" yaml:"severity"`
}

