package awsec2


// The type of VPC endpoint.
type VpcEndpointType string

const (
	// Interface.
	//
	// An interface endpoint is an elastic network interface with a private IP
	// address that serves as an entry point for traffic destined to a supported
	// service.
	VpcEndpointType_INTERFACE VpcEndpointType = "INTERFACE"
	// Gateway.
	//
	// A gateway endpoint is a gateway that is a target for a specified route in
	// your route table, used for traffic destined to a supported AWS service.
	VpcEndpointType_GATEWAY VpcEndpointType = "GATEWAY"
	// A Gateway Load Balancer (GWLB) endpoint is an entry/exit point in your VPC that allows traffic to flow between your VPC and Gateway Load Balancer appliances (like firewalls, intrusion detection systems, or other security appliances) deployed in a separate VPC.
	VpcEndpointType_GATEWAYLOADBALANCER VpcEndpointType = "GATEWAYLOADBALANCER"
	// A ServiceNetwork VPC endpoint is a feature to connect your VPC to an AWS Cloud WAN (Wide Area Network) or Amazon VPC Lattice service.
	VpcEndpointType_SERVICENETWORK VpcEndpointType = "SERVICENETWORK"
	// A Resource VPC endpoint in AWS is specifically designed to connect to AWS Resource Access Manager (RAM) service privately within your VPC, without requiring access through the public internet.
	VpcEndpointType_RESOURCE VpcEndpointType = "RESOURCE"
)

