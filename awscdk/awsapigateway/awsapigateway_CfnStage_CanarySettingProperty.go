package awsapigateway


// The `CanarySetting` property type specifies settings for the canary deployment in this stage.
//
// `CanarySetting` is a property of the [AWS::ApiGateway::Stage](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html) resource.
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
	// The identifier of the deployment that the stage points to.
	DeploymentId *string `field:"optional" json:"deploymentId" yaml:"deploymentId"`
	// The percentage (0-100) of traffic diverted to a canary deployment.
	PercentTraffic *float64 `field:"optional" json:"percentTraffic" yaml:"percentTraffic"`
	// Stage variables overridden for a canary release deployment, including new stage variables introduced in the canary.
	//
	// These stage variables are represented as a string-to-string map between stage variable names and their values.
	//
	// Duplicates are not allowed.
	StageVariableOverrides interface{} `field:"optional" json:"stageVariableOverrides" yaml:"stageVariableOverrides"`
	// Whether the canary deployment uses the stage cache or not.
	UseStageCache interface{} `field:"optional" json:"useStageCache" yaml:"useStageCache"`
}

