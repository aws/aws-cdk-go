package awsmediaconnectalpha


// Attributes for importing an existing Flow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import mediaconnect_alpha "github.com/aws/aws-cdk-go/awsmediaconnectalpha"
//
//   flowAttributes := &FlowAttributes{
//   	FlowArn: jsii.String("flowArn"),
//
//   	// the properties below are optional
//   	EgressIp: jsii.String("egressIp"),
//   	FlowAvailabilityZone: jsii.String("flowAvailabilityZone"),
//   	IsFailoverEnabled: jsii.Boolean(false),
//   	SourceArn: jsii.String("sourceArn"),
//   	SourceIngestIp: jsii.String("sourceIngestIp"),
//   	SourceIngestPort: jsii.String("sourceIngestPort"),
//   }
//
// Experimental.
type FlowAttributes struct {
	// The ARN of the flow.
	// Experimental.
	FlowArn *string `field:"required" json:"flowArn" yaml:"flowArn"`
	// The IP address that the flow uses to send outbound content.
	// Default: - accessing `egressIp` on the imported flow throws; only provide when available.
	//
	// Experimental.
	EgressIp *string `field:"optional" json:"egressIp" yaml:"egressIp"`
	// The Availability Zone that the flow was created in.
	// Default: - `flowAvailabilityZone` is undefined on the imported flow.
	//
	// Experimental.
	FlowAvailabilityZone *string `field:"optional" json:"flowAvailabilityZone" yaml:"flowAvailabilityZone"`
	// Indicates whether failover configured.
	// Default: false.
	//
	// Experimental.
	IsFailoverEnabled *bool `field:"optional" json:"isFailoverEnabled" yaml:"isFailoverEnabled"`
	// ARN of the source defined on the flow.
	//
	// Not encoded in the flow ARN, so must be provided explicitly when you need
	// access to `sourceArn` on the imported construct.
	// Default: - sourceArn is not available on the imported construct.
	//
	// Experimental.
	SourceArn *string `field:"optional" json:"sourceArn" yaml:"sourceArn"`
	// The IP address that the flow listens on for incoming content.
	// Default: - accessing `sourceIngestIp` on the imported flow throws; only provide when available.
	//
	// Experimental.
	SourceIngestIp *string `field:"optional" json:"sourceIngestIp" yaml:"sourceIngestIp"`
	// The port that the flow listens on for incoming content.
	// Default: - accessing `sourceIngestPort` on the imported flow throws; only provide when available.
	//
	// Experimental.
	SourceIngestPort *string `field:"optional" json:"sourceIngestPort" yaml:"sourceIngestPort"`
}

