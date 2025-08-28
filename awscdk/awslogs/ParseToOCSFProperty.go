package awslogs


// Processor to parse events from CloudTrail, Route53Resolver, VPCFlow, EKSAudit and AWSWAF into OCSF V1.1 format.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parseToOCSFProperty := &ParseToOCSFProperty{
//   	EventSource: awscdk.Aws_logs.OCSFSourceType_CLOUD_TRAIL,
//   	OcsfVersion: awscdk.*Aws_logs.OCSFVersion_V1_1,
//
//   	// the properties below are optional
//   	Source: jsii.String("source"),
//   }
//
type ParseToOCSFProperty struct {
	// Type of input log event source to convert to OCSF format.
	EventSource OCSFSourceType `field:"required" json:"eventSource" yaml:"eventSource"`
	// Version of OCSF schema to convert to.
	OcsfVersion OCSFVersion `field:"required" json:"ocsfVersion" yaml:"ocsfVersion"`
	// Path to the field in the log event that will be parsed.
	//
	// Use dot notation to access child fields.
	// Default: '@message'.
	//
	Source *string `field:"optional" json:"source" yaml:"source"`
}

