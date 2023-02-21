// The CDK Construct Library for AWS::GameLift
package awscdkgameliftalpha


// Protocol for use in Connection Rules.
//
// https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml
// Experimental.
type Protocol string

const (
	// Experimental.
	Protocol_TCP Protocol = "TCP"
	// Experimental.
	Protocol_UDP Protocol = "UDP"
)

