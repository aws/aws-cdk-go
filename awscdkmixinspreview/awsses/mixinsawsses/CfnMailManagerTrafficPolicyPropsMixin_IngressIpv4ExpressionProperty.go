package mixinsawsses


// The union type representing the allowed types for the left hand side of an IP condition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ingressIpv4ExpressionProperty := &IngressIpv4ExpressionProperty{
//   	Evaluate: &IngressIpToEvaluateProperty{
//   		Attribute: jsii.String("attribute"),
//   	},
//   	Operator: jsii.String("operator"),
//   	Values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagertrafficpolicy-ingressipv4expression.html
//
type CfnMailManagerTrafficPolicyPropsMixin_IngressIpv4ExpressionProperty struct {
	// The left hand side argument of an IP condition expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagertrafficpolicy-ingressipv4expression.html#cfn-ses-mailmanagertrafficpolicy-ingressipv4expression-evaluate
	//
	Evaluate interface{} `field:"optional" json:"evaluate" yaml:"evaluate"`
	// The matching operator for an IP condition expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagertrafficpolicy-ingressipv4expression.html#cfn-ses-mailmanagertrafficpolicy-ingressipv4expression-operator
	//
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// The right hand side argument of an IP condition expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagertrafficpolicy-ingressipv4expression.html#cfn-ses-mailmanagertrafficpolicy-ingressipv4expression-values
	//
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

