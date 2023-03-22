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
//   cfnSamplingRuleProps := &cfnSamplingRuleProps{
//   	ruleName: jsii.String("ruleName"),
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
//   	samplingRuleRecord: &samplingRuleRecordProperty{
//   		createdAt: jsii.String("createdAt"),
//   		modifiedAt: jsii.String("modifiedAt"),
//   		samplingRule: &samplingRuleProperty{
//   			attributes: map[string]*string{
//   				"attributesKey": jsii.String("attributes"),
//   			},
//   			fixedRate: jsii.Number(123),
//   			host: jsii.String("host"),
//   			httpMethod: jsii.String("httpMethod"),
//   			priority: jsii.Number(123),
//   			reservoirSize: jsii.Number(123),
//   			resourceArn: jsii.String("resourceArn"),
//   			ruleArn: jsii.String("ruleArn"),
//   			ruleName: jsii.String("ruleName"),
//   			serviceName: jsii.String("serviceName"),
//   			serviceType: jsii.String("serviceType"),
//   			urlPath: jsii.String("urlPath"),
//   			version: jsii.Number(123),
//   		},
//   	},
//   	samplingRuleUpdate: &samplingRuleUpdateProperty{
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
//   	},
//   	tags: []interface{}{
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
	// A document specifying changes to a sampling rule's configuration.
	//
	// Must be provided if updating an existing sampling rule. Not valid when creating a new sampling rule.
	//
	// > The `Version` of a sampling rule cannot be updated, and is not part of `SamplingRuleUpdate` .
	SamplingRuleUpdate interface{} `field:"optional" json:"samplingRuleUpdate" yaml:"samplingRuleUpdate"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]interface{} `field:"optional" json:"tags" yaml:"tags"`
}

