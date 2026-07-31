package awsmediaconnectalpha


// Attributes for importing an existing Flow Source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import mediaconnect_alpha "github.com/aws/aws-cdk-go/awsmediaconnectalpha"
//
//   flowSourceAttributes := &FlowSourceAttributes{
//   	FlowSourceArn: jsii.String("flowSourceArn"),
//
//   	// the properties below are optional
//   	FlowSourceName: jsii.String("flowSourceName"),
//   	IngestIp: jsii.String("ingestIp"),
//   	SourceIngestPort: jsii.String("sourceIngestPort"),
//   }
//
// Experimental.
type FlowSourceAttributes struct {
	// The Amazon Resource Name (ARN) of the flow source.
	// Experimental.
	FlowSourceArn *string `field:"required" json:"flowSourceArn" yaml:"flowSourceArn"`
	// The name of the flow source.
	// Default: - accessing `flowSourceName` on the imported source throws; only provide when available.
	//
	// Experimental.
	FlowSourceName *string `field:"optional" json:"flowSourceName" yaml:"flowSourceName"`
	// The IP address that the flow will be listening on for incoming content.
	// Default: - accessing `ingestIp` on the imported source throws; only provide when available.
	//
	// Experimental.
	IngestIp *string `field:"optional" json:"ingestIp" yaml:"ingestIp"`
	// The port that the flow will be listening on for incoming content.
	// Default: - accessing `sourceIngestPort` on the imported source throws; only provide when available.
	//
	// Experimental.
	SourceIngestPort *string `field:"optional" json:"sourceIngestPort" yaml:"sourceIngestPort"`
}

