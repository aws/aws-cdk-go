package awss3


// Specifies the Amazon S3 object key name to filter on and whether to filter on the suffix or prefix of the key name.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterRuleProperty := &FilterRuleProperty{
//   	Name: jsii.String("name"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-filterrule.html
//
type CfnBucket_FilterRuleProperty struct {
	// The object key name prefix or suffix identifying one or more objects to which the filtering rule applies.
	//
	// The maximum length is 1,024 characters. Overlapping prefixes and suffixes are not supported. For more information, see [Configuring Event Notifications](https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html) in the *Amazon S3 User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-filterrule.html#cfn-s3-bucket-filterrule-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The value that the filter searches for in object key names.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-filterrule.html#cfn-s3-bucket-filterrule-value
	//
	Value *string `field:"required" json:"value" yaml:"value"`
}

