package awsdevopsagent


// MCP server details returned after registration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   registeredMCPServerDetailsProperty := &RegisteredMCPServerDetailsProperty{
//   	AuthorizationMethod: jsii.String("authorizationMethod"),
//   	Endpoint: jsii.String("endpoint"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	ApiKeyHeader: jsii.String("apiKeyHeader"),
//   	Description: jsii.String("description"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredmcpserverdetails.html
//
type CfnService_RegisteredMCPServerDetailsProperty struct {
	// MCP server authorization method.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredmcpserverdetails.html#cfn-devopsagent-service-registeredmcpserverdetails-authorizationmethod
	//
	AuthorizationMethod *string `field:"required" json:"authorizationMethod" yaml:"authorizationMethod"`
	// MCP server endpoint URL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredmcpserverdetails.html#cfn-devopsagent-service-registeredmcpserverdetails-endpoint
	//
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// MCP server name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredmcpserverdetails.html#cfn-devopsagent-service-registeredmcpserverdetails-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// API key header name if using API key authentication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredmcpserverdetails.html#cfn-devopsagent-service-registeredmcpserverdetails-apikeyheader
	//
	ApiKeyHeader *string `field:"optional" json:"apiKeyHeader" yaml:"apiKeyHeader"`
	// Optional description for the MCP server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredmcpserverdetails.html#cfn-devopsagent-service-registeredmcpserverdetails-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

