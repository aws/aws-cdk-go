package awsglue


// A reference to a Connection resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectionReference := &ConnectionReference{
//   	ConnectionId: jsii.String("connectionId"),
//   }
//
type ConnectionReference struct {
	// The Id of the Connection resource.
	ConnectionId *string `field:"required" json:"connectionId" yaml:"connectionId"`
}

