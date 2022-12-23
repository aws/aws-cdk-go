package awss3outposts


// A container for an Amazon S3 on Outposts bucket lifecycle rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var filter interface{}
//
//   ruleProperty := &ruleProperty{
//   	status: jsii.String("status"),
//
//   	// the properties below are optional
//   	abortIncompleteMultipartUpload: &abortIncompleteMultipartUploadProperty{
//   		daysAfterInitiation: jsii.Number(123),
//   	},
//   	expirationDate: jsii.String("expirationDate"),
//   	expirationInDays: jsii.Number(123),
//   	filter: filter,
//   	id: jsii.String("id"),
//   }
//
type CfnBucket_RuleProperty struct {
	// If `Enabled` , the rule is currently being applied.
	//
	// If `Disabled` , the rule is not currently being applied.
	Status *string `field:"required" json:"status" yaml:"status"`
	// The container for the abort incomplete multipart upload rule.
	AbortIncompleteMultipartUpload interface{} `field:"optional" json:"abortIncompleteMultipartUpload" yaml:"abortIncompleteMultipartUpload"`
	// Specifies the expiration for the lifecycle of the object by specifying an expiry date.
	ExpirationDate *string `field:"optional" json:"expirationDate" yaml:"expirationDate"`
	// Specifies the expiration for the lifecycle of the object in the form of days that the object has been in the S3 on Outposts bucket.
	ExpirationInDays *float64 `field:"optional" json:"expirationInDays" yaml:"expirationInDays"`
	// The container for the filter of the lifecycle rule.
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
	// The unique identifier for the lifecycle rule.
	//
	// The value can't be longer than 255 characters.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

