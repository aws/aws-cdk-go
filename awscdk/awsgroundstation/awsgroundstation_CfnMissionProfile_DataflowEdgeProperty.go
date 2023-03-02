package awsgroundstation


// A dataflow edge defines from where and to where data will flow during a contact.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataflowEdgeProperty := &dataflowEdgeProperty{
//   	destination: jsii.String("destination"),
//   	source: jsii.String("source"),
//   }
//
type CfnMissionProfile_DataflowEdgeProperty struct {
	// The ARN of the destination for this dataflow edge.
	//
	// For example, specify the ARN of a dataflow endpoint config for a downlink edge or an antenna uplink config for an uplink edge.
	Destination *string `field:"optional" json:"destination" yaml:"destination"`
	// The ARN of the source for this dataflow edge.
	//
	// For example, specify the ARN of an antenna downlink config for a downlink edge or a dataflow endpoint config for an uplink edge.
	Source *string `field:"optional" json:"source" yaml:"source"`
}

