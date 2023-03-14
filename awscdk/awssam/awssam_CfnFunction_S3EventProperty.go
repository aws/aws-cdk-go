package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3EventProperty := &s3EventProperty{
//   	bucket: jsii.String("bucket"),
//   	events: jsii.String("events"),
//
//   	// the properties below are optional
//   	filter: &s3NotificationFilterProperty{
//   		s3Key: &s3KeyFilterProperty{
//   			rules: []interface{}{
//   				&s3KeyFilterRuleProperty{
//   					name: jsii.String("name"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnFunction_S3EventProperty struct {
	// `CfnFunction.S3EventProperty.Bucket`.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// `CfnFunction.S3EventProperty.Events`.
	Events interface{} `field:"required" json:"events" yaml:"events"`
	// `CfnFunction.S3EventProperty.Filter`.
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
}

