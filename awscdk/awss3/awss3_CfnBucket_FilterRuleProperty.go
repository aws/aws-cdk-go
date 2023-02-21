package awss3


// Specifies the Amazon S3 object key name to filter on and whether to filter on the suffix or prefix of the key name.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterRuleProperty := &filterRuleProperty{
//   	name: jsii.String("name"),
//   	value: jsii.String("value"),
//   }
//
type CfnBucket_FilterRuleProperty struct {
	// The object key name prefix or suffix identifying one or more objects to which the filtering rule applies.
	//
	// The maximum length is 1,024 characters. Overlapping prefixes and suffixes are not supported. For more information, see [Configuring Event Notifications](https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html) in the *Amazon S3 User Guide* .
	Name *string `field:"required" json:"name" yaml:"name"`
	// The value that the filter searches for in object key names.
	Value *string `field:"required" json:"value" yaml:"value"`
}

