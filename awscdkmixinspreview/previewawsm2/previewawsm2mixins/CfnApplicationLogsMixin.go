package previewawsm2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslogs"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsm2/previewawsm2mixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies a new application with given parameters. Requires an existing runtime environment and application definition file.
//
// For information about application definitions, see the [AWS Mainframe Modernization User Guide](https://docs.aws.amazon.com/m2/latest/userguide/applications-m2-definition.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var logsDelivery ILogsDelivery
//
//   cfnApplicationLogsMixin := awscdkmixinspreview.Mixins.NewCfnApplicationLogsMixin(jsii.String("logType"), logsDelivery)
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-m2-application.html
//
type CfnApplicationLogsMixin interface {
	core.Mixin
	core.IMixin
	LogDelivery() previewawslogs.ILogsDelivery
	LogType() *string
	// Apply vended logs configuration to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct (has vendedLogs property).
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnApplicationLogsMixin
type jsiiProxy_CfnApplicationLogsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnApplicationLogsMixin) LogDelivery() previewawslogs.ILogsDelivery {
	var returns previewawslogs.ILogsDelivery
	_jsii_.Get(
		j,
		"logDelivery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationLogsMixin) LogType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logType",
		&returns,
	)
	return returns
}


// Create a mixin to enable vended logs for `AWS::M2::Application`.
func NewCfnApplicationLogsMixin(logType *string, logDelivery previewawslogs.ILogsDelivery) CfnApplicationLogsMixin {
	_init_.Initialize()

	if err := validateNewCfnApplicationLogsMixinParameters(logType, logDelivery); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnApplicationLogsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnApplicationLogsMixin",
		[]interface{}{logType, logDelivery},
		&j,
	)

	return &j
}

// Create a mixin to enable vended logs for `AWS::M2::Application`.
func NewCfnApplicationLogsMixin_Override(c CfnApplicationLogsMixin, logType *string, logDelivery previewawslogs.ILogsDelivery) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnApplicationLogsMixin",
		[]interface{}{logType, logDelivery},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnApplicationLogsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnApplicationLogsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnApplicationLogsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnApplicationLogsMixin_BATCH_JOB_LOGS() CfnApplicationBatchJobLogs {
	_init_.Initialize()
	var returns CfnApplicationBatchJobLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnApplicationLogsMixin",
		"BATCH_JOB_LOGS",
		&returns,
	)
	return returns
}

func CfnApplicationLogsMixin_CONFIG_LOGS() CfnApplicationConfigLogs {
	_init_.Initialize()
	var returns CfnApplicationConfigLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnApplicationLogsMixin",
		"CONFIG_LOGS",
		&returns,
	)
	return returns
}

func CfnApplicationLogsMixin_CONSOLE_LOGS() CfnApplicationConsoleLogs {
	_init_.Initialize()
	var returns CfnApplicationConsoleLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnApplicationLogsMixin",
		"CONSOLE_LOGS",
		&returns,
	)
	return returns
}

func CfnApplicationLogsMixin_DATASET_IMPORT_LOGS() CfnApplicationDatasetImportLogs {
	_init_.Initialize()
	var returns CfnApplicationDatasetImportLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnApplicationLogsMixin",
		"DATASET_IMPORT_LOGS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnApplicationLogsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnApplicationLogsMixin) Supports(construct constructs.IConstruct) *bool {
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

