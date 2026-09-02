package awslicensemanager


// License asset rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   licenseAssetRuleProperty := &LicenseAssetRuleProperty{
//   	RuleStatement: &RuleStatementProperty{
//   		InstanceRuleStatement: &InstanceRuleStatementProperty{
//   			AndRuleStatement: &AndRuleStatementProperty{
//   				MatchingRuleStatements: []interface{}{
//   					&MatchingRuleStatementProperty{
//   						Constraint: jsii.String("constraint"),
//   						KeyToMatch: jsii.String("keyToMatch"),
//   						ValueToMatch: []*string{
//   							jsii.String("valueToMatch"),
//   						},
//   					},
//   				},
//   			},
//   			MatchingRuleStatement: &MatchingRuleStatementProperty{
//   				Constraint: jsii.String("constraint"),
//   				KeyToMatch: jsii.String("keyToMatch"),
//   				ValueToMatch: []*string{
//   					jsii.String("valueToMatch"),
//   				},
//   			},
//   			OrRuleStatement: &OrRuleStatementProperty{
//   				MatchingRuleStatements: []interface{}{
//   					&MatchingRuleStatementProperty{
//   						Constraint: jsii.String("constraint"),
//   						KeyToMatch: jsii.String("keyToMatch"),
//   						ValueToMatch: []*string{
//   							jsii.String("valueToMatch"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		LicenseConfigurationRuleStatement: &LicenseConfigurationRuleStatementProperty{
//   			AndRuleStatement: &AndRuleStatementProperty{
//   				MatchingRuleStatements: []interface{}{
//   					&MatchingRuleStatementProperty{
//   						Constraint: jsii.String("constraint"),
//   						KeyToMatch: jsii.String("keyToMatch"),
//   						ValueToMatch: []*string{
//   							jsii.String("valueToMatch"),
//   						},
//   					},
//   				},
//   			},
//   			MatchingRuleStatement: &MatchingRuleStatementProperty{
//   				Constraint: jsii.String("constraint"),
//   				KeyToMatch: jsii.String("keyToMatch"),
//   				ValueToMatch: []*string{
//   					jsii.String("valueToMatch"),
//   				},
//   			},
//   			OrRuleStatement: &OrRuleStatementProperty{
//   				MatchingRuleStatements: []interface{}{
//   					&MatchingRuleStatementProperty{
//   						Constraint: jsii.String("constraint"),
//   						KeyToMatch: jsii.String("keyToMatch"),
//   						ValueToMatch: []*string{
//   							jsii.String("valueToMatch"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		LicenseRuleStatement: &LicenseRuleStatementProperty{
//   			AndRuleStatement: &AndRuleStatementProperty{
//   				MatchingRuleStatements: []interface{}{
//   					&MatchingRuleStatementProperty{
//   						Constraint: jsii.String("constraint"),
//   						KeyToMatch: jsii.String("keyToMatch"),
//   						ValueToMatch: []*string{
//   							jsii.String("valueToMatch"),
//   						},
//   					},
//   				},
//   			},
//   			MatchingRuleStatement: &MatchingRuleStatementProperty{
//   				Constraint: jsii.String("constraint"),
//   				KeyToMatch: jsii.String("keyToMatch"),
//   				ValueToMatch: []*string{
//   					jsii.String("valueToMatch"),
//   				},
//   			},
//   			OrRuleStatement: &OrRuleStatementProperty{
//   				MatchingRuleStatements: []interface{}{
//   					&MatchingRuleStatementProperty{
//   						Constraint: jsii.String("constraint"),
//   						KeyToMatch: jsii.String("keyToMatch"),
//   						ValueToMatch: []*string{
//   							jsii.String("valueToMatch"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-licenseassetruleset-licenseassetrule.html
//
type CfnLicenseAssetRuleSetPropsMixin_LicenseAssetRuleProperty struct {
	// Rule statement.
	//
	// Specify exactly one of InstanceRuleStatement, LicenseRuleStatement, or LicenseConfigurationRuleStatement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-licenseassetruleset-licenseassetrule.html#cfn-licensemanager-licenseassetruleset-licenseassetrule-rulestatement
	//
	RuleStatement interface{} `field:"optional" json:"ruleStatement" yaml:"ruleStatement"`
}

