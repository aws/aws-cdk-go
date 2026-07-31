package awsmediaconnectalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for creating a Router Input.
//
// Example:
//   var stack Stack
//   var networkInterface RouterNetworkInterface
//
//
//   input := awsmediaconnectalpha.NewRouterInput(stack, jsii.String("FailoverInput"), &RouterInputProps{
//   	RouterInputName: jsii.String("failover-input"),
//   	MaximumBitrate: awscdk.Bitrate_Mbps(jsii.Number(10)),
//   	RoutingScope: awsmediaconnectalpha.RoutingScope_REGIONAL(),
//   	Tier: awsmediaconnectalpha.RouterInputTier_INPUT_50(),
//   	Configuration: awsmediaconnectalpha.RouterInputConfiguration_Failover(&FailoverConfigurationProps{
//   		NetworkInterface: networkInterface,
//   		Protocols: []RouterInputProtocol{
//   			awsmediaconnectalpha.RouterInputProtocol_Rist(&RistProtocolProps{
//   				Port: jsii.Number(5000),
//   				RecoveryLatency: awscdk.Duration_Millis(jsii.Number(1000)),
//   			}),
//   			awsmediaconnectalpha.RouterInputProtocol_*Rist(&RistProtocolProps{
//   				Port: jsii.Number(5002),
//   				 // Must not be consecutive with primary port
//   				RecoveryLatency: awscdk.Duration_*Millis(jsii.Number(1000)),
//   			}),
//   		},
//   		SourcePriority: awsmediaconnectalpha.SourcePriorityConfig_PrimarySecondary(awsmediaconnectalpha.PrimarySource_FIRST_SOURCE),
//   	}),
//   })
//
// Experimental.
type RouterInputProps struct {
	// Configuration for the Router Input (standard, failover, merge, or MediaConnect flow).
	// Experimental.
	Configuration RouterInputConfiguration `field:"required" json:"configuration" yaml:"configuration"`
	// Maximum bitrate in bits per second that the Router Input can handle.
	// Experimental.
	MaximumBitrate awscdk.Bitrate `field:"required" json:"maximumBitrate" yaml:"maximumBitrate"`
	// Routing scope for the Router Input.
	// Experimental.
	RoutingScope RoutingScope `field:"required" json:"routingScope" yaml:"routingScope"`
	// Maintenance window configuration.
	// Default: - Default maintenance window will be used.
	//
	// Experimental.
	MaintenanceConfiguration *MaintenanceConfiguration `field:"optional" json:"maintenanceConfiguration" yaml:"maintenanceConfiguration"`
	// AWS region where the Router Input will be created (i.e. us-east-1).
	//
	// Must match the region of the flows, flow outputs, and network interfaces it connects to —
	// MediaConnect rejects a cross-region connection at deploy.
	// Default: - Same as the stack's region.
	//
	// Experimental.
	RegionName *string `field:"optional" json:"regionName" yaml:"regionName"`
	// Name of the Router Input.
	// Default: - Generated automatically.
	//
	// Experimental.
	RouterInputName *string `field:"optional" json:"routerInputName" yaml:"routerInputName"`
	// Tags to add to the Router Input.
	// Default: - No tagging.
	//
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Select a tier based on your maximum bitrate requirements.
	// Default: RouterInputTier.INPUT_20
	//
	// Experimental.
	Tier RouterInputTier `field:"optional" json:"tier" yaml:"tier"`
	// Transit encryption configuration using AWS Secrets Manager.
	//
	// When provided without a role, a scoped IAM role is automatically created with read
	// access to the secret.
	// Default: - Automatic encryption will be configured.
	//
	// Experimental.
	TransitEncryption *TransitEncryption `field:"optional" json:"transitEncryption" yaml:"transitEncryption"`
}

