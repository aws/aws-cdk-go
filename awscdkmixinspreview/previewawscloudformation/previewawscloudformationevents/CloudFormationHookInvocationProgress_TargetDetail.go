package previewawscloudformationevents


// Type definition for Target-detail.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   targetDetail := &TargetDetail{
//   	TargetAction: []*string{
//   		jsii.String("targetAction"),
//   	},
//   	TargetId: []*string{
//   		jsii.String("targetId"),
//   	},
//   	TargetInvocationPoint: []*string{
//   		jsii.String("targetInvocationPoint"),
//   	},
//   	TargetName: []*string{
//   		jsii.String("targetName"),
//   	},
//   	TargetType: []*string{
//   		jsii.String("targetType"),
//   	},
//   }
//
// Experimental.
type CloudFormationHookInvocationProgress_TargetDetail struct {
	// target-action property.
	//
	// Specify an array of string values to match this event if the actual value of target-action is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TargetAction *[]*string `field:"optional" json:"targetAction" yaml:"targetAction"`
	// target-id property.
	//
	// Specify an array of string values to match this event if the actual value of target-id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TargetId *[]*string `field:"optional" json:"targetId" yaml:"targetId"`
	// target-invocation-point property.
	//
	// Specify an array of string values to match this event if the actual value of target-invocation-point is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TargetInvocationPoint *[]*string `field:"optional" json:"targetInvocationPoint" yaml:"targetInvocationPoint"`
	// target-name property.
	//
	// Specify an array of string values to match this event if the actual value of target-name is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TargetName *[]*string `field:"optional" json:"targetName" yaml:"targetName"`
	// target-type property.
	//
	// Specify an array of string values to match this event if the actual value of target-type is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TargetType *[]*string `field:"optional" json:"targetType" yaml:"targetType"`
}

