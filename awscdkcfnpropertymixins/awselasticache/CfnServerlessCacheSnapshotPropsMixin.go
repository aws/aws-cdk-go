package awselasticache

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awselasticache/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::ElastiCache::ServerlessCacheSnapshot.
//
// A serverless cache snapshot is a point-in-time backup of an ElastiCache serverless cache. Available for Valkey, Redis OSS and Serverless Memcached only.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnServerlessCacheSnapshotPropsMixin := awscdkcfnpropertymixins.Aws_elasticache.NewCfnServerlessCacheSnapshotPropsMixin(&CfnServerlessCacheSnapshotMixinProps{
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	ServerlessCacheName: jsii.String("serverlessCacheName"),
//   	ServerlessCacheSnapshotName: jsii.String("serverlessCacheSnapshotName"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-serverlesscachesnapshot.html
//
type CfnServerlessCacheSnapshotPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnServerlessCacheSnapshotMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnServerlessCacheSnapshotPropsMixin
type jsiiProxy_CfnServerlessCacheSnapshotPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnServerlessCacheSnapshotPropsMixin) Props() *CfnServerlessCacheSnapshotMixinProps {
	var returns *CfnServerlessCacheSnapshotMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServerlessCacheSnapshotPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ElastiCache::ServerlessCacheSnapshot`.
func NewCfnServerlessCacheSnapshotPropsMixin(props *CfnServerlessCacheSnapshotMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnServerlessCacheSnapshotPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnServerlessCacheSnapshotPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnServerlessCacheSnapshotPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_elasticache.CfnServerlessCacheSnapshotPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ElastiCache::ServerlessCacheSnapshot`.
func NewCfnServerlessCacheSnapshotPropsMixin_Override(c CfnServerlessCacheSnapshotPropsMixin, props *CfnServerlessCacheSnapshotMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_elasticache.CfnServerlessCacheSnapshotPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnServerlessCacheSnapshotPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnServerlessCacheSnapshotPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_elasticache.CfnServerlessCacheSnapshotPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnServerlessCacheSnapshotPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_elasticache.CfnServerlessCacheSnapshotPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnServerlessCacheSnapshotPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnServerlessCacheSnapshotPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

