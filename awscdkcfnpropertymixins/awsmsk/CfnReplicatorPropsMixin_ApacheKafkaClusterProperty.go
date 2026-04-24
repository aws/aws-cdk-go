package awsmsk


// Details of an Apache Kafka cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   apacheKafkaClusterProperty := &ApacheKafkaClusterProperty{
//   	ApacheKafkaClusterId: jsii.String("apacheKafkaClusterId"),
//   	BootstrapBrokerString: jsii.String("bootstrapBrokerString"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-apachekafkacluster.html
//
type CfnReplicatorPropsMixin_ApacheKafkaClusterProperty struct {
	// The ID of the Apache Kafka cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-apachekafkacluster.html#cfn-msk-replicator-apachekafkacluster-apachekafkaclusterid
	//
	ApacheKafkaClusterId *string `field:"optional" json:"apacheKafkaClusterId" yaml:"apacheKafkaClusterId"`
	// The bootstrap broker string of the Apache Kafka cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-apachekafkacluster.html#cfn-msk-replicator-apachekafkacluster-bootstrapbrokerstring
	//
	BootstrapBrokerString *string `field:"optional" json:"bootstrapBrokerString" yaml:"bootstrapBrokerString"`
}

