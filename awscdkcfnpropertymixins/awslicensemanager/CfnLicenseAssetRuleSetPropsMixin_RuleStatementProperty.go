package awslicensemanager


// Rule statement.
//
// Specify exactly one of InstanceRuleStatement, LicenseRuleStatement, or LicenseConfigurationRuleStatement.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   ruleStatementProperty := &RuleStatementProperty{
//   	InstanceRuleStatement: &InstanceRuleStatementProperty{
//   		AndRuleStatement: &AndRuleStatementProperty{
//   			MatchingRuleStatements: []interface{}{
//   				&MatchingRuleStatementProperty{
//   					Constraint: jsii.String("constraint"),
//   					KeyToMatch: jsii.String("keyToMatch"),
//   					ValueToMatch: []*string{
//   						jsii.String("valueToMatch"),
//   					},
//   				},
//   			},
//   		},
//   		MatchingRuleStatement: &MatchingRuleStatementProperty{
//   			Constraint: jsii.String("constraint"),
//   			KeyToMatch: jsii.String("keyToMatch"),
//   			ValueToMatch: []*string{
//   				jsii.String("valueToMatch"),
//   			},
//   		},
//   		OrRuleStatement: &OrRuleStatementProperty{
//   			MatchingRuleStatements: []interface{}{
//   				&MatchingRuleStatementProperty{
//   					Constraint: jsii.String("constraint"),
//   					KeyToMatch: jsii.String("keyToMatch"),
//   					ValueToMatch: []*string{
//   						jsii.String("valueToMatch"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	LicenseConfigurationRuleStatement: &LicenseConfigurationRuleStatementProperty{
//   		AndRuleStatement: &AndRuleStatementProperty{
//   			MatchingRuleStatements: []interface{}{
//   				&MatchingRuleStatementProperty{
//   					Constraint: jsii.String("constraint"),
//   					KeyToMatch: jsii.String("keyToMatch"),
//   					ValueToMatch: []*string{
//   						jsii.String("valueToMatch"),
//   					},
//   				},
//   			},
//   		},
//   		MatchingRuleStatement: &MatchingRuleStatementProperty{
//   			Constraint: jsii.String("constraint"),
//   			KeyToMatch: jsii.String("keyToMatch"),
//   			ValueToMatch: []*string{
//   				jsii.String("valueToMatch"),
//   			},
//   		},
//   		OrRuleStatement: &OrRuleStatementProperty{
//   			MatchingRuleStatements: []interface{}{
//   				&MatchingRuleStatementProperty{
//   					Constraint: jsii.String("constraint"),
//   					KeyToMatch: jsii.String("keyToMatch"),
//   					ValueToMatch: []*string{
//   						jsii.String("valueToMatch"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	LicenseRuleStatement: &LicenseRuleStatementProperty{
//   		AndRuleStatement: &AndRuleStatementProperty{
//   			MatchingRuleStatements: []interface{}{
//   				&MatchingRuleStatementProperty{
//   					Constraint: jsii.String("constraint"),
//   					KeyToMatch: jsii.String("keyToMatch"),
//   					ValueToMatch: []*string{
//   						jsii.String("valueToMatch"),
//   					},
//   				},
//   			},
//   		},
//   		MatchingRuleStatement: &MatchingRuleStatementProperty{
//   			Constraint: jsii.String("constraint"),
//   			KeyToMatch: jsii.String("keyToMatch"),
//   			ValueToMatch: []*string{
//   				jsii.String("valueToMatch"),
//   			},
//   		},
//   		OrRuleStatement: &OrRuleStatementProperty{
//   			MatchingRuleStatements: []interface{}{
//   				&MatchingRuleStatementProperty{
//   					Constraint: jsii.String("constraint"),
//   					KeyToMatch: jsii.String("keyToMatch"),
//   					ValueToMatch: []*string{
//   						jsii.String("valueToMatch"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-licenseassetruleset-rulestatement.html
//
type CfnLicenseAssetRuleSetPropsMixin_RuleStatementProperty struct {
	// Instance rule statement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-licenseassetruleset-rulestatement.html#cfn-licensemanager-licenseassetruleset-rulestatement-instancerulestatement
	//
	InstanceRuleStatement interface{} `field:"optional" json:"instanceRuleStatement" yaml:"instanceRuleStatement"`
	// License configuration rule statement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-licenseassetruleset-rulestatement.html#cfn-licensemanager-licenseassetruleset-rulestatement-licenseconfigurationrulestatement
	//
	LicenseConfigurationRuleStatement interface{} `field:"optional" json:"licenseConfigurationRuleStatement" yaml:"licenseConfigurationRuleStatement"`
	// License rule statement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-licenseassetruleset-rulestatement.html#cfn-licensemanager-licenseassetruleset-rulestatement-licenserulestatement
	//
	LicenseRuleStatement interface{} `field:"optional" json:"licenseRuleStatement" yaml:"licenseRuleStatement"`
}

