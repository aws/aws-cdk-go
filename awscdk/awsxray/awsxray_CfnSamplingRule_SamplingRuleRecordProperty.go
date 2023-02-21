package awsxray


// A [SamplingRule](https://docs.aws.amazon.com//xray/latest/api/API_SamplingRule.html) and its metadata.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   samplingRuleRecordProperty := &samplingRuleRecordProperty{
//   	createdAt: jsii.String("createdAt"),
//   	modifiedAt: jsii.String("modifiedAt"),
//   	samplingRule: &samplingRuleProperty{
//   		attributes: map[string]*string{
//   			"attributesKey": jsii.String("attributes"),
//   		},
//   		fixedRate: jsii.Number(123),
//   		host: jsii.String("host"),
//   		httpMethod: jsii.String("httpMethod"),
//   		priority: jsii.Number(123),
//   		reservoirSize: jsii.Number(123),
//   		resourceArn: jsii.String("resourceArn"),
//   		ruleArn: jsii.String("ruleArn"),
//   		ruleName: jsii.String("ruleName"),
//   		serviceName: jsii.String("serviceName"),
//   		serviceType: jsii.String("serviceType"),
//   		urlPath: jsii.String("urlPath"),
//   		version: jsii.Number(123),
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

