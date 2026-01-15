package previewawsquicksightmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsquicksight/previewawsquicksightmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Definition of the AWS::QuickSight::ActionConnector Resource Type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnActionConnectorPropsMixin := awscdkmixinspreview.Mixins.NewCfnActionConnectorPropsMixin(&CfnActionConnectorMixinProps{
//   	ActionConnectorId: jsii.String("actionConnectorId"),
//   	AuthenticationConfig: &AuthConfigProperty{
//   		AuthenticationMetadata: &AuthenticationMetadataProperty{
//   			ApiKeyConnectionMetadata: &APIKeyConnectionMetadataProperty{
//   				ApiKey: jsii.String("apiKey"),
//   				BaseEndpoint: jsii.String("baseEndpoint"),
//   				Email: jsii.String("email"),
//   			},
//   			AuthorizationCodeGrantMetadata: &AuthorizationCodeGrantMetadataProperty{
//   				AuthorizationCodeGrantCredentialsDetails: &AuthorizationCodeGrantCredentialsDetailsProperty{
//   					AuthorizationCodeGrantDetails: &AuthorizationCodeGrantDetailsProperty{
//   						AuthorizationEndpoint: jsii.String("authorizationEndpoint"),
//   						ClientId: jsii.String("clientId"),
//   						ClientSecret: jsii.String("clientSecret"),
//   						TokenEndpoint: jsii.String("tokenEndpoint"),
//   					},
//   				},
//   				AuthorizationCodeGrantCredentialsSource: jsii.String("authorizationCodeGrantCredentialsSource"),
//   				BaseEndpoint: jsii.String("baseEndpoint"),
//   				RedirectUrl: jsii.String("redirectUrl"),
//   			},
//   			BasicAuthConnectionMetadata: &BasicAuthConnectionMetadataProperty{
//   				BaseEndpoint: jsii.String("baseEndpoint"),
//   				Password: jsii.String("password"),
//   				Username: jsii.String("username"),
//   			},
//   			ClientCredentialsGrantMetadata: &ClientCredentialsGrantMetadataProperty{
//   				BaseEndpoint: jsii.String("baseEndpoint"),
//   				ClientCredentialsDetails: &ClientCredentialsDetailsProperty{
//   					ClientCredentialsGrantDetails: &ClientCredentialsGrantDetailsProperty{
//   						ClientId: jsii.String("clientId"),
//   						ClientSecret: jsii.String("clientSecret"),
//   						TokenEndpoint: jsii.String("tokenEndpoint"),
//   					},
//   				},
//   				ClientCredentialsSource: jsii.String("clientCredentialsSource"),
//   			},
//   			IamConnectionMetadata: map[string]*string{
//   				"roleArn": jsii.String("roleArn"),
//   			},
//   			NoneConnectionMetadata: &NoneConnectionMetadataProperty{
//   				BaseEndpoint: jsii.String("baseEndpoint"),
//   			},
//   		},
//   		AuthenticationType: jsii.String("authenticationType"),
//   	},
//   	AwsAccountId: jsii.String("awsAccountId"),
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Permissions: []interface{}{
//   		&ResourcePermissionProperty{
//   			Actions: []*string{
//   				jsii.String("actions"),
//   			},
//   			Principal: jsii.String("principal"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   	VpcConnectionArn: jsii.String("vpcConnectionArn"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-actionconnector.html
//
type CfnActionConnectorPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnActionConnectorMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnActionConnectorPropsMixin
type jsiiProxy_CfnActionConnectorPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnActionConnectorPropsMixin) Props() *CfnActionConnectorMixinProps {
	var returns *CfnActionConnectorMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnActionConnectorPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::QuickSight::ActionConnector`.
func NewCfnActionConnectorPropsMixin(props *CfnActionConnectorMixinProps, options *mixins.CfnPropertyMixinOptions) CfnActionConnectorPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnActionConnectorPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnActionConnectorPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_quicksight.mixins.CfnActionConnectorPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::QuickSight::ActionConnector`.
func NewCfnActionConnectorPropsMixin_Override(c CfnActionConnectorPropsMixin, props *CfnActionConnectorMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_quicksight.mixins.CfnActionConnectorPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnActionConnectorPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnActionConnectorPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_quicksight.mixins.CfnActionConnectorPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnActionConnectorPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_quicksight.mixins.CfnActionConnectorPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnActionConnectorPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnActionConnectorPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

