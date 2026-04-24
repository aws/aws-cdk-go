package awsdevopsagent


// SigV4-authenticated MCP server details returned after registration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   registeredMCPServerSigV4DetailsProperty := &RegisteredMCPServerSigV4DetailsProperty{
//   	Endpoint: jsii.String("endpoint"),
//   	Name: jsii.String("name"),
//   	Region: jsii.String("region"),
//   	RoleArn: jsii.String("roleArn"),
//   	Service: jsii.String("service"),
//
//   	// the properties below are optional
//   	CustomHeaders: map[string]*string{
//   		"customHeadersKey": jsii.String("customHeaders"),
//   	},
//   	Description: jsii.String("description"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredmcpserversigv4details.html
//
type CfnService_RegisteredMCPServerSigV4DetailsProperty struct {
	// The MCP server endpoint URL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredmcpserversigv4details.html#cfn-devopsagent-service-registeredmcpserversigv4details-endpoint
	//
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// The MCP server name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredmcpserversigv4details.html#cfn-devopsagent-service-registeredmcpserversigv4details-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// AWS region for SigV4 signing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredmcpserversigv4details.html#cfn-devopsagent-service-registeredmcpserversigv4details-region
	//
	Region *string `field:"required" json:"region" yaml:"region"`
	// IAM role ARN for SigV4 signing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredmcpserversigv4details.html#cfn-devopsagent-service-registeredmcpserversigv4details-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// AWS service name for SigV4 signing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredmcpserversigv4details.html#cfn-devopsagent-service-registeredmcpserversigv4details-service
	//
	Service *string `field:"required" json:"service" yaml:"service"`
	// Custom headers for the SigV4 MCP server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredmcpserversigv4details.html#cfn-devopsagent-service-registeredmcpserversigv4details-customheaders
	//
	CustomHeaders interface{} `field:"optional" json:"customHeaders" yaml:"customHeaders"`
	// Optional description for the MCP server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredmcpserversigv4details.html#cfn-devopsagent-service-registeredmcpserversigv4details-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

