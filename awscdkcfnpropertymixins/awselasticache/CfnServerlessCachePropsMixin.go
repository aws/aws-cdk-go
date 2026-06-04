package awselasticache

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awselasticache/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The resource representing a serverless cache.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnServerlessCachePropsMixin := awscdkcfnpropertymixins.Aws_elasticache.NewCfnServerlessCachePropsMixin(&CfnServerlessCacheMixinProps{
//   	CacheUsageLimits: &CacheUsageLimitsProperty{
//   		DataStorage: &DataStorageProperty{
//   			Maximum: jsii.Number(123),
//   			Minimum: jsii.Number(123),
//   			Unit: jsii.String("unit"),
//   		},
//   		EcpuPerSecond: &ECPUPerSecondProperty{
//   			Maximum: jsii.Number(123),
//   			Minimum: jsii.Number(123),
//   		},
//   	},
//   	DailySnapshotTime: jsii.String("dailySnapshotTime"),
//   	Description: jsii.String("description"),
//   	Endpoint: &EndpointProperty{
//   		Address: jsii.String("address"),
//   		Port: jsii.String("port"),
//   	},
//   	Engine: jsii.String("engine"),
//   	FinalSnapshotName: jsii.String("finalSnapshotName"),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	MajorEngineVersion: jsii.String("majorEngineVersion"),
//   	NetworkType: jsii.String("networkType"),
//   	ReaderEndpoint: &EndpointProperty{
//   		Address: jsii.String("address"),
//   		Port: jsii.String("port"),
//   	},
//   	SecurityGroupIds: []interface{}{
//   		jsii.String("securityGroupIds"),
//   	},
//   	ServerlessCacheName: jsii.String("serverlessCacheName"),
//   	SnapshotArnsToRestore: []*string{
//   		jsii.String("snapshotArnsToRestore"),
//   	},
//   	SnapshotRetentionLimit: jsii.Number(123),
//   	SubnetIds: []interface{}{
//   		jsii.String("subnetIds"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	UserGroupId: jsii.String("userGroupId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-serverlesscache.html
//
type CfnServerlessCachePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnServerlessCacheMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnServerlessCachePropsMixin
type jsiiProxy_CfnServerlessCachePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnServerlessCachePropsMixin) Props() *CfnServerlessCacheMixinProps {
	var returns *CfnServerlessCacheMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServerlessCachePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ElastiCache::ServerlessCache`.
func NewCfnServerlessCachePropsMixin(props *CfnServerlessCacheMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnServerlessCachePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnServerlessCachePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnServerlessCachePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_elasticache.CfnServerlessCachePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ElastiCache::ServerlessCache`.
func NewCfnServerlessCachePropsMixin_Override(c CfnServerlessCachePropsMixin, props *CfnServerlessCacheMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_elasticache.CfnServerlessCachePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnServerlessCachePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnServerlessCachePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_elasticache.CfnServerlessCachePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnServerlessCachePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_elasticache.CfnServerlessCachePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnServerlessCachePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnServerlessCachePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

