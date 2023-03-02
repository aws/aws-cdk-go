package awsssm


// `S3OutputLocation` is a property of the [AWS::SSM::Association](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html) resource that specifies an Amazon S3 bucket where you want to store the results of this association request.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3OutputLocationProperty := &s3OutputLocationProperty{
//   	outputS3BucketName: jsii.String("outputS3BucketName"),
//   	outputS3KeyPrefix: jsii.String("outputS3KeyPrefix"),
//   	outputS3Region: jsii.String("outputS3Region"),
//   }
//
type CfnAssociation_S3OutputLocationProperty struct {
	// The name of the S3 bucket.
	OutputS3BucketName *string `field:"optional" json:"outputS3BucketName" yaml:"outputS3BucketName"`
	// The S3 bucket subfolder.
	OutputS3KeyPrefix *string `field:"optional" json:"outputS3KeyPrefix" yaml:"outputS3KeyPrefix"`
	// The AWS Region of the S3 bucket.
	OutputS3Region *string `field:"optional" json:"outputS3Region" yaml:"outputS3Region"`
}

