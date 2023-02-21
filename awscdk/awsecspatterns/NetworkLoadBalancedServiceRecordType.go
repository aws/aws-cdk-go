package awsecspatterns


// Describes the type of DNS record the service should create.
type NetworkLoadBalancedServiceRecordType string

const (
	// Create Route53 A Alias record.
	NetworkLoadBalancedServiceRecordType_ALIAS NetworkLoadBalancedServiceRecordType = "ALIAS"
	// Create a CNAME record.
	NetworkLoadBalancedServiceRecordType_CNAME NetworkLoadBalancedServiceRecordType = "CNAME"
	// Do not create any DNS records.
	NetworkLoadBalancedServiceRecordType_NONE NetworkLoadBalancedServiceRecordType = "NONE"
)

