package previewawss3mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawss3/previewawss3mixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::S3::MultiRegionAccessPoint` resource creates an Amazon S3 Multi-Region Access Point.
//
// To learn more about Multi-Region Access Points, see [Multi-Region Access Points in Amazon S3](https://docs.aws.amazon.com/AmazonS3/latest/userguide/MultiRegionAccessPoints.html) in the in the *Amazon S3 User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMultiRegionAccessPointPropsMixin := awscdkmixinspreview.Mixins.NewCfnMultiRegionAccessPointPropsMixin(&CfnMultiRegionAccessPointMixinProps{
//   	Name: jsii.String("name"),
//   	PublicAccessBlockConfiguration: &PublicAccessBlockConfigurationProperty{
//   		BlockPublicAcls: jsii.Boolean(false),
//   		BlockPublicPolicy: jsii.Boolean(false),
//   		IgnorePublicAcls: jsii.Boolean(false),
//   		RestrictPublicBuckets: jsii.Boolean(false),
//   	},
//   	Regions: []interface{}{
//   		&RegionProperty{
//   			Bucket: jsii.String("bucket"),
//   			BucketAccountId: jsii.String("bucketAccountId"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-multiregionaccesspoint.html
//
type CfnMultiRegionAccessPointPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnMultiRegionAccessPointMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnMultiRegionAccessPointPropsMixin
type jsiiProxy_CfnMultiRegionAccessPointPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnMultiRegionAccessPointPropsMixin) Props() *CfnMultiRegionAccessPointMixinProps {
	var returns *CfnMultiRegionAccessPointMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMultiRegionAccessPointPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::S3::MultiRegionAccessPoint`.
func NewCfnMultiRegionAccessPointPropsMixin(props *CfnMultiRegionAccessPointMixinProps, options *mixins.CfnPropertyMixinOptions) CfnMultiRegionAccessPointPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnMultiRegionAccessPointPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnMultiRegionAccessPointPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnMultiRegionAccessPointPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::S3::MultiRegionAccessPoint`.
func NewCfnMultiRegionAccessPointPropsMixin_Override(c CfnMultiRegionAccessPointPropsMixin, props *CfnMultiRegionAccessPointMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnMultiRegionAccessPointPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnMultiRegionAccessPointPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnMultiRegionAccessPointPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnMultiRegionAccessPointPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMultiRegionAccessPointPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnMultiRegionAccessPointPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMultiRegionAccessPointPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnMultiRegionAccessPointPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

