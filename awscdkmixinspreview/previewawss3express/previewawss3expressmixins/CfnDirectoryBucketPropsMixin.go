package previewawss3expressmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawss3express/previewawss3expressmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::S3Express::DirectoryBucket` resource defines an Amazon S3 directory bucket in the same AWS Region where you create the AWS CloudFormation stack.
//
// To control how AWS CloudFormation handles the bucket when the stack is deleted, you can set a deletion policy for your bucket. You can choose to *retain* the bucket or to *delete* the bucket. For more information, see [DeletionPolicy attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html) .
//
// > You can only delete empty buckets. Deletion fails for buckets that have contents.
//
// - **Permissions** - The required permissions for CloudFormation to use are based on the operations that are performed on the stack.
//
// - Create
//
// - s3express:CreateBucket
// - s3express:ListAllMyDirectoryBuckets
// - Read
//
// - s3express:ListAllMyDirectoryBuckets
// - ec2:DescribeAvailabilityZones
// - Delete
//
// - s3express:DeleteBucket
// - s3express:ListAllMyDirectoryBuckets
// - List
//
// - s3express:ListAllMyDirectoryBuckets
// - PutBucketEncryption
//
// - s3express:PutEncryptionConfiguration
// - To set a directory bucket default encryption with SSE-KMS, you must also have the kms:GenerateDataKey and kms:Decrypt permissions in IAM identity-based policies and AWS KMS key policies for the target AWS KMS key.
// - GetBucketEncryption
//
// - s3express:GetBucketEncryption
// - DeleteBucketEncryption
//
// - s3express:PutEncryptionConfiguration
//
// The following operations are related to `AWS::S3Express::DirectoryBucket` :
//
// - [CreateBucket](https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucket.html)
// - [ListDirectoryBuckets](https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListDirectoryBuckets.html)
// - [DeleteBucket](https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucket.html)
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDirectoryBucketPropsMixin := awscdkmixinspreview.Mixins.NewCfnDirectoryBucketPropsMixin(&CfnDirectoryBucketMixinProps{
//   	BucketEncryption: &BucketEncryptionProperty{
//   		ServerSideEncryptionConfiguration: []interface{}{
//   			&ServerSideEncryptionRuleProperty{
//   				BucketKeyEnabled: jsii.Boolean(false),
//   				ServerSideEncryptionByDefault: &ServerSideEncryptionByDefaultProperty{
//   					KmsMasterKeyId: jsii.String("kmsMasterKeyId"),
//   					SseAlgorithm: jsii.String("sseAlgorithm"),
//   				},
//   			},
//   		},
//   	},
//   	BucketName: jsii.String("bucketName"),
//   	DataRedundancy: jsii.String("dataRedundancy"),
//   	LifecycleConfiguration: &LifecycleConfigurationProperty{
//   		Rules: []interface{}{
//   			&RuleProperty{
//   				AbortIncompleteMultipartUpload: &AbortIncompleteMultipartUploadProperty{
//   					DaysAfterInitiation: jsii.Number(123),
//   				},
//   				ExpirationInDays: jsii.Number(123),
//   				Id: jsii.String("id"),
//   				ObjectSizeGreaterThan: jsii.String("objectSizeGreaterThan"),
//   				ObjectSizeLessThan: jsii.String("objectSizeLessThan"),
//   				Prefix: jsii.String("prefix"),
//   				Status: jsii.String("status"),
//   			},
//   		},
//   	},
//   	LocationName: jsii.String("locationName"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3express-directorybucket.html
//
type CfnDirectoryBucketPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnDirectoryBucketMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDirectoryBucketPropsMixin
type jsiiProxy_CfnDirectoryBucketPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnDirectoryBucketPropsMixin) Props() *CfnDirectoryBucketMixinProps {
	var returns *CfnDirectoryBucketMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDirectoryBucketPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::S3Express::DirectoryBucket`.
func NewCfnDirectoryBucketPropsMixin(props *CfnDirectoryBucketMixinProps, options *mixins.CfnPropertyMixinOptions) CfnDirectoryBucketPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDirectoryBucketPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDirectoryBucketPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3express.mixins.CfnDirectoryBucketPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::S3Express::DirectoryBucket`.
func NewCfnDirectoryBucketPropsMixin_Override(c CfnDirectoryBucketPropsMixin, props *CfnDirectoryBucketMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3express.mixins.CfnDirectoryBucketPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnDirectoryBucketPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDirectoryBucketPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_s3express.mixins.CfnDirectoryBucketPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDirectoryBucketPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_s3express.mixins.CfnDirectoryBucketPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDirectoryBucketPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnDirectoryBucketPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

