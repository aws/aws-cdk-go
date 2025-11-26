package previewawsopsworkscmmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsopsworkscm/previewawsopsworkscmmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::OpsWorksCM::Server` resource creates an AWS OpsWorks for Chef Automate or OpsWorks for Puppet Enterprise configuration management server.
//
// For more information, see [Create a Chef Automate Server in CloudFormation](https://docs.aws.amazon.com/opsworks/latest/userguide/opscm-create-server-cfn.html) or [Create a Puppet Enterprise Master in CloudFormation](https://docs.aws.amazon.com/opsworks/latest/userguide/opspup-create-server-cfn.html) in the *OpsWorks User Guide* , and [CreateServer](https://docs.aws.amazon.com/opsworks-cm/latest/APIReference/API_CreateServer.html) in the *AWS OpsWorks CM API Reference* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnServerPropsMixin := awscdkmixinspreview.Mixins.NewCfnServerPropsMixin(&CfnServerMixinProps{
//   	AssociatePublicIpAddress: jsii.Boolean(false),
//   	BackupId: jsii.String("backupId"),
//   	BackupRetentionCount: jsii.Number(123),
//   	CustomCertificate: jsii.String("customCertificate"),
//   	CustomDomain: jsii.String("customDomain"),
//   	CustomPrivateKey: jsii.String("customPrivateKey"),
//   	DisableAutomatedBackup: jsii.Boolean(false),
//   	Engine: jsii.String("engine"),
//   	EngineAttributes: []interface{}{
//   		&EngineAttributeProperty{
//   			Name: jsii.String("name"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	EngineModel: jsii.String("engineModel"),
//   	EngineVersion: jsii.String("engineVersion"),
//   	InstanceProfileArn: jsii.String("instanceProfileArn"),
//   	InstanceType: jsii.String("instanceType"),
//   	KeyPair: jsii.String("keyPair"),
//   	PreferredBackupWindow: jsii.String("preferredBackupWindow"),
//   	PreferredMaintenanceWindow: jsii.String("preferredMaintenanceWindow"),
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	ServiceRoleArn: jsii.String("serviceRoleArn"),
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworkscm-server.html
//
type CfnServerPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnServerMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnServerPropsMixin
type jsiiProxy_CfnServerPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnServerPropsMixin) Props() *CfnServerMixinProps {
	var returns *CfnServerMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServerPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::OpsWorksCM::Server`.
func NewCfnServerPropsMixin(props *CfnServerMixinProps, options *mixins.CfnPropertyMixinOptions) CfnServerPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnServerPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnServerPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_opsworkscm.mixins.CfnServerPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::OpsWorksCM::Server`.
func NewCfnServerPropsMixin_Override(c CfnServerPropsMixin, props *CfnServerMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_opsworkscm.mixins.CfnServerPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnServerPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnServerPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_opsworkscm.mixins.CfnServerPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnServerPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_opsworkscm.mixins.CfnServerPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnServerPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnServerPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

