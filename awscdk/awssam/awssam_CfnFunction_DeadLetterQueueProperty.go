package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deadLetterQueueProperty := &deadLetterQueueProperty{
//   	targetArn: jsii.String("targetArn"),
//   	type: jsii.String("type"),
//   }
//
type CfnFunction_DeadLetterQueueProperty struct {
	// `CfnFunction.DeadLetterQueueProperty.TargetArn`.
	TargetArn *string `field:"required" json:"targetArn" yaml:"targetArn"`
	// `CfnFunction.DeadLetterQueueProperty.Type`.
	Type *string `field:"required" json:"type" yaml:"type"`
}

