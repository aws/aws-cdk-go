# AWS::MediaConnect Construct Library

<!--BEGIN STABILITY BANNER-->---


![cdk-constructs: Experimental](https://img.shields.io/badge/cdk--constructs-experimental-important.svg?style=for-the-badge)

> The APIs of higher level constructs in this module are experimental and under active development.
> They are subject to non-backward compatible changes or removal in any future version. These are
> not subject to the [Semantic Versioning](https://semver.org/) model and breaking changes will be
> announced in the release notes. This means that while you may use them, you may need to update
> your source code when upgrading to a newer version of this package.

---
<!--END STABILITY BANNER-->

## AWS Elemental MediaConnect

AWS Elemental MediaConnect is a high-quality transport service for live video. It provides the reliability and security of satellite and fiber-optic combined with the flexibility, agility, and economics of IP-based networks. MediaConnect enables you to build mission-critical live video workflows in a fraction of the time and cost of satellite or fiber services.

This package contains constructs for working with AWS Elemental MediaConnect, allowing you to define Flows, Bridges, Gateways, Router Inputs, Router Outputs, and Router Network Interfaces for transporting live video streams.

For further information on AWS Elemental MediaConnect, see [the documentation](https://aws.amazon.com/mediaconnect/).

## Table of Contents

* [Router Resources](#router-resources)

  * [Router Network Interfaces](#router-network-interfaces)
  * [Router Inputs](#router-inputs)
  * [Router Outputs](#router-outputs)
* [Flows](#flows)

  * [Creating a Flow](#creating-a-flow)
  * [Flow Sources](#flow-sources)
  * [VPC Interfaces](#vpc-interfaces)
  * [Media Streams](#media-streams)
  * [Flow Sizes](#flow-sizes)
* [Gateways](#gateways)

  * [Creating a Gateway](#creating-a-gateway)
  * [Importing an Existing Gateway](#importing-an-existing-gateway)
* [Bridges](#bridges)

  * [Creating a Bridge](#creating-a-bridge)
  * [Bridge Sources](#bridge-sources)
  * [Bridge Outputs](#bridge-outputs)
* [Encryption](#encryption)
* [CloudWatch Metrics](#cloudwatch-metrics)
* [Public CIDR warnings](#public-cidr-warnings)

## Router Resources

MediaConnect routers provide high-performance, low-latency video routing capabilities for building complex live video workflows. Router resources include network interfaces, inputs, and outputs.

### How Router Resources Relate

Router resources work together in a pipeline:

* A **RouterNetworkInterface** defines the network connectivity (public internet or VPC). Required for standard protocol-based inputs and outputs, but not needed when connecting to MediaLive inputs or MediaConnect flows directly.
* A **RouterInput** is the entry point — it receives video from a source via a protocol (RTP, SRT, RIST), or from a MediaConnect flow.
* A **RouterOutput** is the exit point — it sends video to a destination via a protocol, to a MediaLive input, or to a MediaConnect flow.

A typical camera-to-cloud workflow looks like:

```text
Camera → RouterNetworkInterface → RouterInput (SRT) → [Router] → RouterOutput (MediaLive) → MediaLive Channel
```

The router sits in the middle, routing content from inputs to outputs. You can have many outputs per input — for example, one camera feed going to both a MediaLive channel and a MediaConnect flow simultaneously.

### End-to-End Example: SRT Source to MediaLive

Here's a complete example showing how to connect an SRT source through a router to MediaLive:

```go
var stack Stack
var mediaLiveInput IInput


// 1. A public network interface for the SRT input
networkInterface := awsmediaconnectalpha.NewRouterNetworkInterface(stack, jsii.String("NetworkInterface"), &RouterNetworkInterfaceProps{
	RouterNetworkInterfaceName: jsii.String("camera-network"),
	Configuration: awsmediaconnectalpha.RouterNetworkConfiguration_PublicNetwork(&PublicNetworkConfigurationProps{
		Cidr: []*string{
			jsii.String("203.0.113.0/24"),
		},
	}),
})

// 2. A router input receiving SRT from an upstream encoder
input := awsmediaconnectalpha.NewRouterInput(stack, jsii.String("Input"), &RouterInputProps{
	RouterInputName: jsii.String("camera-input"),
	MaximumBitrate: awscdk.Bitrate_Mbps(jsii.Number(10)),
	RoutingScope: awsmediaconnectalpha.RoutingScope_REGIONAL(),
	Tier: awsmediaconnectalpha.RouterInputTier_INPUT_20(),
	Configuration: awsmediaconnectalpha.RouterInputConfiguration_Standard(&StandardConfigurationProps{
		NetworkInterface: *NetworkInterface,
		Protocol: awsmediaconnectalpha.RouterInputProtocol_SrtListener(&SrtListenerProtocolProps{
			Port: jsii.Number(9000),
			MinimumLatency: awscdk.Duration_Millis(jsii.Number(200)),
		}),
	}),
})

// 3. A router output delivering to MediaLive
output := awsmediaconnectalpha.NewRouterOutput(stack, jsii.String("Output"), &RouterOutputProps{
	RouterOutputName: jsii.String("medialive-output"),
	MaximumBitrate: awscdk.Bitrate_*Mbps(jsii.Number(10)),
	RoutingScope: awsmediaconnectalpha.RoutingScope_REGIONAL(),
	Tier: awsmediaconnectalpha.RouterOutputTier_OUTPUT_20(),
	Configuration: awsmediaconnectalpha.RouterOutputConfiguration_MediaLiveInput(&MediaLiveInputConnectionProps{
		Input: mediaLiveInput,
		Pipeline: awsmediaconnectalpha.MediaLivePipeline_PIPELINE_0,
	}),
})
```

This gives you a complete pipeline: the encoder pushes SRT to the network interface, the router receives it as an input, and the router output delivers it to MediaLive. For routing rules — including how bitrate affects which outputs can receive which inputs — see the [documentation](https://docs.aws.amazon.com/mediaconnect/latest/ug/using-router-control-panel.html).

### Router Network Interfaces

Network interfaces define the network connectivity for router inputs and outputs:

#### Public Network Interface

```go
var stack Stack


publicInterface := awsmediaconnectalpha.NewRouterNetworkInterface(stack, jsii.String("PublicInterface"), &RouterNetworkInterfaceProps{
	RouterNetworkInterfaceName: jsii.String("public-interface"),
	Configuration: awsmediaconnectalpha.RouterNetworkConfiguration_PublicNetwork(&PublicNetworkConfigurationProps{
		Cidr: []*string{
			jsii.String("203.0.113.0/24"),
		},
	}),
})
```

#### Private Network Interface

```go
var stack Stack
var securityGroup ISecurityGroup
var subnet ISubnet


privateInterface := awsmediaconnectalpha.NewRouterNetworkInterface(stack, jsii.String("PrivateInterface"), &RouterNetworkInterfaceProps{
	RouterNetworkInterfaceName: jsii.String("private-interface"),
	Configuration: awsmediaconnectalpha.RouterNetworkConfiguration_Vpc(&VpcNetworkConfigurationProps{
		SecurityGroups: []ISecurityGroup{
			securityGroup,
		},
		Subnet: subnet,
	}),
})
```

### Router Inputs

Router inputs receive live video streams from various sources and make them available for routing:

#### Standard Input with RTP Protocol

```go
var stack Stack
var networkInterface RouterNetworkInterface


input := awsmediaconnectalpha.NewRouterInput(stack, jsii.String("RtpInput"), &RouterInputProps{
	RouterInputName: jsii.String("rtp-input"),
	MaximumBitrate: awscdk.Bitrate_Mbps(jsii.Number(10)),
	RoutingScope: awsmediaconnectalpha.RoutingScope_REGIONAL(),
	// tier defaults to RouterInputTier.INPUT_20 (lowest cost)
	Configuration: awsmediaconnectalpha.RouterInputConfiguration_Standard(&StandardConfigurationProps{
		NetworkInterface: networkInterface,
		Protocol: awsmediaconnectalpha.RouterInputProtocol_Rtp(&RtpProtocolProps{
			Port: jsii.Number(5000),
		}),
	}),
})
```

#### Failover Input Configuration

```go
var stack Stack
var networkInterface RouterNetworkInterface


input := awsmediaconnectalpha.NewRouterInput(stack, jsii.String("FailoverInput"), &RouterInputProps{
	RouterInputName: jsii.String("failover-input"),
	MaximumBitrate: awscdk.Bitrate_Mbps(jsii.Number(10)),
	RoutingScope: awsmediaconnectalpha.RoutingScope_REGIONAL(),
	Tier: awsmediaconnectalpha.RouterInputTier_INPUT_50(),
	Configuration: awsmediaconnectalpha.RouterInputConfiguration_Failover(&FailoverConfigurationProps{
		NetworkInterface: networkInterface,
		Protocols: []RouterInputProtocol{
			awsmediaconnectalpha.RouterInputProtocol_Rist(&RistProtocolProps{
				Port: jsii.Number(5000),
				RecoveryLatency: awscdk.Duration_Millis(jsii.Number(1000)),
			}),
			awsmediaconnectalpha.RouterInputProtocol_*Rist(&RistProtocolProps{
				Port: jsii.Number(5002),
				 // Must not be consecutive with primary port
				RecoveryLatency: awscdk.Duration_*Millis(jsii.Number(1000)),
			}),
		},
		SourcePriority: awsmediaconnectalpha.SourcePriorityConfig_PrimarySecondary(awsmediaconnectalpha.PrimarySource_FIRST_SOURCE),
	}),
})
```

#### MediaLive Channel Input

Connect a router input to a MediaLive channel. The `outputName` must match the name of an output configured in the channel's MediaConnect Router output group.

```go
var stack Stack
var mediaLiveChannel IChannel


input := awsmediaconnectalpha.NewRouterInput(stack, jsii.String("ChannelInput"), &RouterInputProps{
	RouterInputName: jsii.String("channel-input"),
	MaximumBitrate: awscdk.Bitrate_Mbps(jsii.Number(20)),
	RoutingScope: awsmediaconnectalpha.RoutingScope_REGIONAL(),
	Tier: awsmediaconnectalpha.RouterInputTier_INPUT_50(),
	Configuration: awsmediaconnectalpha.RouterInputConfiguration_MediaLiveChannel(&MediaLiveChannelConfigurationProps{
		Channel: mediaLiveChannel,
		OutputName: jsii.String("router-ts"),
		Pipeline: awsmediaconnectalpha.MediaLivePipeline_PIPELINE_0,
	}),
})
```

> **Tip:** For full examples of wiring a MediaLive channel to a MediaConnect Router (including
> transit encryption), see the [MediaLive L2 README — MediaConnect Router section](./../aws-medialive-alpha/README.md#aws-elemental-mediaconnect-router).

```go
var stack Stack
var mediaLiveChannel IChannel
var transitSecret Secret
// must hold the same value as the channel's MediaConnectRouterSettings.shared() secret

input := awsmediaconnectalpha.NewRouterInput(stack, jsii.String("ChannelInput"), &RouterInputProps{
	RouterInputName: jsii.String("channel-input"),
	MaximumBitrate: awscdk.Bitrate_Mbps(jsii.Number(20)),
	RoutingScope: awsmediaconnectalpha.RoutingScope_REGIONAL(),
	Tier: awsmediaconnectalpha.RouterInputTier_INPUT_50(),
	Configuration: awsmediaconnectalpha.RouterInputConfiguration_MediaLiveChannel(&MediaLiveChannelConfigurationProps{
		Channel: mediaLiveChannel,
		OutputName: jsii.String("router-ts"),
		Pipeline: awsmediaconnectalpha.MediaLivePipeline_PIPELINE_0,
		SourceTransitDecryption: &TransitEncryption{
			Secret: transitSecret,
		},
	}),
})
```

Or prepare a router input for a MediaLive connection without specifying the channel (requires explicit availability zone):

```go
var stack Stack


input := awsmediaconnectalpha.NewRouterInput(stack, jsii.String("ChannelInputNoConnection"), &RouterInputProps{
	RouterInputName: jsii.String("channel-input-no-connection"),
	MaximumBitrate: awscdk.Bitrate_Mbps(jsii.Number(20)),
	RoutingScope: awsmediaconnectalpha.RoutingScope_REGIONAL(),
	Tier: awsmediaconnectalpha.RouterInputTier_INPUT_50(),
	Configuration: awsmediaconnectalpha.RouterInputConfiguration_MediaLiveChannelWithoutConnection(&MediaLiveChannelConfigurationWithoutConnectionProps{
		AvailabilityZone: jsii.String("us-east-1a"),
	}),
})
```

#### MediaConnect Flow Input

Connect a router input to an existing MediaConnect flow:

```go
var stack Stack
var flow Flow
var flowOutput FlowOutput


input := awsmediaconnectalpha.NewRouterInput(stack, jsii.String("FlowInput"), &RouterInputProps{
	RouterInputName: jsii.String("flow-input"),
	MaximumBitrate: awscdk.Bitrate_Mbps(jsii.Number(20)),
	RoutingScope: awsmediaconnectalpha.RoutingScope_REGIONAL(),
	Tier: awsmediaconnectalpha.RouterInputTier_INPUT_50(),
	Configuration: awsmediaconnectalpha.RouterInputConfiguration_MediaConnectFlow(&MediaConnectFlowConfigurationProps{
		Flow: flow,
		FlowOutput: flowOutput,
	}),
})
```

Or prepare a router input for a flow connection without specifying the flow (requires explicit availability zone):

```go
var stack Stack


input := awsmediaconnectalpha.NewRouterInput(stack, jsii.String("FlowInputNoConnection"), &RouterInputProps{
	RouterInputName: jsii.String("flow-input-no-connection"),
	MaximumBitrate: awscdk.Bitrate_Mbps(jsii.Number(20)),
	RoutingScope: awsmediaconnectalpha.RoutingScope_REGIONAL(),
	Tier: awsmediaconnectalpha.RouterInputTier_INPUT_50(),
	Configuration: awsmediaconnectalpha.RouterInputConfiguration_MediaConnectFlowWithoutConnection(&MediaConnectFlowConfigurationWithoutConnectionProps{
		AvailabilityZone: jsii.String("us-east-1a"),
	}),
})
```

### Router Outputs

Router outputs send video streams to various destinations including standard protocols, MediaLive inputs, and MediaConnect flows:

#### Standard Output with SRT Protocol

```go
var stack Stack
var networkInterface RouterNetworkInterface


output := awsmediaconnectalpha.NewRouterOutput(stack, jsii.String("SrtOutput"), &RouterOutputProps{
	RouterOutputName: jsii.String("srt-output"),
	MaximumBitrate: awscdk.Bitrate_Mbps(jsii.Number(10)),
	RoutingScope: awsmediaconnectalpha.RoutingScope_REGIONAL(),
	// tier defaults to RouterOutputTier.OUTPUT_20 (lowest cost)
	Configuration: awsmediaconnectalpha.RouterOutputConfiguration_Standard(&StandardOutputConfigurationProps{
		Protocol: awsmediaconnectalpha.RouterOutputProtocol_SrtListener(&SrtListenerOutputProtocolProps{
			Port: jsii.Number(9001),
			MinimumLatency: awscdk.Duration_Millis(jsii.Number(200)),
		}),
		NetworkInterface: networkInterface,
	}),
})
```

> **Note:** The `tier` property defaults to the lowest (and cheapest) tier: `INPUT_20` for Router Inputs and `OUTPUT_20` for Router Outputs. The construct validates that `maximumBitrate` does not exceed the tier's capacity (20, 50, or 100 Mbps) at synth time. Per the [documentation](https://docs.aws.amazon.com/mediaconnect/latest/ug/using-router-control-panel.html), if an input is 20 Mbps you can't route it to an output set up for less than 20 Mbps.

#### MediaLive Output

Connect a router output to a MediaLive input (the input must be a MediaConnect Router type):

```go
var stack Stack
var mediaLiveInput IInput


output := awsmediaconnectalpha.NewRouterOutput(stack, jsii.String("MediaLiveOutput"), &RouterOutputProps{
	RouterOutputName: jsii.String("medialive-output"),
	MaximumBitrate: awscdk.Bitrate_Mbps(jsii.Number(15)),
	RoutingScope: awsmediaconnectalpha.RoutingScope_GLOBAL(),
	Tier: awsmediaconnectalpha.RouterOutputTier_OUTPUT_50(),
	Configuration: awsmediaconnectalpha.RouterOutputConfiguration_MediaLiveInput(&MediaLiveInputConnectionProps{
		Input: mediaLiveInput,
		Pipeline: awsmediaconnectalpha.MediaLivePipeline_PIPELINE_0,
	}),
})
```

Or prepare a router output for a MediaLive connection without specifying the input (requires explicit availability zone):

```go
var stack Stack


output := awsmediaconnectalpha.NewRouterOutput(stack, jsii.String("MediaLiveOutputNoConnection"), &RouterOutputProps{
	RouterOutputName: jsii.String("medialive-output-no-connection"),
	MaximumBitrate: awscdk.Bitrate_Mbps(jsii.Number(15)),
	RoutingScope: awsmediaconnectalpha.RoutingScope_GLOBAL(),
	Tier: awsmediaconnectalpha.RouterOutputTier_OUTPUT_50(),
	Configuration: awsmediaconnectalpha.RouterOutputConfiguration_MediaLiveInputWithoutConnection(&MediaLiveNoInputConnectionProps{
		AvailabilityZone: jsii.String("us-east-1a"),
	}),
})
```

#### MediaConnect Flow Output

Connect a router output to an existing MediaConnect flow:

```go
var stack Stack
var flow Flow


output := awsmediaconnectalpha.NewRouterOutput(stack, jsii.String("FlowOutput"), &RouterOutputProps{
	RouterOutputName: jsii.String("flow-output"),
	MaximumBitrate: awscdk.Bitrate_Mbps(jsii.Number(20)),
	RoutingScope: awsmediaconnectalpha.RoutingScope_REGIONAL(),
	Tier: awsmediaconnectalpha.RouterOutputTier_OUTPUT_100(),
	Configuration: awsmediaconnectalpha.RouterOutputConfiguration_MediaConnectFlow(&MediaConnectFlowConnectionProps{
		Flow: flow,
	}),
})
```

Or prepare a router output for a flow connection without specifying the flow (requires explicit availability zone):

```go
var stack Stack


output := awsmediaconnectalpha.NewRouterOutput(stack, jsii.String("FlowOutputNoConnection"), &RouterOutputProps{
	RouterOutputName: jsii.String("flow-output-no-connection"),
	MaximumBitrate: awscdk.Bitrate_Mbps(jsii.Number(20)),
	RoutingScope: awsmediaconnectalpha.RoutingScope_REGIONAL(),
	Tier: awsmediaconnectalpha.RouterOutputTier_OUTPUT_100(),
	Configuration: awsmediaconnectalpha.RouterOutputConfiguration_MediaConnectFlowWithoutConnection(&MediaConnectFlowNoConnectionProps{
		AvailabilityZone: jsii.String("us-east-1a"),
	}),
})
```

#### Output with Encryption

```go
var stack Stack
var networkInterface RouterNetworkInterface
var role IRole
var secret ISecret


output := awsmediaconnectalpha.NewRouterOutput(stack, jsii.String("EncryptedOutput"), &RouterOutputProps{
	RouterOutputName: jsii.String("encrypted-output"),
	MaximumBitrate: awscdk.Bitrate_Mbps(jsii.Number(10)),
	RoutingScope: awsmediaconnectalpha.RoutingScope_REGIONAL(),
	Tier: awsmediaconnectalpha.RouterOutputTier_OUTPUT_50(),
	Configuration: awsmediaconnectalpha.RouterOutputConfiguration_Standard(&StandardOutputConfigurationProps{
		Protocol: awsmediaconnectalpha.RouterOutputProtocol_SrtCaller(&SrtCallerOutputProtocolProps{
			DestinationAddress: jsii.String("203.0.113.100"),
			DestinationPort: jsii.Number(9001),
			MinimumLatency: awscdk.Duration_Millis(jsii.Number(200)),
			EncryptionConfiguration: &RouterSrtEncryption{
				Role: *Role,
				Secret: *Secret,
			},
		}),
		NetworkInterface: networkInterface,
	}),
})
```

## Flows

A MediaConnect flow represents a transport stream connection between a source and one or more outputs. Flows are the primary resource for transporting live video content.

### Creating a Flow

The following example creates a basic MediaConnect flow with an RTP source:

```go
var stack Stack


flow := awsmediaconnectalpha.NewFlow(stack, jsii.String("MyFlow"), &FlowProps{
	FlowName: jsii.String("my-live-stream"),
	Source: awsmediaconnectalpha.SourceConfiguration_Rtp(&SourceRtp{
		FlowSourceName: jsii.String("my-source"),
		Port: jsii.Number(5000),
		Network: awsmediaconnectalpha.NetworkConfiguration_PublicNetwork(jsii.String("203.0.113.0/24")),
	}),
})
```

### Flow Sources

MediaConnect supports multiple source types for ingesting content into a flow. The examples below use `NetworkConfiguration.publicNetwork()` for simplicity, but all protocol-based sources can also use `NetworkConfiguration.vpc()` with a VPC interface for private connectivity.

#### SRT Listener Source

SRT (Secure Reliable Transport) in listener mode configures MediaConnect to listen on a specific port for incoming content. The upstream device connects to MediaConnect as a caller.

```go
var stack Stack


flow := awsmediaconnectalpha.NewFlow(stack, jsii.String("MyFlow"), &FlowProps{
	Source: awsmediaconnectalpha.SourceConfiguration_SrtListener(&SourceSrtListener{
		FlowSourceName: jsii.String("live-encoder-source"),
		Description: jsii.String("Live encoder feed"),
		Port: jsii.Number(5000),
		MinLatency: awscdk.Duration_Millis(jsii.Number(2000)),
		Network: awsmediaconnectalpha.NetworkConfiguration_PublicNetwork(jsii.String("203.0.113.0/24")),
	}),
})
```

#### SRT Caller Source

SRT in caller mode configures MediaConnect to connect to a remote SRT listener. Use this when the source device is listening for incoming connections rather than pushing content.

```go
var stack Stack


flow := awsmediaconnectalpha.NewFlow(stack, jsii.String("MyFlow"), &FlowProps{
	Source: awsmediaconnectalpha.SourceConfiguration_SrtCaller(&SourceSrtCaller{
		FlowSourceName: jsii.String("remote-source"),
		SourceListenerAddress: jsii.String("203.0.113.50"),
		SourceListenerPort: jsii.Number(5000),
		MinLatency: awscdk.Duration_Millis(jsii.Number(200)),
	}),
})
```

#### RTP Source

RTP (Real-time Transport Protocol) is a standard protocol for delivering audio and video over IP networks.

```go
var stack Stack


flow := awsmediaconnectalpha.NewFlow(stack, jsii.String("MyFlow"), &FlowProps{
	Source: awsmediaconnectalpha.SourceConfiguration_Rtp(&SourceRtp{
		FlowSourceName: jsii.String("rtp-source"),
		Port: jsii.Number(5000),
		Network: awsmediaconnectalpha.NetworkConfiguration_PublicNetwork(jsii.String("203.0.113.0/24")),
	}),
})
```

#### RTP-FEC Source

RTP with Forward Error Correction adds redundancy to recover lost packets without retransmission. Use this when contributing via RTP and you need packet recovery.

```go
var stack Stack


flow := awsmediaconnectalpha.NewFlow(stack, jsii.String("MyFlow"), &FlowProps{
	Source: awsmediaconnectalpha.SourceConfiguration_RtpFec(&SourceRtp{
		FlowSourceName: jsii.String("rtp-fec-source"),
		Port: jsii.Number(5000),
		Network: awsmediaconnectalpha.NetworkConfiguration_PublicNetwork(jsii.String("203.0.113.0/24")),
	}),
})
```

#### RIST Source

RIST (Reliable Internet Stream Transport) provides reliable video transport with packet recovery.

```go
var stack Stack


flow := awsmediaconnectalpha.NewFlow(stack, jsii.String("MyFlow"), &FlowProps{
	Source: awsmediaconnectalpha.SourceConfiguration_Rist(&SourceRist{
		FlowSourceName: jsii.String("rist-source"),
		Port: jsii.Number(5000),
		MaxLatency: awscdk.Duration_Millis(jsii.Number(2000)),
		Network: awsmediaconnectalpha.NetworkConfiguration_PublicNetwork(jsii.String("203.0.113.0/24")),
	}),
})
```

#### Zixi Push Source

Zixi Push uses the Zixi protocol for reliable video transport. Content is pushed to MediaConnect from a Zixi-compatible component upstream.

```go
var stack Stack


flow := awsmediaconnectalpha.NewFlow(stack, jsii.String("MyFlow"), &FlowProps{
	Source: awsmediaconnectalpha.SourceConfiguration_ZixiPush(&SourceZixiPush{
		FlowSourceName: jsii.String("zixi-source"),
		MaxLatency: awscdk.Duration_Millis(jsii.Number(2000)),
		Network: awsmediaconnectalpha.NetworkConfiguration_PublicNetwork(jsii.String("203.0.113.0/24")),
	}),
})
```

> Zixi Push ports are assigned by MediaConnect: 2088 for public sources, 2090-2099 for VPC sources. See [Source port assignments](https://docs.aws.amazon.com/mediaconnect/latest/ug/source-ports.html).

#### Router Source

Use a router source when the flow's source comes from a MediaConnect Router rather than a direct connection.

```go
var stack Stack


flow := awsmediaconnectalpha.NewFlow(stack, jsii.String("MyFlow"), &FlowProps{
	Source: awsmediaconnectalpha.SourceConfiguration_Router(),
})
```

#### VPC Source

Use a VPC-based source when you need a connection between a flow and your Amazon VPC. This enables private connectivity for receiving content from on-premises equipment via AWS Direct Connect or VPN, or from other AWS services running in your VPC.

```go
var stack Stack
var securityGroup ISecurityGroup
var subnet ISubnet
var role IRole


vpcInterface := awsmediaconnectalpha.VpcInterface_Define(&VpcInterfaceDefineProps{
	VpcInterfaceName: jsii.String("my-vpc-interface"),
	Role: role,
	SecurityGroups: []ISecurityGroup{
		securityGroup,
	},
	Subnet: subnet,
})

flow := awsmediaconnectalpha.NewFlow(stack, jsii.String("MyFlow"), &FlowProps{
	Source: awsmediaconnectalpha.SourceConfiguration_Rist(&SourceRist{
		FlowSourceName: jsii.String("vpc-source"),
		Description: jsii.String("VPC-based source"),
		Port: jsii.Number(5000),
		MaxLatency: awscdk.Duration_Millis(jsii.Number(2000)),
		Network: awsmediaconnectalpha.NetworkConfiguration_Vpc(vpcInterface),
	}),
	VpcInterfaces: []VpcInterfaceConfig{
		vpcInterface,
	},
})
```

#### Entitled Source (From Another AWS Account)

Entitlements allow you to subscribe to content from another AWS account. The entitlement is created by the content originator in their AWS account, and you import it using the entitlement ARN they provide.

```go
var stack Stack


// Import an entitlement from another AWS account
entitlement := awsmediaconnectalpha.FlowEntitlement_FromFlowEntitlementArn(stack, jsii.String("ImportedEntitlement"), jsii.String("arn:aws:mediaconnect:us-west-2:111122223333:entitlement:1-11111111111111111111111111111111:MyEntitlement"))

flow := awsmediaconnectalpha.NewFlow(stack, jsii.String("MyFlow"), &FlowProps{
	Source: awsmediaconnectalpha.SourceConfiguration_Entitlement(&EntitlementSource{
		Entitlement: entitlement,
	}),
})
```

#### Gateway Bridge Source

Use a gateway bridge source when ingesting content from on-premises equipment through a MediaConnect gateway and bridge. Gateways define the network infrastructure that bridges use to transport video between on-premises and cloud environments.

```go
var stack Stack
var bridge Bridge
var role IRole
var securityGroup ISecurityGroup
var subnet ISubnet


vpcInterface := awsmediaconnectalpha.VpcInterface_Define(&VpcInterfaceDefineProps{
	VpcInterfaceName: jsii.String("bridge-interface"),
	Role: role,
	SecurityGroups: []ISecurityGroup{
		securityGroup,
	},
	Subnet: subnet,
})

flow := awsmediaconnectalpha.NewFlow(stack, jsii.String("MyFlow"), &FlowProps{
	Source: awsmediaconnectalpha.SourceConfiguration_GatewayBridge(&GatewayBridgeSource{
		Bridge: bridge,
		VpcInterface: vpcInterface,
	}),
	VpcInterfaces: []VpcInterfaceConfig{
		vpcInterface,
	},
})
```

### VPC Interfaces

VPC interfaces allow MediaConnect to send or receive content within your VPC. Create VPC interfaces using `VpcInterface.define()` and add them to the flow's `vpcInterfaces` array. The same interface can then be referenced in sources and outputs:

```go
var stack Stack
var role IRole
var securityGroup ISecurityGroup
var subnet ISubnet


// Create VPC interface
vpcInterface := awsmediaconnectalpha.VpcInterface_Define(&VpcInterfaceDefineProps{
	VpcInterfaceName: jsii.String("my-vpc-interface"),
	Role: role,
	SecurityGroups: []ISecurityGroup{
		securityGroup,
	},
	Subnet: subnet,
	NetworkInterfaceType: awsmediaconnectalpha.NetworkInterface_ENA(),
})

// Add to flow and reference in source
flow := awsmediaconnectalpha.NewFlow(stack, jsii.String("MyFlow"), &FlowProps{
	VpcInterfaces: []VpcInterfaceConfig{
		vpcInterface,
	},
	 // Declare at flow level
	Source: awsmediaconnectalpha.SourceConfiguration_Rist(&SourceRist{
		FlowSourceName: jsii.String("vpc-source"),
		Port: jsii.Number(5000),
		MaxLatency: awscdk.Duration_Millis(jsii.Number(2000)),
		Network: awsmediaconnectalpha.NetworkConfiguration_Vpc(vpcInterface),
	}),
})
```

#### CDI and JPEG XS with EFA Interfaces

For high-performance CDI or JPEG XS workflows, use EFA (Elastic Fabric Adapter) interfaces. Note that flows can have a maximum of 1 EFA interface:

```go
var stack Stack
var role IRole
var securityGroup ISecurityGroup
var subnet ISubnet


efaInterface := awsmediaconnectalpha.VpcInterface_Define(&VpcInterfaceDefineProps{
	VpcInterfaceName: jsii.String("efa-interface"),
	Role: role,
	SecurityGroups: []ISecurityGroup{
		securityGroup,
	},
	Subnet: subnet,
	NetworkInterfaceType: awsmediaconnectalpha.NetworkInterface_EFA(),
})

videoStream := awsmediaconnectalpha.MediaStream_Video(&MediaStreamVideo{
	MediaStreamId: jsii.Number(1),
	MediaStreamName: jsii.String("video"),
	VideoFormat: awsmediaconnectalpha.MediaVideoFormat_HD_1080P(),
	Fmtp: &FmtpVideo{
		ExactFramerate: awsmediaconnectalpha.Framerate_FPS_29_97(),
		Par: awsmediaconnectalpha.PixelAspectRatio_SQUARE(),
	},
})

flow := awsmediaconnectalpha.NewFlow(stack, jsii.String("MyCdiFlow"), &FlowProps{
	FlowSize: awsmediaconnectalpha.FlowSize_LARGE_4X(),
	 // Required for CDI and JPEG XS
	VpcInterfaces: []VpcInterfaceConfig{
		efaInterface,
	},
	MediaStreams: []MediaStream{
		videoStream,
	},
	Source: awsmediaconnectalpha.SourceConfiguration_Cdi(&SourceCdi{
		FlowSourceName: jsii.String("cdi-source"),
		VpcInterface: efaInterface,
		Port: jsii.Number(5000),
		MaxSyncBuffer: jsii.Number(100),
		MediaStreamSourceConfigurations: []MediaStreamSourceConfigurationCdi{
			&MediaStreamSourceConfigurationCdi{
				Encoding: awsmediaconnectalpha.Encoding_RAW(),
				MediaStream: videoStream,
			},
		},
	}),
})
```

#### JPEG XS with Redundant Interfaces

JPEG XS requires exactly 2 input interfaces per media stream for redundancy. Typically one EFA and one ENA interface:

```go
var stack Stack
var role IRole
var sg1 ISecurityGroup
var sg2 ISecurityGroup
var subnet ISubnet


efaInterface := awsmediaconnectalpha.VpcInterface_Define(&VpcInterfaceDefineProps{
	VpcInterfaceName: jsii.String("efa-interface"),
	Role: role,
	SecurityGroups: []ISecurityGroup{
		sg1,
	},
	Subnet: subnet,
	NetworkInterfaceType: awsmediaconnectalpha.NetworkInterface_EFA(),
})

enaInterface := awsmediaconnectalpha.VpcInterface_Define(&VpcInterfaceDefineProps{
	VpcInterfaceName: jsii.String("ena-interface"),
	Role: role,
	SecurityGroups: []ISecurityGroup{
		sg2,
	},
	Subnet: subnet,
	NetworkInterfaceType: awsmediaconnectalpha.NetworkInterface_ENA(),
})

videoStream := awsmediaconnectalpha.MediaStream_Video(&MediaStreamVideo{
	MediaStreamId: jsii.Number(1),
	MediaStreamName: jsii.String("video"),
	VideoFormat: awsmediaconnectalpha.MediaVideoFormat_UHD_2160P(),
	Fmtp: &FmtpVideo{
		ExactFramerate: awsmediaconnectalpha.Framerate_FPS_59_94(),
		Par: awsmediaconnectalpha.PixelAspectRatio_SQUARE(),
		Colorimetry: awsmediaconnectalpha.Colorimetry_BT2020(),
		VideoRange: awsmediaconnectalpha.VideoRange_FULL(),
		ScanMode: awsmediaconnectalpha.ScanMode_PROGRESSIVE(),
		Tcs: awsmediaconnectalpha.Tcs_PQ(),
	},
})

flow := awsmediaconnectalpha.NewFlow(stack, jsii.String("MyJpegXsFlow"), &FlowProps{
	FlowSize: awsmediaconnectalpha.FlowSize_LARGE_4X(),
	 // Required for JPEG XS
	VpcInterfaces: []VpcInterfaceConfig{
		efaInterface,
		enaInterface,
	},
	MediaStreams: []MediaStream{
		videoStream,
	},
	Source: awsmediaconnectalpha.SourceConfiguration_JpegXs(&SourceJpegXs{
		FlowSourceName: jsii.String("jpegxs-source"),
		MaxSyncBuffer: jsii.Number(100),
		MediaStreamSourceConfigurations: []MediaStreamSourceConfigurationJpegXs{
			&MediaStreamSourceConfigurationJpegXs{
				Encoding: awsmediaconnectalpha.Encoding_JXSV(),
				Port: jsii.Number(5000),
				InputInterface: []VpcInterfaceConfig{
					efaInterface,
					enaInterface,
				},
				 // 2 interfaces for redundancy
				MediaStream: videoStream,
			},
		},
	}),
})
```

### Media Streams

Media streams represent individual components of your content (video, audio, ancillary data) for ST 2110 JPEG XS or CDI workflows. Create media streams using the static factory methods and add them to the flow's `mediaStreams` array:

```go
var stack Stack


// Create media streams
videoStream := awsmediaconnectalpha.MediaStream_Video(&MediaStreamVideo{
	MediaStreamId: jsii.Number(1),
	MediaStreamName: jsii.String("video-stream"),
	VideoFormat: awsmediaconnectalpha.MediaVideoFormat_HD_1080P(),
	Fmtp: &FmtpVideo{
		Colorimetry: awsmediaconnectalpha.Colorimetry_BT709(),
		ExactFramerate: awsmediaconnectalpha.Framerate_FPS_29_97(),
		Par: awsmediaconnectalpha.PixelAspectRatio_SQUARE(),
		VideoRange: awsmediaconnectalpha.VideoRange_NARROW(),
		ScanMode: awsmediaconnectalpha.ScanMode_PROGRESSIVE(),
		Tcs: awsmediaconnectalpha.Tcs_SDR(),
	},
})

audioStream := awsmediaconnectalpha.MediaStream_Audio(&MediaStreamAudio{
	MediaStreamId: jsii.Number(2),
	MediaStreamName: jsii.String("audio-stream"),
	ChannelOrder: awsmediaconnectalpha.AudioStreamOrderOptions_STANDARD_STEREO(),
})

// Add to flow
flow := awsmediaconnectalpha.NewFlow(stack, jsii.String("MyFlow"), &FlowProps{
	Source: awsmediaconnectalpha.SourceConfiguration_Router(),
	MediaStreams: []MediaStream{
		videoStream,
		audioStream,
	},
})
```

#### Audio Channel Order

For audio media streams, use the `AudioStreamOrderOptions` enum to specify the SMPTE 2110-30 channel order:

```go
// Available channel order options
awsmediaconnectalpha.AudioStreamOrderOptions_MONO() // SMPTE2110.(M)
awsmediaconnectalpha.AudioStreamOrderOptions_DUAL_MONO() // SMPTE2110.(DM)
awsmediaconnectalpha.AudioStreamOrderOptions_STANDARD_STEREO() // SMPTE2110.(ST)
awsmediaconnectalpha.AudioStreamOrderOptions_LTRT_MATRIX_STEREO() // SMPTE2110.(LtRt)
awsmediaconnectalpha.AudioStreamOrderOptions_SURROUND_5_1() // SMPTE2110.(51)
awsmediaconnectalpha.AudioStreamOrderOptions_SURROUND_7_1() // SMPTE2110.(71)
awsmediaconnectalpha.AudioStreamOrderOptions_SURROUND_22_2() // SMPTE2110.(222)
awsmediaconnectalpha.AudioStreamOrderOptions_ONE_SDI_AUDIO_GROUP() // SMPTE2110.(SGRP)

// Example with 5.1 surround
surroundAudio := awsmediaconnectalpha.MediaStream_Audio(&MediaStreamAudio{
	MediaStreamId: jsii.Number(3),
	MediaStreamName: jsii.String("surround-audio"),
	ChannelOrder: awsmediaconnectalpha.AudioStreamOrderOptions_SURROUND_5_1(),
})
```

Media streams can be referenced in source configurations (for CDI and JPEG XS) and output configurations.

### Flow Sizes

MediaConnect offers three flow sizes that determine feature support:

| Flow Size | Transport Streams | NDI | CDI / JPEG XS |
|-----------|------------------|-----|---------------|
| `MEDIUM` (default) | ✅ | ❌ | ❌ |
| `LARGE` | ✅ | ✅ | ❌ |
| `LARGE_4X` | ❌ | ❌ | ✅ |

This table maps each `FlowSize` to the capabilities the construct validates. For output counts, throughput limits, and other per-size details, see [Flow sizes and capabilities](https://docs.aws.amazon.com/mediaconnect/latest/ug/flow-sizes-capabilities.html).

The construct validates flow size constraints at synthesis time based on the source protocol and NDI configuration:

* `MEDIUM` supports transport stream protocols (RTP, SRT, RIST, etc.) but not NDI or CDI
* `LARGE` supports transport streams and NDI, and is required when NDI is enabled
* `LARGE_4X` is required for CDI and JPEG XS protocols, and does not support transport streams or NDI

These are mutually exclusive — CDI/JPEG XS and NDI cannot coexist on the same flow because they require different flow sizes.

```go
var stack Stack
var ndiVpcInterface VpcInterfaceConfig
var efaInterface VpcInterfaceConfig
var videoStream MediaStream


// NDI requires LARGE, an encoding profile, and at least one discovery server
// NDI requires LARGE, an encoding profile, and at least one discovery server
awsmediaconnectalpha.NewFlow(stack, jsii.String("NdiFlow"), &FlowProps{
	FlowSize: awsmediaconnectalpha.FlowSize_LARGE(),
	NdiConfig: &NdiConfig{
		NdiState: awsmediaconnectalpha.State_ENABLED,
		NdiDiscoveryServers: []NdiDiscoveryServerConfig{
			&NdiDiscoveryServerConfig{
				DiscoveryServerAddress: jsii.String("10.0.0.10"),
				VpcInterface: ndiVpcInterface,
			},
		},
	},
	EncodingConfig: &EncodingConfig{
		EncodingProfile: awsmediaconnectalpha.EncodingProfile_CONTRIBUTION_H264_DEFAULT,
	},
	Source: awsmediaconnectalpha.SourceConfiguration_Ndi(&SourceNdi{
		FlowSourceName: jsii.String("ndi-source"),
	}),
})

// CDI and JPEG XS require LARGE_4X
// CDI and JPEG XS require LARGE_4X
awsmediaconnectalpha.NewFlow(stack, jsii.String("CdiFlow"), &FlowProps{
	FlowSize: awsmediaconnectalpha.FlowSize_LARGE_4X(),
	VpcInterfaces: []VpcInterfaceConfig{
		efaInterface,
	},
	MediaStreams: []MediaStream{
		videoStream,
	},
	Source: awsmediaconnectalpha.SourceConfiguration_Cdi(&SourceCdi{
		FlowSourceName: jsii.String("cdi-source"),
		VpcInterface: efaInterface,
		Port: jsii.Number(5000),
		MaxSyncBuffer: jsii.Number(100),
		MediaStreamSourceConfigurations: []MediaStreamSourceConfigurationCdi{
			&MediaStreamSourceConfigurationCdi{
				Encoding: awsmediaconnectalpha.Encoding_RAW(),
				MediaStream: videoStream,
			},
		},
	}),
})
```

For more information, see [Flow sizes and capabilities](https://docs.aws.amazon.com/mediaconnect/latest/ug/flow-sizes-capabilities.html).

## Gateways

MediaConnect Gateways enable the deployment of on-premises resources for transporting live video to and from the AWS Cloud. Gateways are required for creating bridges.

### Creating a Gateway

```go
var stack Stack


productionNetwork := awsmediaconnectalpha.GatewayNetwork_Define(&GatewayNetworkDefineProps{
	CidrBlock: jsii.String("192.168.1.0/24"),
	Name: jsii.String("production-network"),
})

gateway := awsmediaconnectalpha.NewGateway(stack, jsii.String("MyGateway"), &GatewayProps{
	GatewayName: jsii.String("my-gateway"),
	EgressCidrBlocks: []*string{
		jsii.String("10.0.0.0/16"),
	},
	Networks: []GatewayNetwork{
		productionNetwork,
	},
})
```

### Importing an Existing Gateway

```go
var stack Stack


gateway := awsmediaconnectalpha.Gateway_FromGatewayArn(stack, jsii.String("ImportedGateway"), jsii.String("arn:aws:mediaconnect:us-west-2:123456789012:gateway:1-XXXXXX"))
```

## Bridges

MediaConnect bridges enable you to interconnect on-premises equipment with cloud-based workflows. Bridges support both ingress (on-premises to cloud) and egress (cloud to on-premises) scenarios.

### Creating a Bridge

#### Ingress Bridge (On-premises to Cloud)

An ingress bridge receives content from on-premises equipment and makes it available in the cloud:

```go
var stack Stack


productionNetwork := awsmediaconnectalpha.GatewayNetwork_Define(&GatewayNetworkDefineProps{
	CidrBlock: jsii.String("192.168.1.0/24"),
	Name: jsii.String("production-network"),
})

gateway := awsmediaconnectalpha.NewGateway(stack, jsii.String("MyGateway"), &GatewayProps{
	GatewayName: jsii.String("my-gateway"),
	EgressCidrBlocks: []*string{
		jsii.String("10.0.0.0/16"),
	},
	Networks: []GatewayNetwork{
		productionNetwork,
	},
})

ingressBridge := awsmediaconnectalpha.NewBridge(stack, jsii.String("MyIngressBridge"), &BridgeProps{
	BridgeName: jsii.String("my-ingress-bridge"),
	Config: awsmediaconnectalpha.BridgeConfiguration_Ingress(&IngressBridgeConfiguration{
		MaxBitrate: awscdk.Bitrate_Mbps(jsii.Number(10)),
		MaxOutputs: jsii.Number(2),
		NetworkSources: []BridgeNetworkInput{
			&BridgeNetworkInput{
				Name: jsii.String("on-prem-source"),
				Source: &BridgeNetworkSource{
					Protocol: awsmediaconnectalpha.BridgeProtocol_RTP(),
					Network: productionNetwork,
					MulticastIp: jsii.String("239.1.1.1"),
					Port: jsii.Number(5000),
				},
			},
		},
	}),
	Gateway: gateway,
})
```

#### Egress Bridge (Cloud to On-premises)

An egress bridge sends content from MediaConnect flows to on-premises equipment:

```go
var stack Stack
var gateway Gateway
var flow Flow
var vpcInterface VpcInterfaceConfig
var productionNetwork GatewayNetwork


egressBridge := awsmediaconnectalpha.NewBridge(stack, jsii.String("MyEgressBridge"), &BridgeProps{
	BridgeName: jsii.String("my-egress-bridge"),
	Config: awsmediaconnectalpha.BridgeConfiguration_Egress(&EgressBridgeConfiguration{
		MaxBitrate: awscdk.Bitrate_Mbps(jsii.Number(10)),
		FlowSources: []BridgeFlowInput{
			&BridgeFlowInput{
				Name: jsii.String("cloud-source"),
				Source: &BridgeFlowSource{
					Flow: flow,
					VpcInterface: vpcInterface,
				},
			},
		},
		NetworkOutputs: []BridgeNetworkOutput{
			&BridgeNetworkOutput{
				Name: jsii.String("on-prem-output"),
				Output: awsmediaconnectalpha.BridgeOutputConfiguration_Network(&BridgeNetworkOutputProps{
					IpAddress: jsii.String("192.168.1.200"),
					Port: jsii.Number(5001),
					Network: productionNetwork,
					Protocol: awsmediaconnectalpha.BridgeProtocol_RTP(),
					Ttl: jsii.Number(50),
				}),
			},
		},
	}),
	Gateway: gateway,
})
```

### Bridge Sources

For failover scenarios, you can add additional sources to an existing bridge using the `BridgeSource` construct:

```go
var stack Stack
var bridge Bridge
var flow Flow


// Add a flow source to an egress bridge (requires failover to be enabled)
additionalSource := awsmediaconnectalpha.NewBridgeSource(stack, jsii.String("AdditionalSource"), &BridgeSourceProps{
	BridgeSourceName: jsii.String("backup-source"),
	Bridge: bridge,
	Source: awsmediaconnectalpha.BridgeSourceConfiguration_Flow(&BridgeFlowSource{
		Flow: flow,
	}),
})
```

### Bridge Outputs

Bridge outputs are configured as part of the bridge configuration for egress bridges. They define where content exits the bridge to on-premises equipment.

```go
var productionNetwork GatewayNetwork


networkConfig := awsmediaconnectalpha.BridgeOutputConfiguration_Network(&BridgeNetworkOutputProps{
	IpAddress: jsii.String("192.168.1.200"),
	Port: jsii.Number(5001),
	Network: productionNetwork,
	Protocol: awsmediaconnectalpha.BridgeProtocol_RTP(),
	Ttl: jsii.Number(50),
})

namedOutput := map[string]interface{}{
	"name": jsii.String("on-prem-output"),
	"output": networkConfig,
}
```

## Encryption

MediaConnect supports encryption for sources, outputs, and entitlements. This package provides type-safe encryption configuration structs that match the encryption requirements for different protocols.

### Encryption Types

MediaConnect supports two types of encryption:

1. **Static Key Encryption** - Used for Zixi Push/Pull protocols and entitlements
2. **SRT Password Encryption** - Used for SRT Listener and SRT Caller protocols

Note: CFN exposes only `static-key` and `srt-password` for flow output encryption today; SPEKE is not currently part of the surface.

**Auto-created IAM role.** Every encryption struct accepts an optional `role`. Omit it and the consuming construct creates a scoped role for you: trust policy for `mediaconnect.amazonaws.com` with `aws:SourceAccount` + `aws:SourceArn` conditions (confused-deputy protection), and just enough permission to read the provided secret (including `kms:Decrypt` when the secret uses a customer-managed KMS key).

**Providing your own role.** If you supply a `role`, it's used as-is — the construct does **not** grant it any permissions. You must grant it the necessary permissions yourself. Provide your own role when you need stricter control or a shared identity.

**Trust-policy scope: flows vs. routers.** When the L2 auto-creates a role, it also pins `aws:SourceArn` to the consuming resource:

* **Flows** — the trust policy pins the flow ARN (`arn:...:flow:*:<flow-name>`). The `*` wildcards the service-assigned id segment; the flow name is fixed at create time.
* **Routers** — the trust policy can only pin a wildcarded ARN (`arn:...:routerInput:*` / `arn:...:routerOutput:*`). Router I/O ARNs use a service-generated id segment that is unknown at synth time, and using the live ARN attribute would create a CloudFormation dependency cycle (role → router → role). If you need per-resource pinning, supply your own `role` with a trust condition that pins the exact ARN — you can compute it from the resource's `routerInputId` / `routerOutputId` after first deploy and apply it on a follow-up deploy.

### Static Key Encryption

Use a static-key encryption struct for Zixi protocols and entitlements. This requires an encryption algorithm (AES128, AES192, or AES256). Pass it inline where the construct asks for it:

```go
var stack Stack
var flow Flow
var role IRole
var secret ISecret


awsmediaconnectalpha.NewFlowEntitlement(stack, jsii.String("MyEntitlement"), &FlowEntitlementProps{
	Flow: Flow,
	Subscribers: []*string{
		jsii.String("111122223333"),
	},
	Description: jsii.String("Grant partner access to live feed"),
	Encryption: &StaticKeyEncryption{
		Role: *Role,
		Secret: *Secret,
		Algorithm: awsmediaconnectalpha.EncryptionAlgorithm_AES256(),
	},
})
```

### SRT Password Encryption

SRT protocols take an encryption struct with just `role` and `secret` — no algorithm:

```go
var stack Stack
var flow Flow
var role IRole
var secret ISecret


awsmediaconnectalpha.NewFlowOutput(stack, jsii.String("SrtOutput"), &FlowOutputProps{
	Flow: Flow,
	Output: awsmediaconnectalpha.OutputConfiguration_SrtCaller(&SrtCallerOutputConfig{
		Destination: jsii.String("203.0.113.100"),
		Port: jsii.Number(7000),
		Encryption: &SrtPasswordEncryption{
			Role: *Role,
			Secret: *Secret,
		},
	}),
})
```

### Using Encryption with Sources

Apply encryption when configuring flow sources:

```go
var stack Stack
var role IRole
var secret ISecret


// SRT Listener source with encryption
flow := awsmediaconnectalpha.NewFlow(stack, jsii.String("MyFlow"), &FlowProps{
	Source: awsmediaconnectalpha.SourceConfiguration_SrtListener(&SourceSrtListener{
		FlowSourceName: jsii.String("encrypted-source"),
		Port: jsii.Number(5000),
		Network: awsmediaconnectalpha.NetworkConfiguration_PublicNetwork(jsii.String("203.0.113.0/24")),
		Decryption: &SrtPasswordEncryption{
			Role: *Role,
			Secret: *Secret,
		},
	}),
})

// Zixi Push source with encryption
flow2 := awsmediaconnectalpha.NewFlow(stack, jsii.String("MyFlow2"), &FlowProps{
	Source: awsmediaconnectalpha.SourceConfiguration_ZixiPush(&SourceZixiPush{
		FlowSourceName: jsii.String("encrypted-zixi-source"),
		MaxLatency: awscdk.Duration_Millis(jsii.Number(2000)),
		Network: awsmediaconnectalpha.NetworkConfiguration_*PublicNetwork(jsii.String("203.0.113.0/24")),
		Decryption: &StaticKeyEncryption{
			Role: *Role,
			Secret: *Secret,
			Algorithm: awsmediaconnectalpha.EncryptionAlgorithm_AES256(),
		},
	}),
})
```

### Using Encryption with Outputs

Apply encryption when configuring flow outputs:

```go
var stack Stack
var flow Flow
var role IRole
var secret ISecret


// SRT Caller output with encryption
output := awsmediaconnectalpha.NewFlowOutput(stack, jsii.String("EncryptedOutput"), &FlowOutputProps{
	Flow: flow,
	Description: jsii.String("Encrypted SRT output"),
	Output: awsmediaconnectalpha.OutputConfiguration_SrtCaller(&SrtCallerOutputConfig{
		Destination: jsii.String("203.0.113.100"),
		Port: jsii.Number(7000),
		Encryption: &SrtPasswordEncryption{
			Role: *Role,
			Secret: *Secret,
		},
	}),
})
```

### Using Encryption with Entitlements

Entitlements use static key encryption:

```go
var stack Stack
var flow Flow
var role IRole
var secret ISecret


entitlement := awsmediaconnectalpha.NewFlowEntitlement(stack, jsii.String("MyEntitlement"), &FlowEntitlementProps{
	Flow: flow,
	Description: jsii.String("Grant partner access to live feed"),
	Subscribers: []*string{
		jsii.String("111122223333"),
	},
	Encryption: &StaticKeyEncryption{
		Role: *Role,
		Secret: *Secret,
		Algorithm: awsmediaconnectalpha.EncryptionAlgorithm_AES256(),
	},
})
```

### Router Transit Encryption

When integrating flows with routers, use transit encryption to secure the connection between the flow and router:

```go
var stack Stack
var flow Flow
var role IRole
var secret ISecret
var existingRouterOutput RouterOutput


// Flow output to router with transit encryption
routerOutput := awsmediaconnectalpha.NewFlowOutput(stack, jsii.String("RouterOutput"), &FlowOutputProps{
	Flow: flow,
	Output: awsmediaconnectalpha.OutputConfiguration_Router(&RouterTransitConfig{
		Encryption: &TransitEncryption{
			Role: *Role,
			Secret: *Secret,
		},
	}),
})

// Flow source from router with transit encryption
flowFromRouter := awsmediaconnectalpha.NewFlow(stack, jsii.String("FlowFromRouter"), &FlowProps{
	Source: awsmediaconnectalpha.SourceConfiguration_Router(&RouterSource{
		RouterOutput: existingRouterOutput,
		Decryption: &TransitEncryption{
			Role: *Role,
			Secret: *Secret,
		},
	}),
})
```

### Router SRT Encryption

Router outputs using SRT protocols use `RouterSrtEncryption` for encryption:

```go
var stack Stack
var networkInterface RouterNetworkInterface
var role IRole
var secret ISecret


output := awsmediaconnectalpha.NewRouterOutput(stack, jsii.String("EncryptedSrtOutput"), &RouterOutputProps{
	RouterOutputName: jsii.String("encrypted-srt-output"),
	MaximumBitrate: awscdk.Bitrate_Mbps(jsii.Number(10)),
	RoutingScope: awsmediaconnectalpha.RoutingScope_REGIONAL(),
	Tier: awsmediaconnectalpha.RouterOutputTier_OUTPUT_50(),
	Configuration: awsmediaconnectalpha.RouterOutputConfiguration_Standard(&StandardOutputConfigurationProps{
		Protocol: awsmediaconnectalpha.RouterOutputProtocol_SrtCaller(&SrtCallerOutputProtocolProps{
			DestinationAddress: jsii.String("203.0.113.100"),
			DestinationPort: jsii.Number(9001),
			MinimumLatency: awscdk.Duration_Millis(jsii.Number(200)),
			EncryptionConfiguration: &RouterSrtEncryption{
				Role: *Role,
				Secret: *Secret,
			},
		}),
		NetworkInterface: networkInterface,
	}),
})
```

Note: `RouterSrtEncryption` is distinct from `SrtPasswordEncryption` (used on flow sources/outputs) — router outputs use a simpler CFN shape without a `keyType` discriminator.

## CloudWatch Metrics

Flows and Bridges expose CloudWatch metric helpers for monitoring. You can create alarms and dashboards using these metrics:

```go
var flow Flow
var stack Stack


// Create a CloudWatch alarm on source bitrate
alarm := flow.metricSourceBitrate().CreateAlarm(stack, jsii.String("LowBitrate"), &CreateAlarmOptions{
	Threshold: jsii.Number(1000000),
	EvaluationPeriods: jsii.Number(1),
})

// Monitor unrecovered packets
flow.metricSourceNotRecoveredPackets().CreateAlarm(stack, jsii.String("PacketLoss"), &CreateAlarmOptions{
	Threshold: jsii.Number(100),
	EvaluationPeriods: jsii.Number(2),
})

// Track total packets with custom options
totalPackets := flow.metricSourceTotalPackets(&MetricOptions{
	Statistic: jsii.String("sum"),
	Period: awscdk.Duration_Minutes(jsii.Number(5)),
})
```

### Flow metrics

* `metricSourceBitrate()` - Bitrate of content ingested into the flow (average)
* `metricSourceNotRecoveredPackets()` - Packets lost in transit that were not recovered by error correction (sum)
* `metricSourceTotalPackets()` - Total packets received by the flow sources (sum)
* `metricSourceSelected()` - Indicates which source is being used under Failover mode (max; 1 = active, 0 = standby)
* `metricSourceConnected()` - Source connection state for Zixi, SRT, and RIST (min; 1 = connected, 0 = disconnected)
* `metricSourceDisconnections()` - Number of times the source transitioned from connected to disconnected (sum)
* `metricSourceDroppedPackets()` - Packets lost before any error correction took place (sum)
* `metricSourcePacketLossPercent()` - Percentage of packets lost during transit, even if they were recovered (average)
* `metricSourceRoundTripTime()` - Round-trip time to the source for RIST, Zixi, and SRT (average, milliseconds)
* `metricSourceJitter()` - Current network jitter of the source (average, milliseconds)
* `metric(metricName)` - Create a custom metric by name

### Bridge metrics

Bridge metrics are dimensioned by `BridgeARN`. The underlying CloudWatch metric name is chosen automatically based on whether the bridge is ingress or egress.

* `metricSourceBitrate(bridgeSourceName)` - Bitrate of a specific bridge source (average)
* `metricSourcePacketLossPercent(bridgeSourceName)` - Percentage of packets lost on a specific bridge source (average)
* `metricFailoverSwitches()` - Total number of times the bridge switches between sources under `FAILOVER` failover mode (sum)
* `metric(metricName)` - Create a custom metric by name

### Router Input metrics

Router input metrics are dimensioned by `RouterInputARN`.

* `metricBitrate()` - Bitrate of the router input's payload (average)
* `metricNotRecoveredPackets()` - Packets lost in transit that were not recovered by error correction (sum)
* `metricTotalPackets()` - Total number of packets received by the router input (sum)
* `metricConnected()` - Connection state for SRT sources (min; 1 = connected, 0 = disconnected)
* `metricContinuityCounterErrors()` - Continuity counter errors in the transport stream (sum)
* `metricLatency()` - Recovery latency of the input stream for RIST, SRT, and RTP-FEC (average, milliseconds)
* `metricFailoverSwitches()` - Total times the router input switched sources under Failover mode (sum)
* `metric(metricName)` - Create a custom metric by name

### Router Output metrics

Router output metrics are dimensioned by `RouterOutputARN`.

* `metricBitrate()` - Bitrate of the router output's payload (average)
* `metricTotalPackets()` - Total number of packets sent by the router output (sum)
* `metricConnected()` - Connection state for SRT outputs (min; 1 = connected, 0 = disconnected)
* `metricArqRequests()` - Retransmitted packets requested through ARQ for RIST and SRT outputs (sum)
* `metric(metricName)` - Create a custom metric by name

### Gateway metrics

Gateway metrics are dimensioned by `GatewayARN`. Pass extra dimensions such as `NetworkName`, `InstanceId`, or `BridgeSourceName` via `props.dimensionsMap` to narrow to a specific network, appliance, or bridge source.

* `metricEgressBridgeTotalPackets()` - Total packets sent from egress bridges hosted on the gateway (sum)
* `metricEgressBridgeDroppedPackets()` - Packets dropped by egress bridges hosted on the gateway (sum)
* `metricIngressBridgeTotalPackets()` - Total packets received by ingress bridges hosted on the gateway (sum)
* `metricIngressBridgeDroppedPackets()` - Packets dropped by ingress bridges hosted on the gateway (sum)
* `metric(metricName)` - Create a custom metric by name (e.g. `IngressBridgeBitRate`, `EgressBridgeBitRate`, `IngressBridgeSourcePacketLossPercent`)

Pair the total + dropped helpers to build a dropped-packet percentage chart — for example, divide `metricEgressBridgeDroppedPackets()` by `metricEgressBridgeTotalPackets()` in a math expression.

All metrics support standard CloudWatch metric options for customizing period, statistic, and dimensions.

## Public CIDR warnings

Several constructs accept CIDR ranges that determine who can contribute content or pull outputs. Passing an open range (`0.0.0.0/0` or any `/0` prefix) makes the resource reachable from anywhere on the public internet, which is rarely what you want. The module emits synthesis-time warnings when it detects an open range on:

* `GatewayProps.egressCidrBlocks`
* `NetworkConfiguration.publicNetwork(cidr)` used on flow sources
* `cidrAllowList` on Zixi Push / Zixi Pull / SRT Listener flow outputs
* `RouterNetworkConfiguration.publicNetwork({ cidr })`

Restrict each range to the narrowest set of addresses that actually need access.
