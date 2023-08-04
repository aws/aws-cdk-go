package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lookbackWindowProperty := &LookbackWindowProperty{
//   	ColumnName: jsii.String("columnName"),
//   	Size: jsii.Number(123),
//   	SizeUnit: jsii.String("sizeUnit"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-lookbackwindow.html
//
type CfnDataSet_LookbackWindowProperty struct {
	// <p>Column Name</p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-lookbackwindow.html#cfn-quicksight-dataset-lookbackwindow-columnname
	//
	ColumnName *string `field:"optional" json:"columnName" yaml:"columnName"`
	// <p>Size</p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-lookbackwindow.html#cfn-quicksight-dataset-lookbackwindow-size
	//
	Size *float64 `field:"optional" json:"size" yaml:"size"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-lookbackwindow.html#cfn-quicksight-dataset-lookbackwindow-sizeunit
	//
	SizeUnit *string `field:"optional" json:"sizeUnit" yaml:"sizeUnit"`
}

