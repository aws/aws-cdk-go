package awsconnect


// Indicates a field that is read-only to an agent.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   readOnlyFieldInfoProperty := &readOnlyFieldInfoProperty{
//   	id: &fieldIdentifierProperty{
//   		name: jsii.String("name"),
//   	},
//   }
//
type CfnTaskTemplate_ReadOnlyFieldInfoProperty struct {
	// Identifier of the read-only field.
	Id interface{} `field:"required" json:"id" yaml:"id"`
}

