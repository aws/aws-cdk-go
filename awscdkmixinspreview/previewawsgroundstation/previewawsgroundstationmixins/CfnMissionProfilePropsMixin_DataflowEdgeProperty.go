package previewawsgroundstationmixins


// A dataflow edge defines from where and to where data will flow during a contact.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataflowEdgeProperty := &DataflowEdgeProperty{
//   	Destination: jsii.String("destination"),
//   	Source: jsii.String("source"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-missionprofile-dataflowedge.html
//
type CfnMissionProfilePropsMixin_DataflowEdgeProperty struct {
	// The ARN of the destination for this dataflow edge.
	//
	// For example, specify the ARN of a dataflow endpoint config for a downlink edge or an antenna uplink config for an uplink edge.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-missionprofile-dataflowedge.html#cfn-groundstation-missionprofile-dataflowedge-destination
	//
	Destination *string `field:"optional" json:"destination" yaml:"destination"`
	// The ARN of the source for this dataflow edge.
	//
	// For example, specify the ARN of an antenna downlink config for a downlink edge or a dataflow endpoint config for an uplink edge.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-missionprofile-dataflowedge.html#cfn-groundstation-missionprofile-dataflowedge-source
	//
	Source *string `field:"optional" json:"source" yaml:"source"`
}

