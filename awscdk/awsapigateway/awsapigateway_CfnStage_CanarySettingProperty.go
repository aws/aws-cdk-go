package awsapigateway


// Configuration settings of a canary deployment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   canarySettingProperty := &canarySettingProperty{
//   	deploymentId: jsii.String("deploymentId"),
//   	percentTraffic: jsii.Number(123),
//   	stageVariableOverrides: map[string]*string{
//   		"stageVariableOverridesKey": jsii.String("stageVariableOverrides"),
//   	},
//   	useStageCache: jsii.Boolean(false),
//   }
//
type CfnStage_CanarySettingProperty struct {
	// The ID of the canary deployment.
	DeploymentId *string `field:"optional" json:"deploymentId" yaml:"deploymentId"`
	// The percent (0-100) of traffic diverted to a canary deployment.
	PercentTraffic *float64 `field:"optional" json:"percentTraffic" yaml:"percentTraffic"`
	// Stage variables overridden for a canary release deployment, including new stage variables introduced in the canary.
	//
	// These stage variables are represented as a string-to-string map between stage variable names and their values.
	StageVariableOverrides interface{} `field:"optional" json:"stageVariableOverrides" yaml:"stageVariableOverrides"`
	// A Boolean flag to indicate whether the canary deployment uses the stage cache or not.
	UseStageCache interface{} `field:"optional" json:"useStageCache" yaml:"useStageCache"`
}

