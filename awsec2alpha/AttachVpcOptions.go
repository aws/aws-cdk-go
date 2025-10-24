package awsec2alpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Options for creating an Attachment via the attachVpc() method.
//
// Example:
//   transitGateway := awsec2alpha.NewTransitGateway(this, jsii.String("MyTransitGateway"))
//   routeTable := transitGateway.addRouteTable(jsii.String("CustomRouteTable"))
//
//   myVpc := awsec2alpha.NewVpcV2(this, jsii.String("Vpc"))
//   subnet := awsec2alpha.NewSubnetV2(this, jsii.String("Subnet"), &SubnetV2Props{
//   	Vpc: myVpc,
//   	AvailabilityZone: jsii.String("eu-west-2a"),
//   	Ipv4CidrBlock: awsec2alpha.NewIpCidr(jsii.String("10.0.0.0/24")),
//   	SubnetType: awscdk.SubnetType_PUBLIC,
//   })
//
//   attachment := transitGateway.attachVpc(jsii.String("VpcAttachment"), &AttachVpcOptions{
//   	Vpc: myVpc,
//   	Subnets: []ISubnet{
//   		subnet,
//   	},
//   })
//
//   // Add a static route to direct traffic
//   routeTable.AddRoute(jsii.String("StaticRoute"), attachment, jsii.String("10.0.0.0/16"))
//
//   // Block unwanted traffic with a blackhole route
//   routeTable.AddBlackholeRoute(jsii.String("BlackholeRoute"), jsii.String("172.16.0.0/16"))
//
// Experimental.
type AttachVpcOptions struct {
	// A list of one or more subnets to place the attachment in.
	//
	// It is recommended to specify more subnets for better availability.
	// Experimental.
	Subnets *[]awsec2.ISubnet `field:"required" json:"subnets" yaml:"subnets"`
	// A VPC attachment(s) will get assigned to.
	// Experimental.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// An optional route table to associate with this VPC attachment.
	// Default: - No associations will be created unless it is for the default route table and automatic association is enabled.
	//
	// Experimental.
	AssociationRouteTable ITransitGatewayRouteTable `field:"optional" json:"associationRouteTable" yaml:"associationRouteTable"`
	// A list of optional route tables to propagate routes to.
	// Default: - No propagations will be created unless it is for the default route table and automatic propagation is enabled.
	//
	// Experimental.
	PropagationRouteTables *[]ITransitGatewayRouteTable `field:"optional" json:"propagationRouteTables" yaml:"propagationRouteTables"`
	// Physical name of this Transit Gateway VPC Attachment.
	// Default: - Assigned by CloudFormation.
	//
	// Experimental.
	TransitGatewayAttachmentName *string `field:"optional" json:"transitGatewayAttachmentName" yaml:"transitGatewayAttachmentName"`
	// The VPC attachment options.
	// Default: - All options are disabled.
	//
	// Experimental.
	VpcAttachmentOptions ITransitGatewayVpcAttachmentOptions `field:"optional" json:"vpcAttachmentOptions" yaml:"vpcAttachmentOptions"`
}

