package awss3outposts

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awss3outposts/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The AWS::S3Outposts::Bucket resource specifies a new Amazon S3 on Outposts bucket.
//
// To create an S3 on Outposts bucket, you must have S3 on Outposts capacity provisioned on your Outpost. For more information, see [Using Amazon S3 on Outposts](https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html) .
//
// S3 on Outposts buckets support the following:
//
// - Tags
// - Lifecycle configuration rules for deleting expired objects
//
// For a complete list of restrictions and Amazon S3 feature limitations on S3 on Outposts, see [Amazon S3 on Outposts Restrictions and Limitations](https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3OnOutpostsRestrictionsLimitations.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var filter interface{}
//   var mergeStrategy IMergeStrategy
//
//   cfnBucketPropsMixin := awscdkcfnpropertymixins.Aws_s3outposts.NewCfnBucketPropsMixin(&CfnBucketMixinProps{
//   	BucketName: jsii.String("bucketName"),
//   	LifecycleConfiguration: &LifecycleConfigurationProperty{
//   		Rules: []interface{}{
//   			&RuleProperty{
//   				AbortIncompleteMultipartUpload: &AbortIncompleteMultipartUploadProperty{
//   					DaysAfterInitiation: jsii.Number(123),
//   				},
//   				ExpirationDate: jsii.String("expirationDate"),
//   				ExpirationInDays: jsii.Number(123),
//   				Filter: filter,
//   				Id: jsii.String("id"),
//   				Status: jsii.String("status"),
//   			},
//   		},
//   	},
//   	OutpostId: jsii.String("outpostId"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3outposts-bucket.html
//
type CfnBucketPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnBucketMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnBucketPropsMixin
type jsiiProxy_CfnBucketPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnBucketPropsMixin) Props() *CfnBucketMixinProps {
	var returns *CfnBucketMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucketPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::S3Outposts::Bucket`.
func NewCfnBucketPropsMixin(props *CfnBucketMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnBucketPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnBucketPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnBucketPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_s3outposts.CfnBucketPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::S3Outposts::Bucket`.
func NewCfnBucketPropsMixin_Override(c CfnBucketPropsMixin, props *CfnBucketMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_s3outposts.CfnBucketPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnBucketPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnBucketPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_s3outposts.CfnBucketPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnBucketPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_s3outposts.CfnBucketPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnBucketPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnBucketPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

