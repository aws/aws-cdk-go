package mixinsawsamplify

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsamplify/mixinsawsamplify/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// The AWS::Amplify::App resource specifies Apps in Amplify Hosting.
//
// An App is a collection of branches.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAppPropsMixin := awscdkmixinspreview.Mixins.NewCfnAppPropsMixin(&CfnAppMixinProps{
//   	AccessToken: jsii.String("accessToken"),
//   	AutoBranchCreationConfig: &AutoBranchCreationConfigProperty{
//   		AutoBranchCreationPatterns: []*string{
//   			jsii.String("autoBranchCreationPatterns"),
//   		},
//   		BasicAuthConfig: &BasicAuthConfigProperty{
//   			EnableBasicAuth: jsii.Boolean(false),
//   			Password: jsii.String("password"),
//   			Username: jsii.String("username"),
//   		},
//   		BuildSpec: jsii.String("buildSpec"),
//   		EnableAutoBranchCreation: jsii.Boolean(false),
//   		EnableAutoBuild: jsii.Boolean(false),
//   		EnablePerformanceMode: jsii.Boolean(false),
//   		EnablePullRequestPreview: jsii.Boolean(false),
//   		EnvironmentVariables: []interface{}{
//   			&EnvironmentVariableProperty{
//   				Name: jsii.String("name"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		Framework: jsii.String("framework"),
//   		PullRequestEnvironmentName: jsii.String("pullRequestEnvironmentName"),
//   		Stage: jsii.String("stage"),
//   	},
//   	BasicAuthConfig: &BasicAuthConfigProperty{
//   		EnableBasicAuth: jsii.Boolean(false),
//   		Password: jsii.String("password"),
//   		Username: jsii.String("username"),
//   	},
//   	BuildSpec: jsii.String("buildSpec"),
//   	CacheConfig: &CacheConfigProperty{
//   		Type: jsii.String("type"),
//   	},
//   	ComputeRoleArn: jsii.String("computeRoleArn"),
//   	CustomHeaders: jsii.String("customHeaders"),
//   	CustomRules: []interface{}{
//   		&CustomRuleProperty{
//   			Condition: jsii.String("condition"),
//   			Source: jsii.String("source"),
//   			Status: jsii.String("status"),
//   			Target: jsii.String("target"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	EnableBranchAutoDeletion: jsii.Boolean(false),
//   	EnvironmentVariables: []interface{}{
//   		&EnvironmentVariableProperty{
//   			Name: jsii.String("name"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	IamServiceRole: jsii.String("iamServiceRole"),
//   	JobConfig: &JobConfigProperty{
//   		BuildComputeType: jsii.String("buildComputeType"),
//   	},
//   	Name: jsii.String("name"),
//   	OauthToken: jsii.String("oauthToken"),
//   	Platform: jsii.String("platform"),
//   	Repository: jsii.String("repository"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplify-app.html
//
type CfnAppPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnAppMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAppPropsMixin
type jsiiProxy_CfnAppPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnAppPropsMixin) Props() *CfnAppMixinProps {
	var returns *CfnAppMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAppPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Amplify::App`.
func NewCfnAppPropsMixin(props *CfnAppMixinProps, options *mixins.CfnPropertyMixinOptions) CfnAppPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAppPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAppPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_amplify.mixins.CfnAppPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Amplify::App`.
func NewCfnAppPropsMixin_Override(c CfnAppPropsMixin, props *CfnAppMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_amplify.mixins.CfnAppPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnAppPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAppPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_amplify.mixins.CfnAppPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAppPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_amplify.mixins.CfnAppPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAppPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnAppPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

