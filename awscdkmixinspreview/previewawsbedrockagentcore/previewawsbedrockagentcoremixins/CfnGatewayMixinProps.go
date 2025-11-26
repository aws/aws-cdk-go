package previewawsbedrockagentcoremixins


// Properties for CfnGatewayPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnGatewayMixinProps := &CfnGatewayMixinProps{
//   	AuthorizerConfiguration: &AuthorizerConfigurationProperty{
//   		CustomJwtAuthorizer: &CustomJWTAuthorizerConfigurationProperty{
//   			AllowedAudience: []*string{
//   				jsii.String("allowedAudience"),
//   			},
//   			AllowedClients: []*string{
//   				jsii.String("allowedClients"),
//   			},
//   			DiscoveryUrl: jsii.String("discoveryUrl"),
//   		},
//   	},
//   	AuthorizerType: jsii.String("authorizerType"),
//   	Description: jsii.String("description"),
//   	ExceptionLevel: jsii.String("exceptionLevel"),
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	Name: jsii.String("name"),
//   	ProtocolConfiguration: &GatewayProtocolConfigurationProperty{
//   		Mcp: &MCPGatewayConfigurationProperty{
//   			Instructions: jsii.String("instructions"),
//   			SearchType: jsii.String("searchType"),
//   			SupportedVersions: []*string{
//   				jsii.String("supportedVersions"),
//   			},
//   		},
//   	},
//   	ProtocolType: jsii.String("protocolType"),
//   	RoleArn: jsii.String("roleArn"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gateway.html
//
type CfnGatewayMixinProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gateway.html#cfn-bedrockagentcore-gateway-authorizerconfiguration
	//
	AuthorizerConfiguration interface{} `field:"optional" json:"authorizerConfiguration" yaml:"authorizerConfiguration"`
	// The authorizer type for the gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gateway.html#cfn-bedrockagentcore-gateway-authorizertype
	//
	AuthorizerType *string `field:"optional" json:"authorizerType" yaml:"authorizerType"`
	// The description for the gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gateway.html#cfn-bedrockagentcore-gateway-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The exception level for the gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gateway.html#cfn-bedrockagentcore-gateway-exceptionlevel
	//
	ExceptionLevel *string `field:"optional" json:"exceptionLevel" yaml:"exceptionLevel"`
	// The KMS key ARN for the gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gateway.html#cfn-bedrockagentcore-gateway-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// The name for the gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gateway.html#cfn-bedrockagentcore-gateway-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The protocol configuration for the gateway target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gateway.html#cfn-bedrockagentcore-gateway-protocolconfiguration
	//
	ProtocolConfiguration interface{} `field:"optional" json:"protocolConfiguration" yaml:"protocolConfiguration"`
	// The protocol type for the gateway target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gateway.html#cfn-bedrockagentcore-gateway-protocoltype
	//
	ProtocolType *string `field:"optional" json:"protocolType" yaml:"protocolType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gateway.html#cfn-bedrockagentcore-gateway-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The tags for the gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gateway.html#cfn-bedrockagentcore-gateway-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

