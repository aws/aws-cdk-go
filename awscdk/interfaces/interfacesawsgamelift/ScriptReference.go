package interfacesawsgamelift


// A reference to a Script resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scriptReference := &ScriptReference{
//   	ScriptArn: jsii.String("scriptArn"),
//   	ScriptId: jsii.String("scriptId"),
//   }
//
type ScriptReference struct {
	// The ARN of the Script resource.
	ScriptArn *string `field:"required" json:"scriptArn" yaml:"scriptArn"`
	// The Id of the Script resource.
	ScriptId *string `field:"required" json:"scriptId" yaml:"scriptId"`
}

