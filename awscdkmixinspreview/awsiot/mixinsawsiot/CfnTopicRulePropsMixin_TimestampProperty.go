package mixinsawsiot


// Describes how to interpret an application-defined timestamp value from an MQTT message payload and the precision of that value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   timestampProperty := &TimestampProperty{
//   	Unit: jsii.String("unit"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-timestamp.html
//
type CfnTopicRulePropsMixin_TimestampProperty struct {
	// The precision of the timestamp value that results from the expression described in `value` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-timestamp.html#cfn-iot-topicrule-timestamp-unit
	//
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
	// An expression that returns a long epoch time value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-timestamp.html#cfn-iot-topicrule-timestamp-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

