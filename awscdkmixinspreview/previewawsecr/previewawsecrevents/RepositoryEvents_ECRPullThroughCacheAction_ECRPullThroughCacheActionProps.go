package previewawsecrevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for Repository aws.ecr@ECRPullThroughCacheAction event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eCRPullThroughCacheActionProps := &ECRPullThroughCacheActionProps{
//   	EcrRepositoryPrefix: []*string{
//   		jsii.String("ecrRepositoryPrefix"),
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
//   	FailureCode: []*string{
//   		jsii.String("failureCode"),
//   	},
//   	FailureReason: []*string{
//   		jsii.String("failureReason"),
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
//   	SyncStatus: []*string{
//   		jsii.String("syncStatus"),
//   	},
//   	UpstreamRegistryUrl: []*string{
//   		jsii.String("upstreamRegistryUrl"),
//   	},
//   }
//
// Experimental.
type RepositoryEvents_ECRPullThroughCacheAction_ECRPullThroughCacheActionProps struct {
	// ecr-repository-prefix property.
	//
	// Specify an array of string values to match this event if the actual value of ecr-repository-prefix is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EcrRepositoryPrefix *[]*string `field:"optional" json:"ecrRepositoryPrefix" yaml:"ecrRepositoryPrefix"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// failure-code property.
	//
	// Specify an array of string values to match this event if the actual value of failure-code is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FailureCode *[]*string `field:"optional" json:"failureCode" yaml:"failureCode"`
	// failure-reason property.
	//
	// Specify an array of string values to match this event if the actual value of failure-reason is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FailureReason *[]*string `field:"optional" json:"failureReason" yaml:"failureReason"`
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
	// sync-status property.
	//
	// Specify an array of string values to match this event if the actual value of sync-status is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SyncStatus *[]*string `field:"optional" json:"syncStatus" yaml:"syncStatus"`
	// upstream-registry-url property.
	//
	// Specify an array of string values to match this event if the actual value of upstream-registry-url is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	UpstreamRegistryUrl *[]*string `field:"optional" json:"upstreamRegistryUrl" yaml:"upstreamRegistryUrl"`
}

