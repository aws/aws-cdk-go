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
//   		Athena: &AthenaTableReferenceProperty{
//   			DatabaseName: jsii.String("databaseName"),
//   			TableName: jsii.String("tableName"),
//   			WorkGroup: jsii.String("workGroup"),
//
//   			// the properties below are optional
//   			OutputLocation: jsii.String("outputLocation"),
//   		},
//   		Glue: &GlueTableReferenceProperty{
//   			DatabaseName: jsii.String("databaseName"),
//   			TableName: jsii.String("tableName"),
//   		},
//   		Snowflake: &SnowflakeTableReferenceProperty{
//   			AccountIdentifier: jsii.String("accountIdentifier"),
//   			DatabaseName: jsii.String("databaseName"),
//   			SchemaName: jsii.String("schemaName"),
//   			SecretArn: jsii.String("secretArn"),
//   			TableName: jsii.String("tableName"),
//   			TableSchema: &SnowflakeTableSchemaProperty{
//   				V1: []interface{}{
//   					&SnowflakeTableSchemaV1Property{
//   						ColumnName: jsii.String("columnName"),
//   						ColumnType: jsii.String("columnType"),
//   					},
//   				},
//   			},
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
//   	SelectedAnalysisMethods: []*string{
//   		jsii.String("selectedAnalysisMethods"),
//   	},
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
	// `DIRECT_QUERY` allows SQL queries to be run directly on this table.
	//
	// `DIRECT_JOB` allows PySpark jobs to be run directly on this table.
	//
	// `MULTIPLE` allows both SQL queries and PySpark jobs to be run directly on this table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-configuredtable.html#cfn-cleanrooms-configuredtable-analysismethod
	//
	AnalysisMethod *string `field:"required" json:"analysisMethod" yaml:"analysisMethod"`
	// A name for the configured table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-configuredtable.html#cfn-cleanrooms-configuredtable-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The table that this configured table represents.
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
	// The selected analysis methods for the configured table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-configuredtable.html#cfn-cleanrooms-configuredtable-selectedanalysismethods
	//
	SelectedAnalysisMethods *[]*string `field:"optional" json:"selectedAnalysisMethods" yaml:"selectedAnalysisMethods"`
	// An optional label that you can assign to a resource when you create it.
	//
	// Each tag consists of a key and an optional value, both of which you define. When you use tagging, you can also use tag-based access control in IAM policies to control access to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-configuredtable.html#cfn-cleanrooms-configuredtable-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

