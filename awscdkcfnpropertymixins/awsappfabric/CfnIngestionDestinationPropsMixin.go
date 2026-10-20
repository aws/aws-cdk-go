package awsappfabric

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsappfabric/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an ingestion destination, which specifies how an application's ingested data is processed by Amazon Web Services AppFabric and where it's delivered.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnIngestionDestinationPropsMixin := awscdkcfnpropertymixins.Aws_appfabric.NewCfnIngestionDestinationPropsMixin(&CfnIngestionDestinationMixinProps{
//   	AppBundleArn: jsii.String("appBundleArn"),
//   	DestinationConfiguration: &DestinationConfigurationProperty{
//   		AuditLog: &AuditLogDestinationConfigurationProperty{
//   			Destination: &DestinationProperty{
//   				FirehoseStream: &FirehoseStreamProperty{
//   					StreamName: jsii.String("streamName"),
//   				},
//   				S3Bucket: &S3BucketProperty{
//   					BucketName: jsii.String("bucketName"),
//   					Prefix: jsii.String("prefix"),
//   				},
//   			},
//   		},
//   	},
//   	IngestionArn: jsii.String("ingestionArn"),
//   	ProcessingConfiguration: &ProcessingConfigurationProperty{
//   		AuditLog: &AuditLogProcessingConfigurationProperty{
//   			Format: jsii.String("format"),
//   			Schema: jsii.String("schema"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appfabric-ingestiondestination.html
//
type CfnIngestionDestinationPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnIngestionDestinationMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnIngestionDestinationPropsMixin
type jsiiProxy_CfnIngestionDestinationPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnIngestionDestinationPropsMixin) Props() *CfnIngestionDestinationMixinProps {
	var returns *CfnIngestionDestinationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIngestionDestinationPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::AppFabric::IngestionDestination`.
func NewCfnIngestionDestinationPropsMixin(props *CfnIngestionDestinationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnIngestionDestinationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnIngestionDestinationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnIngestionDestinationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_appfabric.CfnIngestionDestinationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::AppFabric::IngestionDestination`.
func NewCfnIngestionDestinationPropsMixin_Override(c CfnIngestionDestinationPropsMixin, props *CfnIngestionDestinationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_appfabric.CfnIngestionDestinationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnIngestionDestinationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnIngestionDestinationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_appfabric.CfnIngestionDestinationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnIngestionDestinationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_appfabric.CfnIngestionDestinationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnIngestionDestinationPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnIngestionDestinationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

