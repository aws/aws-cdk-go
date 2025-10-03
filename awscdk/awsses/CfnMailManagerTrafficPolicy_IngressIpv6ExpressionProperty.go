package awsses


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ingressIpv6ExpressionProperty := &IngressIpv6ExpressionProperty{
//   	Evaluate: &IngressIpv6ToEvaluateProperty{
//   		Attribute: jsii.String("attribute"),
//   	},
//   	Operator: jsii.String("operator"),
//   	Values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagertrafficpolicy-ingressipv6expression.html
//
type CfnMailManagerTrafficPolicy_IngressIpv6ExpressionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagertrafficpolicy-ingressipv6expression.html#cfn-ses-mailmanagertrafficpolicy-ingressipv6expression-evaluate
	//
	Evaluate interface{} `field:"required" json:"evaluate" yaml:"evaluate"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagertrafficpolicy-ingressipv6expression.html#cfn-ses-mailmanagertrafficpolicy-ingressipv6expression-operator
	//
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagertrafficpolicy-ingressipv6expression.html#cfn-ses-mailmanagertrafficpolicy-ingressipv6expression-values
	//
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

