package interfacesawscloudfront


// A reference to a ConnectionFunction resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectionFunctionReference := &ConnectionFunctionReference{
//   	ConnectionFunctionArn: jsii.String("connectionFunctionArn"),
//   	ConnectionFunctionId: jsii.String("connectionFunctionId"),
//   }
//
type ConnectionFunctionReference struct {
	// The ARN of the ConnectionFunction resource.
	ConnectionFunctionArn *string `field:"required" json:"connectionFunctionArn" yaml:"connectionFunctionArn"`
	// The Id of the ConnectionFunction resource.
	ConnectionFunctionId *string `field:"required" json:"connectionFunctionId" yaml:"connectionFunctionId"`
}

