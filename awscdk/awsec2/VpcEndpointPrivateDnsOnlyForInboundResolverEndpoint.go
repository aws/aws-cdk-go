package awsec2


// Indicates whether to enable private DNS only for inbound endpoints.
//
// This option is available only for services that support both gateway and interface endpoints.
// It routes traffic that originates from the VPC to the gateway endpoint and traffic that
// originates from on-premises to the interface endpoint.
type VpcEndpointPrivateDnsOnlyForInboundResolverEndpoint string

const (
	// Enable private DNS for all resolvers.
	VpcEndpointPrivateDnsOnlyForInboundResolverEndpoint_ALL_RESOLVERS VpcEndpointPrivateDnsOnlyForInboundResolverEndpoint = "ALL_RESOLVERS"
	// Enable private DNS only for inbound endpoints.
	VpcEndpointPrivateDnsOnlyForInboundResolverEndpoint_ONLY_INBOUND_RESOLVER VpcEndpointPrivateDnsOnlyForInboundResolverEndpoint = "ONLY_INBOUND_RESOLVER"
)

