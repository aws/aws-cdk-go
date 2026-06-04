package awsdevopsagent


// Grafana MCP server authorization configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   mCPServerGrafanaAuthorizationConfigProperty := &MCPServerGrafanaAuthorizationConfigProperty{
//   	BearerToken: &BearerTokenDetailsProperty{
//   		AuthorizationHeader: jsii.String("authorizationHeader"),
//   		TokenName: jsii.String("tokenName"),
//   		TokenValue: jsii.String("tokenValue"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-mcpservergrafanaauthorizationconfig.html
//
type CfnServicePropsMixin_MCPServerGrafanaAuthorizationConfigProperty struct {
	// Bearer token authentication details.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-mcpservergrafanaauthorizationconfig.html#cfn-devopsagent-service-mcpservergrafanaauthorizationconfig-bearertoken
	//
	BearerToken interface{} `field:"optional" json:"bearerToken" yaml:"bearerToken"`
}

