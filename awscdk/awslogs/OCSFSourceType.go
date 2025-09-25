package awslogs


// Types of event sources supported to convert to OCSF format.
type OCSFSourceType string

const (
	// Log events from CloudTrail.
	OCSFSourceType_CLOUD_TRAIL OCSFSourceType = "CLOUD_TRAIL"
	// Log events from Route53Resolver.
	OCSFSourceType_ROUTE53_RESOLVER OCSFSourceType = "ROUTE53_RESOLVER"
	// Log events from VPCFlow.
	OCSFSourceType_VPC_FLOW OCSFSourceType = "VPC_FLOW"
	// Log events from EKSAudit.
	OCSFSourceType_EKS_AUDIT OCSFSourceType = "EKS_AUDIT"
	// Log events from AWSWAF.
	OCSFSourceType_AWS_WAF OCSFSourceType = "AWS_WAF"
)

