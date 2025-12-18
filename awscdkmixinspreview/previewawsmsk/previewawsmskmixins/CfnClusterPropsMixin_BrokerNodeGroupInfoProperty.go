package previewawsmskmixins


// Describes the setup to be used for the broker nodes in the cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   brokerNodeGroupInfoProperty := &BrokerNodeGroupInfoProperty{
//   	BrokerAzDistribution: jsii.String("brokerAzDistribution"),
//   	ClientSubnets: []*string{
//   		jsii.String("clientSubnets"),
//   	},
//   	ConnectivityInfo: &ConnectivityInfoProperty{
//   		NetworkType: jsii.String("networkType"),
//   		PublicAccess: &PublicAccessProperty{
//   			Type: jsii.String("type"),
//   		},
//   		VpcConnectivity: &VpcConnectivityProperty{
//   			ClientAuthentication: &VpcConnectivityClientAuthenticationProperty{
//   				Sasl: &VpcConnectivitySaslProperty{
//   					Iam: &VpcConnectivityIamProperty{
//   						Enabled: jsii.Boolean(false),
//   					},
//   					Scram: &VpcConnectivityScramProperty{
//   						Enabled: jsii.Boolean(false),
//   					},
//   				},
//   				Tls: &VpcConnectivityTlsProperty{
//   					Enabled: jsii.Boolean(false),
//   				},
//   			},
//   		},
//   	},
//   	InstanceType: jsii.String("instanceType"),
//   	SecurityGroups: []*string{
//   		jsii.String("securityGroups"),
//   	},
//   	StorageInfo: &StorageInfoProperty{
//   		EbsStorageInfo: &EBSStorageInfoProperty{
//   			ProvisionedThroughput: &ProvisionedThroughputProperty{
//   				Enabled: jsii.Boolean(false),
//   				VolumeThroughput: jsii.Number(123),
//   			},
//   			VolumeSize: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-brokernodegroupinfo.html
//
type CfnClusterPropsMixin_BrokerNodeGroupInfoProperty struct {
	// This parameter is currently not in use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-brokernodegroupinfo.html#cfn-msk-cluster-brokernodegroupinfo-brokerazdistribution
	//
	BrokerAzDistribution *string `field:"optional" json:"brokerAzDistribution" yaml:"brokerAzDistribution"`
	// The list of subnets to connect to in the client virtual private cloud (VPC).
	//
	// Amazon creates elastic network interfaces (ENIs) inside these subnets. Client applications use ENIs to produce and consume data.
	//
	// If you use the US West (N. California) Region, specify exactly two subnets. For other Regions where Amazon MSK is available, you can specify either two or three subnets. The subnets that you specify must be in distinct Availability Zones. When you create a cluster, Amazon MSK distributes the broker nodes evenly across the subnets that you specify.
	//
	// Client subnets can't occupy the Availability Zone with ID `use1-az3` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-brokernodegroupinfo.html#cfn-msk-cluster-brokernodegroupinfo-clientsubnets
	//
	ClientSubnets *[]*string `field:"optional" json:"clientSubnets" yaml:"clientSubnets"`
	// Information about the cluster's connectivity setting.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-brokernodegroupinfo.html#cfn-msk-cluster-brokernodegroupinfo-connectivityinfo
	//
	ConnectivityInfo interface{} `field:"optional" json:"connectivityInfo" yaml:"connectivityInfo"`
	// The type of Amazon EC2 instances to use for brokers.
	//
	// Depending on the [broker type](https://docs.aws.amazon.com/msk/latest/developerguide/broker-instance-types.html) , Amazon MSK supports the following broker sizes:
	//
	// *Standard broker sizes*
	//
	// - kafka.t3.small
	//
	// > You can't select the kafka.t3.small instance type when the metadata mode is KRaft.
	// - kafka.m5.large, kafka.m5.xlarge, kafka.m5.2xlarge, kafka.m5.4xlarge, kafka.m5.8xlarge, kafka.m5.12xlarge, kafka.m5.16xlarge, kafka.m5.24xlarge
	// - kafka.m7g.large, kafka.m7g.xlarge, kafka.m7g.2xlarge, kafka.m7g.4xlarge, kafka.m7g.8xlarge, kafka.m7g.12xlarge, kafka.m7g.16xlarge
	//
	// *Express broker sizes*
	//
	// - express.m7g.large, express.m7g.xlarge, express.m7g.2xlarge, express.m7g.4xlarge, express.m7g.8xlarge, express.m7g.12xlarge, express.m7g.16xlarge
	//
	// > Some broker sizes might not be available in certian AWS Regions. See the updated [Pricing tools](https://docs.aws.amazon.com/msk/pricing/) section on the Amazon MSK pricing page for the latest list of available instances by Region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-brokernodegroupinfo.html#cfn-msk-cluster-brokernodegroupinfo-instancetype
	//
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// The security groups to associate with the ENIs in order to specify who can connect to and communicate with the Amazon MSK cluster.
	//
	// If you don't specify a security group, Amazon MSK uses the default security group associated with the VPC. If you specify security groups that were shared with you, you must ensure that you have permissions to them. Specifically, you need the `ec2:DescribeSecurityGroups` permission.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-brokernodegroupinfo.html#cfn-msk-cluster-brokernodegroupinfo-securitygroups
	//
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Contains information about storage volumes attached to Amazon MSK broker nodes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-brokernodegroupinfo.html#cfn-msk-cluster-brokernodegroupinfo-storageinfo
	//
	StorageInfo interface{} `field:"optional" json:"storageInfo" yaml:"storageInfo"`
}

