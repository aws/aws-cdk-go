package awsmsk


// The setup to be used for brokers in the cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   brokerNodeGroupInfoProperty := &brokerNodeGroupInfoProperty{
//   	clientSubnets: []*string{
//   		jsii.String("clientSubnets"),
//   	},
//   	instanceType: jsii.String("instanceType"),
//
//   	// the properties below are optional
//   	brokerAzDistribution: jsii.String("brokerAzDistribution"),
//   	connectivityInfo: &connectivityInfoProperty{
//   		publicAccess: &publicAccessProperty{
//   			type: jsii.String("type"),
//   		},
//   		vpcConnectivity: &vpcConnectivityProperty{
//   			clientAuthentication: &vpcConnectivityClientAuthenticationProperty{
//   				sasl: &vpcConnectivitySaslProperty{
//   					iam: &vpcConnectivityIamProperty{
//   						enabled: jsii.Boolean(false),
//   					},
//   					scram: &vpcConnectivityScramProperty{
//   						enabled: jsii.Boolean(false),
//   					},
//   				},
//   				tls: &vpcConnectivityTlsProperty{
//   					enabled: jsii.Boolean(false),
//   				},
//   			},
//   		},
//   	},
//   	securityGroups: []*string{
//   		jsii.String("securityGroups"),
//   	},
//   	storageInfo: &storageInfoProperty{
//   		ebsStorageInfo: &eBSStorageInfoProperty{
//   			provisionedThroughput: &provisionedThroughputProperty{
//   				enabled: jsii.Boolean(false),
//   				volumeThroughput: jsii.Number(123),
//   			},
//   			volumeSize: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnCluster_BrokerNodeGroupInfoProperty struct {
	// The list of subnets to connect to in the client virtual private cloud (VPC).
	//
	// Amazon creates elastic network interfaces inside these subnets. Client applications use elastic network interfaces to produce and consume data.
	//
	// Specify exactly two subnets if you are using the US West (N. California) Region. For other Regions where Amazon MSK is available, you can specify either two or three subnets. The subnets that you specify must be in distinct Availability Zones. When you create a cluster, Amazon MSK distributes the broker nodes evenly across the subnets that you specify.
	//
	// Client subnets can't occupy the Availability Zone with ID `use1-az3` .
	ClientSubnets *[]*string `field:"required" json:"clientSubnets" yaml:"clientSubnets"`
	// The type of Amazon EC2 instances to use for brokers.
	//
	// The following instance types are allowed: kafka.m5.large, kafka.m5.xlarge, kafka.m5.2xlarge, kafka.m5.4xlarge, kafka.m5.8xlarge, kafka.m5.12xlarge, kafka.m5.16xlarge, and kafka.m5.24xlarge.
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// This parameter is currently not in use.
	BrokerAzDistribution *string `field:"optional" json:"brokerAzDistribution" yaml:"brokerAzDistribution"`
	// Information about the cluster's connectivity setting.
	ConnectivityInfo interface{} `field:"optional" json:"connectivityInfo" yaml:"connectivityInfo"`
	// The security groups to associate with the elastic network interfaces in order to specify who can connect to and communicate with the Amazon MSK cluster.
	//
	// If you don't specify a security group, Amazon MSK uses the default security group associated with the VPC. If you specify security groups that were shared with you, you must ensure that you have permissions to them. Specifically, you need the `ec2:DescribeSecurityGroups` permission.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Contains information about storage volumes attached to MSK broker nodes.
	StorageInfo interface{} `field:"optional" json:"storageInfo" yaml:"storageInfo"`
}

