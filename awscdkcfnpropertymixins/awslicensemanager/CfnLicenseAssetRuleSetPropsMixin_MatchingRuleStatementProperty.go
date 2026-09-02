package awslicensemanager


// Matching rule statement.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   matchingRuleStatementProperty := &MatchingRuleStatementProperty{
//   	Constraint: jsii.String("constraint"),
//   	KeyToMatch: jsii.String("keyToMatch"),
//   	ValueToMatch: []*string{
//   		jsii.String("valueToMatch"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-licenseassetruleset-matchingrulestatement.html
//
type CfnLicenseAssetRuleSetPropsMixin_MatchingRuleStatementProperty struct {
	// Constraint (e.g. Equals, Not_Equals).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-licenseassetruleset-matchingrulestatement.html#cfn-licensemanager-licenseassetruleset-matchingrulestatement-constraint
	//
	Constraint *string `field:"optional" json:"constraint" yaml:"constraint"`
	// Key to match.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-licenseassetruleset-matchingrulestatement.html#cfn-licensemanager-licenseassetruleset-matchingrulestatement-keytomatch
	//
	KeyToMatch *string `field:"optional" json:"keyToMatch" yaml:"keyToMatch"`
	// Values to match.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-licenseassetruleset-matchingrulestatement.html#cfn-licensemanager-licenseassetruleset-matchingrulestatement-valuetomatch
	//
	ValueToMatch *[]*string `field:"optional" json:"valueToMatch" yaml:"valueToMatch"`
}

