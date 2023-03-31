package awsquicksight


// Output column.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   outputColumnProperty := &outputColumnProperty{
//   	description: jsii.String("description"),
//   	name: jsii.String("name"),
//   	type: jsii.String("type"),
//   }
//
type CfnDataSet_OutputColumnProperty struct {
	// A description for a column.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A display name for the dataset.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Type.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

