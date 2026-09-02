package awsmgn


// Configuration for the target network topology and addressing.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetNetworkProperty := &TargetNetworkProperty{
//   	Topology: jsii.String("topology"),
//
//   	// the properties below are optional
//   	InboundCidr: jsii.String("inboundCidr"),
//   	InspectionCidr: jsii.String("inspectionCidr"),
//   	OutboundCidr: jsii.String("outboundCidr"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mgn-networkmigrationdefinition-targetnetwork.html
//
type CfnNetworkMigrationDefinition_TargetNetworkProperty struct {
	// The network topology type for the target environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mgn-networkmigrationdefinition-targetnetwork.html#cfn-mgn-networkmigrationdefinition-targetnetwork-topology
	//
	Topology *string `field:"required" json:"topology" yaml:"topology"`
	// The CIDR block for inbound traffic in the target network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mgn-networkmigrationdefinition-targetnetwork.html#cfn-mgn-networkmigrationdefinition-targetnetwork-inboundcidr
	//
	InboundCidr *string `field:"optional" json:"inboundCidr" yaml:"inboundCidr"`
	// The CIDR block for inspection traffic in the target network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mgn-networkmigrationdefinition-targetnetwork.html#cfn-mgn-networkmigrationdefinition-targetnetwork-inspectioncidr
	//
	InspectionCidr *string `field:"optional" json:"inspectionCidr" yaml:"inspectionCidr"`
	// The CIDR block for outbound traffic in the target network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mgn-networkmigrationdefinition-targetnetwork.html#cfn-mgn-networkmigrationdefinition-targetnetwork-outboundcidr
	//
	OutboundCidr *string `field:"optional" json:"outboundCidr" yaml:"outboundCidr"`
}

