package awslicensemanager


// Instance rule statement.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   instanceRuleStatementProperty := &InstanceRuleStatementProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-licenseassetruleset-instancerulestatement.html
//
type CfnLicenseAssetRuleSetPropsMixin_InstanceRuleStatementProperty struct {
	// AND rule statement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-licenseassetruleset-instancerulestatement.html#cfn-licensemanager-licenseassetruleset-instancerulestatement-andrulestatement
	//
	AndRuleStatement interface{} `field:"optional" json:"andRuleStatement" yaml:"andRuleStatement"`
	// Matching rule statement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-licenseassetruleset-instancerulestatement.html#cfn-licensemanager-licenseassetruleset-instancerulestatement-matchingrulestatement
	//
	MatchingRuleStatement interface{} `field:"optional" json:"matchingRuleStatement" yaml:"matchingRuleStatement"`
	// OR rule statement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-licenseassetruleset-instancerulestatement.html#cfn-licensemanager-licenseassetruleset-instancerulestatement-orrulestatement
	//
	OrRuleStatement interface{} `field:"optional" json:"orRuleStatement" yaml:"orRuleStatement"`
}

