package awscdk


// Properties for AWS EventBridge event metadata.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var bucket Bucket
//
//   bucketEvents := awscdkmixinspreview.BucketEvents_FromBucket(bucket)
//
//   pattern := bucketEvents.ObjectCreatedPattern(&ObjectCreatedProps{
//   	EventMetadata: &AWSEventMetadataProps{
//   		Region: []*string{
//   			jsii.String("us-east-1"),
//   			jsii.String("us-west-2"),
//   		},
//   		Version: []*string{
//   			jsii.String("0"),
//   		},
//   	},
//   })
//
type AWSEventMetadataProps struct {
	// Identifies the AWS region where the event originated.
	// Default: - No filtering on region.
	//
	Region *[]*string `field:"optional" json:"region" yaml:"region"`
	// This JSON array contains ARNs that identify resources that are involved in the event.
	//
	// Inclusion of these ARNs is at the discretion of the
	// service.
	//
	// For example, Amazon EC2 instance state-changes include Amazon EC2
	// instance ARNs, Auto Scaling events include ARNs for both instances and
	// Auto Scaling groups, but API calls with AWS CloudTrail do not include
	// resource ARNs.
	// Default: - No filtering on resource.
	//
	Resources *[]*string `field:"optional" json:"resources" yaml:"resources"`
	// By default, this is set to 0 (zero) in all events.
	// Default: - No filtering on version.
	//
	Version *[]*string `field:"optional" json:"version" yaml:"version"`
}

