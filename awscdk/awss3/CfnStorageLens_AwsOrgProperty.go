package awss3


// This resource contains the details of the AWS Organization for Amazon S3 Storage Lens.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   awsOrgProperty := &AwsOrgProperty{
//   	Arn: jsii.String("arn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-awsorg.html
//
type CfnStorageLens_AwsOrgProperty struct {
	// This resource contains the ARN of the AWS Organization.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-awsorg.html#cfn-s3-storagelens-awsorg-arn
	//
	Arn *string `field:"required" json:"arn" yaml:"arn"`
}

