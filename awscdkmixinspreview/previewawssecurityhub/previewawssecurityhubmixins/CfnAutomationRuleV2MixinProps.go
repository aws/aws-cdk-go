package previewawssecurityhubmixins


// Properties for CfnAutomationRuleV2PropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAutomationRuleV2MixinProps := &CfnAutomationRuleV2MixinProps{
//   	Actions: []interface{}{
//   		&AutomationRulesActionV2Property{
//   			ExternalIntegrationConfiguration: &ExternalIntegrationConfigurationProperty{
//   				ConnectorArn: jsii.String("connectorArn"),
//   			},
//   			FindingFieldsUpdate: &AutomationRulesFindingFieldsUpdateV2Property{
//   				Comment: jsii.String("comment"),
//   				SeverityId: jsii.Number(123),
//   				StatusId: jsii.Number(123),
//   			},
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	Criteria: &CriteriaProperty{
//   		OcsfFindingCriteria: &OcsfFindingFiltersProperty{
//   			CompositeFilters: []interface{}{
//   				&CompositeFilterProperty{
//   					BooleanFilters: []interface{}{
//   						&OcsfBooleanFilterProperty{
//   							FieldName: jsii.String("fieldName"),
//   							Filter: &BooleanFilterProperty{
//   								Value: jsii.Boolean(false),
//   							},
//   						},
//   					},
//   					DateFilters: []interface{}{
//   						&OcsfDateFilterProperty{
//   							FieldName: jsii.String("fieldName"),
//   							Filter: &DateFilterProperty{
//   								DateRange: &DateRangeProperty{
//   									Unit: jsii.String("unit"),
//   									Value: jsii.Number(123),
//   								},
//   								End: jsii.String("end"),
//   								Start: jsii.String("start"),
//   							},
//   						},
//   					},
//   					MapFilters: []interface{}{
//   						&OcsfMapFilterProperty{
//   							FieldName: jsii.String("fieldName"),
//   							Filter: &MapFilterProperty{
//   								Comparison: jsii.String("comparison"),
//   								Key: jsii.String("key"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   					},
//   					NumberFilters: []interface{}{
//   						&OcsfNumberFilterProperty{
//   							FieldName: jsii.String("fieldName"),
//   							Filter: &NumberFilterProperty{
//   								Eq: jsii.Number(123),
//   								Gte: jsii.Number(123),
//   								Lte: jsii.Number(123),
//   							},
//   						},
//   					},
//   					Operator: jsii.String("operator"),
//   					StringFilters: []interface{}{
//   						&OcsfStringFilterProperty{
//   							FieldName: jsii.String("fieldName"),
//   							Filter: &StringFilterProperty{
//   								Comparison: jsii.String("comparison"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   			CompositeOperator: jsii.String("compositeOperator"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	RuleName: jsii.String("ruleName"),
//   	RuleOrder: jsii.Number(123),
//   	RuleStatus: jsii.String("ruleStatus"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-automationrulev2.html
//
type CfnAutomationRuleV2MixinProps struct {
	// A list of actions to be performed when the rule criteria is met.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-automationrulev2.html#cfn-securityhub-automationrulev2-actions
	//
	Actions interface{} `field:"optional" json:"actions" yaml:"actions"`
	// The filtering type and configuration of the automation rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-automationrulev2.html#cfn-securityhub-automationrulev2-criteria
	//
	Criteria interface{} `field:"optional" json:"criteria" yaml:"criteria"`
	// A description of the V2 automation rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-automationrulev2.html#cfn-securityhub-automationrulev2-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the V2 automation rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-automationrulev2.html#cfn-securityhub-automationrulev2-rulename
	//
	RuleName *string `field:"optional" json:"ruleName" yaml:"ruleName"`
	// The value for the rule priority.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-automationrulev2.html#cfn-securityhub-automationrulev2-ruleorder
	//
	RuleOrder *float64 `field:"optional" json:"ruleOrder" yaml:"ruleOrder"`
	// The status of the V2 automation rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-automationrulev2.html#cfn-securityhub-automationrulev2-rulestatus
	//
	RuleStatus *string `field:"optional" json:"ruleStatus" yaml:"ruleStatus"`
	// A list of key-value pairs associated with the V2 automation rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-automationrulev2.html#cfn-securityhub-automationrulev2-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

