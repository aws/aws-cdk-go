package awss3vectors


// Properties for defining a `CfnVectorBucketPolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var policy interface{}
//
//   cfnVectorBucketPolicyProps := &CfnVectorBucketPolicyProps{
//   	Policy: policy,
//
//   	// the properties below are optional
//   	VectorBucketArn: jsii.String("vectorBucketArn"),
//   	VectorBucketName: jsii.String("vectorBucketName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3vectors-vectorbucketpolicy.html
//
type CfnVectorBucketPolicyProps struct {
	// A policy document containing permissions to add to the specified vector bucket.
	//
	// In IAM , you must provide policy documents in JSON format. However, in CloudFormation you can provide the policy in JSON or YAML format because CloudFormation converts YAML to JSON before submitting it to IAM .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3vectors-vectorbucketpolicy.html#cfn-s3vectors-vectorbucketpolicy-policy
	//
	Policy interface{} `field:"required" json:"policy" yaml:"policy"`
	// The Amazon Resource Name (ARN) of the S3 vector bucket to which the policy applies.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3vectors-vectorbucketpolicy.html#cfn-s3vectors-vectorbucketpolicy-vectorbucketarn
	//
	VectorBucketArn *string `field:"optional" json:"vectorBucketArn" yaml:"vectorBucketArn"`
	// The name of the S3 vector bucket to which the policy applies.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3vectors-vectorbucketpolicy.html#cfn-s3vectors-vectorbucketpolicy-vectorbucketname
	//
	VectorBucketName *string `field:"optional" json:"vectorBucketName" yaml:"vectorBucketName"`
}

