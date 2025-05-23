package awsiot


// The value to use for the entry's timestamp.
//
// If blank, the time that the entry was processed is used.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   timestreamTimestampProperty := &TimestreamTimestampProperty{
//   	Unit: jsii.String("unit"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-timestreamtimestamp.html
//
type CfnTopicRule_TimestreamTimestampProperty struct {
	// The precision of the timestamp value that results from the expression described in `value` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-timestreamtimestamp.html#cfn-iot-topicrule-timestreamtimestamp-unit
	//
	Unit *string `field:"required" json:"unit" yaml:"unit"`
	// An expression that returns a long epoch time value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-timestreamtimestamp.html#cfn-iot-topicrule-timestreamtimestamp-value
	//
	Value *string `field:"required" json:"value" yaml:"value"`
}

