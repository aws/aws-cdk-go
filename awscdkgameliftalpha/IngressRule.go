// The CDK Construct Library for AWS::GameLift
package awscdkgameliftalpha


// A range of IP addresses and port settings that allow inbound traffic to connect to server processes on an instance in a fleet.
//
// New game sessions are assigned an IP address/port number combination, which must fall into the fleet's allowed ranges.
//
// Fleets with custom game builds must have permissions explicitly set.
// For Realtime Servers fleets, GameLift automatically opens two port ranges, one for TCP messaging and one for UDP.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import gamelift_alpha "github.com/aws/aws-cdk-go/awscdkgameliftalpha"
//
//   var peer iPeer
//   var port port
//
//   ingressRule := &IngressRule{
//   	Port: port,
//   	Source: peer,
//   }
//
// Experimental.
type IngressRule struct {
	// The port range used for ingress traffic.
	// Experimental.
	Port Port `field:"required" json:"port" yaml:"port"`
	// A range of allowed IP addresses .
	// Experimental.
	Source IPeer `field:"required" json:"source" yaml:"source"`
}

