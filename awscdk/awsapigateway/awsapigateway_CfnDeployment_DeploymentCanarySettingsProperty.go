package awsapigateway


// The `DeploymentCanarySettings` property type specifies settings for the canary deployment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deploymentCanarySettingsProperty := &deploymentCanarySettingsProperty{
//   	percentTraffic: jsii.Number(123),
//   	stageVariableOverrides: map[string]*string{
//   		"stageVariableOverridesKey": jsii.String("stageVariableOverrides"),
//   	},
//   	useStageCache: jsii.Boolean(false),
//   }
//
type CfnDeployment_DeploymentCanarySettingsProperty struct {
	// The percentage (0.0-100.0) of traffic routed to the canary deployment.
	PercentTraffic *float64 `field:"optional" json:"percentTraffic" yaml:"percentTraffic"`
	// A stage variable overrides used for the canary release deployment.
	//
	// They can override existing stage variables or add new stage variables for the canary release deployment. These stage variables are represented as a string-to-string map between stage variable names and their values.
	StageVariableOverrides interface{} `field:"optional" json:"stageVariableOverrides" yaml:"stageVariableOverrides"`
	// A Boolean flag to indicate whether the canary release deployment uses the stage cache or not.
	UseStageCache interface{} `field:"optional" json:"useStageCache" yaml:"useStageCache"`
}

