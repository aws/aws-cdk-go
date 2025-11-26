package previewawsworkspaceswebmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsworkspacesweb/previewawsworkspaceswebmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The session logger resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var all interface{}
//
//   cfnSessionLoggerPropsMixin := awscdkmixinspreview.Mixins.NewCfnSessionLoggerPropsMixin(&CfnSessionLoggerMixinProps{
//   	AdditionalEncryptionContext: map[string]*string{
//   		"additionalEncryptionContextKey": jsii.String("additionalEncryptionContext"),
//   	},
//   	CustomerManagedKey: jsii.String("customerManagedKey"),
//   	DisplayName: jsii.String("displayName"),
//   	EventFilter: &EventFilterProperty{
//   		All: all,
//   		Include: []*string{
//   			jsii.String("include"),
//   		},
//   	},
//   	LogConfiguration: &LogConfigurationProperty{
//   		S3: &S3LogConfigurationProperty{
//   			Bucket: jsii.String("bucket"),
//   			BucketOwner: jsii.String("bucketOwner"),
//   			FolderStructure: jsii.String("folderStructure"),
//   			KeyPrefix: jsii.String("keyPrefix"),
//   			LogFileFormat: jsii.String("logFileFormat"),
//   		},
//   	},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-sessionlogger.html
//
type CfnSessionLoggerPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnSessionLoggerMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnSessionLoggerPropsMixin
type jsiiProxy_CfnSessionLoggerPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnSessionLoggerPropsMixin) Props() *CfnSessionLoggerMixinProps {
	var returns *CfnSessionLoggerMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSessionLoggerPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::WorkSpacesWeb::SessionLogger`.
func NewCfnSessionLoggerPropsMixin(props *CfnSessionLoggerMixinProps, options *mixins.CfnPropertyMixinOptions) CfnSessionLoggerPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnSessionLoggerPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnSessionLoggerPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_workspacesweb.mixins.CfnSessionLoggerPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::WorkSpacesWeb::SessionLogger`.
func NewCfnSessionLoggerPropsMixin_Override(c CfnSessionLoggerPropsMixin, props *CfnSessionLoggerMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_workspacesweb.mixins.CfnSessionLoggerPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnSessionLoggerPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSessionLoggerPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_workspacesweb.mixins.CfnSessionLoggerPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSessionLoggerPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_workspacesweb.mixins.CfnSessionLoggerPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSessionLoggerPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnSessionLoggerPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

