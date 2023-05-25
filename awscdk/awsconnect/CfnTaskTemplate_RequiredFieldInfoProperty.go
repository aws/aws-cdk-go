package awsconnect


// Information about a required field.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   requiredFieldInfoProperty := &RequiredFieldInfoProperty{
//   	Id: &FieldIdentifierProperty{
//   		Name: jsii.String("name"),
//   	},
//   }
//
type CfnTaskTemplate_RequiredFieldInfoProperty struct {
	// The unique identifier for the field.
	Id interface{} `field:"required" json:"id" yaml:"id"`
}

