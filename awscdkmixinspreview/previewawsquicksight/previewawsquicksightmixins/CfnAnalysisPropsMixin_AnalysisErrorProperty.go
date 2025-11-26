package previewawsquicksightmixins


// Analysis error.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   analysisErrorProperty := &AnalysisErrorProperty{
//   	Message: jsii.String("message"),
//   	Type: jsii.String("type"),
//   	ViolatedEntities: []interface{}{
//   		&EntityProperty{
//   			Path: jsii.String("path"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-analysiserror.html
//
type CfnAnalysisPropsMixin_AnalysisErrorProperty struct {
	// The message associated with the analysis error.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-analysiserror.html#cfn-quicksight-analysis-analysiserror-message
	//
	Message *string `field:"optional" json:"message" yaml:"message"`
	// The type of the analysis error.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-analysiserror.html#cfn-quicksight-analysis-analysiserror-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Lists the violated entities that caused the analysis error.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-analysiserror.html#cfn-quicksight-analysis-analysiserror-violatedentities
	//
	ViolatedEntities interface{} `field:"optional" json:"violatedEntities" yaml:"violatedEntities"`
}

