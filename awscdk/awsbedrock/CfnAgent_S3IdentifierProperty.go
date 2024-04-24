package awsbedrock


// Contains information about the S3 object containing the resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3IdentifierProperty := &S3IdentifierProperty{
//   	S3BucketName: jsii.String("s3BucketName"),
//   	S3ObjectKey: jsii.String("s3ObjectKey"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-s3identifier.html
//
type CfnAgent_S3IdentifierProperty struct {
	// The name of the S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-s3identifier.html#cfn-bedrock-agent-s3identifier-s3bucketname
	//
	S3BucketName *string `field:"optional" json:"s3BucketName" yaml:"s3BucketName"`
	// The S3 object key containing the resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-s3identifier.html#cfn-bedrock-agent-s3identifier-s3objectkey
	//
	S3ObjectKey *string `field:"optional" json:"s3ObjectKey" yaml:"s3ObjectKey"`
}

