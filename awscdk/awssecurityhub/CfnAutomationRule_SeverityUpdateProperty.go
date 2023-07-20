package awssecurityhub


// Updates to the severity information for a finding.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   severityUpdateProperty := &SeverityUpdateProperty{
//   	Label: jsii.String("label"),
//   	Normalized: jsii.Number(123),
//   	Product: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-severityupdate.html
//
type CfnAutomationRule_SeverityUpdateProperty struct {
	// The severity value of the finding. The allowed values are the following.
	//
	// - `INFORMATIONAL` - No issue was found.
	// - `LOW` - The issue does not require action on its own.
	// - `MEDIUM` - The issue must be addressed but not urgently.
	// - `HIGH` - The issue must be addressed as a priority.
	// - `CRITICAL` - The issue must be remediated immediately to avoid it escalating.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-severityupdate.html#cfn-securityhub-automationrule-severityupdate-label
	//
	Label *string `field:"optional" json:"label" yaml:"label"`
	// The normalized severity for the finding. This attribute is to be deprecated in favor of `Label` .
	//
	// If you provide `Normalized` and do not provide `Label` , `Label` is set automatically as follows.
	//
	// - 0 - `INFORMATIONAL`
	// - 1–39 - `LOW`
	// - 40–69 - `MEDIUM`
	// - 70–89 - `HIGH`
	// - 90–100 - `CRITICAL`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-severityupdate.html#cfn-securityhub-automationrule-severityupdate-normalized
	//
	Normalized *float64 `field:"optional" json:"normalized" yaml:"normalized"`
	// The native severity as defined by the AWS service or integrated partner product that generated the finding.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-severityupdate.html#cfn-securityhub-automationrule-severityupdate-product
	//
	Product *float64 `field:"optional" json:"product" yaml:"product"`
}

