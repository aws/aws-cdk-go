package awssecurityagent

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awssecurityagent/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::SecurityAgent::Integration.
//
// Creates a connection between Security Agent and a third-party provider for code review and remediation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnIntegrationPropsMixin := awscdkcfnpropertymixins.Aws_securityagent.NewCfnIntegrationPropsMixin(&CfnIntegrationMixinProps{
//   	DisplayName: jsii.String("displayName"),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	PrivateConnectionName: jsii.String("privateConnectionName"),
//   	Provider: jsii.String("provider"),
//   	ProviderInput: &ProviderInputProperty{
//   		Bitbucket: &BitbucketInputProperty{
//   			Code: jsii.String("code"),
//   			InstallationId: jsii.String("installationId"),
//   			State: jsii.String("state"),
//   			Workspace: jsii.String("workspace"),
//   		},
//   		Confluence: &ConfluenceInputProperty{
//   			Code: jsii.String("code"),
//   			InstallationId: jsii.String("installationId"),
//   			SiteUrl: jsii.String("siteUrl"),
//   			State: jsii.String("state"),
//   		},
//   		Github: &GitHubInputProperty{
//   			Code: jsii.String("code"),
//   			InstallationId: jsii.String("installationId"),
//   			OrganizationName: jsii.String("organizationName"),
//   			State: jsii.String("state"),
//   			TargetUrl: jsii.String("targetUrl"),
//   		},
//   		Gitlab: &GitLabInputProperty{
//   			AccessToken: jsii.String("accessToken"),
//   			GroupId: jsii.String("groupId"),
//   			TargetUrl: jsii.String("targetUrl"),
//   			TokenType: jsii.String("tokenType"),
//   		},
//   	},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityagent-integration.html
//
type CfnIntegrationPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnIntegrationMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnIntegrationPropsMixin
type jsiiProxy_CfnIntegrationPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnIntegrationPropsMixin) Props() *CfnIntegrationMixinProps {
	var returns *CfnIntegrationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegrationPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SecurityAgent::Integration`.
func NewCfnIntegrationPropsMixin(props *CfnIntegrationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnIntegrationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnIntegrationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnIntegrationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_securityagent.CfnIntegrationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SecurityAgent::Integration`.
func NewCfnIntegrationPropsMixin_Override(c CfnIntegrationPropsMixin, props *CfnIntegrationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_securityagent.CfnIntegrationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnIntegrationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnIntegrationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_securityagent.CfnIntegrationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnIntegrationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_securityagent.CfnIntegrationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnIntegrationPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnIntegrationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

