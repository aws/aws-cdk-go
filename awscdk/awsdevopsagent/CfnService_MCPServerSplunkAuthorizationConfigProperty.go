package awsdevopsagent


// MCP server splunk authorization configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mCPServerSplunkAuthorizationConfigProperty := &MCPServerSplunkAuthorizationConfigProperty{
//   	BearerToken: &BearerTokenDetailsProperty{
//   		TokenName: jsii.String("tokenName"),
//   		TokenValue: jsii.String("tokenValue"),
//
//   		// the properties below are optional
//   		AuthorizationHeader: jsii.String("authorizationHeader"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-mcpserversplunkauthorizationconfig.html
//
type CfnService_MCPServerSplunkAuthorizationConfigProperty struct {
	// Bearer token authentication details.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-mcpserversplunkauthorizationconfig.html#cfn-devopsagent-service-mcpserversplunkauthorizationconfig-bearertoken
	//
	BearerToken interface{} `field:"required" json:"bearerToken" yaml:"bearerToken"`
}

