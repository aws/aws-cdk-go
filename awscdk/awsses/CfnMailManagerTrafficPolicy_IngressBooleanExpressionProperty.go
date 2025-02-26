package awsses


// The structure for a boolean condition matching on the incoming mail.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ingressBooleanExpressionProperty := &IngressBooleanExpressionProperty{
//   	Evaluate: &IngressBooleanToEvaluateProperty{
//   		Analysis: &IngressAnalysisProperty{
//   			Analyzer: jsii.String("analyzer"),
//   			ResultField: jsii.String("resultField"),
//   		},
//   	},
//   	Operator: jsii.String("operator"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagertrafficpolicy-ingressbooleanexpression.html
//
type CfnMailManagerTrafficPolicy_IngressBooleanExpressionProperty struct {
	// The operand on which to perform a boolean condition operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagertrafficpolicy-ingressbooleanexpression.html#cfn-ses-mailmanagertrafficpolicy-ingressbooleanexpression-evaluate
	//
	Evaluate interface{} `field:"required" json:"evaluate" yaml:"evaluate"`
	// The matching operator for a boolean condition expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagertrafficpolicy-ingressbooleanexpression.html#cfn-ses-mailmanagertrafficpolicy-ingressbooleanexpression-operator
	//
	Operator *string `field:"required" json:"operator" yaml:"operator"`
}

