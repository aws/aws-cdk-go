package previewawssecurityhubevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.securityhub@SecurityHubFindingsCustomAction event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   securityHubFindingsCustomActionProps := &SecurityHubFindingsCustomActionProps{
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
//   	Findings: []*string{
//   		jsii.String("findings"),
//   	},
//   }
//
// Experimental.
type SecurityHubFindingsCustomAction_SecurityHubFindingsCustomActionProps struct {
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// findings property.
	//
	// Specify an array of string values to match this event if the actual value of findings is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Findings *[]*string `field:"optional" json:"findings" yaml:"findings"`
}

