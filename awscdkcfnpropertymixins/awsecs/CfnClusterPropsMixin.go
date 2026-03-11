package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsecs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::ECS::Cluster` resource creates an Amazon Elastic Container Service (Amazon ECS) cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnClusterPropsMixin := awscdkcfnpropertymixins.Aws_ecs.NewCfnClusterPropsMixin(&CfnClusterMixinProps{
//   	CapacityProviders: []interface{}{
//   		jsii.String("capacityProviders"),
//   	},
//   	ClusterName: jsii.String("clusterName"),
//   	ClusterSettings: []interface{}{
//   		&ClusterSettingsProperty{
//   			Name: jsii.String("name"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Configuration: &ClusterConfigurationProperty{
//   		ExecuteCommandConfiguration: &ExecuteCommandConfigurationProperty{
//   			KmsKeyId: jsii.String("kmsKeyId"),
//   			LogConfiguration: &ExecuteCommandLogConfigurationProperty{
//   				CloudWatchEncryptionEnabled: jsii.Boolean(false),
//   				CloudWatchLogGroupName: jsii.String("cloudWatchLogGroupName"),
//   				S3BucketName: jsii.String("s3BucketName"),
//   				S3EncryptionEnabled: jsii.Boolean(false),
//   				S3KeyPrefix: jsii.String("s3KeyPrefix"),
//   			},
//   			Logging: jsii.String("logging"),
//   		},
//   		ManagedStorageConfiguration: &ManagedStorageConfigurationProperty{
//   			FargateEphemeralStorageKmsKeyId: jsii.String("fargateEphemeralStorageKmsKeyId"),
//   			KmsKeyId: jsii.String("kmsKeyId"),
//   		},
//   	},
//   	DefaultCapacityProviderStrategy: []interface{}{
//   		&CapacityProviderStrategyItemProperty{
//   			Base: jsii.Number(123),
//   			CapacityProvider: jsii.String("capacityProvider"),
//   			Weight: jsii.Number(123),
//   		},
//   	},
//   	ServiceConnectDefaults: &ServiceConnectDefaultsProperty{
//   		Namespace: jsii.String("namespace"),
//   	},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-cluster.html
//
type CfnClusterPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnClusterMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnClusterPropsMixin
type jsiiProxy_CfnClusterPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnClusterPropsMixin) Props() *CfnClusterMixinProps {
	var returns *CfnClusterMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ECS::Cluster`.
func NewCfnClusterPropsMixin(props *CfnClusterMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnClusterPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnClusterPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnClusterPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnClusterPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ECS::Cluster`.
func NewCfnClusterPropsMixin_Override(c CfnClusterPropsMixin, props *CfnClusterMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnClusterPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnClusterPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnClusterPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnClusterPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnClusterPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnClusterPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnClusterPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnClusterPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

