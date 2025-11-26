package previewawss3outpostsmixins


// A container for an Amazon S3 on Outposts bucket lifecycle rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var filter interface{}
//
//   ruleProperty := &RuleProperty{
//   	AbortIncompleteMultipartUpload: &AbortIncompleteMultipartUploadProperty{
//   		DaysAfterInitiation: jsii.Number(123),
//   	},
//   	ExpirationDate: jsii.String("expirationDate"),
//   	ExpirationInDays: jsii.Number(123),
//   	Filter: filter,
//   	Id: jsii.String("id"),
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-bucket-rule.html
//
type CfnBucketPropsMixin_RuleProperty struct {
	// The container for the abort incomplete multipart upload rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-bucket-rule.html#cfn-s3outposts-bucket-rule-abortincompletemultipartupload
	//
	AbortIncompleteMultipartUpload interface{} `field:"optional" json:"abortIncompleteMultipartUpload" yaml:"abortIncompleteMultipartUpload"`
	// Specifies the expiration for the lifecycle of the object by specifying an expiry date.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-bucket-rule.html#cfn-s3outposts-bucket-rule-expirationdate
	//
	ExpirationDate *string `field:"optional" json:"expirationDate" yaml:"expirationDate"`
	// Specifies the expiration for the lifecycle of the object in the form of days that the object has been in the S3 on Outposts bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-bucket-rule.html#cfn-s3outposts-bucket-rule-expirationindays
	//
	ExpirationInDays *float64 `field:"optional" json:"expirationInDays" yaml:"expirationInDays"`
	// The container for the filter of the lifecycle rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-bucket-rule.html#cfn-s3outposts-bucket-rule-filter
	//
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
	// Unique identifier for the lifecycle rule.
	//
	// The value can't be longer than 255 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-bucket-rule.html#cfn-s3outposts-bucket-rule-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
	// If `Enabled` , the rule is currently being applied.
	//
	// If `Disabled` , the rule is not currently being applied.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-bucket-rule.html#cfn-s3outposts-bucket-rule-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

