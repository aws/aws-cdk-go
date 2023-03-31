package awsxray


// A document specifying changes to a sampling rule's configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   samplingRuleUpdateProperty := &samplingRuleUpdateProperty{
//   	attributes: map[string]*string{
//   		"attributesKey": jsii.String("attributes"),
//   	},
//   	fixedRate: jsii.Number(123),
//   	host: jsii.String("host"),
//   	httpMethod: jsii.String("httpMethod"),
//   	priority: jsii.Number(123),
//   	reservoirSize: jsii.Number(123),
//   	resourceArn: jsii.String("resourceArn"),
//   	ruleArn: jsii.String("ruleArn"),
//   	ruleName: jsii.String("ruleName"),
//   	serviceName: jsii.String("serviceName"),
//   	serviceType: jsii.String("serviceType"),
//   	urlPath: jsii.String("urlPath"),
//   }
//
type CfnSamplingRule_SamplingRuleUpdateProperty struct {
	// Matches attributes derived from the request.
	//
	// *Map Entries:* Maximum number of 5 items.
	//
	// *Key Length Constraints:* Minimum length of 1. Maximum length of 32.
	//
	// *Value Length Constraints:* Minimum length of 1. Maximum length of 32.
	Attributes interface{} `field:"optional" json:"attributes" yaml:"attributes"`
	// The percentage of matching requests to instrument, after the reservoir is exhausted.
	FixedRate *float64 `field:"optional" json:"fixedRate" yaml:"fixedRate"`
	// Matches the hostname from a request URL.
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Matches the HTTP method of a request.
	HttpMethod *string `field:"optional" json:"httpMethod" yaml:"httpMethod"`
	// The priority of the sampling rule.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// A fixed number of matching requests to instrument per second, prior to applying the fixed rate.
	//
	// The reservoir is not used directly by services, but applies to all services using the rule collectively.
	ReservoirSize *float64 `field:"optional" json:"reservoirSize" yaml:"reservoirSize"`
	// Matches the ARN of the AWS resource on which the service runs.
	ResourceArn *string `field:"optional" json:"resourceArn" yaml:"resourceArn"`
	// The ARN of the sampling rule.
	//
	// You must specify either RuleARN or RuleName, but not both.
	RuleArn *string `field:"optional" json:"ruleArn" yaml:"ruleArn"`
	// The name of the sampling rule.
	//
	// You must specify either RuleARN or RuleName, but not both.
	RuleName *string `field:"optional" json:"ruleName" yaml:"ruleName"`
	// Matches the `name` that the service uses to identify itself in segments.
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
	// Matches the `origin` that the service uses to identify its type in segments.
	ServiceType *string `field:"optional" json:"serviceType" yaml:"serviceType"`
	// Matches the path from a request URL.
	UrlPath *string `field:"optional" json:"urlPath" yaml:"urlPath"`
}

