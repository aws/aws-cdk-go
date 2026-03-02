package previewawsdevopsagentmixins


// MCP server details returned after registration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   registeredMCPServerDetailsProperty := &RegisteredMCPServerDetailsProperty{
//   	ApiKeyHeader: jsii.String("apiKeyHeader"),
//   	AuthorizationMethod: jsii.String("authorizationMethod"),
//   	Description: jsii.String("description"),
//   	Endpoint: jsii.String("endpoint"),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredmcpserverdetails.html
//
type CfnServicePropsMixin_RegisteredMCPServerDetailsProperty struct {
	// API key header name if using API key authentication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredmcpserverdetails.html#cfn-devopsagent-service-registeredmcpserverdetails-apikeyheader
	//
	ApiKeyHeader *string `field:"optional" json:"apiKeyHeader" yaml:"apiKeyHeader"`
	// MCP server authorization method.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredmcpserverdetails.html#cfn-devopsagent-service-registeredmcpserverdetails-authorizationmethod
	//
	AuthorizationMethod *string `field:"optional" json:"authorizationMethod" yaml:"authorizationMethod"`
	// Optional description for the MCP server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredmcpserverdetails.html#cfn-devopsagent-service-registeredmcpserverdetails-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// MCP server endpoint URL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredmcpserverdetails.html#cfn-devopsagent-service-registeredmcpserverdetails-endpoint
	//
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// MCP server name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredmcpserverdetails.html#cfn-devopsagent-service-registeredmcpserverdetails-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

