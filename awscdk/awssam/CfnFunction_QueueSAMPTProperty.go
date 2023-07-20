package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   queueSAMPTProperty := &QueueSAMPTProperty{
//   	QueueName: jsii.String("queueName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-queuesampt.html
//
type CfnFunction_QueueSAMPTProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-queuesampt.html#cfn-serverless-function-queuesampt-queuename
	//
	QueueName *string `field:"required" json:"queueName" yaml:"queueName"`
}

