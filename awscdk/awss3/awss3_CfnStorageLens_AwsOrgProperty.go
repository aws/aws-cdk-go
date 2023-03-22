package awss3


// This resource contains the details of the AWS Organization for Amazon S3 Storage Lens.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   awsOrgProperty := &awsOrgProperty{
//   	arn: jsii.String("arn"),
//   }
//
type CfnStorageLens_AwsOrgProperty struct {
	// This resource contains the ARN of the AWS Organization.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
}

