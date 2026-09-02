package awsagentregistry

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsagentregistry/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Definition of AWS::AgentRegistry::RegistryRecord Resource Type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnRegistryRecordPropsMixin := awscdkcfnpropertymixins.Aws_agentregistry.NewCfnRegistryRecordPropsMixin(&CfnRegistryRecordMixinProps{
//   	Description: jsii.String("description"),
//   	Descriptors: &DescriptorsProperty{
//   		A2AAgentCard: &A2aAgentCardDescriptorProperty{
//   			Data: jsii.String("data"),
//   			DataSchemaVersion: jsii.String("dataSchemaVersion"),
//   			Source: &DescriptorSourceProperty{
//   				FromUrl: &DescriptorSourceFromUrlProperty{
//   					CredentialProviderConfigurations: []interface{}{
//   						&RegistryRecordCredentialProviderConfigurationProperty{
//   							CredentialProvider: &RegistryRecordCredentialProviderUnionProperty{
//   								IamCredentialProvider: &RegistryRecordIamCredentialProviderProperty{
//   									Region: jsii.String("region"),
//   									RoleArn: jsii.String("roleArn"),
//   									Service: jsii.String("service"),
//   								},
//   								OauthCredentialProvider: &RegistryRecordOAuthCredentialProviderProperty{
//   									CustomParameters: map[string]*string{
//   										"customParametersKey": jsii.String("customParameters"),
//   									},
//   									GrantType: jsii.String("grantType"),
//   									ProviderArn: jsii.String("providerArn"),
//   									Scopes: []*string{
//   										jsii.String("scopes"),
//   									},
//   								},
//   							},
//   							CredentialProviderType: jsii.String("credentialProviderType"),
//   						},
//   					},
//   					Url: jsii.String("url"),
//   				},
//   			},
//   		},
//   		AgentSkillsDefinition: &AgentSkillsDefinitionDescriptorProperty{
//   			AdditionalData: &AgentSkillsAdditionalDataProperty{
//   				SkillMd: &AgentSkillsMdDescriptorProperty{
//   					Data: jsii.String("data"),
//   					DataSchemaVersion: jsii.String("dataSchemaVersion"),
//   					Source: &SkillMdSourceProperty{
//   						FromUrl: &SkillMdSourceFromUrlProperty{
//   							Url: jsii.String("url"),
//   						},
//   					},
//   				},
//   			},
//   			Data: jsii.String("data"),
//   			DataSchemaVersion: jsii.String("dataSchemaVersion"),
//   		},
//   		Custom: &CustomDescriptorProperty{
//   			Data: jsii.String("data"),
//   		},
//   		McpServer: &McpServerDescriptorProperty{
//   			AdditionalData: &McpServerAdditionalDataProperty{
//   				Tools: &McpToolsDescriptorProperty{
//   					Data: jsii.String("data"),
//   					DataSchemaVersion: jsii.String("dataSchemaVersion"),
//   				},
//   			},
//   			Data: jsii.String("data"),
//   			DataSchemaVersion: jsii.String("dataSchemaVersion"),
//   			Source: &DescriptorSourceProperty{
//   				FromUrl: &DescriptorSourceFromUrlProperty{
//   					CredentialProviderConfigurations: []interface{}{
//   						&RegistryRecordCredentialProviderConfigurationProperty{
//   							CredentialProvider: &RegistryRecordCredentialProviderUnionProperty{
//   								IamCredentialProvider: &RegistryRecordIamCredentialProviderProperty{
//   									Region: jsii.String("region"),
//   									RoleArn: jsii.String("roleArn"),
//   									Service: jsii.String("service"),
//   								},
//   								OauthCredentialProvider: &RegistryRecordOAuthCredentialProviderProperty{
//   									CustomParameters: map[string]*string{
//   										"customParametersKey": jsii.String("customParameters"),
//   									},
//   									GrantType: jsii.String("grantType"),
//   									ProviderArn: jsii.String("providerArn"),
//   									Scopes: []*string{
//   										jsii.String("scopes"),
//   									},
//   								},
//   							},
//   							CredentialProviderType: jsii.String("credentialProviderType"),
//   						},
//   					},
//   					Url: jsii.String("url"),
//   				},
//   			},
//   		},
//   	},
//   	DisplayName: jsii.String("displayName"),
//   	Name: jsii.String("name"),
//   	RecordType: jsii.String("recordType"),
//   	RecordVersion: jsii.String("recordVersion"),
//   	RegistryId: jsii.String("registryId"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-agentregistry-registryrecord.html
//
type CfnRegistryRecordPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnRegistryRecordMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnRegistryRecordPropsMixin
type jsiiProxy_CfnRegistryRecordPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnRegistryRecordPropsMixin) Props() *CfnRegistryRecordMixinProps {
	var returns *CfnRegistryRecordMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegistryRecordPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::AgentRegistry::RegistryRecord`.
func NewCfnRegistryRecordPropsMixin(props *CfnRegistryRecordMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnRegistryRecordPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnRegistryRecordPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnRegistryRecordPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_agentregistry.CfnRegistryRecordPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::AgentRegistry::RegistryRecord`.
func NewCfnRegistryRecordPropsMixin_Override(c CfnRegistryRecordPropsMixin, props *CfnRegistryRecordMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_agentregistry.CfnRegistryRecordPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnRegistryRecordPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnRegistryRecordPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_agentregistry.CfnRegistryRecordPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRegistryRecordPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_agentregistry.CfnRegistryRecordPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnRegistryRecordPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnRegistryRecordPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

