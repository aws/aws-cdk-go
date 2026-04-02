package previewawseksmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawseks/previewawseksmixins/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslogs"
	"github.com/aws/constructs-go/constructs/v10"
)

// An object representing a managed capability in an Amazon EKS cluster.
//
// This includes all configuration, status, and health information for the capability.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var logsDelivery ILogsDelivery
//
//   cfnCapabilityLogsMixin := awscdkmixinspreview.Mixins.NewCfnCapabilityLogsMixin(jsii.String("logType"), logsDelivery)
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-capability.html
//
type CfnCapabilityLogsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	LogDelivery() previewawslogs.ILogsDelivery
	LogType() *string
	// Apply vended logs configuration to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct (has vendedLogs property).
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnCapabilityLogsMixin
type jsiiProxy_CfnCapabilityLogsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnCapabilityLogsMixin) LogDelivery() previewawslogs.ILogsDelivery {
	var returns previewawslogs.ILogsDelivery
	_jsii_.Get(
		j,
		"logDelivery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCapabilityLogsMixin) LogType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logType",
		&returns,
	)
	return returns
}


// Create a mixin to enable vended logs for `AWS::EKS::Capability`.
func NewCfnCapabilityLogsMixin(logType *string, logDelivery previewawslogs.ILogsDelivery) CfnCapabilityLogsMixin {
	_init_.Initialize()

	if err := validateNewCfnCapabilityLogsMixinParameters(logType, logDelivery); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCapabilityLogsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityLogsMixin",
		[]interface{}{logType, logDelivery},
		&j,
	)

	return &j
}

// Create a mixin to enable vended logs for `AWS::EKS::Capability`.
func NewCfnCapabilityLogsMixin_Override(c CfnCapabilityLogsMixin, logType *string, logDelivery previewawslogs.ILogsDelivery) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityLogsMixin",
		[]interface{}{logType, logDelivery},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnCapabilityLogsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCapabilityLogsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityLogsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCapabilityLogsMixin_EKS_CAPABILITY_ACK_LOGS() CfnCapabilityEksCapabilityAckLogs {
	_init_.Initialize()
	var returns CfnCapabilityEksCapabilityAckLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityLogsMixin",
		"EKS_CAPABILITY_ACK_LOGS",
		&returns,
	)
	return returns
}

func CfnCapabilityLogsMixin_EKS_CAPABILITY_ARGOCD_APPLICATION_LOGS() CfnCapabilityEksCapabilityArgocdApplicationLogs {
	_init_.Initialize()
	var returns CfnCapabilityEksCapabilityArgocdApplicationLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityLogsMixin",
		"EKS_CAPABILITY_ARGOCD_APPLICATION_LOGS",
		&returns,
	)
	return returns
}

func CfnCapabilityLogsMixin_EKS_CAPABILITY_ARGOCD_APPLICATIONSET_LOGS() CfnCapabilityEksCapabilityArgocdApplicationsetLogs {
	_init_.Initialize()
	var returns CfnCapabilityEksCapabilityArgocdApplicationsetLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityLogsMixin",
		"EKS_CAPABILITY_ARGOCD_APPLICATIONSET_LOGS",
		&returns,
	)
	return returns
}

func CfnCapabilityLogsMixin_EKS_CAPABILITY_ARGOCD_COMMITSERVER_LOGS() CfnCapabilityEksCapabilityArgocdCommitserverLogs {
	_init_.Initialize()
	var returns CfnCapabilityEksCapabilityArgocdCommitserverLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityLogsMixin",
		"EKS_CAPABILITY_ARGOCD_COMMITSERVER_LOGS",
		&returns,
	)
	return returns
}

func CfnCapabilityLogsMixin_EKS_CAPABILITY_ARGOCD_REPOSERVER_LOGS() CfnCapabilityEksCapabilityArgocdReposerverLogs {
	_init_.Initialize()
	var returns CfnCapabilityEksCapabilityArgocdReposerverLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityLogsMixin",
		"EKS_CAPABILITY_ARGOCD_REPOSERVER_LOGS",
		&returns,
	)
	return returns
}

func CfnCapabilityLogsMixin_EKS_CAPABILITY_ARGOCD_SERVER_LOGS() CfnCapabilityEksCapabilityArgocdServerLogs {
	_init_.Initialize()
	var returns CfnCapabilityEksCapabilityArgocdServerLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityLogsMixin",
		"EKS_CAPABILITY_ARGOCD_SERVER_LOGS",
		&returns,
	)
	return returns
}

func CfnCapabilityLogsMixin_EKS_CAPABILITY_KRO_LOGS() CfnCapabilityEksCapabilityKroLogs {
	_init_.Initialize()
	var returns CfnCapabilityEksCapabilityKroLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityLogsMixin",
		"EKS_CAPABILITY_KRO_LOGS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCapabilityLogsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnCapabilityLogsMixin) Supports(construct constructs.IConstruct) *bool {
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

