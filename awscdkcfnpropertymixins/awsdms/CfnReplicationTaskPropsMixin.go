package awsdms

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsdms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::DMS::ReplicationTask` resource creates an AWS DMS replication task.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnReplicationTaskPropsMixin := awscdkcfnpropertymixins.Aws_dms.NewCfnReplicationTaskPropsMixin(&CfnReplicationTaskMixinProps{
//   	CdcStartPosition: jsii.String("cdcStartPosition"),
//   	CdcStartTime: jsii.Number(123),
//   	CdcStopPosition: jsii.String("cdcStopPosition"),
//   	MigrationType: jsii.String("migrationType"),
//   	ReplicationInstanceArn: jsii.String("replicationInstanceArn"),
//   	ReplicationTaskIdentifier: jsii.String("replicationTaskIdentifier"),
//   	ReplicationTaskSettings: jsii.String("replicationTaskSettings"),
//   	ResourceIdentifier: jsii.String("resourceIdentifier"),
//   	SourceEndpointArn: jsii.String("sourceEndpointArn"),
//   	TableMappings: jsii.String("tableMappings"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TargetEndpointArn: jsii.String("targetEndpointArn"),
//   	TaskData: jsii.String("taskData"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationtask.html
//
type CfnReplicationTaskPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnReplicationTaskMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnReplicationTaskPropsMixin
type jsiiProxy_CfnReplicationTaskPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnReplicationTaskPropsMixin) Props() *CfnReplicationTaskMixinProps {
	var returns *CfnReplicationTaskMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationTaskPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::DMS::ReplicationTask`.
func NewCfnReplicationTaskPropsMixin(props *CfnReplicationTaskMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnReplicationTaskPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnReplicationTaskPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnReplicationTaskPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_dms.CfnReplicationTaskPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::DMS::ReplicationTask`.
func NewCfnReplicationTaskPropsMixin_Override(c CfnReplicationTaskPropsMixin, props *CfnReplicationTaskMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_dms.CfnReplicationTaskPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnReplicationTaskPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnReplicationTaskPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_dms.CfnReplicationTaskPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnReplicationTaskPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_dms.CfnReplicationTaskPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnReplicationTaskPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnReplicationTaskPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

