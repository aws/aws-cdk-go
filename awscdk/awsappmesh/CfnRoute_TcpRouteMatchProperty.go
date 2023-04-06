package awsappmesh


// An object representing the TCP route to match.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tcpRouteMatchProperty := &TcpRouteMatchProperty{
//   	Port: jsii.Number(123),
//   }
//
type CfnRoute_TcpRouteMatchProperty struct {
	// The port number to match on.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

