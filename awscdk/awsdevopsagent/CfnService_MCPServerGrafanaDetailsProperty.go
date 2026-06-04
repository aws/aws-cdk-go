package awsdevopsagent


// Grafana MCP server configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mCPServerGrafanaDetailsProperty := &MCPServerGrafanaDetailsProperty{
//   	AuthorizationConfig: &MCPServerGrafanaAuthorizationConfigProperty{
//   		BearerToken: &BearerTokenDetailsProperty{
//   			TokenName: jsii.String("tokenName"),
//   			TokenValue: jsii.String("tokenValue"),
//
//   			// the properties below are optional
//   			AuthorizationHeader: jsii.String("authorizationHeader"),
//   		},
//   	},
//   	Endpoint: jsii.String("endpoint"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-mcpservergrafanadetails.html
//
type CfnService_MCPServerGrafanaDetailsProperty struct {
	// Grafana MCP server authorization configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-mcpservergrafanadetails.html#cfn-devopsagent-service-mcpservergrafanadetails-authorizationconfig
	//
	AuthorizationConfig interface{} `field:"required" json:"authorizationConfig" yaml:"authorizationConfig"`
	// MCP server endpoint URL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-mcpservergrafanadetails.html#cfn-devopsagent-service-mcpservergrafanadetails-endpoint
	//
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// MCP server name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-mcpservergrafanadetails.html#cfn-devopsagent-service-mcpservergrafanadetails-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Optional description for the MCP server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-mcpservergrafanadetails.html#cfn-devopsagent-service-mcpservergrafanadetails-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

