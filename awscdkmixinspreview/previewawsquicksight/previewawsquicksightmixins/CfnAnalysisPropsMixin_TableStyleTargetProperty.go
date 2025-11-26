package previewawsquicksightmixins


// The table style target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   tableStyleTargetProperty := &TableStyleTargetProperty{
//   	CellType: jsii.String("cellType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-tablestyletarget.html
//
type CfnAnalysisPropsMixin_TableStyleTargetProperty struct {
	// The cell type of the table style target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-tablestyletarget.html#cfn-quicksight-analysis-tablestyletarget-celltype
	//
	CellType *string `field:"optional" json:"cellType" yaml:"cellType"`
}

