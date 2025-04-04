package awsses


// The structure for a string based condition matching on the incoming mail.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ingressStringExpressionProperty := &IngressStringExpressionProperty{
//   	Evaluate: &IngressStringToEvaluateProperty{
//   		Attribute: jsii.String("attribute"),
//   	},
//   	Operator: jsii.String("operator"),
//   	Values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagertrafficpolicy-ingressstringexpression.html
//
type CfnMailManagerTrafficPolicy_IngressStringExpressionProperty struct {
	// The left hand side argument of a string condition expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagertrafficpolicy-ingressstringexpression.html#cfn-ses-mailmanagertrafficpolicy-ingressstringexpression-evaluate
	//
	Evaluate interface{} `field:"required" json:"evaluate" yaml:"evaluate"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagertrafficpolicy-ingressstringexpression.html#cfn-ses-mailmanagertrafficpolicy-ingressstringexpression-operator
	//
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// The right hand side argument of a string condition expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagertrafficpolicy-ingressstringexpression.html#cfn-ses-mailmanagertrafficpolicy-ingressstringexpression-values
	//
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

