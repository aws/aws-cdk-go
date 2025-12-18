package awscleanrooms


// Parameters that control the generation of synthetic data for machine learning, including privacy settings and column classification details.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mLSyntheticDataParametersProperty := &MLSyntheticDataParametersProperty{
//   	ColumnClassification: &ColumnClassificationDetailsProperty{
//   		ColumnMapping: []interface{}{
//   			&SyntheticDataColumnPropertiesProperty{
//   				ColumnName: jsii.String("columnName"),
//   				ColumnType: jsii.String("columnType"),
//   				IsPredictiveValue: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	Epsilon: jsii.Number(123),
//   	MaxMembershipInferenceAttackScore: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-analysistemplate-mlsyntheticdataparameters.html
//
type CfnAnalysisTemplate_MLSyntheticDataParametersProperty struct {
	// Classification details for data columns that specify how each column should be treated during synthetic data generation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-analysistemplate-mlsyntheticdataparameters.html#cfn-cleanrooms-analysistemplate-mlsyntheticdataparameters-columnclassification
	//
	ColumnClassification interface{} `field:"required" json:"columnClassification" yaml:"columnClassification"`
	// The epsilon value for differential privacy when generating synthetic data.
	//
	// Lower values provide stronger privacy guarantees but may reduce data utility.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-analysistemplate-mlsyntheticdataparameters.html#cfn-cleanrooms-analysistemplate-mlsyntheticdataparameters-epsilon
	//
	Epsilon *float64 `field:"required" json:"epsilon" yaml:"epsilon"`
	// The maximum acceptable score for membership inference attack vulnerability.
	//
	// Synthetic data generation fails if the score for the resulting data exceeds this threshold.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-analysistemplate-mlsyntheticdataparameters.html#cfn-cleanrooms-analysistemplate-mlsyntheticdataparameters-maxmembershipinferenceattackscore
	//
	MaxMembershipInferenceAttackScore *float64 `field:"required" json:"maxMembershipInferenceAttackScore" yaml:"maxMembershipInferenceAttackScore"`
}

