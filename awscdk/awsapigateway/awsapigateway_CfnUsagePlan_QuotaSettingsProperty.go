package awsapigateway


// `QuotaSettings` is a property of the [AWS::ApiGateway::UsagePlan](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-usageplan.html) resource that specifies a target for the maximum number of requests users can make to your REST APIs.
//
// In some cases clients can exceed the targets that you set. Don’t rely on usage plans to control costs. Consider using [AWS Budgets](https://docs.aws.amazon.com/cost-management/latest/userguide/budgets-managing-costs.html) to monitor costs and [AWS WAF](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) to manage API requests.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   quotaSettingsProperty := &quotaSettingsProperty{
//   	limit: jsii.Number(123),
//   	offset: jsii.Number(123),
//   	period: jsii.String("period"),
//   }
//
type CfnUsagePlan_QuotaSettingsProperty struct {
	// The target maximum number of requests that can be made in a given time period.
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// The day that a time period starts.
	//
	// For example, with a time period of `WEEK` , an offset of `0` starts on Sunday, and an offset of `1` starts on Monday.
	Offset *float64 `field:"optional" json:"offset" yaml:"offset"`
	// The time period for which the target maximum limit of requests applies, such as `DAY` or `WEEK` .
	//
	// For valid values, see the period property for the [UsagePlan](https://docs.aws.amazon.com/apigateway/api-reference/resource/usage-plan) resource in the *Amazon API Gateway REST API Reference* .
	Period *string `field:"optional" json:"period" yaml:"period"`
}

