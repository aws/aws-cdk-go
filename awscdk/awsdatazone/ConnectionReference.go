package awsdatazone


// A reference to a Connection resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectionReference := &ConnectionReference{
//   	ConnectionId: jsii.String("connectionId"),
//   	DomainId: jsii.String("domainId"),
//   }
//
type ConnectionReference struct {
	// The ConnectionId of the Connection resource.
	ConnectionId *string `field:"required" json:"connectionId" yaml:"connectionId"`
	// The DomainId of the Connection resource.
	DomainId *string `field:"required" json:"domainId" yaml:"domainId"`
}

