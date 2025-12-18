package awscleanrooms


// The parameters that control how synthetic data is generated, including privacy settings, column classifications, and other configuration options that affect the data synthesis process.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   syntheticDataParametersProperty := &SyntheticDataParametersProperty{
//   	MlSyntheticDataParameters: &MLSyntheticDataParametersProperty{
//   		ColumnClassification: &ColumnClassificationDetailsProperty{
//   			ColumnMapping: []interface{}{
//   				&SyntheticDataColumnPropertiesProperty{
//   					ColumnName: jsii.String("columnName"),
//   					ColumnType: jsii.String("columnType"),
//   					IsPredictiveValue: jsii.Boolean(false),
//   				},
//   			},
//   		},
//   		Epsilon: jsii.Number(123),
//   		MaxMembershipInferenceAttackScore: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-analysistemplate-syntheticdataparameters.html
//
type CfnAnalysisTemplate_SyntheticDataParametersProperty struct {
	// The machine learning-specific parameters for synthetic data generation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-analysistemplate-syntheticdataparameters.html#cfn-cleanrooms-analysistemplate-syntheticdataparameters-mlsyntheticdataparameters
	//
	MlSyntheticDataParameters interface{} `field:"required" json:"mlSyntheticDataParameters" yaml:"mlSyntheticDataParameters"`
}

