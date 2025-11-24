package mixinsawss3

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awss3/mixinsawss3/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::S3::StorageLensGroup` resource creates an S3 Storage Lens group.
//
// A Storage Lens group is a custom grouping of objects that include filters for prefixes, suffixes, object tags, object size, or object age. You can create an S3 Storage Lens group that includes a single filter or multiple filter conditions. To specify multiple filter conditions, you use `AND` or `OR` logical operators. For more information about S3 Storage Lens groups, see [Working with S3 Storage Lens groups](https://docs.aws.amazon.com/AmazonS3/latest/userguide/storage-lens-groups-overview.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnStorageLensGroupPropsMixin := awscdkmixinspreview.Mixins.NewCfnStorageLensGroupPropsMixin(&CfnStorageLensGroupMixinProps{
//   	Filter: &FilterProperty{
//   		And: &AndProperty{
//   			MatchAnyPrefix: []*string{
//   				jsii.String("matchAnyPrefix"),
//   			},
//   			MatchAnySuffix: []*string{
//   				jsii.String("matchAnySuffix"),
//   			},
//   			MatchAnyTag: []interface{}{
//   				&CfnTag{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			MatchObjectAge: &MatchObjectAgeProperty{
//   				DaysGreaterThan: jsii.Number(123),
//   				DaysLessThan: jsii.Number(123),
//   			},
//   			MatchObjectSize: &MatchObjectSizeProperty{
//   				BytesGreaterThan: jsii.Number(123),
//   				BytesLessThan: jsii.Number(123),
//   			},
//   		},
//   		MatchAnyPrefix: []*string{
//   			jsii.String("matchAnyPrefix"),
//   		},
//   		MatchAnySuffix: []*string{
//   			jsii.String("matchAnySuffix"),
//   		},
//   		MatchAnyTag: []interface{}{
//   			&CfnTag{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		MatchObjectAge: &MatchObjectAgeProperty{
//   			DaysGreaterThan: jsii.Number(123),
//   			DaysLessThan: jsii.Number(123),
//   		},
//   		MatchObjectSize: &MatchObjectSizeProperty{
//   			BytesGreaterThan: jsii.Number(123),
//   			BytesLessThan: jsii.Number(123),
//   		},
//   		Or: &OrProperty{
//   			MatchAnyPrefix: []*string{
//   				jsii.String("matchAnyPrefix"),
//   			},
//   			MatchAnySuffix: []*string{
//   				jsii.String("matchAnySuffix"),
//   			},
//   			MatchAnyTag: []interface{}{
//   				&CfnTag{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			MatchObjectAge: &MatchObjectAgeProperty{
//   				DaysGreaterThan: jsii.Number(123),
//   				DaysLessThan: jsii.Number(123),
//   			},
//   			MatchObjectSize: &MatchObjectSizeProperty{
//   				BytesGreaterThan: jsii.Number(123),
//   				BytesLessThan: jsii.Number(123),
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-storagelensgroup.html
//
type CfnStorageLensGroupPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnStorageLensGroupMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnStorageLensGroupPropsMixin
type jsiiProxy_CfnStorageLensGroupPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnStorageLensGroupPropsMixin) Props() *CfnStorageLensGroupMixinProps {
	var returns *CfnStorageLensGroupMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStorageLensGroupPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::S3::StorageLensGroup`.
func NewCfnStorageLensGroupPropsMixin(props *CfnStorageLensGroupMixinProps, options *mixins.CfnPropertyMixinOptions) CfnStorageLensGroupPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnStorageLensGroupPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnStorageLensGroupPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnStorageLensGroupPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::S3::StorageLensGroup`.
func NewCfnStorageLensGroupPropsMixin_Override(c CfnStorageLensGroupPropsMixin, props *CfnStorageLensGroupMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnStorageLensGroupPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnStorageLensGroupPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnStorageLensGroupPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnStorageLensGroupPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnStorageLensGroupPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnStorageLensGroupPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnStorageLensGroupPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnStorageLensGroupPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

