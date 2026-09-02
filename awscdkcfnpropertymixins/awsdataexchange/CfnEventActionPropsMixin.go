package awsdataexchange

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsdataexchange/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// An event action is an AWS Data Exchange resource that automatically exports data set revisions to Amazon S3 when a revision is published.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnEventActionPropsMixin := awscdkcfnpropertymixins.Aws_dataexchange.NewCfnEventActionPropsMixin(&CfnEventActionMixinProps{
//   	Action: &ActionProperty{
//   		ExportRevisionToS3: &AutoExportRevisionToS3RequestDetailsProperty{
//   			Encryption: &ExportServerSideEncryptionProperty{
//   				KmsKeyArn: jsii.String("kmsKeyArn"),
//   				Type: jsii.String("type"),
//   			},
//   			RevisionDestination: &AutoExportRevisionDestinationEntryProperty{
//   				Bucket: jsii.String("bucket"),
//   				KeyPattern: jsii.String("keyPattern"),
//   			},
//   		},
//   	},
//   	Event: &EventProperty{
//   		RevisionPublished: &RevisionPublishedProperty{
//   			DataSetId: jsii.String("dataSetId"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dataexchange-eventaction.html
//
type CfnEventActionPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnEventActionMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnEventActionPropsMixin
type jsiiProxy_CfnEventActionPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnEventActionPropsMixin) Props() *CfnEventActionMixinProps {
	var returns *CfnEventActionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventActionPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::DataExchange::EventAction`.
func NewCfnEventActionPropsMixin(props *CfnEventActionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnEventActionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnEventActionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnEventActionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_dataexchange.CfnEventActionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::DataExchange::EventAction`.
func NewCfnEventActionPropsMixin_Override(c CfnEventActionPropsMixin, props *CfnEventActionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_dataexchange.CfnEventActionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnEventActionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEventActionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_dataexchange.CfnEventActionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEventActionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_dataexchange.CfnEventActionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnEventActionPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnEventActionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

