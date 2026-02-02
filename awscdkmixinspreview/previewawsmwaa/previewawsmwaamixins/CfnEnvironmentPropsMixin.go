package previewawsmwaamixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsmwaa/previewawsmwaamixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::MWAA::Environment` resource creates an Amazon Managed Workflows for Apache Airflow (MWAA) environment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var airflowConfigurationOptions interface{}
//   var tags interface{}
//
//   cfnEnvironmentPropsMixin := awscdkmixinspreview.Mixins.NewCfnEnvironmentPropsMixin(&CfnEnvironmentMixinProps{
//   	AirflowConfigurationOptions: airflowConfigurationOptions,
//   	AirflowVersion: jsii.String("airflowVersion"),
//   	DagS3Path: jsii.String("dagS3Path"),
//   	EndpointManagement: jsii.String("endpointManagement"),
//   	EnvironmentClass: jsii.String("environmentClass"),
//   	ExecutionRoleArn: jsii.String("executionRoleArn"),
//   	KmsKey: jsii.String("kmsKey"),
//   	LoggingConfiguration: &LoggingConfigurationProperty{
//   		DagProcessingLogs: &ModuleLoggingConfigurationProperty{
//   			CloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   			Enabled: jsii.Boolean(false),
//   			LogLevel: jsii.String("logLevel"),
//   		},
//   		SchedulerLogs: &ModuleLoggingConfigurationProperty{
//   			CloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   			Enabled: jsii.Boolean(false),
//   			LogLevel: jsii.String("logLevel"),
//   		},
//   		TaskLogs: &ModuleLoggingConfigurationProperty{
//   			CloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   			Enabled: jsii.Boolean(false),
//   			LogLevel: jsii.String("logLevel"),
//   		},
//   		WebserverLogs: &ModuleLoggingConfigurationProperty{
//   			CloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   			Enabled: jsii.Boolean(false),
//   			LogLevel: jsii.String("logLevel"),
//   		},
//   		WorkerLogs: &ModuleLoggingConfigurationProperty{
//   			CloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   			Enabled: jsii.Boolean(false),
//   			LogLevel: jsii.String("logLevel"),
//   		},
//   	},
//   	MaxWebservers: jsii.Number(123),
//   	MaxWorkers: jsii.Number(123),
//   	MinWebservers: jsii.Number(123),
//   	MinWorkers: jsii.Number(123),
//   	Name: jsii.String("name"),
//   	NetworkConfiguration: &NetworkConfigurationProperty{
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   	PluginsS3ObjectVersion: jsii.String("pluginsS3ObjectVersion"),
//   	PluginsS3Path: jsii.String("pluginsS3Path"),
//   	RequirementsS3ObjectVersion: jsii.String("requirementsS3ObjectVersion"),
//   	RequirementsS3Path: jsii.String("requirementsS3Path"),
//   	Schedulers: jsii.Number(123),
//   	SourceBucketArn: jsii.String("sourceBucketArn"),
//   	StartupScriptS3ObjectVersion: jsii.String("startupScriptS3ObjectVersion"),
//   	StartupScriptS3Path: jsii.String("startupScriptS3Path"),
//   	Tags: tags,
//   	WebserverAccessMode: jsii.String("webserverAccessMode"),
//   	WeeklyMaintenanceWindowStart: jsii.String("weeklyMaintenanceWindowStart"),
//   	WorkerReplacementStrategy: jsii.String("workerReplacementStrategy"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html
//
type CfnEnvironmentPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnEnvironmentMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnEnvironmentPropsMixin
type jsiiProxy_CfnEnvironmentPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnEnvironmentPropsMixin) Props() *CfnEnvironmentMixinProps {
	var returns *CfnEnvironmentMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironmentPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::MWAA::Environment`.
func NewCfnEnvironmentPropsMixin(props *CfnEnvironmentMixinProps, options *mixins.CfnPropertyMixinOptions) CfnEnvironmentPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnEnvironmentPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnEnvironmentPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mwaa.mixins.CfnEnvironmentPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::MWAA::Environment`.
func NewCfnEnvironmentPropsMixin_Override(c CfnEnvironmentPropsMixin, props *CfnEnvironmentMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mwaa.mixins.CfnEnvironmentPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnEnvironmentPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEnvironmentPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_mwaa.mixins.CfnEnvironmentPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEnvironmentPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_mwaa.mixins.CfnEnvironmentPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnEnvironmentPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnEnvironmentPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

