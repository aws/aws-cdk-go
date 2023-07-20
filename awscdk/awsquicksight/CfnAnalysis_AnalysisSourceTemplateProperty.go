package awsquicksight


// The source template of an analysis.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   analysisSourceTemplateProperty := &AnalysisSourceTemplateProperty{
//   	Arn: jsii.String("arn"),
//   	DataSetReferences: []interface{}{
//   		&DataSetReferenceProperty{
//   			DataSetArn: jsii.String("dataSetArn"),
//   			DataSetPlaceholder: jsii.String("dataSetPlaceholder"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-analysissourcetemplate.html
//
type CfnAnalysis_AnalysisSourceTemplateProperty struct {
	// The Amazon Resource Name (ARN) of the source template of an analysis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-analysissourcetemplate.html#cfn-quicksight-analysis-analysissourcetemplate-arn
	//
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// The dataset references of the source template of an analysis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-analysissourcetemplate.html#cfn-quicksight-analysis-analysissourcetemplate-datasetreferences
	//
	DataSetReferences interface{} `field:"required" json:"dataSetReferences" yaml:"dataSetReferences"`
}

