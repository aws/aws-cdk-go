package previewawscleanroomsmixins


// Contains classification information for data columns, including mappings that specify how columns should be handled during synthetic data generation and privacy analysis.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   columnClassificationDetailsProperty := &ColumnClassificationDetailsProperty{
//   	ColumnMapping: []interface{}{
//   		&SyntheticDataColumnPropertiesProperty{
//   			ColumnName: jsii.String("columnName"),
//   			ColumnType: jsii.String("columnType"),
//   			IsPredictiveValue: jsii.Boolean(false),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-analysistemplate-columnclassificationdetails.html
//
type CfnAnalysisTemplatePropsMixin_ColumnClassificationDetailsProperty struct {
	// A mapping that defines the classification of data columns for synthetic data generation and specifies how each column should be handled during the privacy-preserving data synthesis process.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-analysistemplate-columnclassificationdetails.html#cfn-cleanrooms-analysistemplate-columnclassificationdetails-columnmapping
	//
	ColumnMapping interface{} `field:"optional" json:"columnMapping" yaml:"columnMapping"`
}

