package awss3


// A container for object key name prefix and suffix filtering rules.
//
// For more information about object key name filtering, see [Configuring event notifications using object key name filtering](https://docs.aws.amazon.com/AmazonS3/latest/userguide/notification-how-to-filtering.html) in the *Amazon S3 User Guide* .
//
// > The same type of filter rule cannot be used more than once. For example, you cannot specify two prefix rules.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3KeyFilterProperty := &s3KeyFilterProperty{
//   	rules: []interface{}{
//   		&filterRuleProperty{
//   			name: jsii.String("name"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnBucket_S3KeyFilterProperty struct {
	// A list of containers for the key-value pair that defines the criteria for the filter rule.
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
}

