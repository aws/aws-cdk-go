package awsecspatterns


// Describes the type of DNS record the service should create.
type ApplicationLoadBalancedServiceRecordType string

const (
	// Create Route53 A Alias record.
	ApplicationLoadBalancedServiceRecordType_ALIAS ApplicationLoadBalancedServiceRecordType = "ALIAS"
	// Create a CNAME record.
	ApplicationLoadBalancedServiceRecordType_CNAME ApplicationLoadBalancedServiceRecordType = "CNAME"
	// Do not create any DNS records.
	ApplicationLoadBalancedServiceRecordType_NONE ApplicationLoadBalancedServiceRecordType = "NONE"
)

