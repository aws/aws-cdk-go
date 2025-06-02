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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-xray-samplingrule-samplingruleupdate.html
//
type CfnSamplingRule_SamplingRuleUpdateProperty struct {
	// Matches attributes derived from the request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-xray-samplingrule-samplingruleupdate.html#cfn-xray-samplingrule-samplingruleupdate-attributes
	//
	Attributes interface{} `field:"optional" json:"attributes" yaml:"attributes"`
	// The percentage of matching requests to instrument, after the reservoir is exhausted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-xray-samplingrule-samplingruleupdate.html#cfn-xray-samplingrule-samplingruleupdate-fixedrate
	//
	FixedRate *float64 `field:"optional" json:"fixedRate" yaml:"fixedRate"`
	// Matches the hostname from a request URL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-xray-samplingrule-samplingruleupdate.html#cfn-xray-samplingrule-samplingruleupdate-host
	//
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Matches the HTTP method from a request URL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-xray-samplingrule-samplingruleupdate.html#cfn-xray-samplingrule-samplingruleupdate-httpmethod
	//
	HttpMethod *string `field:"optional" json:"httpMethod" yaml:"httpMethod"`
	// The priority of the sampling rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-xray-samplingrule-samplingruleupdate.html#cfn-xray-samplingrule-samplingruleupdate-priority
	//
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// A fixed number of matching requests to instrument per second, prior to applying the fixed rate.
	//
	// The reservoir is not used directly by services, but applies to all services using the rule collectively.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-xray-samplingrule-samplingruleupdate.html#cfn-xray-samplingrule-samplingruleupdate-reservoirsize
	//
	ReservoirSize *float64 `field:"optional" json:"reservoirSize" yaml:"reservoirSize"`
	// Matches the ARN of the AWS resource on which the service runs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-xray-samplingrule-samplingruleupdate.html#cfn-xray-samplingrule-samplingruleupdate-resourcearn
	//
	ResourceArn *string `field:"optional" json:"resourceArn" yaml:"resourceArn"`
	// The ARN of the sampling rule.
	//
	// Specify a rule by either name or ARN, but not both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-xray-samplingrule-samplingruleupdate.html#cfn-xray-samplingrule-samplingruleupdate-rulearn
	//
	RuleArn *string `field:"optional" json:"ruleArn" yaml:"ruleArn"`
	// The ARN of the sampling rule.
	//
	// Specify a rule by either name or ARN, but not both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-xray-samplingrule-samplingruleupdate.html#cfn-xray-samplingrule-samplingruleupdate-rulename
	//
	RuleName *string `field:"optional" json:"ruleName" yaml:"ruleName"`
	// Matches the name that the service uses to identify itself in segments.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-xray-samplingrule-samplingruleupdate.html#cfn-xray-samplingrule-samplingruleupdate-servicename
	//
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
	// Matches the origin that the service uses to identify its type in segments.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-xray-samplingrule-samplingruleupdate.html#cfn-xray-samplingrule-samplingruleupdate-servicetype
	//
	ServiceType *string `field:"optional" json:"serviceType" yaml:"serviceType"`
	// Matches the path from a request URL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-xray-samplingrule-samplingruleupdate.html#cfn-xray-samplingrule-samplingruleupdate-urlpath
	//
	UrlPath *string `field:"optional" json:"urlPath" yaml:"urlPath"`
}

