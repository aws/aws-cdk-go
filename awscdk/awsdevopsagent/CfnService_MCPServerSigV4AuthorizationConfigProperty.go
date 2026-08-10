package awsdevopsagent


// SigV4 authorization configuration for MCP server.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mCPServerSigV4AuthorizationConfigProperty := &MCPServerSigV4AuthorizationConfigProperty{
//   	Region: jsii.String("region"),
//   	Service: jsii.String("service"),
//
//   	// the properties below are optional
//   	CustomHeaders: map[string]*string{
//   		"customHeadersKey": jsii.String("customHeaders"),
//   	},
//   	McpRoleArn: jsii.String("mcpRoleArn"),
//   	RoleArn: jsii.String("roleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-mcpserversigv4authorizationconfig.html
//
type CfnService_MCPServerSigV4AuthorizationConfigProperty struct {
	// AWS region for SigV4 signing.
	//
	// Use '*' for SigV4a multi-region signing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-mcpserversigv4authorizationconfig.html#cfn-devopsagent-service-mcpserversigv4authorizationconfig-region
	//
	Region *string `field:"required" json:"region" yaml:"region"`
	// AWS service name for SigV4 signing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-mcpserversigv4authorizationconfig.html#cfn-devopsagent-service-mcpserversigv4authorizationconfig-service
	//
	Service *string `field:"required" json:"service" yaml:"service"`
	// Custom headers for the SigV4 MCP server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-mcpserversigv4authorizationconfig.html#cfn-devopsagent-service-mcpserversigv4authorizationconfig-customheaders
	//
	CustomHeaders interface{} `field:"optional" json:"customHeaders" yaml:"customHeaders"`
	// IAM role ARN to assume for SigV4 signing.
	//
	// Optional - when omitted, credentials are resolved at runtime via a monitor account association.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-mcpserversigv4authorizationconfig.html#cfn-devopsagent-service-mcpserversigv4authorizationconfig-mcprolearn
	//
	McpRoleArn *string `field:"optional" json:"mcpRoleArn" yaml:"mcpRoleArn"`
	// Deprecated - use McpRoleArn instead.
	//
	// IAM role ARN to assume for SigV4 signing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-mcpserversigv4authorizationconfig.html#cfn-devopsagent-service-mcpserversigv4authorizationconfig-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

