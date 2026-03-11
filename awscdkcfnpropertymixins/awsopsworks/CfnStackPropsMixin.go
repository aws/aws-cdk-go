package awsopsworks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsopsworks/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var customJson interface{}
//   var mergeStrategy IMergeStrategy
//
//   cfnStackPropsMixin := awscdkcfnpropertymixins.Aws_opsworks.NewCfnStackPropsMixin(&CfnStackMixinProps{
//   	AgentVersion: jsii.String("agentVersion"),
//   	Attributes: map[string]*string{
//   		"attributesKey": jsii.String("attributes"),
//   	},
//   	ChefConfiguration: &ChefConfigurationProperty{
//   		BerkshelfVersion: jsii.String("berkshelfVersion"),
//   		ManageBerkshelf: jsii.Boolean(false),
//   	},
//   	CloneAppIds: []*string{
//   		jsii.String("cloneAppIds"),
//   	},
//   	ClonePermissions: jsii.Boolean(false),
//   	ConfigurationManager: &StackConfigurationManagerProperty{
//   		Name: jsii.String("name"),
//   		Version: jsii.String("version"),
//   	},
//   	CustomCookbooksSource: &SourceProperty{
//   		Password: jsii.String("password"),
//   		Revision: jsii.String("revision"),
//   		SshKey: jsii.String("sshKey"),
//   		Type: jsii.String("type"),
//   		Url: jsii.String("url"),
//   		Username: jsii.String("username"),
//   	},
//   	CustomJson: customJson,
//   	DefaultAvailabilityZone: jsii.String("defaultAvailabilityZone"),
//   	DefaultInstanceProfileArn: jsii.String("defaultInstanceProfileArn"),
//   	DefaultOs: jsii.String("defaultOs"),
//   	DefaultRootDeviceType: jsii.String("defaultRootDeviceType"),
//   	DefaultSshKeyName: jsii.String("defaultSshKeyName"),
//   	DefaultSubnetId: jsii.String("defaultSubnetId"),
//   	EcsClusterArn: jsii.String("ecsClusterArn"),
//   	ElasticIps: []interface{}{
//   		&ElasticIpProperty{
//   			Ip: jsii.String("ip"),
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	HostnameTheme: jsii.String("hostnameTheme"),
//   	Name: jsii.String("name"),
//   	RdsDbInstances: []interface{}{
//   		&RdsDbInstanceProperty{
//   			DbPassword: jsii.String("dbPassword"),
//   			DbUser: jsii.String("dbUser"),
//   			RdsDbInstanceArn: jsii.String("rdsDbInstanceArn"),
//   		},
//   	},
//   	ServiceRoleArn: jsii.String("serviceRoleArn"),
//   	SourceStackId: jsii.String("sourceStackId"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	UseCustomCookbooks: jsii.Boolean(false),
//   	UseOpsworksSecurityGroups: jsii.Boolean(false),
//   	VpcId: jsii.String("vpcId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html
//
type CfnStackPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnStackMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnStackPropsMixin
type jsiiProxy_CfnStackPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnStackPropsMixin) Props() *CfnStackMixinProps {
	var returns *CfnStackMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::OpsWorks::Stack`.
func NewCfnStackPropsMixin(props *CfnStackMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnStackPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnStackPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnStackPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_opsworks.CfnStackPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::OpsWorks::Stack`.
func NewCfnStackPropsMixin_Override(c CfnStackPropsMixin, props *CfnStackMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_opsworks.CfnStackPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnStackPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnStackPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_opsworks.CfnStackPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnStackPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_opsworks.CfnStackPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnStackPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnStackPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

