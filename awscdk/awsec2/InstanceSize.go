package awsec2


// What size of instance to use.
//
// Example:
//   var build build
//
//
//   fleet := gamelift.NewBuildFleet(this, jsii.String("Game server fleet"), &BuildFleetProps{
//   	FleetName: jsii.String("test-fleet"),
//   	Content: build,
//   	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_C4, ec2.InstanceSize_LARGE),
//   	RuntimeConfiguration: &RuntimeConfiguration{
//   		ServerProcesses: []serverProcess{
//   			&serverProcess{
//   				LaunchPath: jsii.String("/local/game/GameLiftExampleServer.x86_64"),
//   			},
//   		},
//   	},
//   	IngressRules: []ingressRule{
//   		&ingressRule{
//   			Source: gamelift.Peer_AnyIpv4(),
//   			Port: gamelift.Port_TcpRange(jsii.Number(100), jsii.Number(200)),
//   		},
//   	},
//   })
//   // Allowing a specific CIDR for port 1111 on UDP Protocol
//   fleet.AddIngressRule(gamelift.Peer_Ipv4(jsii.String("1.2.3.4/32")), gamelift.Port_Udp(jsii.Number(1111)))
//
type InstanceSize string

const (
	// Instance size NANO (nano).
	InstanceSize_NANO InstanceSize = "NANO"
	// Instance size MICRO (micro).
	InstanceSize_MICRO InstanceSize = "MICRO"
	// Instance size SMALL (small).
	InstanceSize_SMALL InstanceSize = "SMALL"
	// Instance size MEDIUM (medium).
	InstanceSize_MEDIUM InstanceSize = "MEDIUM"
	// Instance size LARGE (large).
	InstanceSize_LARGE InstanceSize = "LARGE"
	// Instance size XLARGE (xlarge).
	InstanceSize_XLARGE InstanceSize = "XLARGE"
	// Instance size XLARGE2 (2xlarge).
	InstanceSize_XLARGE2 InstanceSize = "XLARGE2"
	// Instance size XLARGE3 (3xlarge).
	InstanceSize_XLARGE3 InstanceSize = "XLARGE3"
	// Instance size XLARGE4 (4xlarge).
	InstanceSize_XLARGE4 InstanceSize = "XLARGE4"
	// Instance size XLARGE6 (6xlarge).
	InstanceSize_XLARGE6 InstanceSize = "XLARGE6"
	// Instance size XLARGE8 (8xlarge).
	InstanceSize_XLARGE8 InstanceSize = "XLARGE8"
	// Instance size XLARGE9 (9xlarge).
	InstanceSize_XLARGE9 InstanceSize = "XLARGE9"
	// Instance size XLARGE10 (10xlarge).
	InstanceSize_XLARGE10 InstanceSize = "XLARGE10"
	// Instance size XLARGE12 (12xlarge).
	InstanceSize_XLARGE12 InstanceSize = "XLARGE12"
	// Instance size XLARGE16 (16xlarge).
	InstanceSize_XLARGE16 InstanceSize = "XLARGE16"
	// Instance size XLARGE18 (18xlarge).
	InstanceSize_XLARGE18 InstanceSize = "XLARGE18"
	// Instance size XLARGE24 (24xlarge).
	InstanceSize_XLARGE24 InstanceSize = "XLARGE24"
	// Instance size XLARGE32 (32xlarge).
	InstanceSize_XLARGE32 InstanceSize = "XLARGE32"
	// Instance size XLARGE48 (48xlarge).
	InstanceSize_XLARGE48 InstanceSize = "XLARGE48"
	// Instance size XLARGE56 (56xlarge).
	InstanceSize_XLARGE56 InstanceSize = "XLARGE56"
	// Instance size XLARGE56 (112xlarge).
	InstanceSize_XLARGE112 InstanceSize = "XLARGE112"
	// Instance size METAL (metal).
	InstanceSize_METAL InstanceSize = "METAL"
)

