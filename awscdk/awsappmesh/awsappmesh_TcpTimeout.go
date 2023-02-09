package awsappmesh

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Represents timeouts for TCP protocols.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var duration duration
//
//   tcpTimeout := &tcpTimeout{
//   	idle: duration,
//   }
//
// Experimental.
type TcpTimeout struct {
	// Represents an idle timeout.
	//
	// The amount of time that a connection may be idle.
	// Experimental.
	Idle awscdk.Duration `field:"optional" json:"idle" yaml:"idle"`
}

