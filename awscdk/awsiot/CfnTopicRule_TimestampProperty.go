package awsiot


// Describes how to interpret an application-defined timestamp value from an MQTT message payload and the precision of that value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   timestampProperty := &TimestampProperty{
//   	Value: jsii.String("value"),
//
//   	// the properties below are optional
//   	Unit: jsii.String("unit"),
//   }
//
type CfnTopicRule_TimestampProperty struct {
	// An expression that returns a long epoch time value.
	Value *string `field:"required" json:"value" yaml:"value"`
	// The precision of the timestamp value that results from the expression described in `value` .
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

