package awsconnect


// The identifier of the task template field.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fieldIdentifierProperty := &FieldIdentifierProperty{
//   	Name: jsii.String("name"),
//   }
//
type CfnTaskTemplate_FieldIdentifierProperty struct {
	// The name of the task template field.
	Name *string `field:"required" json:"name" yaml:"name"`
}

