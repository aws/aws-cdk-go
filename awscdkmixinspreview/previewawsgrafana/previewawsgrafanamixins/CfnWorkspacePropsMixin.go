package previewawsgrafanamixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsgrafana/previewawsgrafanamixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies a *workspace* .
//
// In a workspace, you can create Grafana dashboards and visualizations to analyze your metrics, logs, and traces. You don't have to build, package, or deploy any hardware to run the Grafana server.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnWorkspacePropsMixin := awscdkmixinspreview.Mixins.NewCfnWorkspacePropsMixin(&CfnWorkspaceMixinProps{
//   	AccountAccessType: jsii.String("accountAccessType"),
//   	AuthenticationProviders: []*string{
//   		jsii.String("authenticationProviders"),
//   	},
//   	ClientToken: jsii.String("clientToken"),
//   	DataSources: []*string{
//   		jsii.String("dataSources"),
//   	},
//   	Description: jsii.String("description"),
//   	GrafanaVersion: jsii.String("grafanaVersion"),
//   	Name: jsii.String("name"),
//   	NetworkAccessControl: &NetworkAccessControlProperty{
//   		PrefixListIds: []*string{
//   			jsii.String("prefixListIds"),
//   		},
//   		VpceIds: []*string{
//   			jsii.String("vpceIds"),
//   		},
//   	},
//   	NotificationDestinations: []*string{
//   		jsii.String("notificationDestinations"),
//   	},
//   	OrganizationalUnits: []*string{
//   		jsii.String("organizationalUnits"),
//   	},
//   	OrganizationRoleName: jsii.String("organizationRoleName"),
//   	PermissionType: jsii.String("permissionType"),
//   	PluginAdminEnabled: jsii.Boolean(false),
//   	RoleArn: jsii.String("roleArn"),
//   	SamlConfiguration: &SamlConfigurationProperty{
//   		AllowedOrganizations: []*string{
//   			jsii.String("allowedOrganizations"),
//   		},
//   		AssertionAttributes: &AssertionAttributesProperty{
//   			Email: jsii.String("email"),
//   			Groups: jsii.String("groups"),
//   			Login: jsii.String("login"),
//   			Name: jsii.String("name"),
//   			Org: jsii.String("org"),
//   			Role: jsii.String("role"),
//   		},
//   		IdpMetadata: &IdpMetadataProperty{
//   			Url: jsii.String("url"),
//   			Xml: jsii.String("xml"),
//   		},
//   		LoginValidityDuration: jsii.Number(123),
//   		RoleValues: &RoleValuesProperty{
//   			Admin: []*string{
//   				jsii.String("admin"),
//   			},
//   			Editor: []*string{
//   				jsii.String("editor"),
//   			},
//   		},
//   	},
//   	StackSetName: jsii.String("stackSetName"),
//   	VpcConfiguration: &VpcConfigurationProperty{
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-grafana-workspace.html
//
type CfnWorkspacePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnWorkspaceMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnWorkspacePropsMixin
type jsiiProxy_CfnWorkspacePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnWorkspacePropsMixin) Props() *CfnWorkspaceMixinProps {
	var returns *CfnWorkspaceMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspacePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Grafana::Workspace`.
func NewCfnWorkspacePropsMixin(props *CfnWorkspaceMixinProps, options *mixins.CfnPropertyMixinOptions) CfnWorkspacePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnWorkspacePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnWorkspacePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_grafana.mixins.CfnWorkspacePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Grafana::Workspace`.
func NewCfnWorkspacePropsMixin_Override(c CfnWorkspacePropsMixin, props *CfnWorkspaceMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_grafana.mixins.CfnWorkspacePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnWorkspacePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnWorkspacePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_grafana.mixins.CfnWorkspacePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnWorkspacePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_grafana.mixins.CfnWorkspacePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnWorkspacePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnWorkspacePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

