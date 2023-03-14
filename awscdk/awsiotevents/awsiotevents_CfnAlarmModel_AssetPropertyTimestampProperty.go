package awsiotevents


// A structure that contains timestamp information. For more information, see [TimeInNanos](https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_TimeInNanos.html) in the *AWS IoT SiteWise API Reference* .
//
// You must use expressions for all parameters in `AssetPropertyTimestamp` . The expressions accept literals, operators, functions, references, and substitution templates.
//
// **Examples** - For literal values, the expressions must contain single quotes. For example, the value for the `timeInSeconds` parameter can be `'1586400675'` .
// - For references, you must specify either variables or input values. For example, the value for the `offsetInNanos` parameter can be `$variable.time` .
// - For a substitution template, you must use `${}` , and the template must be in single quotes. A substitution template can also contain a combination of literals, operators, functions, references, and substitution templates.
//
// In the following example, the value for the `timeInSeconds` parameter uses a substitution template.
//
// `'${$input.TemperatureInput.sensorData.timestamp / 1000}'`
//
// For more information, see [Expressions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-expressions.html) in the *AWS IoT Events Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assetPropertyTimestampProperty := &assetPropertyTimestampProperty{
//   	timeInSeconds: jsii.String("timeInSeconds"),
//
//   	// the properties below are optional
//   	offsetInNanos: jsii.String("offsetInNanos"),
//   }
//
type CfnAlarmModel_AssetPropertyTimestampProperty struct {
	// The timestamp, in seconds, in the Unix epoch format.
	//
	// The valid range is between 1-31556889864403199.
	TimeInSeconds *string `field:"required" json:"timeInSeconds" yaml:"timeInSeconds"`
	// The nanosecond offset converted from `timeInSeconds` .
	//
	// The valid range is between 0-999999999.
	OffsetInNanos *string `field:"optional" json:"offsetInNanos" yaml:"offsetInNanos"`
}

