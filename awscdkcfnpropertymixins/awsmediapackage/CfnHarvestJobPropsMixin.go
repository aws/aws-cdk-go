package awsmediapackage

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsmediapackage/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource schema for AWS::MediaPackage::HarvestJob.
//
// A HarvestJob extracts a video on demand (VOD) clip from a live content stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnHarvestJobPropsMixin := awscdkcfnpropertymixins.Aws_mediapackage.NewCfnHarvestJobPropsMixin(&CfnHarvestJobMixinProps{
//   	EndTime: jsii.String("endTime"),
//   	Id: jsii.String("id"),
//   	OriginEndpointId: jsii.String("originEndpointId"),
//   	S3Destination: &S3DestinationProperty{
//   		BucketName: jsii.String("bucketName"),
//   		ManifestKey: jsii.String("manifestKey"),
//   		RoleArn: jsii.String("roleArn"),
//   	},
//   	StartTime: jsii.String("startTime"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-harvestjob.html
//
type CfnHarvestJobPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnHarvestJobMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnHarvestJobPropsMixin
type jsiiProxy_CfnHarvestJobPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnHarvestJobPropsMixin) Props() *CfnHarvestJobMixinProps {
	var returns *CfnHarvestJobMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHarvestJobPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::MediaPackage::HarvestJob`.
func NewCfnHarvestJobPropsMixin(props *CfnHarvestJobMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnHarvestJobPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnHarvestJobPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnHarvestJobPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_mediapackage.CfnHarvestJobPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::MediaPackage::HarvestJob`.
func NewCfnHarvestJobPropsMixin_Override(c CfnHarvestJobPropsMixin, props *CfnHarvestJobMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_mediapackage.CfnHarvestJobPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnHarvestJobPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnHarvestJobPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_mediapackage.CfnHarvestJobPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnHarvestJobPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_mediapackage.CfnHarvestJobPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnHarvestJobPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnHarvestJobPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

