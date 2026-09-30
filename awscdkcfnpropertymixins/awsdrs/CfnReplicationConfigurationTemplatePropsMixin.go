package awsdrs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsdrs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A replication configuration template for AWS Elastic Disaster Recovery Service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnReplicationConfigurationTemplatePropsMixin := awscdkcfnpropertymixins.Aws_drs.NewCfnReplicationConfigurationTemplatePropsMixin(&CfnReplicationConfigurationTemplateMixinProps{
//   	AssociateDefaultSecurityGroup: jsii.Boolean(false),
//   	AutoReplicateNewDisks: jsii.Boolean(false),
//   	BandwidthThrottling: jsii.Number(123),
//   	CreatePublicIp: jsii.Boolean(false),
//   	DataPlaneRouting: jsii.String("dataPlaneRouting"),
//   	DefaultLargeStagingDiskType: jsii.String("defaultLargeStagingDiskType"),
//   	EbsEncryption: jsii.String("ebsEncryption"),
//   	EbsEncryptionKeyArn: jsii.String("ebsEncryptionKeyArn"),
//   	InternetProtocol: jsii.String("internetProtocol"),
//   	PitPolicy: []interface{}{
//   		&PITPolicyRuleProperty{
//   			Enabled: jsii.Boolean(false),
//   			Interval: jsii.Number(123),
//   			RetentionDuration: jsii.Number(123),
//   			RuleId: jsii.Number(123),
//   			Units: jsii.String("units"),
//   		},
//   	},
//   	ReplicationServerInstanceType: jsii.String("replicationServerInstanceType"),
//   	ReplicationServersSecurityGroupsIDs: []*string{
//   		jsii.String("replicationServersSecurityGroupsIDs"),
//   	},
//   	StagingAreaSubnetId: jsii.String("stagingAreaSubnetId"),
//   	StagingAreaTags: map[string]*string{
//   		"stagingAreaTagsKey": jsii.String("stagingAreaTags"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	UseDedicatedReplicationServer: jsii.Boolean(false),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-drs-replicationconfigurationtemplate.html
//
type CfnReplicationConfigurationTemplatePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnReplicationConfigurationTemplateMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnReplicationConfigurationTemplatePropsMixin
type jsiiProxy_CfnReplicationConfigurationTemplatePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnReplicationConfigurationTemplatePropsMixin) Props() *CfnReplicationConfigurationTemplateMixinProps {
	var returns *CfnReplicationConfigurationTemplateMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationConfigurationTemplatePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::DRS::ReplicationConfigurationTemplate`.
func NewCfnReplicationConfigurationTemplatePropsMixin(props *CfnReplicationConfigurationTemplateMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnReplicationConfigurationTemplatePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnReplicationConfigurationTemplatePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnReplicationConfigurationTemplatePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_drs.CfnReplicationConfigurationTemplatePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::DRS::ReplicationConfigurationTemplate`.
func NewCfnReplicationConfigurationTemplatePropsMixin_Override(c CfnReplicationConfigurationTemplatePropsMixin, props *CfnReplicationConfigurationTemplateMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_drs.CfnReplicationConfigurationTemplatePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnReplicationConfigurationTemplatePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnReplicationConfigurationTemplatePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_drs.CfnReplicationConfigurationTemplatePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnReplicationConfigurationTemplatePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_drs.CfnReplicationConfigurationTemplatePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnReplicationConfigurationTemplatePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnReplicationConfigurationTemplatePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

