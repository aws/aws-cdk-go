package awsses


// The structure for a TLS related condition matching on the incoming mail.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ingressTlsProtocolExpressionProperty := &IngressTlsProtocolExpressionProperty{
//   	Evaluate: &IngressTlsProtocolToEvaluateProperty{
//   		Attribute: jsii.String("attribute"),
//   	},
//   	Operator: jsii.String("operator"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagertrafficpolicy-ingresstlsprotocolexpression.html
//
type CfnMailManagerTrafficPolicy_IngressTlsProtocolExpressionProperty struct {
	// The left hand side argument of a TLS condition expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagertrafficpolicy-ingresstlsprotocolexpression.html#cfn-ses-mailmanagertrafficpolicy-ingresstlsprotocolexpression-evaluate
	//
	Evaluate interface{} `field:"required" json:"evaluate" yaml:"evaluate"`
	// The matching operator for a TLS condition expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagertrafficpolicy-ingresstlsprotocolexpression.html#cfn-ses-mailmanagertrafficpolicy-ingresstlsprotocolexpression-operator
	//
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// The right hand side argument of a TLS condition expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagertrafficpolicy-ingresstlsprotocolexpression.html#cfn-ses-mailmanagertrafficpolicy-ingresstlsprotocolexpression-value
	//
	Value *string `field:"required" json:"value" yaml:"value"`
}

