package awsdevopsagent


// Grafana MCP server configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   mCPServerGrafanaDetailsProperty := &MCPServerGrafanaDetailsProperty{
//   	AuthorizationConfig: &MCPServerGrafanaAuthorizationConfigProperty{
//   		BearerToken: &BearerTokenDetailsProperty{
//   			AuthorizationHeader: jsii.String("authorizationHeader"),
//   			TokenName: jsii.String("tokenName"),
//   			TokenValue: jsii.String("tokenValue"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	Endpoint: jsii.String("endpoint"),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-mcpservergrafanadetails.html
//
type CfnServicePropsMixin_MCPServerGrafanaDetailsProperty struct {
	// Grafana MCP server authorization configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-mcpservergrafanadetails.html#cfn-devopsagent-service-mcpservergrafanadetails-authorizationconfig
	//
	AuthorizationConfig interface{} `field:"optional" json:"authorizationConfig" yaml:"authorizationConfig"`
	// Optional description for the MCP server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-mcpservergrafanadetails.html#cfn-devopsagent-service-mcpservergrafanadetails-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// MCP server endpoint URL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-mcpservergrafanadetails.html#cfn-devopsagent-service-mcpservergrafanadetails-endpoint
	//
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// MCP server name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-mcpservergrafanadetails.html#cfn-devopsagent-service-mcpservergrafanadetails-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

