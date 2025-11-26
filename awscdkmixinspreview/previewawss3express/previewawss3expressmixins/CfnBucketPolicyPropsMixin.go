package previewawss3expressmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawss3express/previewawss3expressmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::S3Express::BucketPolicy` resource defines an Amazon S3 bucket policy to an Amazon S3 directory bucket.
//
// - **Permissions** - If you are using an identity other than the root user of the AWS account that owns the bucket, the calling identity must both have the required permissions on the specified bucket and belong to the bucket owner's account in order to use this operation. For more information about directory bucket policies and permissions, see [AWS Identity and Access Management (IAM) for S3 Express One Zone](https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-security-iam.html) in the *Amazon S3 User Guide* .
//
// > To ensure that bucket owners don't inadvertently lock themselves out of their own buckets, the root principal in a bucket owner's AWS account can perform the `GetBucketPolicy` , `PutBucketPolicy` , and `DeleteBucketPolicy` API actions, even if their bucket policy explicitly denies the root principal's access. Bucket owner root principals can only be blocked from performing these API actions by VPC endpoint policies and AWS Organizations policies.
//
// The required permissions for CloudFormation to use are based on the operations that are performed on the stack.
//
// - Create
//
// - s3express:GetBucketPolicy
// - s3express:PutBucketPolicy
// - Read
//
// - s3express:GetBucketPolicy
// - Update
//
// - s3express:GetBucketPolicy
// - s3express:PutBucketPolicy
// - Delete
//
// - s3express:GetBucketPolicy
// - s3express:DeleteBucketPolicy
// - List
//
// - s3express:GetBucketPolicy
// - s3express:ListAllMyDirectoryBuckets
//
// For more information about example bucket policies, see [Example bucket policies for S3 Express One Zone](https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-security-iam-example-bucket-policies.html) in the *Amazon S3 User Guide* .
//
// The following operations are related to `AWS::S3Express::BucketPolicy` :
//
// - [PutBucketPolicy](https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketPolicy.html)
// - [GetBucketPolicy](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketPolicy.html)
// - [DeleteBucketPolicy](https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketPolicy.html)
// - [ListDirectoryBuckets](https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListDirectoryBuckets.html)
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var policyDocument interface{}
//
//   cfnBucketPolicyPropsMixin := awscdkmixinspreview.Mixins.NewCfnBucketPolicyPropsMixin(&CfnBucketPolicyMixinProps{
//   	Bucket: jsii.String("bucket"),
//   	PolicyDocument: policyDocument,
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3express-bucketpolicy.html
//
type CfnBucketPolicyPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnBucketPolicyMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnBucketPolicyPropsMixin
type jsiiProxy_CfnBucketPolicyPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnBucketPolicyPropsMixin) Props() *CfnBucketPolicyMixinProps {
	var returns *CfnBucketPolicyMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucketPolicyPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::S3Express::BucketPolicy`.
func NewCfnBucketPolicyPropsMixin(props *CfnBucketPolicyMixinProps, options *mixins.CfnPropertyMixinOptions) CfnBucketPolicyPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnBucketPolicyPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnBucketPolicyPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3express.mixins.CfnBucketPolicyPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::S3Express::BucketPolicy`.
func NewCfnBucketPolicyPropsMixin_Override(c CfnBucketPolicyPropsMixin, props *CfnBucketPolicyMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3express.mixins.CfnBucketPolicyPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnBucketPolicyPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnBucketPolicyPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_s3express.mixins.CfnBucketPolicyPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnBucketPolicyPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_s3express.mixins.CfnBucketPolicyPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnBucketPolicyPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnBucketPolicyPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

