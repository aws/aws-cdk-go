package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tableStyleTargetProperty := &TableStyleTargetProperty{
//   	CellType: jsii.String("cellType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-tablestyletarget.html
//
type CfnAnalysis_TableStyleTargetProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-tablestyletarget.html#cfn-quicksight-analysis-tablestyletarget-celltype
	//
	CellType *string `field:"required" json:"cellType" yaml:"cellType"`
}

