package previewawssesmixins


// The email traffic filtering conditions which are contained in a traffic policy resource.
//
// > This data type is a UNION, so only one of the following members can be specified when used or returned.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   policyConditionProperty := &PolicyConditionProperty{
//   	BooleanExpression: &IngressBooleanExpressionProperty{
//   		Evaluate: &IngressBooleanToEvaluateProperty{
//   			Analysis: &IngressAnalysisProperty{
//   				Analyzer: jsii.String("analyzer"),
//   				ResultField: jsii.String("resultField"),
//   			},
//   			IsInAddressList: &IngressIsInAddressListProperty{
//   				AddressLists: []*string{
//   					jsii.String("addressLists"),
//   				},
//   				Attribute: jsii.String("attribute"),
//   			},
//   		},
//   		Operator: jsii.String("operator"),
//   	},
//   	IpExpression: &IngressIpv4ExpressionProperty{
//   		Evaluate: &IngressIpToEvaluateProperty{
//   			Attribute: jsii.String("attribute"),
//   		},
//   		Operator: jsii.String("operator"),
//   		Values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	Ipv6Expression: &IngressIpv6ExpressionProperty{
//   		Evaluate: &IngressIpv6ToEvaluateProperty{
//   			Attribute: jsii.String("attribute"),
//   		},
//   		Operator: jsii.String("operator"),
//   		Values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	StringExpression: &IngressStringExpressionProperty{
//   		Evaluate: &IngressStringToEvaluateProperty{
//   			Analysis: &IngressAnalysisProperty{
//   				Analyzer: jsii.String("analyzer"),
//   				ResultField: jsii.String("resultField"),
//   			},
//   			Attribute: jsii.String("attribute"),
//   		},
//   		Operator: jsii.String("operator"),
//   		Values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	TlsExpression: &IngressTlsProtocolExpressionProperty{
//   		Evaluate: &IngressTlsProtocolToEvaluateProperty{
//   			Attribute: jsii.String("attribute"),
//   		},
//   		Operator: jsii.String("operator"),
//   		Value: jsii.String("value"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagertrafficpolicy-policycondition.html
//
type CfnMailManagerTrafficPolicyPropsMixin_PolicyConditionProperty struct {
	// This represents a boolean type condition matching on the incoming mail.
	//
	// It performs the boolean operation configured in 'Operator' and evaluates the 'Protocol' object against the 'Value'.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagertrafficpolicy-policycondition.html#cfn-ses-mailmanagertrafficpolicy-policycondition-booleanexpression
	//
	BooleanExpression interface{} `field:"optional" json:"booleanExpression" yaml:"booleanExpression"`
	// This represents an IP based condition matching on the incoming mail.
	//
	// It performs the operation configured in 'Operator' and evaluates the 'Protocol' object against the 'Value'.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagertrafficpolicy-policycondition.html#cfn-ses-mailmanagertrafficpolicy-policycondition-ipexpression
	//
	IpExpression interface{} `field:"optional" json:"ipExpression" yaml:"ipExpression"`
	// This represents an IPv6 based condition matching on the incoming mail.
	//
	// It performs the operation configured in 'Operator' and evaluates the 'Protocol' object against the 'Value'.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagertrafficpolicy-policycondition.html#cfn-ses-mailmanagertrafficpolicy-policycondition-ipv6expression
	//
	Ipv6Expression interface{} `field:"optional" json:"ipv6Expression" yaml:"ipv6Expression"`
	// This represents a string based condition matching on the incoming mail.
	//
	// It performs the string operation configured in 'Operator' and evaluates the 'Protocol' object against the 'Value'.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagertrafficpolicy-policycondition.html#cfn-ses-mailmanagertrafficpolicy-policycondition-stringexpression
	//
	StringExpression interface{} `field:"optional" json:"stringExpression" yaml:"stringExpression"`
	// This represents a TLS based condition matching on the incoming mail.
	//
	// It performs the operation configured in 'Operator' and evaluates the 'Protocol' object against the 'Value'.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagertrafficpolicy-policycondition.html#cfn-ses-mailmanagertrafficpolicy-policycondition-tlsexpression
	//
	TlsExpression interface{} `field:"optional" json:"tlsExpression" yaml:"tlsExpression"`
}

