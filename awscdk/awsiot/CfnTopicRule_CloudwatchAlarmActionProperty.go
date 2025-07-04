package awsiot


// Describes an action that updates a CloudWatch alarm.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudwatchAlarmActionProperty := &CloudwatchAlarmActionProperty{
//   	AlarmName: jsii.String("alarmName"),
//   	RoleArn: jsii.String("roleArn"),
//   	StateReason: jsii.String("stateReason"),
//   	StateValue: jsii.String("stateValue"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-cloudwatchalarmaction.html
//
type CfnTopicRule_CloudwatchAlarmActionProperty struct {
	// The CloudWatch alarm name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-cloudwatchalarmaction.html#cfn-iot-topicrule-cloudwatchalarmaction-alarmname
	//
	AlarmName *string `field:"required" json:"alarmName" yaml:"alarmName"`
	// The IAM role that allows access to the CloudWatch alarm.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-cloudwatchalarmaction.html#cfn-iot-topicrule-cloudwatchalarmaction-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The reason for the alarm change.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-cloudwatchalarmaction.html#cfn-iot-topicrule-cloudwatchalarmaction-statereason
	//
	StateReason *string `field:"required" json:"stateReason" yaml:"stateReason"`
	// The value of the alarm state.
	//
	// Acceptable values are: OK, ALARM, INSUFFICIENT_DATA.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-cloudwatchalarmaction.html#cfn-iot-topicrule-cloudwatchalarmaction-statevalue
	//
	StateValue *string `field:"required" json:"stateValue" yaml:"stateValue"`
}

