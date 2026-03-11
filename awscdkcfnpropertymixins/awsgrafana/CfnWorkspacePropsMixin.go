package awsgrafana

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsgrafana/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies a *workspace* .
//
// In a workspace, you can create Grafana dashboards and visualizations to analyze your metrics, logs, and traces. You don't have to build, package, or deploy any hardware to run the Grafana server.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnWorkspacePropsMixin := awscdkcfnpropertymixins.Aws_grafana.NewCfnWorkspacePropsMixin(&CfnWorkspaceMixinProps{
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
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-grafana-workspace.html
//
type CfnWorkspacePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnWorkspaceMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnWorkspacePropsMixin
type jsiiProxy_CfnWorkspacePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
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

func (j *jsiiProxy_CfnWorkspacePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Grafana::Workspace`.
func NewCfnWorkspacePropsMixin(props *CfnWorkspaceMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnWorkspacePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnWorkspacePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnWorkspacePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_grafana.CfnWorkspacePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Grafana::Workspace`.
func NewCfnWorkspacePropsMixin_Override(c CfnWorkspacePropsMixin, props *CfnWorkspaceMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_grafana.CfnWorkspacePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnWorkspacePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnWorkspacePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_grafana.CfnWorkspacePropsMixin",
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
		"@aws-cdk/cfn-property-mixins.aws_grafana.CfnWorkspacePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnWorkspacePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
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

