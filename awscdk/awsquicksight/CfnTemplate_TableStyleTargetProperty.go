package awsquicksight


// The table style target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tableStyleTargetProperty := &TableStyleTargetProperty{
//   	CellType: jsii.String("cellType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-tablestyletarget.html
//
type CfnTemplate_TableStyleTargetProperty struct {
	// The cell type of the table style target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-tablestyletarget.html#cfn-quicksight-template-tablestyletarget-celltype
	//
	CellType *string `field:"required" json:"cellType" yaml:"cellType"`
}

