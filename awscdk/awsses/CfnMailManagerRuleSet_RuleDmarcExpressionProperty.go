package awsses


// A DMARC policy expression.
//
// The condition matches if the given DMARC policy matches that of the incoming email.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ruleDmarcExpressionProperty := &RuleDmarcExpressionProperty{
//   	Operator: jsii.String("operator"),
//   	Values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-ruledmarcexpression.html
//
type CfnMailManagerRuleSet_RuleDmarcExpressionProperty struct {
	// The operator to apply to the DMARC policy of the incoming email.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-ruledmarcexpression.html#cfn-ses-mailmanagerruleset-ruledmarcexpression-operator
	//
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// The values to use for the given DMARC policy operator.
	//
	// For the operator EQUALS, if multiple values are given, they are evaluated as an OR. That is, if any of the given values match, the condition is deemed to match. For the operator NOT_EQUALS, if multiple values are given, they are evaluated as an AND. That is, only if the email's DMARC policy is not equal to any of the given values, then the condition is deemed to match.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-ruledmarcexpression.html#cfn-ses-mailmanagerruleset-ruledmarcexpression-values
	//
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

