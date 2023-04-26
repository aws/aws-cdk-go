package awsxray


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   samplingRuleUpdateProperty := &SamplingRuleUpdateProperty{
//   	Attributes: map[string]*string{
//   		"attributesKey": jsii.String("attributes"),
//   	},
//   	FixedRate: jsii.Number(123),
//   	Host: jsii.String("host"),
//   	HttpMethod: jsii.String("httpMethod"),
//   	Priority: jsii.Number(123),
//   	ReservoirSize: jsii.Number(123),
//   	ResourceArn: jsii.String("resourceArn"),
//   	RuleArn: jsii.String("ruleArn"),
//   	RuleName: jsii.String("ruleName"),
//   	ServiceName: jsii.String("serviceName"),
//   	ServiceType: jsii.String("serviceType"),
//   	UrlPath: jsii.String("urlPath"),
//   }
//
type CfnSamplingRule_SamplingRuleUpdateProperty struct {
	// `CfnSamplingRule.SamplingRuleUpdateProperty.Attributes`.
	Attributes interface{} `field:"optional" json:"attributes" yaml:"attributes"`
	// `CfnSamplingRule.SamplingRuleUpdateProperty.FixedRate`.
	FixedRate *float64 `field:"optional" json:"fixedRate" yaml:"fixedRate"`
	// `CfnSamplingRule.SamplingRuleUpdateProperty.Host`.
	Host *string `field:"optional" json:"host" yaml:"host"`
	// `CfnSamplingRule.SamplingRuleUpdateProperty.HTTPMethod`.
	HttpMethod *string `field:"optional" json:"httpMethod" yaml:"httpMethod"`
	// `CfnSamplingRule.SamplingRuleUpdateProperty.Priority`.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// `CfnSamplingRule.SamplingRuleUpdateProperty.ReservoirSize`.
	ReservoirSize *float64 `field:"optional" json:"reservoirSize" yaml:"reservoirSize"`
	// `CfnSamplingRule.SamplingRuleUpdateProperty.ResourceARN`.
	ResourceArn *string `field:"optional" json:"resourceArn" yaml:"resourceArn"`
	// `CfnSamplingRule.SamplingRuleUpdateProperty.RuleARN`.
	RuleArn *string `field:"optional" json:"ruleArn" yaml:"ruleArn"`
	// `CfnSamplingRule.SamplingRuleUpdateProperty.RuleName`.
	RuleName *string `field:"optional" json:"ruleName" yaml:"ruleName"`
	// `CfnSamplingRule.SamplingRuleUpdateProperty.ServiceName`.
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
	// `CfnSamplingRule.SamplingRuleUpdateProperty.ServiceType`.
	ServiceType *string `field:"optional" json:"serviceType" yaml:"serviceType"`
	// `CfnSamplingRule.SamplingRuleUpdateProperty.URLPath`.
	UrlPath *string `field:"optional" json:"urlPath" yaml:"urlPath"`
}

