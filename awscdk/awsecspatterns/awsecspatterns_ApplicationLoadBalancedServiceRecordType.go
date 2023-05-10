package awsecspatterns


// Describes the type of DNS record the service should create.
// Experimental.
type ApplicationLoadBalancedServiceRecordType string

const (
	// Create Route53 A Alias record.
	// Experimental.
	ApplicationLoadBalancedServiceRecordType_ALIAS ApplicationLoadBalancedServiceRecordType = "ALIAS"
	// Create a CNAME record.
	// Experimental.
	ApplicationLoadBalancedServiceRecordType_CNAME ApplicationLoadBalancedServiceRecordType = "CNAME"
	// Do not create any DNS records.
	// Experimental.
	ApplicationLoadBalancedServiceRecordType_NONE ApplicationLoadBalancedServiceRecordType = "NONE"
)

