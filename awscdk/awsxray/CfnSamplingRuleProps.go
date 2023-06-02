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
//   	SamplingRule: &SamplingRuleProperty{
//   		FixedRate: jsii.Number(123),
//   		Host: jsii.String("host"),
//   		HttpMethod: jsii.String("httpMethod"),
//   		Priority: jsii.Number(123),
//   		ReservoirSize: jsii.Number(123),
//   		ResourceArn: jsii.String("resourceArn"),
//   		ServiceName: jsii.String("serviceName"),
//   		ServiceType: jsii.String("serviceType"),
//   		UrlPath: jsii.String("urlPath"),
//
//   		// the properties below are optional
//   		Attributes: map[string]*string{
//   			"attributesKey": jsii.String("attributes"),
//   		},
//   		RuleArn: jsii.String("ruleArn"),
//   		RuleName: jsii.String("ruleName"),
//   		Version: jsii.Number(123),
//   	},
//   	Tags: []interface{}{
//   		tags,
//   	},
//   }
//
type CfnSamplingRuleProps struct {
	// The sampling rule to be created or updated.
	SamplingRule interface{} `field:"optional" json:"samplingRule" yaml:"samplingRule"`
	// An array of key-value pairs to apply to this resource.
	Tags *[]interface{} `field:"optional" json:"tags" yaml:"tags"`
}

