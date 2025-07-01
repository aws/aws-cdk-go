package awscloudwatchactions


// Properties for Lambda Alarm Action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaActionProps := &LambdaActionProps{
//   	UseUniquePermissionId: jsii.Boolean(false),
//   }
//
type LambdaActionProps struct {
	// Whether to generate unique Lambda Permission id.
	//
	// Use this parameter to resolve id collision in case of multiple alarms triggering the same action.
	// See: https://github.com/aws/aws-cdk/issues/33958
	//
	// Default: - false.
	//
	UseUniquePermissionId *bool `field:"optional" json:"useUniquePermissionId" yaml:"useUniquePermissionId"`
}

