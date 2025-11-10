package awstransfer


// VPC_LATTICE egress configuration that specifies the Resource Configuration ARN and port for connecting to SFTP servers through customer VPCs.
//
// Requires a valid Resource Configuration with appropriate network access.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectorVpcLatticeEgressConfigProperty := &ConnectorVpcLatticeEgressConfigProperty{
//   	ResourceConfigurationArn: jsii.String("resourceConfigurationArn"),
//
//   	// the properties below are optional
//   	PortNumber: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-connector-connectorvpclatticeegressconfig.html
//
type CfnConnector_ConnectorVpcLatticeEgressConfigProperty struct {
	// ARN of the VPC_LATTICE Resource Configuration that defines the target SFTP server location.
	//
	// Must point to a valid Resource Configuration in the customer's VPC with appropriate network connectivity to the SFTP server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-connector-connectorvpclatticeegressconfig.html#cfn-transfer-connector-connectorvpclatticeegressconfig-resourceconfigurationarn
	//
	ResourceConfigurationArn *string `field:"required" json:"resourceConfigurationArn" yaml:"resourceConfigurationArn"`
	// Port number for connecting to the SFTP server through VPC_LATTICE.
	//
	// Defaults to 22 if not specified. Must match the port on which the target SFTP server is listening.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-connector-connectorvpclatticeegressconfig.html#cfn-transfer-connector-connectorvpclatticeegressconfig-portnumber
	//
	PortNumber *float64 `field:"optional" json:"portNumber" yaml:"portNumber"`
}

