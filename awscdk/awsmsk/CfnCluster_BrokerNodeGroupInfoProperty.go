package awsmsk


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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-brokernodegroupinfo.html
//
type CfnCluster_BrokerNodeGroupInfoProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-brokernodegroupinfo.html#cfn-msk-cluster-brokernodegroupinfo-clientsubnets
	//
	ClientSubnets *[]*string `field:"required" json:"clientSubnets" yaml:"clientSubnets"`
	// The type of Amazon EC2 instances to use for brokers.
	//
	// The following instance types are allowed: kafka.m5.large, kafka.m5.xlarge, kafka.m5.2xlarge, kafka.m5.4xlarge, kafka.m5.8xlarge, kafka.m5.12xlarge, kafka.m5.16xlarge, kafka.m5.24xlarge, and kafka.t3.small.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-brokernodegroupinfo.html#cfn-msk-cluster-brokernodegroupinfo-instancetype
	//
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-brokernodegroupinfo.html#cfn-msk-cluster-brokernodegroupinfo-brokerazdistribution
	//
	BrokerAzDistribution *string `field:"optional" json:"brokerAzDistribution" yaml:"brokerAzDistribution"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-brokernodegroupinfo.html#cfn-msk-cluster-brokernodegroupinfo-connectivityinfo
	//
	ConnectivityInfo interface{} `field:"optional" json:"connectivityInfo" yaml:"connectivityInfo"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-brokernodegroupinfo.html#cfn-msk-cluster-brokernodegroupinfo-securitygroups
	//
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-brokernodegroupinfo.html#cfn-msk-cluster-brokernodegroupinfo-storageinfo
	//
	StorageInfo interface{} `field:"optional" json:"storageInfo" yaml:"storageInfo"`
}

