package awsxray


// Properties for defining a `CfnSamplingRule`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnSamplingRuleProps := &CfnSamplingRuleProps{
//   	RuleName: jsii.String("ruleName"),
//   	SamplingRule: &SamplingRuleProperty{
//   		Attributes: map[string]*string{
//   			"attributesKey": jsii.String("attributes"),
//   		},
//   		FixedRate: jsii.Number(123),
//   		Host: jsii.String("host"),
//   		HttpMethod: jsii.String("httpMethod"),
//   		Priority: jsii.Number(123),
//   		ReservoirSize: jsii.Number(123),
//   		ResourceArn: jsii.String("resourceArn"),
//   		RuleArn: jsii.String("ruleArn"),
//   		RuleName: jsii.String("ruleName"),
//   		ServiceName: jsii.String("serviceName"),
//   		ServiceType: jsii.String("serviceType"),
//   		UrlPath: jsii.String("urlPath"),
//   		Version: jsii.Number(123),
//   	},
//   	SamplingRuleRecord: &SamplingRuleRecordProperty{
//   		CreatedAt: jsii.String("createdAt"),
//   		ModifiedAt: jsii.String("modifiedAt"),
//   		SamplingRule: &SamplingRuleProperty{
//   			Attributes: map[string]*string{
//   				"attributesKey": jsii.String("attributes"),
//   			},
//   			FixedRate: jsii.Number(123),
//   			Host: jsii.String("host"),
//   			HttpMethod: jsii.String("httpMethod"),
//   			Priority: jsii.Number(123),
//   			ReservoirSize: jsii.Number(123),
//   			ResourceArn: jsii.String("resourceArn"),
//   			RuleArn: jsii.String("ruleArn"),
//   			RuleName: jsii.String("ruleName"),
//   			ServiceName: jsii.String("serviceName"),
//   			ServiceType: jsii.String("serviceType"),
//   			UrlPath: jsii.String("urlPath"),
//   			Version: jsii.Number(123),
//   		},
//   	},
//   	SamplingRuleUpdate: &SamplingRuleUpdateProperty{
//   		Attributes: map[string]*string{
//   			"attributesKey": jsii.String("attributes"),
//   		},
//   		FixedRate: jsii.Number(123),
//   		Host: jsii.String("host"),
//   		HttpMethod: jsii.String("httpMethod"),
//   		Priority: jsii.Number(123),
//   		ReservoirSize: jsii.Number(123),
//   		ResourceArn: jsii.String("resourceArn"),
//   		RuleArn: jsii.String("ruleArn"),
//   		RuleName: jsii.String("ruleName"),
//   		ServiceName: jsii.String("serviceName"),
//   		ServiceType: jsii.String("serviceType"),
//   		UrlPath: jsii.String("urlPath"),
//   	},
//   	Tags: []interface{}{
//   		tags,
//   	},
//   }
//
type CfnSamplingRuleProps struct {
	// The name of the sampling rule.
	//
	// Specify a rule by either name or ARN, but not both. Used only when deleting a sampling rule. When creating or updating a sampling rule, use the `RuleName` or `RuleARN` properties within `SamplingRule` or `SamplingRuleUpdate` .
	RuleName *string `field:"optional" json:"ruleName" yaml:"ruleName"`
	// The sampling rule to be created.
	//
	// Must be provided if creating a new sampling rule. Not valid when updating an existing sampling rule.
	SamplingRule interface{} `field:"optional" json:"samplingRule" yaml:"samplingRule"`
	// `AWS::XRay::SamplingRule.SamplingRuleRecord`.
	SamplingRuleRecord interface{} `field:"optional" json:"samplingRuleRecord" yaml:"samplingRuleRecord"`
	// `AWS::XRay::SamplingRule.SamplingRuleUpdate`.
	SamplingRuleUpdate interface{} `field:"optional" json:"samplingRuleUpdate" yaml:"samplingRuleUpdate"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]interface{} `field:"optional" json:"tags" yaml:"tags"`
}

