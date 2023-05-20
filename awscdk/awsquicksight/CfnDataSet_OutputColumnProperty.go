package awsquicksight


// Output column.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   outputColumnProperty := &OutputColumnProperty{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Type: jsii.String("type"),
//   }
//
type CfnDataSet_OutputColumnProperty struct {
	// A description for a column.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A display name for the dataset.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The type.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

