package previewawsdatasyncmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsdatasync/previewawsdatasyncmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::DataSync::Task` resource specifies a task.
//
// A task is a set of two locations (source and destination) and a set of `Options` that you use to control the behavior of a task. If you don't specify `Options` when you create a task, AWS DataSync populates them with service defaults.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTaskPropsMixin := awscdkmixinspreview.Mixins.NewCfnTaskPropsMixin(&CfnTaskMixinProps{
//   	CloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   	DestinationLocationArn: jsii.String("destinationLocationArn"),
//   	Excludes: []interface{}{
//   		&FilterRuleProperty{
//   			FilterType: jsii.String("filterType"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Includes: []interface{}{
//   		&FilterRuleProperty{
//   			FilterType: jsii.String("filterType"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ManifestConfig: &ManifestConfigProperty{
//   		Action: jsii.String("action"),
//   		Format: jsii.String("format"),
//   		Source: &SourceProperty{
//   			S3: &ManifestConfigSourceS3Property{
//   				BucketAccessRoleArn: jsii.String("bucketAccessRoleArn"),
//   				ManifestObjectPath: jsii.String("manifestObjectPath"),
//   				ManifestObjectVersionId: jsii.String("manifestObjectVersionId"),
//   				S3BucketArn: jsii.String("s3BucketArn"),
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Options: &OptionsProperty{
//   		Atime: jsii.String("atime"),
//   		BytesPerSecond: jsii.Number(123),
//   		Gid: jsii.String("gid"),
//   		LogLevel: jsii.String("logLevel"),
//   		Mtime: jsii.String("mtime"),
//   		ObjectTags: jsii.String("objectTags"),
//   		OverwriteMode: jsii.String("overwriteMode"),
//   		PosixPermissions: jsii.String("posixPermissions"),
//   		PreserveDeletedFiles: jsii.String("preserveDeletedFiles"),
//   		PreserveDevices: jsii.String("preserveDevices"),
//   		SecurityDescriptorCopyFlags: jsii.String("securityDescriptorCopyFlags"),
//   		TaskQueueing: jsii.String("taskQueueing"),
//   		TransferMode: jsii.String("transferMode"),
//   		Uid: jsii.String("uid"),
//   		VerifyMode: jsii.String("verifyMode"),
//   	},
//   	Schedule: &TaskScheduleProperty{
//   		ScheduleExpression: jsii.String("scheduleExpression"),
//   		Status: jsii.String("status"),
//   	},
//   	SourceLocationArn: jsii.String("sourceLocationArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TaskMode: jsii.String("taskMode"),
//   	TaskReportConfig: &TaskReportConfigProperty{
//   		Destination: &DestinationProperty{
//   			S3: &S3Property{
//   				BucketAccessRoleArn: jsii.String("bucketAccessRoleArn"),
//   				S3BucketArn: jsii.String("s3BucketArn"),
//   				Subdirectory: jsii.String("subdirectory"),
//   			},
//   		},
//   		ObjectVersionIds: jsii.String("objectVersionIds"),
//   		OutputType: jsii.String("outputType"),
//   		Overrides: &OverridesProperty{
//   			Deleted: &DeletedProperty{
//   				ReportLevel: jsii.String("reportLevel"),
//   			},
//   			Skipped: &SkippedProperty{
//   				ReportLevel: jsii.String("reportLevel"),
//   			},
//   			Transferred: &TransferredProperty{
//   				ReportLevel: jsii.String("reportLevel"),
//   			},
//   			Verified: &VerifiedProperty{
//   				ReportLevel: jsii.String("reportLevel"),
//   			},
//   		},
//   		ReportLevel: jsii.String("reportLevel"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-task.html
//
type CfnTaskPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnTaskMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnTaskPropsMixin
type jsiiProxy_CfnTaskPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnTaskPropsMixin) Props() *CfnTaskMixinProps {
	var returns *CfnTaskMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::DataSync::Task`.
func NewCfnTaskPropsMixin(props *CfnTaskMixinProps, options *mixins.CfnPropertyMixinOptions) CfnTaskPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnTaskPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTaskPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnTaskPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::DataSync::Task`.
func NewCfnTaskPropsMixin_Override(c CfnTaskPropsMixin, props *CfnTaskMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnTaskPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnTaskPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTaskPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnTaskPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTaskPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnTaskPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTaskPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnTaskPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

