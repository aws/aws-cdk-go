package previewawsdeadlinemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsdeadline/previewawsdeadlinemixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a queue to coordinate the order in which jobs run on a farm.
//
// A queue can also specify where to pull resources and indicate where to output completed jobs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnQueuePropsMixin := awscdkmixinspreview.Mixins.NewCfnQueuePropsMixin(&CfnQueueMixinProps{
//   	AllowedStorageProfileIds: []*string{
//   		jsii.String("allowedStorageProfileIds"),
//   	},
//   	DefaultBudgetAction: jsii.String("defaultBudgetAction"),
//   	Description: jsii.String("description"),
//   	DisplayName: jsii.String("displayName"),
//   	FarmId: jsii.String("farmId"),
//   	JobAttachmentSettings: &JobAttachmentSettingsProperty{
//   		RootPrefix: jsii.String("rootPrefix"),
//   		S3BucketName: jsii.String("s3BucketName"),
//   	},
//   	JobRunAsUser: &JobRunAsUserProperty{
//   		Posix: &PosixUserProperty{
//   			Group: jsii.String("group"),
//   			User: jsii.String("user"),
//   		},
//   		RunAs: jsii.String("runAs"),
//   		Windows: &WindowsUserProperty{
//   			PasswordArn: jsii.String("passwordArn"),
//   			User: jsii.String("user"),
//   		},
//   	},
//   	RequiredFileSystemLocationNames: []*string{
//   		jsii.String("requiredFileSystemLocationNames"),
//   	},
//   	RoleArn: jsii.String("roleArn"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-queue.html
//
type CfnQueuePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnQueueMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnQueuePropsMixin
type jsiiProxy_CfnQueuePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnQueuePropsMixin) Props() *CfnQueueMixinProps {
	var returns *CfnQueueMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnQueuePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Deadline::Queue`.
func NewCfnQueuePropsMixin(props *CfnQueueMixinProps, options *mixins.CfnPropertyMixinOptions) CfnQueuePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnQueuePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnQueuePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_deadline.mixins.CfnQueuePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Deadline::Queue`.
func NewCfnQueuePropsMixin_Override(c CfnQueuePropsMixin, props *CfnQueueMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_deadline.mixins.CfnQueuePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnQueuePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnQueuePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_deadline.mixins.CfnQueuePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnQueuePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_deadline.mixins.CfnQueuePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnQueuePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnQueuePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

