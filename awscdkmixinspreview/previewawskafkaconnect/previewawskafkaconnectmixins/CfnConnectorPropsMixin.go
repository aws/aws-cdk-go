package previewawskafkaconnectmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawskafkaconnect/previewawskafkaconnectmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a connector using the specified properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnConnectorPropsMixin := awscdkmixinspreview.Mixins.NewCfnConnectorPropsMixin(&CfnConnectorMixinProps{
//   	Capacity: &CapacityProperty{
//   		AutoScaling: &AutoScalingProperty{
//   			MaxWorkerCount: jsii.Number(123),
//   			McuCount: jsii.Number(123),
//   			MinWorkerCount: jsii.Number(123),
//   			ScaleInPolicy: &ScaleInPolicyProperty{
//   				CpuUtilizationPercentage: jsii.Number(123),
//   			},
//   			ScaleOutPolicy: &ScaleOutPolicyProperty{
//   				CpuUtilizationPercentage: jsii.Number(123),
//   			},
//   		},
//   		ProvisionedCapacity: &ProvisionedCapacityProperty{
//   			McuCount: jsii.Number(123),
//   			WorkerCount: jsii.Number(123),
//   		},
//   	},
//   	ConnectorConfiguration: map[string]*string{
//   		"connectorConfigurationKey": jsii.String("connectorConfiguration"),
//   	},
//   	ConnectorDescription: jsii.String("connectorDescription"),
//   	ConnectorName: jsii.String("connectorName"),
//   	KafkaCluster: &KafkaClusterProperty{
//   		ApacheKafkaCluster: &ApacheKafkaClusterProperty{
//   			BootstrapServers: jsii.String("bootstrapServers"),
//   			Vpc: &VpcProperty{
//   				SecurityGroups: []*string{
//   					jsii.String("securityGroups"),
//   				},
//   				Subnets: []*string{
//   					jsii.String("subnets"),
//   				},
//   			},
//   		},
//   	},
//   	KafkaClusterClientAuthentication: &KafkaClusterClientAuthenticationProperty{
//   		AuthenticationType: jsii.String("authenticationType"),
//   	},
//   	KafkaClusterEncryptionInTransit: &KafkaClusterEncryptionInTransitProperty{
//   		EncryptionType: jsii.String("encryptionType"),
//   	},
//   	KafkaConnectVersion: jsii.String("kafkaConnectVersion"),
//   	LogDelivery: &LogDeliveryProperty{
//   		WorkerLogDelivery: &WorkerLogDeliveryProperty{
//   			CloudWatchLogs: &CloudWatchLogsLogDeliveryProperty{
//   				Enabled: jsii.Boolean(false),
//   				LogGroup: jsii.String("logGroup"),
//   			},
//   			Firehose: &FirehoseLogDeliveryProperty{
//   				DeliveryStream: jsii.String("deliveryStream"),
//   				Enabled: jsii.Boolean(false),
//   			},
//   			S3: &S3LogDeliveryProperty{
//   				Bucket: jsii.String("bucket"),
//   				Enabled: jsii.Boolean(false),
//   				Prefix: jsii.String("prefix"),
//   			},
//   		},
//   	},
//   	NetworkType: jsii.String("networkType"),
//   	Plugins: []interface{}{
//   		&PluginProperty{
//   			CustomPlugin: &CustomPluginProperty{
//   				CustomPluginArn: jsii.String("customPluginArn"),
//   				Revision: jsii.Number(123),
//   			},
//   		},
//   	},
//   	ServiceExecutionRoleArn: jsii.String("serviceExecutionRoleArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	WorkerConfiguration: &WorkerConfigurationProperty{
//   		Revision: jsii.Number(123),
//   		WorkerConfigurationArn: jsii.String("workerConfigurationArn"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kafkaconnect-connector.html
//
type CfnConnectorPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnConnectorMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnConnectorPropsMixin
type jsiiProxy_CfnConnectorPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnConnectorPropsMixin) Props() *CfnConnectorMixinProps {
	var returns *CfnConnectorMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::KafkaConnect::Connector`.
func NewCfnConnectorPropsMixin(props *CfnConnectorMixinProps, options *mixins.CfnPropertyMixinOptions) CfnConnectorPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnConnectorPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnConnectorPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_kafkaconnect.mixins.CfnConnectorPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::KafkaConnect::Connector`.
func NewCfnConnectorPropsMixin_Override(c CfnConnectorPropsMixin, props *CfnConnectorMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_kafkaconnect.mixins.CfnConnectorPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnConnectorPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnConnectorPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_kafkaconnect.mixins.CfnConnectorPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnConnectorPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_kafkaconnect.mixins.CfnConnectorPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnConnectorPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnConnectorPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

