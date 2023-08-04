package awsappmesh

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Represents timeouts for TCP protocols.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tcpTimeout := &TcpTimeout{
//   	Idle: cdk.Duration_Minutes(jsii.Number(30)),
//   }
//
type TcpTimeout struct {
	// Represents an idle timeout.
	//
	// The amount of time that a connection may be idle.
	// Default: - none.
	//
	Idle awscdk.Duration `field:"optional" json:"idle" yaml:"idle"`
}

