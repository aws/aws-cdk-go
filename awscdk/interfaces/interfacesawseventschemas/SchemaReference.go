package interfacesawseventschemas


// A reference to a Schema resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   schemaReference := &SchemaReference{
//   	SchemaArn: jsii.String("schemaArn"),
//   }
//
type SchemaReference struct {
	// The SchemaArn of the Schema resource.
	SchemaArn *string `field:"required" json:"schemaArn" yaml:"schemaArn"`
}

