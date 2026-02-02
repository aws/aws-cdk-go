package previewawsdmsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsdms/previewawsdmsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationconfig.html.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var replicationSettings interface{}
//   var supplementalSettings interface{}
//   var tableMappings interface{}
//
//   cfnReplicationConfigPropsMixin := awscdkmixinspreview.Mixins.NewCfnReplicationConfigPropsMixin(&CfnReplicationConfigMixinProps{
//   	ComputeConfig: &ComputeConfigProperty{
//   		AvailabilityZone: jsii.String("availabilityZone"),
//   		DnsNameServers: jsii.String("dnsNameServers"),
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   		MaxCapacityUnits: jsii.Number(123),
//   		MinCapacityUnits: jsii.Number(123),
//   		MultiAz: jsii.Boolean(false),
//   		PreferredMaintenanceWindow: jsii.String("preferredMaintenanceWindow"),
//   		ReplicationSubnetGroupId: jsii.String("replicationSubnetGroupId"),
//   		VpcSecurityGroupIds: []*string{
//   			jsii.String("vpcSecurityGroupIds"),
//   		},
//   	},
//   	ReplicationConfigIdentifier: jsii.String("replicationConfigIdentifier"),
//   	ReplicationSettings: replicationSettings,
//   	ReplicationType: jsii.String("replicationType"),
//   	ResourceIdentifier: jsii.String("resourceIdentifier"),
//   	SourceEndpointArn: jsii.String("sourceEndpointArn"),
//   	SupplementalSettings: supplementalSettings,
//   	TableMappings: tableMappings,
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TargetEndpointArn: jsii.String("targetEndpointArn"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationconfig.html
//
type CfnReplicationConfigPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnReplicationConfigMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnReplicationConfigPropsMixin
type jsiiProxy_CfnReplicationConfigPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnReplicationConfigPropsMixin) Props() *CfnReplicationConfigMixinProps {
	var returns *CfnReplicationConfigMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationConfigPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::DMS::ReplicationConfig`.
func NewCfnReplicationConfigPropsMixin(props *CfnReplicationConfigMixinProps, options *mixins.CfnPropertyMixinOptions) CfnReplicationConfigPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnReplicationConfigPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnReplicationConfigPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnReplicationConfigPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::DMS::ReplicationConfig`.
func NewCfnReplicationConfigPropsMixin_Override(c CfnReplicationConfigPropsMixin, props *CfnReplicationConfigMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnReplicationConfigPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnReplicationConfigPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnReplicationConfigPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnReplicationConfigPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnReplicationConfigPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnReplicationConfigPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnReplicationConfigPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnReplicationConfigPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

