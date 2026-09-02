package awslicensemanager


// License configuration rule statement.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   licenseConfigurationRuleStatementProperty := &LicenseConfigurationRuleStatementProperty{
//   	AndRuleStatement: &AndRuleStatementProperty{
//   		MatchingRuleStatements: []interface{}{
//   			&MatchingRuleStatementProperty{
//   				Constraint: jsii.String("constraint"),
//   				KeyToMatch: jsii.String("keyToMatch"),
//   				ValueToMatch: []*string{
//   					jsii.String("valueToMatch"),
//   				},
//   			},
//   		},
//   	},
//   	MatchingRuleStatement: &MatchingRuleStatementProperty{
//   		Constraint: jsii.String("constraint"),
//   		KeyToMatch: jsii.String("keyToMatch"),
//   		ValueToMatch: []*string{
//   			jsii.String("valueToMatch"),
//   		},
//   	},
//   	OrRuleStatement: &OrRuleStatementProperty{
//   		MatchingRuleStatements: []interface{}{
//   			&MatchingRuleStatementProperty{
//   				Constraint: jsii.String("constraint"),
//   				KeyToMatch: jsii.String("keyToMatch"),
//   				ValueToMatch: []*string{
//   					jsii.String("valueToMatch"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-licenseassetruleset-licenseconfigurationrulestatement.html
//
type CfnLicenseAssetRuleSet_LicenseConfigurationRuleStatementProperty struct {
	// AND rule statement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-licenseassetruleset-licenseconfigurationrulestatement.html#cfn-licensemanager-licenseassetruleset-licenseconfigurationrulestatement-andrulestatement
	//
	AndRuleStatement interface{} `field:"optional" json:"andRuleStatement" yaml:"andRuleStatement"`
	// Matching rule statement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-licenseassetruleset-licenseconfigurationrulestatement.html#cfn-licensemanager-licenseassetruleset-licenseconfigurationrulestatement-matchingrulestatement
	//
	MatchingRuleStatement interface{} `field:"optional" json:"matchingRuleStatement" yaml:"matchingRuleStatement"`
	// OR rule statement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-licenseassetruleset-licenseconfigurationrulestatement.html#cfn-licensemanager-licenseassetruleset-licenseconfigurationrulestatement-orrulestatement
	//
	OrRuleStatement interface{} `field:"optional" json:"orRuleStatement" yaml:"orRuleStatement"`
}

