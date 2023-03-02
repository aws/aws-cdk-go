package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   queueSAMPTProperty := &queueSAMPTProperty{
//   	queueName: jsii.String("queueName"),
//   }
//
type CfnFunction_QueueSAMPTProperty struct {
	// `CfnFunction.QueueSAMPTProperty.QueueName`.
	QueueName *string `field:"required" json:"queueName" yaml:"queueName"`
}

