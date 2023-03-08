package awsquicksight


// Metadata for a column that is used as the input of a transform operation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputColumnProperty := &inputColumnProperty{
//   	name: jsii.String("name"),
//   	type: jsii.String("type"),
//   }
//
type CfnDataSet_InputColumnProperty struct {
	// The name of this column in the underlying data source.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The data type of the column.
	Type *string `field:"required" json:"type" yaml:"type"`
}

