package interfacesawsiot


// A reference to a Command resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   commandReference := &CommandReference{
//   	CommandArn: jsii.String("commandArn"),
//   	CommandId: jsii.String("commandId"),
//   }
//
type CommandReference struct {
	// The ARN of the Command resource.
	CommandArn *string `field:"required" json:"commandArn" yaml:"commandArn"`
	// The CommandId of the Command resource.
	CommandId *string `field:"required" json:"commandId" yaml:"commandId"`
}

