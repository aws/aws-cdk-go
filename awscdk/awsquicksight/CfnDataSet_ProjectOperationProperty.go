package awsquicksight


// A transform operation that projects columns.
//
// Operations that come after a projection can only refer to projected columns.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   projectOperationProperty := &ProjectOperationProperty{
//   	ProjectedColumns: []*string{
//   		jsii.String("projectedColumns"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-projectoperation.html
//
type CfnDataSet_ProjectOperationProperty struct {
	// Projected columns.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-projectoperation.html#cfn-quicksight-dataset-projectoperation-projectedcolumns
	//
	ProjectedColumns *[]*string `field:"optional" json:"projectedColumns" yaml:"projectedColumns"`
}

