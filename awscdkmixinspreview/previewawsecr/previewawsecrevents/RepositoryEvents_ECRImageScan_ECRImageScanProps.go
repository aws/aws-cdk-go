package previewawsecrevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for Repository aws.ecr@ECRImageScan event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eCRImageScanProps := &ECRImageScanProps{
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
//   	FindingSeverityCounts: &FindingSeverityCounts{
//   		Critical: []*string{
//   			jsii.String("critical"),
//   		},
//   		High: []*string{
//   			jsii.String("high"),
//   		},
//   		Informational: []*string{
//   			jsii.String("informational"),
//   		},
//   		Low: []*string{
//   			jsii.String("low"),
//   		},
//   		Medium: []*string{
//   			jsii.String("medium"),
//   		},
//   		Undefined: []*string{
//   			jsii.String("undefined"),
//   		},
//   	},
//   	ImageDigest: []*string{
//   		jsii.String("imageDigest"),
//   	},
//   	ImageTags: []*string{
//   		jsii.String("imageTags"),
//   	},
//   	RepositoryName: []*string{
//   		jsii.String("repositoryName"),
//   	},
//   	ScanStatus: []*string{
//   		jsii.String("scanStatus"),
//   	},
//   }
//
// Experimental.
type RepositoryEvents_ECRImageScan_ECRImageScanProps struct {
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// finding-severity-counts property.
	//
	// Specify an array of string values to match this event if the actual value of finding-severity-counts is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FindingSeverityCounts *RepositoryEvents_ECRImageScan_FindingSeverityCounts `field:"optional" json:"findingSeverityCounts" yaml:"findingSeverityCounts"`
	// image-digest property.
	//
	// Specify an array of string values to match this event if the actual value of image-digest is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ImageDigest *[]*string `field:"optional" json:"imageDigest" yaml:"imageDigest"`
	// image-tags property.
	//
	// Specify an array of string values to match this event if the actual value of image-tags is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ImageTags *[]*string `field:"optional" json:"imageTags" yaml:"imageTags"`
	// repository-name property.
	//
	// Specify an array of string values to match this event if the actual value of repository-name is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Filter with the Repository reference.
	//
	// Experimental.
	RepositoryName *[]*string `field:"optional" json:"repositoryName" yaml:"repositoryName"`
	// scan-status property.
	//
	// Specify an array of string values to match this event if the actual value of scan-status is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ScanStatus *[]*string `field:"optional" json:"scanStatus" yaml:"scanStatus"`
}

