package previewawsdevopsagentmixins


// MCP server splunk authorization configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   mCPServerSplunkAuthorizationConfigProperty := &MCPServerSplunkAuthorizationConfigProperty{
//   	BearerToken: &BearerTokenDetailsProperty{
//   		AuthorizationHeader: jsii.String("authorizationHeader"),
//   		TokenName: jsii.String("tokenName"),
//   		TokenValue: jsii.String("tokenValue"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-mcpserversplunkauthorizationconfig.html
//
type CfnServicePropsMixin_MCPServerSplunkAuthorizationConfigProperty struct {
	// Bearer token authentication details.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-mcpserversplunkauthorizationconfig.html#cfn-devopsagent-service-mcpserversplunkauthorizationconfig-bearertoken
	//
	BearerToken interface{} `field:"optional" json:"bearerToken" yaml:"bearerToken"`
}

