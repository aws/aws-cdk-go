package awss3


// Specifies object key name filtering rules.
//
// For information about key name filtering, see [Configuring event notifications using object key name filtering](https://docs.aws.amazon.com/AmazonS3/latest/userguide/notification-how-to-filtering.html) in the *Amazon S3 User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   notificationFilterProperty := &NotificationFilterProperty{
//   	S3Key: &S3KeyFilterProperty{
//   		Rules: []interface{}{
//   			&FilterRuleProperty{
//   				Name: jsii.String("name"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   }
//
type CfnBucket_NotificationFilterProperty struct {
	// A container for object key name prefix and suffix filtering rules.
	S3Key interface{} `field:"required" json:"s3Key" yaml:"s3Key"`
}

