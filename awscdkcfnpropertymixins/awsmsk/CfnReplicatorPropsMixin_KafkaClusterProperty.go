package awsmsk


// Information about Kafka Cluster to be used as source / target for replication.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   kafkaClusterProperty := &KafkaClusterProperty{
//   	AmazonMskCluster: &AmazonMskClusterProperty{
//   		MskClusterArn: jsii.String("mskClusterArn"),
//   	},
//   	ApacheKafkaCluster: &ApacheKafkaClusterProperty{
//   		ApacheKafkaClusterId: jsii.String("apacheKafkaClusterId"),
//   		BootstrapBrokerString: jsii.String("bootstrapBrokerString"),
//   	},
//   	ClientAuthentication: &KafkaClusterClientAuthenticationProperty{
//   		Mtls: &KafkaClusterMtlsAuthenticationProperty{
//   			SecretArn: jsii.String("secretArn"),
//   		},
//   		SaslScram: &KafkaClusterSaslScramAuthenticationProperty{
//   			Mechanism: jsii.String("mechanism"),
//   			SecretArn: jsii.String("secretArn"),
//   		},
//   	},
//   	EncryptionInTransit: &KafkaClusterEncryptionInTransitProperty{
//   		EncryptionType: jsii.String("encryptionType"),
//   		RootCaCertificate: jsii.String("rootCaCertificate"),
//   	},
//   	VpcConfig: &KafkaClusterClientVpcConfigProperty{
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-kafkacluster.html
//
type CfnReplicatorPropsMixin_KafkaClusterProperty struct {
	// Details of an Amazon MSK Cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-kafkacluster.html#cfn-msk-replicator-kafkacluster-amazonmskcluster
	//
	AmazonMskCluster interface{} `field:"optional" json:"amazonMskCluster" yaml:"amazonMskCluster"`
	// Details of an Apache Kafka cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-kafkacluster.html#cfn-msk-replicator-kafkacluster-apachekafkacluster
	//
	ApacheKafkaCluster interface{} `field:"optional" json:"apacheKafkaCluster" yaml:"apacheKafkaCluster"`
	// Details of the client authentication used by the Apache Kafka cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-kafkacluster.html#cfn-msk-replicator-kafkacluster-clientauthentication
	//
	ClientAuthentication interface{} `field:"optional" json:"clientAuthentication" yaml:"clientAuthentication"`
	// Details of encryption in transit to the Apache Kafka cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-kafkacluster.html#cfn-msk-replicator-kafkacluster-encryptionintransit
	//
	EncryptionInTransit interface{} `field:"optional" json:"encryptionInTransit" yaml:"encryptionInTransit"`
	// Details of an Amazon VPC which has network connectivity to the Apache Kafka cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-kafkacluster.html#cfn-msk-replicator-kafkacluster-vpcconfig
	//
	VpcConfig interface{} `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

