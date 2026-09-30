package awskinesis

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awskinesis/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::Kinesis::Channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnChannelPropsMixin := awscdkcfnpropertymixins.Aws_kinesis.NewCfnChannelPropsMixin(&CfnChannelMixinProps{
//   	ChannelName: jsii.String("channelName"),
//   	EncryptionConfiguration: &EncryptionConfigurationProperty{
//   		EncryptionType: jsii.String("encryptionType"),
//   		KeyId: jsii.String("keyId"),
//   	},
//   	LoggingConfiguration: &LoggingConfigurationProperty{
//   		CloudWatchLogs: &CloudWatchLogsConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   			LogGroupName: jsii.String("logGroupName"),
//   			LogStreamName: jsii.String("logStreamName"),
//   		},
//   	},
//   	S3DestinationConfiguration: &S3DestinationConfigurationProperty{
//   		DataFreshnessInSeconds: jsii.Number(123),
//   		DeadLetterQueueS3Configuration: &DeadLetterQueueS3ConfigurationProperty{
//   			BucketArn: jsii.String("bucketArn"),
//   			ErrorOutputPrefix: jsii.String("errorOutputPrefix"),
//   			ExpectedBucketOwner: jsii.String("expectedBucketOwner"),
//   		},
//   		StorageConfiguration: &S3StorageConfigurationProperty{
//   			BucketArn: jsii.String("bucketArn"),
//   			CompressionType: jsii.String("compressionType"),
//   			ExpectedBucketOwner: jsii.String("expectedBucketOwner"),
//   			OutputKeyTemplate: jsii.String("outputKeyTemplate"),
//   			StorageClass: jsii.String("storageClass"),
//   		},
//   	},
//   	S3TablesDestinationConfiguration: &S3TablesDestinationConfigurationProperty{
//   		DataFreshnessInSeconds: jsii.Number(123),
//   		DeadLetterQueueS3Configuration: &DeadLetterQueueS3ConfigurationProperty{
//   			BucketArn: jsii.String("bucketArn"),
//   			ErrorOutputPrefix: jsii.String("errorOutputPrefix"),
//   			ExpectedBucketOwner: jsii.String("expectedBucketOwner"),
//   		},
//   		S3TablesConfigurationList: []interface{}{
//   			&S3TableConfigurationProperty{
//   				CompressionType: jsii.String("compressionType"),
//   				Namespace: jsii.String("namespace"),
//   				PartitionSpec: &PartitionSpecProperty{
//   					PartitionFields: []interface{}{
//   						&PartitionFieldProperty{
//   							SourceName: jsii.String("sourceName"),
//   							Transform: jsii.String("transform"),
//   						},
//   					},
//   				},
//   				TableBucketArn: jsii.String("tableBucketArn"),
//   				TableName: jsii.String("tableName"),
//   			},
//   		},
//   	},
//   	ServiceExecutionRoleArn: jsii.String("serviceExecutionRoleArn"),
//   	StreamConfigurationList: []interface{}{
//   		&StreamConfigurationProperty{
//   			RecordConfiguration: &RecordConfigurationProperty{
//   				GsrSchemaArn: jsii.String("gsrSchemaArn"),
//   				RecordFormatType: jsii.String("recordFormatType"),
//   			},
//   			StreamArn: jsii.String("streamArn"),
//   		},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesis-channel.html
//
type CfnChannelPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnChannelMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnChannelPropsMixin
type jsiiProxy_CfnChannelPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnChannelPropsMixin) Props() *CfnChannelMixinProps {
	var returns *CfnChannelMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannelPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Kinesis::Channel`.
func NewCfnChannelPropsMixin(props *CfnChannelMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnChannelPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnChannelPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnChannelPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_kinesis.CfnChannelPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Kinesis::Channel`.
func NewCfnChannelPropsMixin_Override(c CfnChannelPropsMixin, props *CfnChannelMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_kinesis.CfnChannelPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnChannelPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnChannelPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_kinesis.CfnChannelPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnChannelPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_kinesis.CfnChannelPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnChannelPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnChannelPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

