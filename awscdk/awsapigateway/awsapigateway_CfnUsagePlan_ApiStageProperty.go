package awsapigateway


// `ApiStage` is a property of the [AWS::ApiGateway::UsagePlan](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-usageplan.html) resource that specifies which stages and APIs to associate with a usage plan.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   apiStageProperty := &apiStageProperty{
//   	apiId: jsii.String("apiId"),
//   	stage: jsii.String("stage"),
//   	throttle: map[string]interface{}{
//   		"throttleKey": &ThrottleSettingsProperty{
//   			"burstLimit": jsii.Number(123),
//   			"rateLimit": jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnUsagePlan_ApiStageProperty struct {
	// The ID of an API that is in the specified `Stage` property that you want to associate with the usage plan.
	ApiId *string `field:"optional" json:"apiId" yaml:"apiId"`
	// The name of the stage to associate with the usage plan.
	Stage *string `field:"optional" json:"stage" yaml:"stage"`
	// Map containing method-level throttling information for an API stage in a usage plan.
	//
	// The key for the map is the path and method for which to configure custom throttling, for example, "/pets/GET".
	//
	// Duplicates are not allowed.
	Throttle interface{} `field:"optional" json:"throttle" yaml:"throttle"`
}

