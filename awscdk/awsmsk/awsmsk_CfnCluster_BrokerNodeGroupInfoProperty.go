package awsmsk


// The setup to be used for brokers in the cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   brokerNodeGroupInfoProperty := &BrokerNodeGroupInfoProperty{
//   	ClientSubnets: []*string{
//   		jsii.String("clientSubnets"),
//   	},
//   	InstanceType: jsii.String("instanceType"),
//
//   	// the properties below are optional
//   	BrokerAzDistribution: jsii.String("brokerAzDistribution"),
//   	ConnectivityInfo: &ConnectivityInfoProperty{
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
type CfnCluster_BrokerNodeGroupInfoProperty struct {
	// The list of subnets to connect to in the client virtual private cloud (VPC).
	//
	// Amazon creates elastic network interfaces inside these subnets. Client applications use elastic network interfaces to produce and consume data.
	//
	// If you use the US West (N. California) Region, specify exactly two subnets. For other Regions where Amazon MSK is available, you can specify either two or three subnets. The subnets that you specify must be in distinct Availability Zones. When you create a cluster, Amazon MSK distributes the broker nodes evenly across the subnets that you specify.
	//
	// Client subnets can't occupy the Availability Zone with ID `use1-az3` .
	ClientSubnets *[]*string `field:"required" json:"clientSubnets" yaml:"clientSubnets"`
	// The type of Amazon EC2 instances to use for brokers.
	//
	// The following instance types are allowed: kafka.m5.large, kafka.m5.xlarge, kafka.m5.2xlarge, kafka.m5.4xlarge, kafka.m5.8xlarge, kafka.m5.12xlarge, kafka.m5.16xlarge, kafka.m5.24xlarge, and kafka.t3.small.
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// This parameter is currently not in use.
	BrokerAzDistribution *string `field:"optional" json:"brokerAzDistribution" yaml:"brokerAzDistribution"`
	// Information about the cluster's connectivity setting.
	ConnectivityInfo interface{} `field:"optional" json:"connectivityInfo" yaml:"connectivityInfo"`
	// The security groups to associate with the elastic network interfaces in order to specify who can connect to and communicate with the Amazon MSK cluster.
	//
	// If you don't specify a security group, Amazon MSK uses the default security group associated with the VPC. If you specify security groups that were shared with you, you must ensure that you have permissions to them. Specifically, you need the `ec2:DescribeSecurityGroups` permission.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Contains information about storage volumes attached to Amazon MSK broker nodes.
	StorageInfo interface{} `field:"optional" json:"storageInfo" yaml:"storageInfo"`
}

