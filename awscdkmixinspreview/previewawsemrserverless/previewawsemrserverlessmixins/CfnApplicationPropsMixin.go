package previewawsemrserverlessmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsemrserverless/previewawsemrserverlessmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::EMRServerless::Application` resource specifies an EMR Serverless application.
//
// An application uses open source analytics frameworks to run jobs that process data. To create an application, you must specify the release version for the open source framework version you want to use and the type of application you want, such as Apache Spark or Apache Hive. After you create an application, you can submit data processing jobs or interactive requests to it.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var configurationObjectProperty_ ConfigurationObjectProperty
//
//   cfnApplicationPropsMixin := awscdkmixinspreview.Mixins.NewCfnApplicationPropsMixin(&CfnApplicationMixinProps{
//   	Architecture: jsii.String("architecture"),
//   	AutoStartConfiguration: &AutoStartConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	AutoStopConfiguration: &AutoStopConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   		IdleTimeoutMinutes: jsii.Number(123),
//   	},
//   	IdentityCenterConfiguration: &IdentityCenterConfigurationProperty{
//   		IdentityCenterInstanceArn: jsii.String("identityCenterInstanceArn"),
//   	},
//   	ImageConfiguration: &ImageConfigurationInputProperty{
//   		ImageUri: jsii.String("imageUri"),
//   	},
//   	InitialCapacity: []interface{}{
//   		&InitialCapacityConfigKeyValuePairProperty{
//   			Key: jsii.String("key"),
//   			Value: &InitialCapacityConfigProperty{
//   				WorkerConfiguration: &WorkerConfigurationProperty{
//   					Cpu: jsii.String("cpu"),
//   					Disk: jsii.String("disk"),
//   					DiskType: jsii.String("diskType"),
//   					Memory: jsii.String("memory"),
//   				},
//   				WorkerCount: jsii.Number(123),
//   			},
//   		},
//   	},
//   	InteractiveConfiguration: &InteractiveConfigurationProperty{
//   		LivyEndpointEnabled: jsii.Boolean(false),
//   		StudioEnabled: jsii.Boolean(false),
//   	},
//   	MaximumCapacity: &MaximumAllowedResourcesProperty{
//   		Cpu: jsii.String("cpu"),
//   		Disk: jsii.String("disk"),
//   		Memory: jsii.String("memory"),
//   	},
//   	MonitoringConfiguration: &MonitoringConfigurationProperty{
//   		CloudWatchLoggingConfiguration: &CloudWatchLoggingConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   			EncryptionKeyArn: jsii.String("encryptionKeyArn"),
//   			LogGroupName: jsii.String("logGroupName"),
//   			LogStreamNamePrefix: jsii.String("logStreamNamePrefix"),
//   			LogTypeMap: []interface{}{
//   				&LogTypeMapKeyValuePairProperty{
//   					Key: jsii.String("key"),
//   					Value: []*string{
//   						jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   		ManagedPersistenceMonitoringConfiguration: &ManagedPersistenceMonitoringConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   			EncryptionKeyArn: jsii.String("encryptionKeyArn"),
//   		},
//   		PrometheusMonitoringConfiguration: &PrometheusMonitoringConfigurationProperty{
//   			RemoteWriteUrl: jsii.String("remoteWriteUrl"),
//   		},
//   		S3MonitoringConfiguration: &S3MonitoringConfigurationProperty{
//   			EncryptionKeyArn: jsii.String("encryptionKeyArn"),
//   			LogUri: jsii.String("logUri"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	NetworkConfiguration: &NetworkConfigurationProperty{
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   	ReleaseLabel: jsii.String("releaseLabel"),
//   	RuntimeConfiguration: []interface{}{
//   		&ConfigurationObjectProperty{
//   			Classification: jsii.String("classification"),
//   			Configurations: []interface{}{
//   				configurationObjectProperty_,
//   			},
//   			Properties: map[string]*string{
//   				"propertiesKey": jsii.String("properties"),
//   			},
//   		},
//   	},
//   	SchedulerConfiguration: &SchedulerConfigurationProperty{
//   		MaxConcurrentRuns: jsii.Number(123),
//   		QueueTimeoutMinutes: jsii.Number(123),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   	WorkerTypeSpecifications: map[string]interface{}{
//   		"workerTypeSpecificationsKey": &WorkerTypeSpecificationInputProperty{
//   			"imageConfiguration": &ImageConfigurationInputProperty{
//   				"imageUri": jsii.String("imageUri"),
//   			},
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrserverless-application.html
//
type CfnApplicationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnApplicationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnApplicationPropsMixin
type jsiiProxy_CfnApplicationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnApplicationPropsMixin) Props() *CfnApplicationMixinProps {
	var returns *CfnApplicationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EMRServerless::Application`.
func NewCfnApplicationPropsMixin(props *CfnApplicationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnApplicationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnApplicationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnApplicationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_emrserverless.mixins.CfnApplicationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EMRServerless::Application`.
func NewCfnApplicationPropsMixin_Override(c CfnApplicationPropsMixin, props *CfnApplicationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_emrserverless.mixins.CfnApplicationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnApplicationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnApplicationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_emrserverless.mixins.CfnApplicationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnApplicationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_emrserverless.mixins.CfnApplicationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnApplicationPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnApplicationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

