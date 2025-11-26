package previewawss3expressmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawss3express/previewawss3expressmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Access points simplify managing data access at scale for shared datasets in Amazon S3 .
//
// Access points are unique hostnames you create to enforce distinct permissions and network controls for all requests made through an access point. You can create hundreds of access points per bucket, each with a distinct name and permissions customized for each application. Each access point works in conjunction with the bucket policy that is attached to the underlying bucket. For more information, see [Managing access to shared datasets in directory buckets with access points](https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-points-directory-buckets.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var policy interface{}
//
//   cfnAccessPointPropsMixin := awscdkmixinspreview.Mixins.NewCfnAccessPointPropsMixin(&CfnAccessPointMixinProps{
//   	Bucket: jsii.String("bucket"),
//   	BucketAccountId: jsii.String("bucketAccountId"),
//   	Name: jsii.String("name"),
//   	Policy: policy,
//   	PublicAccessBlockConfiguration: &PublicAccessBlockConfigurationProperty{
//   		BlockPublicAcls: jsii.Boolean(false),
//   		BlockPublicPolicy: jsii.Boolean(false),
//   		IgnorePublicAcls: jsii.Boolean(false),
//   		RestrictPublicBuckets: jsii.Boolean(false),
//   	},
//   	Scope: &ScopeProperty{
//   		Permissions: []*string{
//   			jsii.String("permissions"),
//   		},
//   		Prefixes: []*string{
//   			jsii.String("prefixes"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VpcConfiguration: &VpcConfigurationProperty{
//   		VpcId: jsii.String("vpcId"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3express-accesspoint.html
//
type CfnAccessPointPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnAccessPointMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAccessPointPropsMixin
type jsiiProxy_CfnAccessPointPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnAccessPointPropsMixin) Props() *CfnAccessPointMixinProps {
	var returns *CfnAccessPointMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccessPointPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::S3Express::AccessPoint`.
func NewCfnAccessPointPropsMixin(props *CfnAccessPointMixinProps, options *mixins.CfnPropertyMixinOptions) CfnAccessPointPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAccessPointPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAccessPointPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3express.mixins.CfnAccessPointPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::S3Express::AccessPoint`.
func NewCfnAccessPointPropsMixin_Override(c CfnAccessPointPropsMixin, props *CfnAccessPointMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3express.mixins.CfnAccessPointPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnAccessPointPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAccessPointPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_s3express.mixins.CfnAccessPointPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAccessPointPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_s3express.mixins.CfnAccessPointPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAccessPointPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnAccessPointPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

