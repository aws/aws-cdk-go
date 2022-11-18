package awsecspatterns


// Describes the type of DNS record the service should create.
// Experimental.
type NetworkLoadBalancedServiceRecordType string

const (
	// Create Route53 A Alias record.
	// Experimental.
	NetworkLoadBalancedServiceRecordType_ALIAS NetworkLoadBalancedServiceRecordType = "ALIAS"
	// Create a CNAME record.
	// Experimental.
	NetworkLoadBalancedServiceRecordType_CNAME NetworkLoadBalancedServiceRecordType = "CNAME"
	// Do not create any DNS records.
	// Experimental.
	NetworkLoadBalancedServiceRecordType_NONE NetworkLoadBalancedServiceRecordType = "NONE"
)

