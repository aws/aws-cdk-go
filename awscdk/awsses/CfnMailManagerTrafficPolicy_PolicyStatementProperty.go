package awsses


// The structure containing traffic policy conditions and actions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   policyStatementProperty := &PolicyStatementProperty{
//   	Action: jsii.String("action"),
//   	Conditions: []interface{}{
//   		&PolicyConditionProperty{
//   			BooleanExpression: &IngressBooleanExpressionProperty{
//   				Evaluate: &IngressBooleanToEvaluateProperty{
//   					Analysis: &IngressAnalysisProperty{
//   						Analyzer: jsii.String("analyzer"),
//   						ResultField: jsii.String("resultField"),
//   					},
//   				},
//   				Operator: jsii.String("operator"),
//   			},
//   			IpExpression: &IngressIpv4ExpressionProperty{
//   				Evaluate: &IngressIpToEvaluateProperty{
//   					Attribute: jsii.String("attribute"),
//   				},
//   				Operator: jsii.String("operator"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			StringExpression: &IngressStringExpressionProperty{
//   				Evaluate: &IngressStringToEvaluateProperty{
//   					Attribute: jsii.String("attribute"),
//   				},
//   				Operator: jsii.String("operator"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			TlsExpression: &IngressTlsProtocolExpressionProperty{
//   				Evaluate: &IngressTlsProtocolToEvaluateProperty{
//   					Attribute: jsii.String("attribute"),
//   				},
//   				Operator: jsii.String("operator"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagertrafficpolicy-policystatement.html
//
type CfnMailManagerTrafficPolicy_PolicyStatementProperty struct {
	// The action that informs a traffic policy resource to either allow or block the email if it matches a condition in the policy statement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagertrafficpolicy-policystatement.html#cfn-ses-mailmanagertrafficpolicy-policystatement-action
	//
	Action *string `field:"required" json:"action" yaml:"action"`
	// The list of conditions to apply to incoming messages for filtering email traffic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagertrafficpolicy-policystatement.html#cfn-ses-mailmanagertrafficpolicy-policystatement-conditions
	//
	Conditions interface{} `field:"required" json:"conditions" yaml:"conditions"`
}

