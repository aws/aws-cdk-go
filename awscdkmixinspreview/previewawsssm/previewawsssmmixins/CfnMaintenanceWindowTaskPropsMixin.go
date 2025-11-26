package previewawsssmmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsssm/previewawsssmmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::SSM::MaintenanceWindowTask` resource defines information about a task for an AWS Systems Manager maintenance window.
//
// For more information, see [RegisterTaskWithMaintenanceWindow](https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_RegisterTaskWithMaintenanceWindow.html) in the *AWS Systems Manager API Reference* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var parameters interface{}
//   var taskParameters interface{}
//
//   cfnMaintenanceWindowTaskPropsMixin := awscdkmixinspreview.Mixins.NewCfnMaintenanceWindowTaskPropsMixin(&CfnMaintenanceWindowTaskMixinProps{
//   	CutoffBehavior: jsii.String("cutoffBehavior"),
//   	Description: jsii.String("description"),
//   	LoggingInfo: &LoggingInfoProperty{
//   		Region: jsii.String("region"),
//   		S3Bucket: jsii.String("s3Bucket"),
//   		S3Prefix: jsii.String("s3Prefix"),
//   	},
//   	MaxConcurrency: jsii.String("maxConcurrency"),
//   	MaxErrors: jsii.String("maxErrors"),
//   	Name: jsii.String("name"),
//   	Priority: jsii.Number(123),
//   	ServiceRoleArn: jsii.String("serviceRoleArn"),
//   	Targets: []interface{}{
//   		&TargetProperty{
//   			Key: jsii.String("key"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   	TaskArn: jsii.String("taskArn"),
//   	TaskInvocationParameters: &TaskInvocationParametersProperty{
//   		MaintenanceWindowAutomationParameters: &MaintenanceWindowAutomationParametersProperty{
//   			DocumentVersion: jsii.String("documentVersion"),
//   			Parameters: parameters,
//   		},
//   		MaintenanceWindowLambdaParameters: &MaintenanceWindowLambdaParametersProperty{
//   			ClientContext: jsii.String("clientContext"),
//   			Payload: jsii.String("payload"),
//   			Qualifier: jsii.String("qualifier"),
//   		},
//   		MaintenanceWindowRunCommandParameters: &MaintenanceWindowRunCommandParametersProperty{
//   			CloudWatchOutputConfig: &CloudWatchOutputConfigProperty{
//   				CloudWatchLogGroupName: jsii.String("cloudWatchLogGroupName"),
//   				CloudWatchOutputEnabled: jsii.Boolean(false),
//   			},
//   			Comment: jsii.String("comment"),
//   			DocumentHash: jsii.String("documentHash"),
//   			DocumentHashType: jsii.String("documentHashType"),
//   			DocumentVersion: jsii.String("documentVersion"),
//   			NotificationConfig: &NotificationConfigProperty{
//   				NotificationArn: jsii.String("notificationArn"),
//   				NotificationEvents: []*string{
//   					jsii.String("notificationEvents"),
//   				},
//   				NotificationType: jsii.String("notificationType"),
//   			},
//   			OutputS3BucketName: jsii.String("outputS3BucketName"),
//   			OutputS3KeyPrefix: jsii.String("outputS3KeyPrefix"),
//   			Parameters: parameters,
//   			ServiceRoleArn: jsii.String("serviceRoleArn"),
//   			TimeoutSeconds: jsii.Number(123),
//   		},
//   		MaintenanceWindowStepFunctionsParameters: &MaintenanceWindowStepFunctionsParametersProperty{
//   			Input: jsii.String("input"),
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	TaskParameters: taskParameters,
//   	TaskType: jsii.String("taskType"),
//   	WindowId: jsii.String("windowId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtask.html
//
type CfnMaintenanceWindowTaskPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnMaintenanceWindowTaskMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnMaintenanceWindowTaskPropsMixin
type jsiiProxy_CfnMaintenanceWindowTaskPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnMaintenanceWindowTaskPropsMixin) Props() *CfnMaintenanceWindowTaskMixinProps {
	var returns *CfnMaintenanceWindowTaskMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTaskPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SSM::MaintenanceWindowTask`.
func NewCfnMaintenanceWindowTaskPropsMixin(props *CfnMaintenanceWindowTaskMixinProps, options *mixins.CfnPropertyMixinOptions) CfnMaintenanceWindowTaskPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnMaintenanceWindowTaskPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnMaintenanceWindowTaskPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnMaintenanceWindowTaskPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SSM::MaintenanceWindowTask`.
func NewCfnMaintenanceWindowTaskPropsMixin_Override(c CfnMaintenanceWindowTaskPropsMixin, props *CfnMaintenanceWindowTaskMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnMaintenanceWindowTaskPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnMaintenanceWindowTaskPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnMaintenanceWindowTaskPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnMaintenanceWindowTaskPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMaintenanceWindowTaskPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnMaintenanceWindowTaskPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMaintenanceWindowTaskPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnMaintenanceWindowTaskPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

