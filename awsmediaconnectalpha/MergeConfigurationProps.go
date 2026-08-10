package awsmediaconnectalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for merge Router Input configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import mediaconnect_alpha "github.com/aws/aws-cdk-go/awsmediaconnectalpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var routerInputProtocol RouterInputProtocol
//   var routerNetworkInterface RouterNetworkInterface
//
//   mergeConfigurationProps := &MergeConfigurationProps{
//   	MergeRecoveryWindow: cdk.Duration_Minutes(jsii.Number(30)),
//   	NetworkInterface: routerNetworkInterface,
//   	Protocols: []RouterInputProtocol{
//   		routerInputProtocol,
//   	},
//
//   	// the properties below are optional
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   }
//
// Experimental.
type MergeConfigurationProps struct {
	// Recovery window for merge operation.
	// Experimental.
	MergeRecoveryWindow awscdk.Duration `field:"required" json:"mergeRecoveryWindow" yaml:"mergeRecoveryWindow"`
	// Network interface for the Router Input.
	// Experimental.
	NetworkInterface IRouterNetworkInterface `field:"required" json:"networkInterface" yaml:"networkInterface"`
	// Array of exactly 2 protocol configurations for merge (must be same non-SRT protocol type).
	// Experimental.
	Protocols *[]RouterInputProtocol `field:"required" json:"protocols" yaml:"protocols"`
	// The availability zone where the router input is located.
	// Default: - assigned by the MediaConnect service.
	//
	// Experimental.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
}

