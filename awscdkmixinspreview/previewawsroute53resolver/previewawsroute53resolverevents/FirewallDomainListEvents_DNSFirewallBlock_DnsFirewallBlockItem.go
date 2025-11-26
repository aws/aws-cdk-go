package previewawsroute53resolverevents


// Type definition for DNSFirewallBlockItem.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dnsFirewallBlockItem := &DnsFirewallBlockItem{
//   	ResolverEndpointDetails: &ResolverEndpointDetails{
//   		Id: []*string{
//   			jsii.String("id"),
//   		},
//   	},
//   	ResolverNetworkInterfaceDetails: &ResolverNetworkInterfaceDetails{
//   		Id: []*string{
//   			jsii.String("id"),
//   		},
//   	},
//   	ResourceType: []*string{
//   		jsii.String("resourceType"),
//   	},
//   }
//
// Experimental.
type FirewallDomainListEvents_DNSFirewallBlock_DnsFirewallBlockItem struct {
	// resolver-endpoint-details property.
	//
	// Specify an array of string values to match this event if the actual value of resolver-endpoint-details is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ResolverEndpointDetails *FirewallDomainListEvents_DNSFirewallBlock_ResolverEndpointDetails `field:"optional" json:"resolverEndpointDetails" yaml:"resolverEndpointDetails"`
	// resolver-network-interface-details property.
	//
	// Specify an array of string values to match this event if the actual value of resolver-network-interface-details is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ResolverNetworkInterfaceDetails *FirewallDomainListEvents_DNSFirewallBlock_ResolverNetworkInterfaceDetails `field:"optional" json:"resolverNetworkInterfaceDetails" yaml:"resolverNetworkInterfaceDetails"`
	// resource-type property.
	//
	// Specify an array of string values to match this event if the actual value of resource-type is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ResourceType *[]*string `field:"optional" json:"resourceType" yaml:"resourceType"`
}

