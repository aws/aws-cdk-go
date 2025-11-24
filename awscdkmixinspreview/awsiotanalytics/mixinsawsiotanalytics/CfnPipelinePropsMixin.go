package mixinsawsiotanalytics

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsiotanalytics/mixinsawsiotanalytics/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// The AWS::IoTAnalytics::Pipeline resource consumes messages from one or more channels and allows you to process the messages before storing them in a data store.
//
// You must specify both a `channel` and a `datastore` activity and, optionally, as many as 23 additional activities in the `pipelineActivities` array. For more information, see [How to Use AWS IoT Analytics](https://docs.aws.amazon.com/iotanalytics/latest/userguide/welcome.html#aws-iot-analytics-how) in the *AWS IoT Analytics User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPipelinePropsMixin := awscdkmixinspreview.Mixins.NewCfnPipelinePropsMixin(&CfnPipelineMixinProps{
//   	PipelineActivities: []interface{}{
//   		&ActivityProperty{
//   			AddAttributes: &AddAttributesProperty{
//   				Attributes: map[string]*string{
//   					"attributesKey": jsii.String("attributes"),
//   				},
//   				Name: jsii.String("name"),
//   				Next: jsii.String("next"),
//   			},
//   			Channel: &ChannelProperty{
//   				ChannelName: jsii.String("channelName"),
//   				Name: jsii.String("name"),
//   				Next: jsii.String("next"),
//   			},
//   			Datastore: &DatastoreProperty{
//   				DatastoreName: jsii.String("datastoreName"),
//   				Name: jsii.String("name"),
//   			},
//   			DeviceRegistryEnrich: &DeviceRegistryEnrichProperty{
//   				Attribute: jsii.String("attribute"),
//   				Name: jsii.String("name"),
//   				Next: jsii.String("next"),
//   				RoleArn: jsii.String("roleArn"),
//   				ThingName: jsii.String("thingName"),
//   			},
//   			DeviceShadowEnrich: &DeviceShadowEnrichProperty{
//   				Attribute: jsii.String("attribute"),
//   				Name: jsii.String("name"),
//   				Next: jsii.String("next"),
//   				RoleArn: jsii.String("roleArn"),
//   				ThingName: jsii.String("thingName"),
//   			},
//   			Filter: &FilterProperty{
//   				Filter: jsii.String("filter"),
//   				Name: jsii.String("name"),
//   				Next: jsii.String("next"),
//   			},
//   			Lambda: &LambdaProperty{
//   				BatchSize: jsii.Number(123),
//   				LambdaName: jsii.String("lambdaName"),
//   				Name: jsii.String("name"),
//   				Next: jsii.String("next"),
//   			},
//   			Math: &MathProperty{
//   				Attribute: jsii.String("attribute"),
//   				Math: jsii.String("math"),
//   				Name: jsii.String("name"),
//   				Next: jsii.String("next"),
//   			},
//   			RemoveAttributes: &RemoveAttributesProperty{
//   				Attributes: []*string{
//   					jsii.String("attributes"),
//   				},
//   				Name: jsii.String("name"),
//   				Next: jsii.String("next"),
//   			},
//   			SelectAttributes: &SelectAttributesProperty{
//   				Attributes: []*string{
//   					jsii.String("attributes"),
//   				},
//   				Name: jsii.String("name"),
//   				Next: jsii.String("next"),
//   			},
//   		},
//   	},
//   	PipelineName: jsii.String("pipelineName"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotanalytics-pipeline.html
//
type CfnPipelinePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnPipelineMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnPipelinePropsMixin
type jsiiProxy_CfnPipelinePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnPipelinePropsMixin) Props() *CfnPipelineMixinProps {
	var returns *CfnPipelineMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipelinePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::IoTAnalytics::Pipeline`.
func NewCfnPipelinePropsMixin(props *CfnPipelineMixinProps, options *mixins.CfnPropertyMixinOptions) CfnPipelinePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnPipelinePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnPipelinePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnPipelinePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::IoTAnalytics::Pipeline`.
func NewCfnPipelinePropsMixin_Override(c CfnPipelinePropsMixin, props *CfnPipelineMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnPipelinePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnPipelinePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPipelinePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnPipelinePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPipelinePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_iotanalytics.mixins.CfnPipelinePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPipelinePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnPipelinePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

