package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   restartPolicyProperty := &RestartPolicyProperty{
//   	Enabled: jsii.Boolean(false),
//   	IgnoredExitCodes: []interface{}{
//   		jsii.Number(123),
//   	},
//   	RestartAttemptPeriod: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-restartpolicy.html
//
type CfnDaemonTaskDefinitionPropsMixin_RestartPolicyProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-restartpolicy.html#cfn-ecs-daemontaskdefinition-restartpolicy-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-restartpolicy.html#cfn-ecs-daemontaskdefinition-restartpolicy-ignoredexitcodes
	//
	IgnoredExitCodes interface{} `field:"optional" json:"ignoredExitCodes" yaml:"ignoredExitCodes"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-restartpolicy.html#cfn-ecs-daemontaskdefinition-restartpolicy-restartattemptperiod
	//
	RestartAttemptPeriod *float64 `field:"optional" json:"restartAttemptPeriod" yaml:"restartAttemptPeriod"`
}

