package awskinesisanalyticsv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awskinesisanalyticsv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::KinesisAnalyticsV2::Snapshot, a point-in-time snapshot of a Managed Service for Apache Flink (Kinesis Analytics V2) application's state.
//
// Snapshots are account-scoped and immutable: every input property except Tags is create-only, so changing any of them forces replacement, while changing Tags updates the snapshot in place. Specify exactly one source: ApplicationName to snapshot a running application (CreateSnapshot), or ImportSourceConfiguration to import a snapshot from Amazon S3 (ImportSnapshot). CloudFormation deletes the snapshot via DeleteSnapshot; to retain a snapshot beyond the stack, set DeletionPolicy: Retain on the resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnSnapshotPropsMixin := awscdkcfnpropertymixins.Aws_kinesisanalyticsv2.NewCfnSnapshotPropsMixin(&CfnSnapshotMixinProps{
//   	ApplicationName: jsii.String("applicationName"),
//   	ImportSourceConfiguration: &ImportSourceConfigurationProperty{
//   		ImportRoleArn: jsii.String("importRoleArn"),
//   		SourceEncryptionConfiguration: &SnapshotEncryptionConfigurationDescriptionProperty{
//   			KeyId: jsii.String("keyId"),
//   			KeyType: jsii.String("keyType"),
//   		},
//   		SourceS3PrefixLocation: &S3PrefixLocationProperty{
//   			BucketArn: jsii.String("bucketArn"),
//   			KeyPrefix: jsii.String("keyPrefix"),
//   		},
//   		SourceType: jsii.String("sourceType"),
//   	},
//   	SnapshotDescription: jsii.String("snapshotDescription"),
//   	SnapshotEncryptionConfigurationDescription: &SnapshotEncryptionConfigurationDescriptionProperty{
//   		KeyId: jsii.String("keyId"),
//   		KeyType: jsii.String("keyType"),
//   	},
//   	SnapshotFormat: jsii.String("snapshotFormat"),
//   	SnapshotName: jsii.String("snapshotName"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalyticsv2-snapshot.html
//
type CfnSnapshotPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnSnapshotMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnSnapshotPropsMixin
type jsiiProxy_CfnSnapshotPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnSnapshotPropsMixin) Props() *CfnSnapshotMixinProps {
	var returns *CfnSnapshotMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSnapshotPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::KinesisAnalyticsV2::Snapshot`.
func NewCfnSnapshotPropsMixin(props *CfnSnapshotMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnSnapshotPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnSnapshotPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnSnapshotPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnSnapshotPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::KinesisAnalyticsV2::Snapshot`.
func NewCfnSnapshotPropsMixin_Override(c CfnSnapshotPropsMixin, props *CfnSnapshotMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnSnapshotPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnSnapshotPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSnapshotPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnSnapshotPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSnapshotPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnSnapshotPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSnapshotPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnSnapshotPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

