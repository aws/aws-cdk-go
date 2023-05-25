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
)

