package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3NotificationFilterProperty := &S3NotificationFilterProperty{
//   	S3Key: &S3KeyFilterProperty{
//   		Rules: []interface{}{
//   			&S3KeyFilterRuleProperty{
//   				Name: jsii.String("name"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-s3notificationfilter.html
//
type CfnFunction_S3NotificationFilterProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-s3notificationfilter.html#cfn-serverless-function-s3notificationfilter-s3key
	//
	S3Key interface{} `field:"required" json:"s3Key" yaml:"s3Key"`
}

