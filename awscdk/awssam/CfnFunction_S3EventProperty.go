package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3EventProperty := &S3EventProperty{
//   	Bucket: jsii.String("bucket"),
//   	Events: jsii.String("events"),
//
//   	// the properties below are optional
//   	Filter: &S3NotificationFilterProperty{
//   		S3Key: &S3KeyFilterProperty{
//   			Rules: []interface{}{
//   				&S3KeyFilterRuleProperty{
//   					Name: jsii.String("name"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-s3event.html
//
type CfnFunction_S3EventProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-s3event.html#cfn-serverless-function-s3event-bucket
	//
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-s3event.html#cfn-serverless-function-s3event-events
	//
	Events *string `field:"required" json:"events" yaml:"events"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-s3event.html#cfn-serverless-function-s3event-filter
	//
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
}

