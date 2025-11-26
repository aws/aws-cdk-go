package previewawsqbusinessmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslogs"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsqbusiness/previewawsqbusinessmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an Amazon Q Business application.
//
// > There are new tiers for Amazon Q Business. Not all features in Amazon Q Business Pro are also available in Amazon Q Business Lite. For information on what's included in Amazon Q Business Lite and what's included in Amazon Q Business Pro, see [Amazon Q Business tiers](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/tiers.html#user-sub-tiers) . You must use the Amazon Q Business console to assign subscription tiers to users.
// >
// > An Amazon Q Apps service linked role will be created if it's absent in the AWS account when `QAppsConfiguration` is enabled in the request. For more information, see [Using service-linked roles for Q Apps](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/using-service-linked-roles-qapps.html) .
// >
// > When you create an application, Amazon Q Business may securely transmit data for processing from your selected AWS region, but within your geography. For more information, see [Cross region inference in Amazon Q Business](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/cross-region-inference.html) .
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-application.html
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


// Create a mixin to enable vended logs for `AWS::QBusiness::Application`.
func NewCfnApplicationLogsMixin(logType *string, logDelivery previewawslogs.ILogsDelivery) CfnApplicationLogsMixin {
	_init_.Initialize()

	if err := validateNewCfnApplicationLogsMixinParameters(logType, logDelivery); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnApplicationLogsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnApplicationLogsMixin",
		[]interface{}{logType, logDelivery},
		&j,
	)

	return &j
}

// Create a mixin to enable vended logs for `AWS::QBusiness::Application`.
func NewCfnApplicationLogsMixin_Override(c CfnApplicationLogsMixin, logType *string, logDelivery previewawslogs.ILogsDelivery) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnApplicationLogsMixin",
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
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnApplicationLogsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnApplicationLogsMixin_EVENT_LOGS() CfnApplicationEventLogs {
	_init_.Initialize()
	var returns CfnApplicationEventLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnApplicationLogsMixin",
		"EVENT_LOGS",
		&returns,
	)
	return returns
}

func CfnApplicationLogsMixin_SYNC_JOB_LOGS() CfnApplicationSyncJobLogs {
	_init_.Initialize()
	var returns CfnApplicationSyncJobLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnApplicationLogsMixin",
		"SYNC_JOB_LOGS",
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

