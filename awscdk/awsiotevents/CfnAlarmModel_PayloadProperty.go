package awsiotevents


// Information needed to configure the payload.
//
// By default, AWS IoT Events generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the detector model instance and the event triggered the action. To configure the action payload, you can use `contentExpression` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   payloadProperty := &PayloadProperty{
//   	ContentExpression: jsii.String("contentExpression"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-payload.html
//
type CfnAlarmModel_PayloadProperty struct {
	// The content of the payload.
	//
	// You can use a string expression that includes quoted strings ( `'<string>'` ), variables ( `$variable.<variable-name>` ), input values ( `$input.<input-name>.<path-to-datum>` ), string concatenations, and quoted strings that contain `${}` as the content. The recommended maximum size of a content expression is 1 KB.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-payload.html#cfn-iotevents-alarmmodel-payload-contentexpression
	//
	ContentExpression *string `field:"required" json:"contentExpression" yaml:"contentExpression"`
	// The value of the payload type can be either `STRING` or `JSON` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-payload.html#cfn-iotevents-alarmmodel-payload-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
}

