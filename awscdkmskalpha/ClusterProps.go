// The CDK Construct Library for AWS::MSK
package awscdkmskalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Properties for a MSK Cluster.
//
// Example:
//   var vpc vpc
//
//   cluster := msk.NewCluster(this, jsii.String("cluster"), &ClusterProps{
//   	ClusterName: jsii.String("myCluster"),
//   	KafkaVersion: msk.KafkaVersion_V2_8_1(),
//   	Vpc: Vpc,
//   	EncryptionInTransit: &EncryptionInTransitConfig{
//   		ClientBroker: msk.ClientBrokerEncryption_TLS,
//   	},
//   	ClientAuthentication: msk.ClientAuthentication_Sasl(&SaslAuthProps{
//   		Scram: jsii.Boolean(true),
//   	}),
//   })
//
// Experimental.
type ClusterProps struct {
	// The physical name of the cluster.
	// Experimental.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// The version of Apache Kafka.
	// Experimental.
	KafkaVersion KafkaVersion `field:"required" json:"kafkaVersion" yaml:"kafkaVersion"`
	// Defines the virtual networking environment for this cluster.
	//
	// Must have at least 2 subnets in two different AZs.
	// Experimental.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// Configuration properties for client authentication.
	//
	// MSK supports using private TLS certificates or SASL/SCRAM to authenticate the identity of clients.
	// Experimental.
	ClientAuthentication ClientAuthentication `field:"optional" json:"clientAuthentication" yaml:"clientAuthentication"`
	// The Amazon MSK configuration to use for the cluster.
	// Experimental.
	ConfigurationInfo *ClusterConfigurationInfo `field:"optional" json:"configurationInfo" yaml:"configurationInfo"`
	// Information about storage volumes attached to MSK broker nodes.
	// Experimental.
	EbsStorageInfo *EbsStorageInfo `field:"optional" json:"ebsStorageInfo" yaml:"ebsStorageInfo"`
	// Config details for encryption in transit.
	// Experimental.
	EncryptionInTransit *EncryptionInTransitConfig `field:"optional" json:"encryptionInTransit" yaml:"encryptionInTransit"`
	// The EC2 instance type that you want Amazon MSK to use when it creates your brokers.
	// See: https://docs.aws.amazon.com/msk/latest/developerguide/msk-create-cluster.html#broker-instance-types
	//
	// Experimental.
	InstanceType awsec2.InstanceType `field:"optional" json:"instanceType" yaml:"instanceType"`
	// Configure your MSK cluster to send broker logs to different destination types.
	// Experimental.
	Logging *BrokerLogging `field:"optional" json:"logging" yaml:"logging"`
	// Cluster monitoring configuration.
	// Experimental.
	Monitoring *MonitoringConfiguration `field:"optional" json:"monitoring" yaml:"monitoring"`
	// Number of Apache Kafka brokers deployed in each Availability Zone.
	// Experimental.
	NumberOfBrokerNodes *float64 `field:"optional" json:"numberOfBrokerNodes" yaml:"numberOfBrokerNodes"`
	// What to do when this resource is deleted from a stack.
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// The AWS security groups to associate with the elastic network interfaces in order to specify who can connect to and communicate with the Amazon MSK cluster.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Where to place the nodes within the VPC.
	//
	// Amazon MSK distributes the broker nodes evenly across the subnets that you specify.
	// The subnets that you specify must be in distinct Availability Zones.
	// Client subnets can't be in Availability Zone us-east-1e.
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

