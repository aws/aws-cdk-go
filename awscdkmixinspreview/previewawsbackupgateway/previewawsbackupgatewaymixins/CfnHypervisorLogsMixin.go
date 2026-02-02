package previewawsbackupgatewaymixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsbackupgateway/previewawsbackupgatewaymixins/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslogs"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents the hypervisor's permissions to which the gateway will connect.
//
// A hypervisor is hardware, software, or firmware that creates and manages virtual machines, and allocates resources to them.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var logsDelivery ILogsDelivery
//
//   cfnHypervisorLogsMixin := awscdkmixinspreview.Mixins.NewCfnHypervisorLogsMixin(jsii.String("logType"), logsDelivery)
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backupgateway-hypervisor.html
//
type CfnHypervisorLogsMixin interface {
	core.Mixin
	core.IMixin
	LogDelivery() previewawslogs.ILogsDelivery
	LogType() *string
	// Apply vended logs configuration to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct (has vendedLogs property).
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnHypervisorLogsMixin
type jsiiProxy_CfnHypervisorLogsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnHypervisorLogsMixin) LogDelivery() previewawslogs.ILogsDelivery {
	var returns previewawslogs.ILogsDelivery
	_jsii_.Get(
		j,
		"logDelivery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHypervisorLogsMixin) LogType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logType",
		&returns,
	)
	return returns
}


// Create a mixin to enable vended logs for `AWS::BackupGateway::Hypervisor`.
func NewCfnHypervisorLogsMixin(logType *string, logDelivery previewawslogs.ILogsDelivery) CfnHypervisorLogsMixin {
	_init_.Initialize()

	if err := validateNewCfnHypervisorLogsMixinParameters(logType, logDelivery); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnHypervisorLogsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_backupgateway.mixins.CfnHypervisorLogsMixin",
		[]interface{}{logType, logDelivery},
		&j,
	)

	return &j
}

// Create a mixin to enable vended logs for `AWS::BackupGateway::Hypervisor`.
func NewCfnHypervisorLogsMixin_Override(c CfnHypervisorLogsMixin, logType *string, logDelivery previewawslogs.ILogsDelivery) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_backupgateway.mixins.CfnHypervisorLogsMixin",
		[]interface{}{logType, logDelivery},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnHypervisorLogsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnHypervisorLogsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_backupgateway.mixins.CfnHypervisorLogsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnHypervisorLogsMixin_BGW_HYPERVISOR_LOGS() CfnHypervisorBgwHypervisorLogs {
	_init_.Initialize()
	var returns CfnHypervisorBgwHypervisorLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_backupgateway.mixins.CfnHypervisorLogsMixin",
		"BGW_HYPERVISOR_LOGS",
		&returns,
	)
	return returns
}

func CfnHypervisorLogsMixin_DATA_ACCESS_LOGS() CfnHypervisorDataAccessLogs {
	_init_.Initialize()
	var returns CfnHypervisorDataAccessLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_backupgateway.mixins.CfnHypervisorLogsMixin",
		"DATA_ACCESS_LOGS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnHypervisorLogsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnHypervisorLogsMixin) Supports(construct constructs.IConstruct) *bool {
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

