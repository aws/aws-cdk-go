package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataPathTypeProperty := &DataPathTypeProperty{
//   	PivotTableDataPathType: jsii.String("pivotTableDataPathType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-datapathtype.html
//
type CfnAnalysis_DataPathTypeProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-datapathtype.html#cfn-quicksight-analysis-datapathtype-pivottabledatapathtype
	//
	PivotTableDataPathType *string `field:"optional" json:"pivotTableDataPathType" yaml:"pivotTableDataPathType"`
}

