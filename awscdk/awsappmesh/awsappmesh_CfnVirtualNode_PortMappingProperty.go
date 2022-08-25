package awsappmesh


// An object representing a virtual node or virtual router listener port mapping.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   portMappingProperty := &portMappingProperty{
//   	port: jsii.Number(123),
//   	protocol: jsii.String("protocol"),
//   }
//
type CfnVirtualNode_PortMappingProperty struct {
	// The port used for the port mapping.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// The protocol used for the port mapping.
	//
	// Specify `http` , `http2` , `grpc` , or `tcp` .
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
}

