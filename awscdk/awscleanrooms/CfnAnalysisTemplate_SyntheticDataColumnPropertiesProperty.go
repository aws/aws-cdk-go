package awscleanrooms


// Properties that define how a specific data column should be handled during synthetic data generation, including its name, type, and role in predictive modeling.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   syntheticDataColumnPropertiesProperty := &SyntheticDataColumnPropertiesProperty{
//   	ColumnName: jsii.String("columnName"),
//   	ColumnType: jsii.String("columnType"),
//   	IsPredictiveValue: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-analysistemplate-syntheticdatacolumnproperties.html
//
type CfnAnalysisTemplate_SyntheticDataColumnPropertiesProperty struct {
	// The name of the data column as it appears in the dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-analysistemplate-syntheticdatacolumnproperties.html#cfn-cleanrooms-analysistemplate-syntheticdatacolumnproperties-columnname
	//
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// The data type of the column, which determines how the synthetic data generation algorithm processes and synthesizes values for this column.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-analysistemplate-syntheticdatacolumnproperties.html#cfn-cleanrooms-analysistemplate-syntheticdatacolumnproperties-columntype
	//
	ColumnType *string `field:"required" json:"columnType" yaml:"columnType"`
	// Indicates if this column contains predictive values that should be treated as target variables in machine learning models.
	//
	// This affects how the synthetic data generation preserves statistical relationships.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-analysistemplate-syntheticdatacolumnproperties.html#cfn-cleanrooms-analysistemplate-syntheticdatacolumnproperties-ispredictivevalue
	//
	IsPredictiveValue interface{} `field:"required" json:"isPredictiveValue" yaml:"isPredictiveValue"`
}

