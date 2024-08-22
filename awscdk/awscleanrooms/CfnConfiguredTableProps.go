package awscleanrooms

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnConfiguredTable`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConfiguredTableProps := &CfnConfiguredTableProps{
//   	AllowedColumns: []*string{
//   		jsii.String("allowedColumns"),
//   	},
//   	AnalysisMethod: jsii.String("analysisMethod"),
//   	Name: jsii.String("name"),
//   	TableReference: &TableReferenceProperty{
//   		Glue: &GlueTableReferenceProperty{
//   			DatabaseName: jsii.String("databaseName"),
//   			TableName: jsii.String("tableName"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	AnalysisRules: []interface{}{
//   		&AnalysisRuleProperty{
//   			Policy: &ConfiguredTableAnalysisRulePolicyProperty{
//   				V1: &ConfiguredTableAnalysisRulePolicyV1Property{
//   					Aggregation: &AnalysisRuleAggregationProperty{
//   						AggregateColumns: []interface{}{
//   							&AggregateColumnProperty{
//   								ColumnNames: []*string{
//   									jsii.String("columnNames"),
//   								},
//   								Function: jsii.String("function"),
//   							},
//   						},
//   						DimensionColumns: []*string{
//   							jsii.String("dimensionColumns"),
//   						},
//   						JoinColumns: []*string{
//   							jsii.String("joinColumns"),
//   						},
//   						OutputConstraints: []interface{}{
//   							&AggregationConstraintProperty{
//   								ColumnName: jsii.String("columnName"),
//   								Minimum: jsii.Number(123),
//   								Type: jsii.String("type"),
//   							},
//   						},
//   						ScalarFunctions: []*string{
//   							jsii.String("scalarFunctions"),
//   						},
//
//   						// the properties below are optional
//   						AdditionalAnalyses: jsii.String("additionalAnalyses"),
//   						AllowedJoinOperators: []*string{
//   							jsii.String("allowedJoinOperators"),
//   						},
//   						JoinRequired: jsii.String("joinRequired"),
//   					},
//   					Custom: &AnalysisRuleCustomProperty{
//   						AllowedAnalyses: []*string{
//   							jsii.String("allowedAnalyses"),
//   						},
//
//   						// the properties below are optional
//   						AdditionalAnalyses: jsii.String("additionalAnalyses"),
//   						AllowedAnalysisProviders: []*string{
//   							jsii.String("allowedAnalysisProviders"),
//   						},
//   						DifferentialPrivacy: &DifferentialPrivacyProperty{
//   							Columns: []interface{}{
//   								&DifferentialPrivacyColumnProperty{
//   									Name: jsii.String("name"),
//   								},
//   							},
//   						},
//   						DisallowedOutputColumns: []*string{
//   							jsii.String("disallowedOutputColumns"),
//   						},
//   					},
//   					List: &AnalysisRuleListProperty{
//   						JoinColumns: []*string{
//   							jsii.String("joinColumns"),
//   						},
//   						ListColumns: []*string{
//   							jsii.String("listColumns"),
//   						},
//
//   						// the properties below are optional
//   						AdditionalAnalyses: jsii.String("additionalAnalyses"),
//   						AllowedJoinOperators: []*string{
//   							jsii.String("allowedJoinOperators"),
//   						},
//   					},
//   				},
//   			},
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-configuredtable.html
//
type CfnConfiguredTableProps struct {
	// The columns within the underlying AWS Glue table that can be utilized within collaborations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-configuredtable.html#cfn-cleanrooms-configuredtable-allowedcolumns
	//
	AllowedColumns *[]*string `field:"required" json:"allowedColumns" yaml:"allowedColumns"`
	// The analysis method for the configured table.
	//
	// The only valid value is currently `DIRECT_QUERY`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-configuredtable.html#cfn-cleanrooms-configuredtable-analysismethod
	//
	AnalysisMethod *string `field:"required" json:"analysisMethod" yaml:"analysisMethod"`
	// A name for the configured table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-configuredtable.html#cfn-cleanrooms-configuredtable-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The AWS Glue table that this configured table represents.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-configuredtable.html#cfn-cleanrooms-configuredtable-tablereference
	//
	TableReference interface{} `field:"required" json:"tableReference" yaml:"tableReference"`
	// The analysis rule that was created for the configured table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-configuredtable.html#cfn-cleanrooms-configuredtable-analysisrules
	//
	AnalysisRules interface{} `field:"optional" json:"analysisRules" yaml:"analysisRules"`
	// A description for the configured table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-configuredtable.html#cfn-cleanrooms-configuredtable-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An optional label that you can assign to a resource when you create it.
	//
	// Each tag consists of a key and an optional value, both of which you define. When you use tagging, you can also use tag-based access control in IAM policies to control access to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-configuredtable.html#cfn-cleanrooms-configuredtable-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

