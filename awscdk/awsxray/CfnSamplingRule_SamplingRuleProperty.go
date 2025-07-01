package awsxray


// A sampling rule that services use to decide whether to instrument a request.
//
// Rule fields can match properties of the service, or properties of a request. The service can ignore rules that don't match its properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   samplingRuleProperty := &SamplingRuleProperty{
//   	FixedRate: jsii.Number(123),
//   	Host: jsii.String("host"),
//   	HttpMethod: jsii.String("httpMethod"),
//   	Priority: jsii.Number(123),
//   	ReservoirSize: jsii.Number(123),
//   	ResourceArn: jsii.String("resourceArn"),
//   	ServiceName: jsii.String("serviceName"),
//   	ServiceType: jsii.String("serviceType"),
//   	UrlPath: jsii.String("urlPath"),
//
//   	// the properties below are optional
//   	Attributes: map[string]*string{
//   		"attributesKey": jsii.String("attributes"),
//   	},
//   	RuleArn: jsii.String("ruleArn"),
//   	RuleName: jsii.String("ruleName"),
//   	Version: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-xray-samplingrule-samplingrule.html
//
type CfnSamplingRule_SamplingRuleProperty struct {
	// The percentage of matching requests to instrument, after the reservoir is exhausted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-xray-samplingrule-samplingrule.html#cfn-xray-samplingrule-samplingrule-fixedrate
	//
	FixedRate *float64 `field:"required" json:"fixedRate" yaml:"fixedRate"`
	// Matches the hostname from a request URL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-xray-samplingrule-samplingrule.html#cfn-xray-samplingrule-samplingrule-host
	//
	Host *string `field:"required" json:"host" yaml:"host"`
	// Matches the HTTP method of a request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-xray-samplingrule-samplingrule.html#cfn-xray-samplingrule-samplingrule-httpmethod
	//
	HttpMethod *string `field:"required" json:"httpMethod" yaml:"httpMethod"`
	// The priority of the sampling rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-xray-samplingrule-samplingrule.html#cfn-xray-samplingrule-samplingrule-priority
	//
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// A fixed number of matching requests to instrument per second, prior to applying the fixed rate.
	//
	// The reservoir is not used directly by services, but applies to all services using the rule collectively.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-xray-samplingrule-samplingrule.html#cfn-xray-samplingrule-samplingrule-reservoirsize
	//
	ReservoirSize *float64 `field:"required" json:"reservoirSize" yaml:"reservoirSize"`
	// Matches the ARN of the AWS resource on which the service runs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-xray-samplingrule-samplingrule.html#cfn-xray-samplingrule-samplingrule-resourcearn
	//
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// Matches the `name` that the service uses to identify itself in segments.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-xray-samplingrule-samplingrule.html#cfn-xray-samplingrule-samplingrule-servicename
	//
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
	// Matches the `origin` that the service uses to identify its type in segments.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-xray-samplingrule-samplingrule.html#cfn-xray-samplingrule-samplingrule-servicetype
	//
	ServiceType *string `field:"required" json:"serviceType" yaml:"serviceType"`
	// Matches the path from a request URL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-xray-samplingrule-samplingrule.html#cfn-xray-samplingrule-samplingrule-urlpath
	//
	UrlPath *string `field:"required" json:"urlPath" yaml:"urlPath"`
	// Matches attributes derived from the request.
	//
	// *Map Entries:* Maximum number of 5 items.
	//
	// *Key Length Constraints:* Minimum length of 1. Maximum length of 32.
	//
	// *Value Length Constraints:* Minimum length of 1. Maximum length of 32.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-xray-samplingrule-samplingrule.html#cfn-xray-samplingrule-samplingrule-attributes
	//
	Attributes interface{} `field:"optional" json:"attributes" yaml:"attributes"`
	// The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.
	//
	// > Specifying a sampling rule by name is recommended, as specifying by ARN will be deprecated in future.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-xray-samplingrule-samplingrule.html#cfn-xray-samplingrule-samplingrule-rulearn
	//
	RuleArn *string `field:"optional" json:"ruleArn" yaml:"ruleArn"`
	// The name of the sampling rule.
	//
	// Specify a rule by either name or ARN, but not both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-xray-samplingrule-samplingrule.html#cfn-xray-samplingrule-samplingrule-rulename
	//
	RuleName *string `field:"optional" json:"ruleName" yaml:"ruleName"`
	// The version of the sampling rule.
	//
	// `Version` can only be set when creating a new sampling rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-xray-samplingrule-samplingrule.html#cfn-xray-samplingrule-samplingrule-version
	//
	Version *float64 `field:"optional" json:"version" yaml:"version"`
}

