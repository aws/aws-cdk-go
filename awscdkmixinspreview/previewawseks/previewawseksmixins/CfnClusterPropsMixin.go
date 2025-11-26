package previewawseksmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawseks/previewawseksmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an Amazon EKS control plane.
//
// The Amazon EKS control plane consists of control plane instances that run the Kubernetes software, such as `etcd` and the API server. The control plane runs in an account managed by AWS , and the Kubernetes API is exposed by the Amazon EKS API server endpoint. Each Amazon EKS cluster control plane is single tenant and unique. It runs on its own set of Amazon EC2 instances.
//
// The cluster control plane is provisioned across multiple Availability Zones and fronted by an ELB Network Load Balancer. Amazon EKS also provisions elastic network interfaces in your VPC subnets to provide connectivity from the control plane instances to the nodes (for example, to support `kubectl exec` , `logs` , and `proxy` data flows).
//
// Amazon EKS nodes run in your AWS account and connect to your cluster's control plane over the Kubernetes API server endpoint and a certificate file that is created for your cluster.
//
// You can use the `endpointPublicAccess` and `endpointPrivateAccess` parameters to enable or disable public and private access to your cluster's Kubernetes API server endpoint. By default, public access is enabled, and private access is disabled. The endpoint domain name and IP address family depends on the value of the `ipFamily` for the cluster. For more information, see [Amazon EKS Cluster Endpoint Access Control](https://docs.aws.amazon.com/eks/latest/userguide/cluster-endpoint.html) in the **Amazon EKS User Guide** .
//
// You can use the `logging` parameter to enable or disable exporting the Kubernetes control plane logs for your cluster to CloudWatch Logs. By default, cluster control plane logs aren't exported to CloudWatch Logs. For more information, see [Amazon EKS Cluster Control Plane Logs](https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html) in the **Amazon EKS User Guide** .
//
// > CloudWatch Logs ingestion, archive storage, and data scanning rates apply to exported control plane logs. For more information, see [CloudWatch Pricing](https://docs.aws.amazon.com/cloudwatch/pricing/) .
//
// In most cases, it takes several minutes to create a cluster. After you create an Amazon EKS cluster, you must configure your Kubernetes tooling to communicate with the API server and launch nodes into your cluster. For more information, see [Allowing users to access your cluster](https://docs.aws.amazon.com/eks/latest/userguide/cluster-auth.html) and [Launching Amazon EKS nodes](https://docs.aws.amazon.com/eks/latest/userguide/launch-workers.html) in the *Amazon EKS User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnClusterPropsMixin := awscdkmixinspreview.Mixins.NewCfnClusterPropsMixin(&CfnClusterMixinProps{
//   	AccessConfig: &AccessConfigProperty{
//   		AuthenticationMode: jsii.String("authenticationMode"),
//   		BootstrapClusterCreatorAdminPermissions: jsii.Boolean(false),
//   	},
//   	BootstrapSelfManagedAddons: jsii.Boolean(false),
//   	ComputeConfig: &ComputeConfigProperty{
//   		Enabled: jsii.Boolean(false),
//   		NodePools: []*string{
//   			jsii.String("nodePools"),
//   		},
//   		NodeRoleArn: jsii.String("nodeRoleArn"),
//   	},
//   	ControlPlaneScalingConfig: &ControlPlaneScalingConfigProperty{
//   		Tier: jsii.String("tier"),
//   	},
//   	DeletionProtection: jsii.Boolean(false),
//   	EncryptionConfig: []interface{}{
//   		&EncryptionConfigProperty{
//   			Provider: &ProviderProperty{
//   				KeyArn: jsii.String("keyArn"),
//   			},
//   			Resources: []*string{
//   				jsii.String("resources"),
//   			},
//   		},
//   	},
//   	Force: jsii.Boolean(false),
//   	KubernetesNetworkConfig: &KubernetesNetworkConfigProperty{
//   		ElasticLoadBalancing: &ElasticLoadBalancingProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   		IpFamily: jsii.String("ipFamily"),
//   		ServiceIpv4Cidr: jsii.String("serviceIpv4Cidr"),
//   		ServiceIpv6Cidr: jsii.String("serviceIpv6Cidr"),
//   	},
//   	Logging: &LoggingProperty{
//   		ClusterLogging: &ClusterLoggingProperty{
//   			EnabledTypes: []interface{}{
//   				&LoggingTypeConfigProperty{
//   					Type: jsii.String("type"),
//   				},
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	OutpostConfig: &OutpostConfigProperty{
//   		ControlPlaneInstanceType: jsii.String("controlPlaneInstanceType"),
//   		ControlPlanePlacement: &ControlPlanePlacementProperty{
//   			GroupName: jsii.String("groupName"),
//   		},
//   		OutpostArns: []*string{
//   			jsii.String("outpostArns"),
//   		},
//   	},
//   	RemoteNetworkConfig: &RemoteNetworkConfigProperty{
//   		RemoteNodeNetworks: []interface{}{
//   			&RemoteNodeNetworkProperty{
//   				Cidrs: []*string{
//   					jsii.String("cidrs"),
//   				},
//   			},
//   		},
//   		RemotePodNetworks: []interface{}{
//   			&RemotePodNetworkProperty{
//   				Cidrs: []*string{
//   					jsii.String("cidrs"),
//   				},
//   			},
//   		},
//   	},
//   	ResourcesVpcConfig: &ResourcesVpcConfigProperty{
//   		EndpointPrivateAccess: jsii.Boolean(false),
//   		EndpointPublicAccess: jsii.Boolean(false),
//   		PublicAccessCidrs: []*string{
//   			jsii.String("publicAccessCidrs"),
//   		},
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   	StorageConfig: &StorageConfigProperty{
//   		BlockStorage: &BlockStorageProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	UpgradePolicy: &UpgradePolicyProperty{
//   		SupportType: jsii.String("supportType"),
//   	},
//   	Version: jsii.String("version"),
//   	ZonalShiftConfig: &ZonalShiftConfigProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-cluster.html
//
type CfnClusterPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnClusterMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
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


// Create a mixin to apply properties to `AWS::EKS::Cluster`.
func NewCfnClusterPropsMixin(props *CfnClusterMixinProps, options *mixins.CfnPropertyMixinOptions) CfnClusterPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnClusterPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnClusterPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EKS::Cluster`.
func NewCfnClusterPropsMixin_Override(c CfnClusterPropsMixin, props *CfnClusterMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterPropsMixin",
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
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterPropsMixin",
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
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnClusterPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

