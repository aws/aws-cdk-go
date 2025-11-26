package previewawssesmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnMailManagerTrafficPolicyPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMailManagerTrafficPolicyMixinProps := &CfnMailManagerTrafficPolicyMixinProps{
//   	DefaultAction: jsii.String("defaultAction"),
//   	MaxMessageSizeBytes: jsii.Number(123),
//   	PolicyStatements: []interface{}{
//   		&PolicyStatementProperty{
//   			Action: jsii.String("action"),
//   			Conditions: []interface{}{
//   				&PolicyConditionProperty{
//   					BooleanExpression: &IngressBooleanExpressionProperty{
//   						Evaluate: &IngressBooleanToEvaluateProperty{
//   							Analysis: &IngressAnalysisProperty{
//   								Analyzer: jsii.String("analyzer"),
//   								ResultField: jsii.String("resultField"),
//   							},
//   							IsInAddressList: &IngressIsInAddressListProperty{
//   								AddressLists: []*string{
//   									jsii.String("addressLists"),
//   								},
//   								Attribute: jsii.String("attribute"),
//   							},
//   						},
//   						Operator: jsii.String("operator"),
//   					},
//   					IpExpression: &IngressIpv4ExpressionProperty{
//   						Evaluate: &IngressIpToEvaluateProperty{
//   							Attribute: jsii.String("attribute"),
//   						},
//   						Operator: jsii.String("operator"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   					Ipv6Expression: &IngressIpv6ExpressionProperty{
//   						Evaluate: &IngressIpv6ToEvaluateProperty{
//   							Attribute: jsii.String("attribute"),
//   						},
//   						Operator: jsii.String("operator"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   					StringExpression: &IngressStringExpressionProperty{
//   						Evaluate: &IngressStringToEvaluateProperty{
//   							Analysis: &IngressAnalysisProperty{
//   								Analyzer: jsii.String("analyzer"),
//   								ResultField: jsii.String("resultField"),
//   							},
//   							Attribute: jsii.String("attribute"),
//   						},
//   						Operator: jsii.String("operator"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   					TlsExpression: &IngressTlsProtocolExpressionProperty{
//   						Evaluate: &IngressTlsProtocolToEvaluateProperty{
//   							Attribute: jsii.String("attribute"),
//   						},
//   						Operator: jsii.String("operator"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TrafficPolicyName: jsii.String("trafficPolicyName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-mailmanagertrafficpolicy.html
//
type CfnMailManagerTrafficPolicyMixinProps struct {
	// Default action instructs the traﬃc policy to either Allow or Deny (block) messages that fall outside of (or not addressed by) the conditions of your policy statements.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-mailmanagertrafficpolicy.html#cfn-ses-mailmanagertrafficpolicy-defaultaction
	//
	DefaultAction *string `field:"optional" json:"defaultAction" yaml:"defaultAction"`
	// The maximum message size in bytes of email which is allowed in by this traffic policy—anything larger will be blocked.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-mailmanagertrafficpolicy.html#cfn-ses-mailmanagertrafficpolicy-maxmessagesizebytes
	//
	MaxMessageSizeBytes *float64 `field:"optional" json:"maxMessageSizeBytes" yaml:"maxMessageSizeBytes"`
	// Conditional statements for filtering email traffic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-mailmanagertrafficpolicy.html#cfn-ses-mailmanagertrafficpolicy-policystatements
	//
	PolicyStatements interface{} `field:"optional" json:"policyStatements" yaml:"policyStatements"`
	// The tags used to organize, track, or control access for the resource.
	//
	// For example, { "tags": {"key1":"value1", "key2":"value2"} }.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-mailmanagertrafficpolicy.html#cfn-ses-mailmanagertrafficpolicy-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The name of the policy.
	//
	// The policy name cannot exceed 64 characters and can only include alphanumeric characters, dashes, and underscores.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-mailmanagertrafficpolicy.html#cfn-ses-mailmanagertrafficpolicy-trafficpolicyname
	//
	TrafficPolicyName *string `field:"optional" json:"trafficPolicyName" yaml:"trafficPolicyName"`
}

