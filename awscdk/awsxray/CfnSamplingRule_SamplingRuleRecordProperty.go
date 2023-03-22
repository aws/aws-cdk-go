package awsxray


// A [SamplingRule](https://docs.aws.amazon.com//xray/latest/api/API_SamplingRule.html) and its metadata.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   samplingRuleRecordProperty := &SamplingRuleRecordProperty{
//   	CreatedAt: jsii.String("createdAt"),
//   	ModifiedAt: jsii.String("modifiedAt"),
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
//   }
//
type CfnSamplingRule_SamplingRuleRecordProperty struct {
	// When the rule was created, in Unix time seconds.
	CreatedAt *string `field:"optional" json:"createdAt" yaml:"createdAt"`
	// When the rule was last modified, in Unix time seconds.
	ModifiedAt *string `field:"optional" json:"modifiedAt" yaml:"modifiedAt"`
	// The sampling rule.
	SamplingRule interface{} `field:"optional" json:"samplingRule" yaml:"samplingRule"`
}

