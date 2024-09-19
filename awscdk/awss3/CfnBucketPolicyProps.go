package awss3


// Properties for defining a `CfnBucketPolicy`.
//
// Example:
//   bucketName := "my-favorite-bucket-name"
//   accessLogsBucket := s3.NewBucket(this, jsii.String("AccessLogsBucket"), &BucketProps{
//   	ObjectOwnership: s3.ObjectOwnership_BUCKET_OWNER_ENFORCED,
//   	BucketName: jsii.String(BucketName),
//   })
//
//   // Creating a bucket policy using L1
//   bucketPolicy := s3.NewCfnBucketPolicy(this, jsii.String("BucketPolicy"), &CfnBucketPolicyProps{
//   	Bucket: bucketName,
//   	PolicyDocument: map[string]interface{}{
//   		"Statement": []map[string]interface{}{
//   			map[string]interface{}{
//   				"Action": jsii.String("s3:*"),
//   				"Effect": jsii.String("Deny"),
//   				"Principal": map[string]*string{
//   					"AWS": jsii.String("*"),
//   				},
//   				"Resource": []*string{
//   					accessLogsBucket.bucketArn,
//   					fmt.Sprintf("%v/*", accessLogsBucket.bucketArn),
//   				},
//   			},
//   		},
//   		"Version": jsii.String("2012-10-17"),
//   	},
//   })
//
//   // 'serverAccessLogsBucket' will create a new L2 bucket policy
//   // to allow log delivery and overwrite the L1 bucket policy.
//   bucket := s3.NewBucket(this, jsii.String("MyBucket"), &BucketProps{
//   	ServerAccessLogsBucket: accessLogsBucket,
//   	ServerAccessLogsPrefix: jsii.String("logs"),
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-bucketpolicy.html
//
type CfnBucketPolicyProps struct {
	// The name of the Amazon S3 bucket to which the policy applies.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-bucketpolicy.html#cfn-s3-bucketpolicy-bucket
	//
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// A policy document containing permissions to add to the specified bucket.
	//
	// In IAM, you must provide policy documents in JSON format. However, in CloudFormation you can provide the policy in JSON or YAML format because CloudFormation converts YAML to JSON before submitting it to IAM. For more information, see the AWS::IAM::Policy [PolicyDocument](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-policy.html#cfn-iam-policy-policydocument) resource description in this guide and [Access Policy Language Overview](https://docs.aws.amazon.com/AmazonS3/latest/dev/access-policy-language-overview.html) in the *Amazon S3 User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-bucketpolicy.html#cfn-s3-bucketpolicy-policydocument
	//
	PolicyDocument interface{} `field:"required" json:"policyDocument" yaml:"policyDocument"`
}

