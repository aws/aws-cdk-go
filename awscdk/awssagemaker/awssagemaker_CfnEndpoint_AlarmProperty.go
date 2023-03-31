package awssagemaker


// An Amazon CloudWatch alarm configured to monitor metrics on an endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   alarmProperty := &alarmProperty{
//   	alarmName: jsii.String("alarmName"),
//   }
//
type CfnEndpoint_AlarmProperty struct {
	// The name of a CloudWatch alarm in your account.
	AlarmName *string `field:"required" json:"alarmName" yaml:"alarmName"`
}

