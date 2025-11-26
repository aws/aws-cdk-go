package previewawsecrevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for Repository aws.ecr@ECRReplicationAction event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eCRReplicationActionProps := &ECRReplicationActionProps{
//   	ActionType: []*string{
//   		jsii.String("actionType"),
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
//   	ImageDigest: []*string{
//   		jsii.String("imageDigest"),
//   	},
//   	ImageTag: []*string{
//   		jsii.String("imageTag"),
//   	},
//   	RepositoryName: []*string{
//   		jsii.String("repositoryName"),
//   	},
//   	Result: []*string{
//   		jsii.String("result"),
//   	},
//   	SourceAccount: []*string{
//   		jsii.String("sourceAccount"),
//   	},
//   	SourceRegion: []*string{
//   		jsii.String("sourceRegion"),
//   	},
//   }
//
// Experimental.
type RepositoryEvents_ECRReplicationAction_ECRReplicationActionProps struct {
	// action-type property.
	//
	// Specify an array of string values to match this event if the actual value of action-type is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ActionType *[]*string `field:"optional" json:"actionType" yaml:"actionType"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// image-digest property.
	//
	// Specify an array of string values to match this event if the actual value of image-digest is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ImageDigest *[]*string `field:"optional" json:"imageDigest" yaml:"imageDigest"`
	// image-tag property.
	//
	// Specify an array of string values to match this event if the actual value of image-tag is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ImageTag *[]*string `field:"optional" json:"imageTag" yaml:"imageTag"`
	// repository-name property.
	//
	// Specify an array of string values to match this event if the actual value of repository-name is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Filter with the Repository reference.
	//
	// Experimental.
	RepositoryName *[]*string `field:"optional" json:"repositoryName" yaml:"repositoryName"`
	// result property.
	//
	// Specify an array of string values to match this event if the actual value of result is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Result *[]*string `field:"optional" json:"result" yaml:"result"`
	// source-account property.
	//
	// Specify an array of string values to match this event if the actual value of source-account is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SourceAccount *[]*string `field:"optional" json:"sourceAccount" yaml:"sourceAccount"`
	// source-region property.
	//
	// Specify an array of string values to match this event if the actual value of source-region is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SourceRegion *[]*string `field:"optional" json:"sourceRegion" yaml:"sourceRegion"`
}

