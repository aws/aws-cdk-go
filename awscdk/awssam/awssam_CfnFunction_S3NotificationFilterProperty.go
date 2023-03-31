package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3NotificationFilterProperty := &s3NotificationFilterProperty{
//   	s3Key: &s3KeyFilterProperty{
//   		rules: []interface{}{
//   			&s3KeyFilterRuleProperty{
//   				name: jsii.String("name"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   	},
//   }
//
type CfnFunction_S3NotificationFilterProperty struct {
	// `CfnFunction.S3NotificationFilterProperty.S3Key`.
	S3Key interface{} `field:"required" json:"s3Key" yaml:"s3Key"`
}

