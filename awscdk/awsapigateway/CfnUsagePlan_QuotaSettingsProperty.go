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
//   quotaSettingsProperty := &QuotaSettingsProperty{
//   	Limit: jsii.Number(123),
//   	Offset: jsii.Number(123),
//   	Period: jsii.String("period"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-quotasettings.html
//
type CfnUsagePlan_QuotaSettingsProperty struct {
	// The target maximum number of requests that can be made in a given time period.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-quotasettings.html#cfn-apigateway-usageplan-quotasettings-limit
	//
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// The number of requests subtracted from the given limit in the initial time period.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-quotasettings.html#cfn-apigateway-usageplan-quotasettings-offset
	//
	Offset *float64 `field:"optional" json:"offset" yaml:"offset"`
	// The time period in which the limit applies.
	//
	// Valid values are "DAY", "WEEK" or "MONTH".
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-quotasettings.html#cfn-apigateway-usageplan-quotasettings-period
	//
	Period *string `field:"optional" json:"period" yaml:"period"`
}

