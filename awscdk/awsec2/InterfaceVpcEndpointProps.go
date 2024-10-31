package awsec2


// Construction properties for an InterfaceVpcEndpoint.
//
// Example:
//   var vpc vpc
//
//
//   ec2.NewInterfaceVpcEndpoint(this, jsii.String("VPC Endpoint"), &InterfaceVpcEndpointProps{
//   	Vpc: Vpc,
//   	Service: ec2.NewInterfaceVpcEndpointService(jsii.String("com.amazonaws.vpce.us-east-1.vpce-svc-uuddlrlrbastrtsvc"), jsii.Number(443)),
//   	Subnets: &SubnetSelection{
//   		SubnetType: ec2.SubnetType_PRIVATE_ISOLATED,
//   		AvailabilityZones: []*string{
//   			jsii.String("us-east-1a"),
//   			jsii.String("us-east-1c"),
//   		},
//   	},
//   })
//
type InterfaceVpcEndpointProps struct {
	// The service to use for this interface VPC endpoint.
	Service IInterfaceVpcEndpointService `field:"required" json:"service" yaml:"service"`
	// Limit to only those availability zones where the endpoint service can be created.
	//
	// Setting this to 'true' requires a lookup to be performed at synthesis time. Account
	// and region must be set on the containing stack for this to work.
	// Default: false.
	//
	LookupSupportedAzs *bool `field:"optional" json:"lookupSupportedAzs" yaml:"lookupSupportedAzs"`
	// Whether to automatically allow VPC traffic to the endpoint.
	//
	// If enabled, all traffic to the endpoint from within the VPC will be
	// automatically allowed. This is done based on the VPC's CIDR range.
	// Default: true.
	//
	Open *bool `field:"optional" json:"open" yaml:"open"`
	// Whether to associate a private hosted zone with the specified VPC.
	//
	// This
	// allows you to make requests to the service using its default DNS hostname.
	// Default: set by the instance of IInterfaceVpcEndpointService, or true if
	// not defined by the instance of IInterfaceVpcEndpointService.
	//
	PrivateDnsEnabled *bool `field:"optional" json:"privateDnsEnabled" yaml:"privateDnsEnabled"`
	// The security groups to associate with this interface VPC endpoint.
	// Default: - a new security group is created.
	//
	SecurityGroups *[]ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The subnets in which to create an endpoint network interface.
	//
	// At most one
	// per availability zone.
	// Default: - private subnets.
	//
	Subnets *SubnetSelection `field:"optional" json:"subnets" yaml:"subnets"`
	// The VPC network in which the interface endpoint will be used.
	Vpc IVpc `field:"required" json:"vpc" yaml:"vpc"`
}

