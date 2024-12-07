package awss3express


// Specifies lifecycle rules for an Amazon S3 bucket.
//
// For more information, see [Put Bucket Lifecycle Configuration](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTlifecycle.html) in the *Amazon S3 API Reference* . For examples, see [Put Bucket Lifecycle Configuration Examples](https://docs.aws.amazon.com//AmazonS3/latest/API/API_PutBucketLifecycleConfiguration.html#API_PutBucketLifecycleConfiguration_Examples) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ruleProperty := &RuleProperty{
//   	Status: jsii.String("status"),
//
//   	// the properties below are optional
//   	AbortIncompleteMultipartUpload: &AbortIncompleteMultipartUploadProperty{
//   		DaysAfterInitiation: jsii.Number(123),
//   	},
//   	ExpirationInDays: jsii.Number(123),
//   	Id: jsii.String("id"),
//   	ObjectSizeGreaterThan: jsii.String("objectSizeGreaterThan"),
//   	ObjectSizeLessThan: jsii.String("objectSizeLessThan"),
//   	Prefix: jsii.String("prefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3express-directorybucket-rule.html
//
type CfnDirectoryBucket_RuleProperty struct {
	// If `Enabled` , the rule is currently being applied.
	//
	// If `Disabled` , the rule is not currently being applied.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3express-directorybucket-rule.html#cfn-s3express-directorybucket-rule-status
	//
	Status *string `field:"required" json:"status" yaml:"status"`
	// Specifies the days since the initiation of an incomplete multipart upload that Amazon S3 will wait before permanently removing all parts of the upload.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3express-directorybucket-rule.html#cfn-s3express-directorybucket-rule-abortincompletemultipartupload
	//
	AbortIncompleteMultipartUpload interface{} `field:"optional" json:"abortIncompleteMultipartUpload" yaml:"abortIncompleteMultipartUpload"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3express-directorybucket-rule.html#cfn-s3express-directorybucket-rule-expirationindays
	//
	ExpirationInDays *float64 `field:"optional" json:"expirationInDays" yaml:"expirationInDays"`
	// Unique identifier for the rule.
	//
	// The value can't be longer than 255 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3express-directorybucket-rule.html#cfn-s3express-directorybucket-rule-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3express-directorybucket-rule.html#cfn-s3express-directorybucket-rule-objectsizegreaterthan
	//
	ObjectSizeGreaterThan *string `field:"optional" json:"objectSizeGreaterThan" yaml:"objectSizeGreaterThan"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3express-directorybucket-rule.html#cfn-s3express-directorybucket-rule-objectsizelessthan
	//
	ObjectSizeLessThan *string `field:"optional" json:"objectSizeLessThan" yaml:"objectSizeLessThan"`
	// Object key prefix that identifies one or more objects to which this rule applies.
	//
	// > Replacement must be made for object keys containing special characters (such as carriage returns) when using XML requests. For more information, see [XML related object key constraints](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-keys.html#object-key-xml-related-constraints) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3express-directorybucket-rule.html#cfn-s3express-directorybucket-rule-prefix
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

