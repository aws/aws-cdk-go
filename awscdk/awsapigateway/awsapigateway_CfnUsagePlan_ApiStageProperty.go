package awsapigateway


// API stage name of the associated API stage in a usage plan.
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
	// API Id of the associated API stage in a usage plan.
	ApiId *string `field:"optional" json:"apiId" yaml:"apiId"`
	// API stage name of the associated API stage in a usage plan.
	Stage *string `field:"optional" json:"stage" yaml:"stage"`
	// Map containing method level throttling information for API stage in a usage plan.
	Throttle interface{} `field:"optional" json:"throttle" yaml:"throttle"`
}

