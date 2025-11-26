package previewawsquicksightmixins


// The source entity of an analysis.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   analysisSourceEntityProperty := &AnalysisSourceEntityProperty{
//   	SourceTemplate: &AnalysisSourceTemplateProperty{
//   		Arn: jsii.String("arn"),
//   		DataSetReferences: []interface{}{
//   			&DataSetReferenceProperty{
//   				DataSetArn: jsii.String("dataSetArn"),
//   				DataSetPlaceholder: jsii.String("dataSetPlaceholder"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-analysissourceentity.html
//
type CfnAnalysisPropsMixin_AnalysisSourceEntityProperty struct {
	// The source template for the source entity of the analysis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-analysissourceentity.html#cfn-quicksight-analysis-analysissourceentity-sourcetemplate
	//
	SourceTemplate interface{} `field:"optional" json:"sourceTemplate" yaml:"sourceTemplate"`
}

