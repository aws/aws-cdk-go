package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sNSEventProperty := &sNSEventProperty{
//   	topic: jsii.String("topic"),
//   }
//
type CfnFunction_SNSEventProperty struct {
	// `CfnFunction.SNSEventProperty.Topic`.
	Topic *string `field:"required" json:"topic" yaml:"topic"`
}

