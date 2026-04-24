package awsdevopsagent


// SigV4-authenticated MCP server details returned after registration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   registeredMCPServerSigV4DetailsProperty := &RegisteredMCPServerSigV4DetailsProperty{
//   	CustomHeaders: map[string]*string{
//   		"customHeadersKey": jsii.String("customHeaders"),
//   	},
//   	Description: jsii.String("description"),
//   	Endpoint: jsii.String("endpoint"),
//   	Name: jsii.String("name"),
//   	Region: jsii.String("region"),
//   	RoleArn: jsii.String("roleArn"),
//   	Service: jsii.String("service"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredmcpserversigv4details.html
//
type CfnServicePropsMixin_RegisteredMCPServerSigV4DetailsProperty struct {
	// Custom headers for the SigV4 MCP server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredmcpserversigv4details.html#cfn-devopsagent-service-registeredmcpserversigv4details-customheaders
	//
	CustomHeaders interface{} `field:"optional" json:"customHeaders" yaml:"customHeaders"`
	// Optional description for the MCP server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredmcpserversigv4details.html#cfn-devopsagent-service-registeredmcpserversigv4details-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The MCP server endpoint URL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredmcpserversigv4details.html#cfn-devopsagent-service-registeredmcpserversigv4details-endpoint
	//
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// The MCP server name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredmcpserversigv4details.html#cfn-devopsagent-service-registeredmcpserversigv4details-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// AWS region for SigV4 signing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredmcpserversigv4details.html#cfn-devopsagent-service-registeredmcpserversigv4details-region
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
	// IAM role ARN for SigV4 signing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredmcpserversigv4details.html#cfn-devopsagent-service-registeredmcpserversigv4details-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// AWS service name for SigV4 signing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredmcpserversigv4details.html#cfn-devopsagent-service-registeredmcpserversigv4details-service
	//
	Service *string `field:"optional" json:"service" yaml:"service"`
}

