// The CDK Construct Library for AWS::Glue
package awscdkgluealpha


// A column of a table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import glue_alpha "github.com/aws/aws-cdk-go/awscdkgluealpha"
//
//   column := &Column{
//   	Name: jsii.String("name"),
//   	Type: &Type{
//   		InputString: jsii.String("inputString"),
//   		IsPrimitive: jsii.Boolean(false),
//   	},
//
//   	// the properties below are optional
//   	Comment: jsii.String("comment"),
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

