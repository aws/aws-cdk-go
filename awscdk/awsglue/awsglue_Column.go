package awsglue


// A column of a table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   column := &column{
//   	name: jsii.String("name"),
//   	type: &type{
//   		inputString: jsii.String("inputString"),
//   		isPrimitive: jsii.Boolean(false),
//   	},
//
//   	// the properties below are optional
//   	comment: jsii.String("comment"),
//   }
//
// Experimental.
type Column struct {
	// Name of the column.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Type of the column.
	// Experimental.
	Type *Type `field:"required" json:"type" yaml:"type"`
	// Coment describing the column.
	// Experimental.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
}

