package previewawsbedrockagentcoremixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsbedrockagentcore/previewawsbedrockagentcoremixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// After creating a gateway, you can add targets, which define the tools that your gateway will host.
//
// For more information about adding gateway targets, see [Add targets to an existing gateway](https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/gateway-building-adding-targets.html) .
//
// See the *Properties* section below for descriptions of both the required and optional properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var schemaDefinitionProperty_ SchemaDefinitionProperty
//
//   cfnGatewayTargetPropsMixin := awscdkmixinspreview.Mixins.NewCfnGatewayTargetPropsMixin(&CfnGatewayTargetMixinProps{
//   	CredentialProviderConfigurations: []interface{}{
//   		&CredentialProviderConfigurationProperty{
//   			CredentialProvider: &CredentialProviderProperty{
//   				ApiKeyCredentialProvider: &ApiKeyCredentialProviderProperty{
//   					CredentialLocation: jsii.String("credentialLocation"),
//   					CredentialParameterName: jsii.String("credentialParameterName"),
//   					CredentialPrefix: jsii.String("credentialPrefix"),
//   					ProviderArn: jsii.String("providerArn"),
//   				},
//   				OauthCredentialProvider: &OAuthCredentialProviderProperty{
//   					CustomParameters: map[string]*string{
//   						"customParametersKey": jsii.String("customParameters"),
//   					},
//   					DefaultReturnUrl: jsii.String("defaultReturnUrl"),
//   					GrantType: jsii.String("grantType"),
//   					ProviderArn: jsii.String("providerArn"),
//   					Scopes: []*string{
//   						jsii.String("scopes"),
//   					},
//   				},
//   			},
//   			CredentialProviderType: jsii.String("credentialProviderType"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	GatewayIdentifier: jsii.String("gatewayIdentifier"),
//   	MetadataConfiguration: &MetadataConfigurationProperty{
//   		AllowedQueryParameters: []*string{
//   			jsii.String("allowedQueryParameters"),
//   		},
//   		AllowedRequestHeaders: []*string{
//   			jsii.String("allowedRequestHeaders"),
//   		},
//   		AllowedResponseHeaders: []*string{
//   			jsii.String("allowedResponseHeaders"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	TargetConfiguration: &TargetConfigurationProperty{
//   		Mcp: &McpTargetConfigurationProperty{
//   			ApiGateway: &ApiGatewayTargetConfigurationProperty{
//   				ApiGatewayToolConfiguration: &ApiGatewayToolConfigurationProperty{
//   					ToolFilters: []interface{}{
//   						&ApiGatewayToolFilterProperty{
//   							FilterPath: jsii.String("filterPath"),
//   							Methods: []*string{
//   								jsii.String("methods"),
//   							},
//   						},
//   					},
//   					ToolOverrides: []interface{}{
//   						&ApiGatewayToolOverrideProperty{
//   							Description: jsii.String("description"),
//   							Method: jsii.String("method"),
//   							Name: jsii.String("name"),
//   							Path: jsii.String("path"),
//   						},
//   					},
//   				},
//   				RestApiId: jsii.String("restApiId"),
//   				Stage: jsii.String("stage"),
//   			},
//   			Lambda: &McpLambdaTargetConfigurationProperty{
//   				LambdaArn: jsii.String("lambdaArn"),
//   				ToolSchema: &ToolSchemaProperty{
//   					InlinePayload: []interface{}{
//   						&ToolDefinitionProperty{
//   							Description: jsii.String("description"),
//   							InputSchema: &SchemaDefinitionProperty{
//   								Description: jsii.String("description"),
//   								Items: schemaDefinitionProperty_,
//   								Properties: map[string]interface{}{
//   									"propertiesKey": schemaDefinitionProperty_,
//   								},
//   								Required: []*string{
//   									jsii.String("required"),
//   								},
//   								Type: jsii.String("type"),
//   							},
//   							Name: jsii.String("name"),
//   							OutputSchema: &SchemaDefinitionProperty{
//   								Description: jsii.String("description"),
//   								Items: schemaDefinitionProperty_,
//   								Properties: map[string]interface{}{
//   									"propertiesKey": schemaDefinitionProperty_,
//   								},
//   								Required: []*string{
//   									jsii.String("required"),
//   								},
//   								Type: jsii.String("type"),
//   							},
//   						},
//   					},
//   					S3: &S3ConfigurationProperty{
//   						BucketOwnerAccountId: jsii.String("bucketOwnerAccountId"),
//   						Uri: jsii.String("uri"),
//   					},
//   				},
//   			},
//   			McpServer: &McpServerTargetConfigurationProperty{
//   				Endpoint: jsii.String("endpoint"),
//   			},
//   			OpenApiSchema: &ApiSchemaConfigurationProperty{
//   				InlinePayload: jsii.String("inlinePayload"),
//   				S3: &S3ConfigurationProperty{
//   					BucketOwnerAccountId: jsii.String("bucketOwnerAccountId"),
//   					Uri: jsii.String("uri"),
//   				},
//   			},
//   			SmithyModel: &ApiSchemaConfigurationProperty{
//   				InlinePayload: jsii.String("inlinePayload"),
//   				S3: &S3ConfigurationProperty{
//   					BucketOwnerAccountId: jsii.String("bucketOwnerAccountId"),
//   					Uri: jsii.String("uri"),
//   				},
//   			},
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gatewaytarget.html
//
type CfnGatewayTargetPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnGatewayTargetMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnGatewayTargetPropsMixin
type jsiiProxy_CfnGatewayTargetPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnGatewayTargetPropsMixin) Props() *CfnGatewayTargetMixinProps {
	var returns *CfnGatewayTargetMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGatewayTargetPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::BedrockAgentCore::GatewayTarget`.
func NewCfnGatewayTargetPropsMixin(props *CfnGatewayTargetMixinProps, options *mixins.CfnPropertyMixinOptions) CfnGatewayTargetPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnGatewayTargetPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnGatewayTargetPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayTargetPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::BedrockAgentCore::GatewayTarget`.
func NewCfnGatewayTargetPropsMixin_Override(c CfnGatewayTargetPropsMixin, props *CfnGatewayTargetMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayTargetPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnGatewayTargetPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnGatewayTargetPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayTargetPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnGatewayTargetPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayTargetPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnGatewayTargetPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGatewayTargetPropsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

