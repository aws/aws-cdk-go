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
	// The percentage (0-100) of traffic diverted to a canary deployment.
	PercentTraffic *float64 `field:"optional" json:"percentTraffic" yaml:"percentTraffic"`
	// Stage variables overridden for a canary release deployment, including new stage variables introduced in the canary.
	//
	// These stage variables are represented as a string-to-string map between stage variable names and their values.
	//
	// Duplicates are not allowed.
	StageVariableOverrides interface{} `field:"optional" json:"stageVariableOverrides" yaml:"stageVariableOverrides"`
	// Whether the canary deployment uses the stage cache.
	UseStageCache interface{} `field:"optional" json:"useStageCache" yaml:"useStageCache"`
}

