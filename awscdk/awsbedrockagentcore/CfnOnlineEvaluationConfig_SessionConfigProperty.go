package awsbedrockagentcore


// The configuration that defines how agent sessions are detected.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sessionConfigProperty := &SessionConfigProperty{
//   	SessionTimeoutMinutes: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-onlineevaluationconfig-sessionconfig.html
//
type CfnOnlineEvaluationConfig_SessionConfigProperty struct {
	// The number of minutes of inactivity after which an agent session is considered complete.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-onlineevaluationconfig-sessionconfig.html#cfn-bedrockagentcore-onlineevaluationconfig-sessionconfig-sessiontimeoutminutes
	//
	SessionTimeoutMinutes *float64 `field:"required" json:"sessionTimeoutMinutes" yaml:"sessionTimeoutMinutes"`
}

