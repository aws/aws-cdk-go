package awsconnect


// A field that is invisible to an agent.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   invisibleFieldInfoProperty := &InvisibleFieldInfoProperty{
//   	Id: &FieldIdentifierProperty{
//   		Name: jsii.String("name"),
//   	},
//   }
//
type CfnTaskTemplate_InvisibleFieldInfoProperty struct {
	// Identifier of the invisible field.
	Id interface{} `field:"required" json:"id" yaml:"id"`
}

