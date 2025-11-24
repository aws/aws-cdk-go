package mixinsawsxray

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnSamplingRulePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSamplingRuleMixinProps := &CfnSamplingRuleMixinProps{
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
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-xray-samplingrule.html
//
type CfnSamplingRuleMixinProps struct {
	// The ARN of the sampling rule.
	//
	// Specify a rule by either name or ARN, but not both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-xray-samplingrule.html#cfn-xray-samplingrule-rulename
	//
	// Deprecated: this property has been deprecated.
	RuleName *string `field:"optional" json:"ruleName" yaml:"ruleName"`
	// The sampling rule to be created or updated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-xray-samplingrule.html#cfn-xray-samplingrule-samplingrule
	//
	SamplingRule interface{} `field:"optional" json:"samplingRule" yaml:"samplingRule"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-xray-samplingrule.html#cfn-xray-samplingrule-samplingrulerecord
	//
	// Deprecated: this property has been deprecated.
	SamplingRuleRecord interface{} `field:"optional" json:"samplingRuleRecord" yaml:"samplingRuleRecord"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-xray-samplingrule.html#cfn-xray-samplingrule-samplingruleupdate
	//
	// Deprecated: this property has been deprecated.
	SamplingRuleUpdate interface{} `field:"optional" json:"samplingRuleUpdate" yaml:"samplingRuleUpdate"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-xray-samplingrule.html#cfn-xray-samplingrule-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

