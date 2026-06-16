package awsrtbfabric

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnLinkRoutingRule`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLinkRoutingRuleProps := &CfnLinkRoutingRuleProps{
//   	Conditions: &RuleConditionProperty{
//   		HostHeader: jsii.String("hostHeader"),
//   		HostHeaderWildcard: jsii.String("hostHeaderWildcard"),
//   		PathExact: jsii.String("pathExact"),
//   		PathPrefix: jsii.String("pathPrefix"),
//   		QueryStringEquals: &QueryStringKeyValuePairProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   		QueryStringExists: jsii.String("queryStringExists"),
//   	},
//   	GatewayId: jsii.String("gatewayId"),
//   	LinkId: jsii.String("linkId"),
//   	Priority: jsii.Number(123),
//
//   	// the properties below are optional
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rtbfabric-linkroutingrule.html
//
type CfnLinkRoutingRuleProps struct {
	// Conditions for a routing rule.
	//
	// All non-null fields must match (AND logic). At least one field must be set. HostHeader and HostHeaderWildcard are mutually exclusive. PathPrefix and PathExact are mutually exclusive.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rtbfabric-linkroutingrule.html#cfn-rtbfabric-linkroutingrule-conditions
	//
	Conditions interface{} `field:"required" json:"conditions" yaml:"conditions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rtbfabric-linkroutingrule.html#cfn-rtbfabric-linkroutingrule-gatewayid
	//
	GatewayId *string `field:"required" json:"gatewayId" yaml:"gatewayId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rtbfabric-linkroutingrule.html#cfn-rtbfabric-linkroutingrule-linkid
	//
	LinkId *string `field:"required" json:"linkId" yaml:"linkId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rtbfabric-linkroutingrule.html#cfn-rtbfabric-linkroutingrule-priority
	//
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Tags to assign to the LinkRoutingRule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rtbfabric-linkroutingrule.html#cfn-rtbfabric-linkroutingrule-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

