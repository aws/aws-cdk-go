package previewawsmskmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsmsk/previewawsmskmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a new MSK cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnClusterPropsMixin := awscdkmixinspreview.Mixins.NewCfnClusterPropsMixin(&CfnClusterMixinProps{
//   	BrokerNodeGroupInfo: &BrokerNodeGroupInfoProperty{
//   		BrokerAzDistribution: jsii.String("brokerAzDistribution"),
//   		ClientSubnets: []*string{
//   			jsii.String("clientSubnets"),
//   		},
//   		ConnectivityInfo: &ConnectivityInfoProperty{
//   			NetworkType: jsii.String("networkType"),
//   			PublicAccess: &PublicAccessProperty{
//   				Type: jsii.String("type"),
//   			},
//   			VpcConnectivity: &VpcConnectivityProperty{
//   				ClientAuthentication: &VpcConnectivityClientAuthenticationProperty{
//   					Sasl: &VpcConnectivitySaslProperty{
//   						Iam: &VpcConnectivityIamProperty{
//   							Enabled: jsii.Boolean(false),
//   						},
//   						Scram: &VpcConnectivityScramProperty{
//   							Enabled: jsii.Boolean(false),
//   						},
//   					},
//   					Tls: &VpcConnectivityTlsProperty{
//   						Enabled: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   		},
//   		InstanceType: jsii.String("instanceType"),
//   		SecurityGroups: []*string{
//   			jsii.String("securityGroups"),
//   		},
//   		StorageInfo: &StorageInfoProperty{
//   			EbsStorageInfo: &EBSStorageInfoProperty{
//   				ProvisionedThroughput: &ProvisionedThroughputProperty{
//   					Enabled: jsii.Boolean(false),
//   					VolumeThroughput: jsii.Number(123),
//   				},
//   				VolumeSize: jsii.Number(123),
//   			},
//   		},
//   	},
//   	ClientAuthentication: &ClientAuthenticationProperty{
//   		Sasl: &SaslProperty{
//   			Iam: &IamProperty{
//   				Enabled: jsii.Boolean(false),
//   			},
//   			Scram: &ScramProperty{
//   				Enabled: jsii.Boolean(false),
//   			},
//   		},
//   		Tls: &TlsProperty{
//   			CertificateAuthorityArnList: []*string{
//   				jsii.String("certificateAuthorityArnList"),
//   			},
//   			Enabled: jsii.Boolean(false),
//   		},
//   		Unauthenticated: &UnauthenticatedProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   	},
//   	ClusterName: jsii.String("clusterName"),
//   	ConfigurationInfo: &ConfigurationInfoProperty{
//   		Arn: jsii.String("arn"),
//   		Revision: jsii.Number(123),
//   	},
//   	CurrentVersion: jsii.String("currentVersion"),
//   	EncryptionInfo: &EncryptionInfoProperty{
//   		EncryptionAtRest: &EncryptionAtRestProperty{
//   			DataVolumeKmsKeyId: jsii.String("dataVolumeKmsKeyId"),
//   		},
//   		EncryptionInTransit: &EncryptionInTransitProperty{
//   			ClientBroker: jsii.String("clientBroker"),
//   			InCluster: jsii.Boolean(false),
//   		},
//   	},
//   	EnhancedMonitoring: jsii.String("enhancedMonitoring"),
//   	KafkaVersion: jsii.String("kafkaVersion"),
//   	LoggingInfo: &LoggingInfoProperty{
//   		BrokerLogs: &BrokerLogsProperty{
//   			CloudWatchLogs: &CloudWatchLogsProperty{
//   				Enabled: jsii.Boolean(false),
//   				LogGroup: jsii.String("logGroup"),
//   			},
//   			Firehose: &FirehoseProperty{
//   				DeliveryStream: jsii.String("deliveryStream"),
//   				Enabled: jsii.Boolean(false),
//   			},
//   			S3: &S3Property{
//   				Bucket: jsii.String("bucket"),
//   				Enabled: jsii.Boolean(false),
//   				Prefix: jsii.String("prefix"),
//   			},
//   		},
//   	},
//   	NumberOfBrokerNodes: jsii.Number(123),
//   	OpenMonitoring: &OpenMonitoringProperty{
//   		Prometheus: &PrometheusProperty{
//   			JmxExporter: &JmxExporterProperty{
//   				EnabledInBroker: jsii.Boolean(false),
//   			},
//   			NodeExporter: &NodeExporterProperty{
//   				EnabledInBroker: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	Rebalancing: &RebalancingProperty{
//   		Status: jsii.String("status"),
//   	},
//   	StorageMode: jsii.String("storageMode"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-cluster.html
//
type CfnClusterPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnClusterMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnClusterPropsMixin
type jsiiProxy_CfnClusterPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnClusterPropsMixin) Props() *CfnClusterMixinProps {
	var returns *CfnClusterMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::MSK::Cluster`.
func NewCfnClusterPropsMixin(props *CfnClusterMixinProps, options *mixins.CfnPropertyMixinOptions) CfnClusterPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnClusterPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnClusterPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnClusterPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::MSK::Cluster`.
func NewCfnClusterPropsMixin_Override(c CfnClusterPropsMixin, props *CfnClusterMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnClusterPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnClusterPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnClusterPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnClusterPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnClusterPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnClusterPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnClusterPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnClusterPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

