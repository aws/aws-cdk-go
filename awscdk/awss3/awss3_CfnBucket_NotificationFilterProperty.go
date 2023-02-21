package awss3


// Specifies object key name filtering rules.
//
// For information about key name filtering, see [Configuring Event Notifications](https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html) in the *Amazon S3 User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   notificationFilterProperty := &notificationFilterProperty{
//   	s3Key: &s3KeyFilterProperty{
//   		rules: []interface{}{
//   			&filterRuleProperty{
//   				name: jsii.String("name"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   	},
//   }
//
type CfnBucket_NotificationFilterProperty struct {
	// A container for object key name prefix and suffix filtering rules.
	S3Key interface{} `field:"required" json:"s3Key" yaml:"s3Key"`
}

