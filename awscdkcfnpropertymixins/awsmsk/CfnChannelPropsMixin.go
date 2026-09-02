package awsmsk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsmsk/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::MSK::Channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnChannelPropsMixin := awscdkcfnpropertymixins.Aws_msk.NewCfnChannelPropsMixin(&CfnChannelMixinProps{
//   	ChannelName: jsii.String("channelName"),
//   	ClusterArn: jsii.String("clusterArn"),
//   	EncryptionConfiguration: &EncryptionConfigurationProperty{
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   	},
//   	IcebergDestinationConfiguration: &IcebergDestinationConfigurationProperty{
//   		AppendOnly: jsii.Boolean(false),
//   		Catalog: &CatalogProperty{
//   			CatalogArn: jsii.String("catalogArn"),
//   			WarehouseLocation: jsii.String("warehouseLocation"),
//   		},
//   		CompressionType: jsii.String("compressionType"),
//   		DataFreshnessInSeconds: jsii.Number(123),
//   		DeadLetterQueueS3: &DeadLetterQueueS3Property{
//   			BucketArn: jsii.String("bucketArn"),
//   			ErrorOutputPrefix: jsii.String("errorOutputPrefix"),
//   			ExpectedBucketOwner: jsii.String("expectedBucketOwner"),
//   		},
//   		DestinationTableList: []interface{}{
//   			&DestinationTableProperty{
//   				DestinationDatabaseName: jsii.String("destinationDatabaseName"),
//   				DestinationTableName: jsii.String("destinationTableName"),
//   				PartitionSpec: &PartitionSpecProperty{
//   					PartitionStrategy: jsii.String("partitionStrategy"),
//   					SourceList: []interface{}{
//   						&PartitionSourceProperty{
//   							SourceName: jsii.String("sourceName"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		SchemaEvolution: &SchemaEvolutionProperty{
//   			EnableSchemaEvolution: jsii.Boolean(false),
//   		},
//   		ServiceExecutionRoleArn: jsii.String("serviceExecutionRoleArn"),
//   		TableCreation: &TableCreationProperty{
//   			EnableTableCreation: jsii.Boolean(false),
//   		},
//   	},
//   	LoggingInfo: &ChannelLoggingInfoProperty{
//   		CloudWatchLogs: &CloudWatchLogsLogDestinationProperty{
//   			Enabled: jsii.Boolean(false),
//   			LogGroup: jsii.String("logGroup"),
//   		},
//   		Firehose: &FirehoseLogDestinationProperty{
//   			DeliveryStream: jsii.String("deliveryStream"),
//   			Enabled: jsii.Boolean(false),
//   		},
//   		S3: &S3LogDestinationProperty{
//   			Bucket: jsii.String("bucket"),
//   			Enabled: jsii.Boolean(false),
//   			Prefix: jsii.String("prefix"),
//   		},
//   	},
//   	S3DestinationConfiguration: &S3DestinationConfigurationProperty{
//   		DataFreshnessInSeconds: jsii.Number(123),
//   		DeadLetterQueueS3: &DeadLetterQueueS3Property{
//   			BucketArn: jsii.String("bucketArn"),
//   			ErrorOutputPrefix: jsii.String("errorOutputPrefix"),
//   			ExpectedBucketOwner: jsii.String("expectedBucketOwner"),
//   		},
//   		ServiceExecutionRoleArn: jsii.String("serviceExecutionRoleArn"),
//   		Storage: &S3StorageProperty{
//   			BucketArn: jsii.String("bucketArn"),
//   			CompressionType: jsii.String("compressionType"),
//   			ExpectedBucketOwner: jsii.String("expectedBucketOwner"),
//   			OutputKeyTemplate: jsii.String("outputKeyTemplate"),
//   			OutputPrefix: jsii.String("outputPrefix"),
//   			StorageClass: jsii.String("storageClass"),
//   		},
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	TopicConfigurationList: []interface{}{
//   		&TopicConfigurationProperty{
//   			RecordConverter: &RecordConverterProperty{
//   				ValueConverter: jsii.String("valueConverter"),
//   			},
//   			RecordSchema: &RecordSchemaProperty{
//   				GsrArn: jsii.String("gsrArn"),
//   			},
//   			TopicArn: jsii.String("topicArn"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-channel.html
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


// Create a mixin to apply properties to `AWS::MSK::Channel`.
func NewCfnChannelPropsMixin(props *CfnChannelMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnChannelPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnChannelPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnChannelPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnChannelPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::MSK::Channel`.
func NewCfnChannelPropsMixin_Override(c CfnChannelPropsMixin, props *CfnChannelMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnChannelPropsMixin",
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
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnChannelPropsMixin",
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
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnChannelPropsMixin",
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

