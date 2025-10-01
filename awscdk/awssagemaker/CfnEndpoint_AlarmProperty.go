package awssagemaker


// An Amazon CloudWatch alarm configured to monitor metrics on an endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   alarmProperty := &AlarmProperty{
//   	AlarmName: jsii.String("alarmName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpoint-alarm.html
//
type CfnEndpoint_AlarmProperty struct {
	// The name of a CloudWatch alarm in your account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpoint-alarm.html#cfn-sagemaker-endpoint-alarm-alarmname
	//
	AlarmName *string `field:"required" json:"alarmName" yaml:"alarmName"`
}

