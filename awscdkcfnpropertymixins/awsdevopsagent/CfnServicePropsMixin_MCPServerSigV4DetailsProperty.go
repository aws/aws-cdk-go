package awsdevopsagent


// SigV4-authenticated MCP server configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   mCPServerSigV4DetailsProperty := &MCPServerSigV4DetailsProperty{
//   	AuthorizationConfig: &MCPServerSigV4AuthorizationConfigProperty{
//   		CustomHeaders: map[string]*string{
//   			"customHeadersKey": jsii.String("customHeaders"),
//   		},
//   		Region: jsii.String("region"),
//   		RoleArn: jsii.String("roleArn"),
//   		Service: jsii.String("service"),
//   	},
//   	Description: jsii.String("description"),
//   	Endpoint: jsii.String("endpoint"),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-mcpserversigv4details.html
//
type CfnServicePropsMixin_MCPServerSigV4DetailsProperty struct {
	// SigV4 authorization configuration for MCP server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-mcpserversigv4details.html#cfn-devopsagent-service-mcpserversigv4details-authorizationconfig
	//
	AuthorizationConfig interface{} `field:"optional" json:"authorizationConfig" yaml:"authorizationConfig"`
	// Optional description for the MCP server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-mcpserversigv4details.html#cfn-devopsagent-service-mcpserversigv4details-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// MCP server endpoint URL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-mcpserversigv4details.html#cfn-devopsagent-service-mcpserversigv4details-endpoint
	//
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// MCP server name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-mcpserversigv4details.html#cfn-devopsagent-service-mcpserversigv4details-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

