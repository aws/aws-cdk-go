package awsapigateway


// API stage name of the associated API stage in a usage plan.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   apiStageProperty := &ApiStageProperty{
//   	ApiId: jsii.String("apiId"),
//   	Stage: jsii.String("stage"),
//   	Throttle: map[string]interface{}{
//   		"throttleKey": &ThrottleSettingsProperty{
//   			"burstLimit": jsii.Number(123),
//   			"rateLimit": jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-apistage.html
//
type CfnUsagePlan_ApiStageProperty struct {
	// API Id of the associated API stage in a usage plan.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-apistage.html#cfn-apigateway-usageplan-apistage-apiid
	//
	ApiId *string `field:"optional" json:"apiId" yaml:"apiId"`
	// API stage name of the associated API stage in a usage plan.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-apistage.html#cfn-apigateway-usageplan-apistage-stage
	//
	Stage *string `field:"optional" json:"stage" yaml:"stage"`
	// Map containing method level throttling information for API stage in a usage plan.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-apistage.html#cfn-apigateway-usageplan-apistage-throttle
	//
	Throttle interface{} `field:"optional" json:"throttle" yaml:"throttle"`
}

