package awstransfer


// Configuration structure that defines how traffic is routed from the connector to the SFTP server.
//
// Contains VPC Lattice settings when using VPC_LATTICE egress type for private connectivity through customer VPCs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectorEgressConfigProperty := &ConnectorEgressConfigProperty{
//   	VpcLattice: &ConnectorVpcLatticeEgressConfigProperty{
//   		ResourceConfigurationArn: jsii.String("resourceConfigurationArn"),
//
//   		// the properties below are optional
//   		PortNumber: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-connector-connectoregressconfig.html
//
type CfnConnector_ConnectorEgressConfigProperty struct {
	// VPC_LATTICE configuration for routing connector traffic through customer VPCs.
	//
	// Enables private connectivity to SFTP servers without requiring public internet access or complex network configurations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-connector-connectoregressconfig.html#cfn-transfer-connector-connectoregressconfig-vpclattice
	//
	VpcLattice interface{} `field:"required" json:"vpcLattice" yaml:"vpcLattice"`
}

