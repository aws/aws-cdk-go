package awsbedrockagentcore


// Properties for defining a `CfnGateway`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGatewayProps := &CfnGatewayProps{
//   	AuthorizerType: jsii.String("authorizerType"),
//   	Name: jsii.String("name"),
//   	ProtocolType: jsii.String("protocolType"),
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	AuthorizerConfiguration: &AuthorizerConfigurationProperty{
//   		CustomJwtAuthorizer: &CustomJWTAuthorizerConfigurationProperty{
//   			DiscoveryUrl: jsii.String("discoveryUrl"),
//
//   			// the properties below are optional
//   			AllowedAudience: []*string{
//   				jsii.String("allowedAudience"),
//   			},
//   			AllowedClients: []*string{
//   				jsii.String("allowedClients"),
//   			},
//   			AllowedScopes: []*string{
//   				jsii.String("allowedScopes"),
//   			},
//   			CustomClaims: []interface{}{
//   				&CustomClaimValidationTypeProperty{
//   					AuthorizingClaimMatchValue: &AuthorizingClaimMatchValueTypeProperty{
//   						ClaimMatchOperator: jsii.String("claimMatchOperator"),
//   						ClaimMatchValue: &ClaimMatchValueTypeProperty{
//   							MatchValueString: jsii.String("matchValueString"),
//   							MatchValueStringList: []*string{
//   								jsii.String("matchValueStringList"),
//   							},
//   						},
//   					},
//   					InboundTokenClaimName: jsii.String("inboundTokenClaimName"),
//   					InboundTokenClaimValueType: jsii.String("inboundTokenClaimValueType"),
//   				},
//   			},
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	ExceptionLevel: jsii.String("exceptionLevel"),
//   	InterceptorConfigurations: []interface{}{
//   		&GatewayInterceptorConfigurationProperty{
//   			InterceptionPoints: []*string{
//   				jsii.String("interceptionPoints"),
//   			},
//   			Interceptor: &InterceptorConfigurationProperty{
//   				Lambda: &LambdaInterceptorConfigurationProperty{
//   					Arn: jsii.String("arn"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			InputConfiguration: &InterceptorInputConfigurationProperty{
//   				PassRequestHeaders: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	ProtocolConfiguration: &GatewayProtocolConfigurationProperty{
//   		Mcp: &MCPGatewayConfigurationProperty{
//   			Instructions: jsii.String("instructions"),
//   			SearchType: jsii.String("searchType"),
//   			SupportedVersions: []*string{
//   				jsii.String("supportedVersions"),
//   			},
//   		},
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gateway.html
//
type CfnGatewayProps struct {
	// The authorizer type for the gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gateway.html#cfn-bedrockagentcore-gateway-authorizertype
	//
	AuthorizerType *string `field:"required" json:"authorizerType" yaml:"authorizerType"`
	// The name for the gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gateway.html#cfn-bedrockagentcore-gateway-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The protocol type for the gateway target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gateway.html#cfn-bedrockagentcore-gateway-protocoltype
	//
	ProtocolType *string `field:"required" json:"protocolType" yaml:"protocolType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gateway.html#cfn-bedrockagentcore-gateway-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gateway.html#cfn-bedrockagentcore-gateway-authorizerconfiguration
	//
	AuthorizerConfiguration interface{} `field:"optional" json:"authorizerConfiguration" yaml:"authorizerConfiguration"`
	// The description for the gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gateway.html#cfn-bedrockagentcore-gateway-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The exception level for the gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gateway.html#cfn-bedrockagentcore-gateway-exceptionlevel
	//
	ExceptionLevel *string `field:"optional" json:"exceptionLevel" yaml:"exceptionLevel"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gateway.html#cfn-bedrockagentcore-gateway-interceptorconfigurations
	//
	InterceptorConfigurations interface{} `field:"optional" json:"interceptorConfigurations" yaml:"interceptorConfigurations"`
	// The KMS key ARN for the gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gateway.html#cfn-bedrockagentcore-gateway-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// The protocol configuration for the gateway target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gateway.html#cfn-bedrockagentcore-gateway-protocolconfiguration
	//
	ProtocolConfiguration interface{} `field:"optional" json:"protocolConfiguration" yaml:"protocolConfiguration"`
	// The tags for the gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gateway.html#cfn-bedrockagentcore-gateway-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

