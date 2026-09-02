package awslicensemanager

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnLicenseAssetRuleSetPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnLicenseAssetRuleSetMixinProps := &CfnLicenseAssetRuleSetMixinProps{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Rules: []interface{}{
//   		&LicenseAssetRuleProperty{
//   			RuleStatement: &RuleStatementProperty{
//   				InstanceRuleStatement: &InstanceRuleStatementProperty{
//   					AndRuleStatement: &AndRuleStatementProperty{
//   						MatchingRuleStatements: []interface{}{
//   							&MatchingRuleStatementProperty{
//   								Constraint: jsii.String("constraint"),
//   								KeyToMatch: jsii.String("keyToMatch"),
//   								ValueToMatch: []*string{
//   									jsii.String("valueToMatch"),
//   								},
//   							},
//   						},
//   					},
//   					MatchingRuleStatement: &MatchingRuleStatementProperty{
//   						Constraint: jsii.String("constraint"),
//   						KeyToMatch: jsii.String("keyToMatch"),
//   						ValueToMatch: []*string{
//   							jsii.String("valueToMatch"),
//   						},
//   					},
//   					OrRuleStatement: &OrRuleStatementProperty{
//   						MatchingRuleStatements: []interface{}{
//   							&MatchingRuleStatementProperty{
//   								Constraint: jsii.String("constraint"),
//   								KeyToMatch: jsii.String("keyToMatch"),
//   								ValueToMatch: []*string{
//   									jsii.String("valueToMatch"),
//   								},
//   							},
//   						},
//   					},
//   				},
//   				LicenseConfigurationRuleStatement: &LicenseConfigurationRuleStatementProperty{
//   					AndRuleStatement: &AndRuleStatementProperty{
//   						MatchingRuleStatements: []interface{}{
//   							&MatchingRuleStatementProperty{
//   								Constraint: jsii.String("constraint"),
//   								KeyToMatch: jsii.String("keyToMatch"),
//   								ValueToMatch: []*string{
//   									jsii.String("valueToMatch"),
//   								},
//   							},
//   						},
//   					},
//   					MatchingRuleStatement: &MatchingRuleStatementProperty{
//   						Constraint: jsii.String("constraint"),
//   						KeyToMatch: jsii.String("keyToMatch"),
//   						ValueToMatch: []*string{
//   							jsii.String("valueToMatch"),
//   						},
//   					},
//   					OrRuleStatement: &OrRuleStatementProperty{
//   						MatchingRuleStatements: []interface{}{
//   							&MatchingRuleStatementProperty{
//   								Constraint: jsii.String("constraint"),
//   								KeyToMatch: jsii.String("keyToMatch"),
//   								ValueToMatch: []*string{
//   									jsii.String("valueToMatch"),
//   								},
//   							},
//   						},
//   					},
//   				},
//   				LicenseRuleStatement: &LicenseRuleStatementProperty{
//   					AndRuleStatement: &AndRuleStatementProperty{
//   						MatchingRuleStatements: []interface{}{
//   							&MatchingRuleStatementProperty{
//   								Constraint: jsii.String("constraint"),
//   								KeyToMatch: jsii.String("keyToMatch"),
//   								ValueToMatch: []*string{
//   									jsii.String("valueToMatch"),
//   								},
//   							},
//   						},
//   					},
//   					MatchingRuleStatement: &MatchingRuleStatementProperty{
//   						Constraint: jsii.String("constraint"),
//   						KeyToMatch: jsii.String("keyToMatch"),
//   						ValueToMatch: []*string{
//   							jsii.String("valueToMatch"),
//   						},
//   					},
//   					OrRuleStatement: &OrRuleStatementProperty{
//   						MatchingRuleStatements: []interface{}{
//   							&MatchingRuleStatementProperty{
//   								Constraint: jsii.String("constraint"),
//   								KeyToMatch: jsii.String("keyToMatch"),
//   								ValueToMatch: []*string{
//   									jsii.String("valueToMatch"),
//   								},
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-licenseassetruleset.html
//
type CfnLicenseAssetRuleSetMixinProps struct {
	// License asset ruleset description.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-licenseassetruleset.html#cfn-licensemanager-licenseassetruleset-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// License asset ruleset name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-licenseassetruleset.html#cfn-licensemanager-licenseassetruleset-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// License asset rules.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-licenseassetruleset.html#cfn-licensemanager-licenseassetruleset-rules
	//
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
	// Tags to add to the license asset ruleset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-licenseassetruleset.html#cfn-licensemanager-licenseassetruleset-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

