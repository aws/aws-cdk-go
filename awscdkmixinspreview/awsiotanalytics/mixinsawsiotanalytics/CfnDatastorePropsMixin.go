package mixinsawsiotanalytics

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsiotanalytics/mixinsawsiotanalytics/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// AWS::IoTAnalytics::Datastore resource is a repository for messages.
//
// For more information, see [How to Use AWS IoT Analytics](https://docs.aws.amazon.com/iotanalytics/latest/userguide/welcome.html#aws-iot-analytics-how) in the *AWS IoT Analytics User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var jsonConfiguration interface{}
//   var serviceManagedS3 interface{}
//
//   cfnDatastorePropsMixin := awscdkmixinspreview.Mixins.NewCfnDatastorePropsMixin(&CfnDatastoreMixinProps{
//   	DatastoreName: jsii.String("datastoreName"),
//   	DatastorePartitions: &DatastorePartitionsProperty{
//   		Partitions: []interface{}{
//   			&DatastorePartitionProperty{
//   				Partition: &PartitionProperty{
//   					AttributeName: jsii.String("attributeName"),
//   				},
//   				TimestampPartition: &TimestampPartitionProperty{
//   					AttributeName: jsii.String("attributeName"),
//   					TimestampFormat: jsii.String("timestampFormat"),
//   				},
//   			},
//   		},
//   	},
//   	DatastoreStorage: &DatastoreStorageProperty{
//   		CustomerManagedS3: &CustomerManagedS3Property{
//   			Bucket: jsii.String("bucket"),
//   			KeyPrefix: jsii.String("keyPrefix"),
//   			RoleArn: jsii.String("roleArn"),
//   		},
//   		IotSiteWiseMultiLayerStorage: &IotSiteWiseMultiLayerStorageProperty{
//   			CustomerManagedS3Storage: &CustomerManagedS3StorageProperty{
//   				Bucket: jsii.String("bucket"),
//   				KeyPrefix: jsii.String("keyPrefix"),
//   			},
//   		},
//   		ServiceManagedS3: serviceManagedS3,
//   	},
//   	FileFormatConfiguration: &FileFormatConfigurationProperty{
//   		JsonConfiguration: jsonConfiguration,
//   		ParquetConfiguration: &ParquetConfigurationProperty{
//   			SchemaDefinition: &SchemaDefinitionProperty{
//   				Columns: []interface{}{
//   					&ColumnProperty{
//   						Name: jsii.String("name"),
//   						Type: jsii.String("type"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	RetentionPeriod: &RetentionPeriodProperty{
//   		NumberOfDays: jsii.Number(123),
//   		Unlimited: jsii.Boolean(false),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotanalytics-datastore.html
//
type CfnDatastorePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnDatastoreMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDatastorePropsMixin
type jsiiProxy_CfnDatastorePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnDatastorePropsMixin) Props() *CfnDatastoreMixinProps {
	var returns *CfnDatastoreMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatastorePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::IoTAnalytics::Datastore`.
func NewCfnDatastorePropsMixin(props *CfnDatastoreMixinProps, options *mixins.CfnPropertyMixinOptions) CfnDatastorePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDatastorePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDatastorePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnDatastorePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::IoTAnalytics::Datastore`.
func NewCfnDatastorePropsMixin_Override(c CfnDatastorePropsMixin, props *CfnDatastoreMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnDatastorePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnDatastorePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDatastorePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnDatastorePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDatastorePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnDatastorePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDatastorePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnDatastorePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

