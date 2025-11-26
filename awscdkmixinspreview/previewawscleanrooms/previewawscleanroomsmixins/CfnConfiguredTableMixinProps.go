package previewawscleanroomsmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnConfiguredTablePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnConfiguredTableMixinProps := &CfnConfiguredTableMixinProps{
//   	AllowedColumns: []*string{
//   		jsii.String("allowedColumns"),
//   	},
//   	AnalysisMethod: jsii.String("analysisMethod"),
//   	AnalysisRules: []interface{}{
//   		&AnalysisRuleProperty{
//   			Policy: &ConfiguredTableAnalysisRulePolicyProperty{
//   				V1: &ConfiguredTableAnalysisRulePolicyV1Property{
//   					Aggregation: &AnalysisRuleAggregationProperty{
//   						AdditionalAnalyses: jsii.String("additionalAnalyses"),
//   						AggregateColumns: []interface{}{
//   							&AggregateColumnProperty{
//   								ColumnNames: []*string{
//   									jsii.String("columnNames"),
//   								},
//   								Function: jsii.String("function"),
//   							},
//   						},
//   						AllowedJoinOperators: []*string{
//   							jsii.String("allowedJoinOperators"),
//   						},
//   						DimensionColumns: []*string{
//   							jsii.String("dimensionColumns"),
//   						},
//   						JoinColumns: []*string{
//   							jsii.String("joinColumns"),
//   						},
//   						JoinRequired: jsii.String("joinRequired"),
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
//   					},
//   					Custom: &AnalysisRuleCustomProperty{
//   						AdditionalAnalyses: jsii.String("additionalAnalyses"),
//   						AllowedAnalyses: []*string{
//   							jsii.String("allowedAnalyses"),
//   						},
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
//   						AdditionalAnalyses: jsii.String("additionalAnalyses"),
//   						AllowedJoinOperators: []*string{
//   							jsii.String("allowedJoinOperators"),
//   						},
//   						JoinColumns: []*string{
//   							jsii.String("joinColumns"),
//   						},
//   						ListColumns: []*string{
//   							jsii.String("listColumns"),
//   						},
//   					},
//   				},
//   			},
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	SelectedAnalysisMethods: []*string{
//   		jsii.String("selectedAnalysisMethods"),
//   	},
//   	TableReference: &TableReferenceProperty{
//   		Athena: &AthenaTableReferenceProperty{
//   			DatabaseName: jsii.String("databaseName"),
//   			OutputLocation: jsii.String("outputLocation"),
//   			Region: jsii.String("region"),
//   			TableName: jsii.String("tableName"),
//   			WorkGroup: jsii.String("workGroup"),
//   		},
//   		Glue: &GlueTableReferenceProperty{
//   			DatabaseName: jsii.String("databaseName"),
//   			Region: jsii.String("region"),
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
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-configuredtable.html
//
type CfnConfiguredTableMixinProps struct {
	// The columns within the underlying AWS Glue table that can be used within collaborations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-configuredtable.html#cfn-cleanrooms-configuredtable-allowedcolumns
	//
	AllowedColumns *[]*string `field:"optional" json:"allowedColumns" yaml:"allowedColumns"`
	// The analysis method for the configured table.
	//
	// `DIRECT_QUERY` allows SQL queries to be run directly on this table.
	//
	// `DIRECT_JOB` allows PySpark jobs to be run directly on this table.
	//
	// `MULTIPLE` allows both SQL queries and PySpark jobs to be run directly on this table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-configuredtable.html#cfn-cleanrooms-configuredtable-analysismethod
	//
	AnalysisMethod *string `field:"optional" json:"analysisMethod" yaml:"analysisMethod"`
	// The analysis rule that was created for the configured table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-configuredtable.html#cfn-cleanrooms-configuredtable-analysisrules
	//
	AnalysisRules interface{} `field:"optional" json:"analysisRules" yaml:"analysisRules"`
	// A description for the configured table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-configuredtable.html#cfn-cleanrooms-configuredtable-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A name for the configured table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-configuredtable.html#cfn-cleanrooms-configuredtable-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The selected analysis methods for the configured table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-configuredtable.html#cfn-cleanrooms-configuredtable-selectedanalysismethods
	//
	SelectedAnalysisMethods *[]*string `field:"optional" json:"selectedAnalysisMethods" yaml:"selectedAnalysisMethods"`
	// The table that this configured table represents.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-configuredtable.html#cfn-cleanrooms-configuredtable-tablereference
	//
	TableReference interface{} `field:"optional" json:"tableReference" yaml:"tableReference"`
	// An optional label that you can assign to a resource when you create it.
	//
	// Each tag consists of a key and an optional value, both of which you define. When you use tagging, you can also use tag-based access control in IAM policies to control access to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-configuredtable.html#cfn-cleanrooms-configuredtable-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

