package awslogs


// Types of AWS vended logs with built-in parsers.
//
// AWS provides specialized parsers for common log formats produced by various AWS services.
type VendedLogType string

const (
	// Parse CloudFront logs.
	VendedLogType_CLOUDFRONT VendedLogType = "CLOUDFRONT"
	// Parse VPC flow logs.
	VendedLogType_VPC VendedLogType = "VPC"
	// Parse AWS WAF logs.
	VendedLogType_WAF VendedLogType = "WAF"
	// Parse Route 53 logs.
	VendedLogType_ROUTE53 VendedLogType = "ROUTE53"
	// Parse PostgreSQL logs.
	VendedLogType_POSTGRES VendedLogType = "POSTGRES"
)

