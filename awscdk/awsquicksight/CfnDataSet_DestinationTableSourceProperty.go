package awsquicksight


// Specifies the source of data for a destination table, including the transform operation and column mappings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   destinationTableSourceProperty := &DestinationTableSourceProperty{
//   	TransformOperationId: jsii.String("transformOperationId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-destinationtablesource.html
//
type CfnDataSet_DestinationTableSourceProperty struct {
	// The identifier of the transform operation that provides data to the destination table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-destinationtablesource.html#cfn-quicksight-dataset-destinationtablesource-transformoperationid
	//
	TransformOperationId *string `field:"required" json:"transformOperationId" yaml:"transformOperationId"`
}

