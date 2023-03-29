package awskafkaconnect


// The details of the Apache Kafka cluster to which the connector is connected.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kafkaClusterProperty := &KafkaClusterProperty{
//   	ApacheKafkaCluster: &ApacheKafkaClusterProperty{
//   		BootstrapServers: jsii.String("bootstrapServers"),
//   		Vpc: &VpcProperty{
//   			SecurityGroups: []*string{
//   				jsii.String("securityGroups"),
//   			},
//   			Subnets: []*string{
//   				jsii.String("subnets"),
//   			},
//   		},
//   	},
//   }
//
type CfnConnector_KafkaClusterProperty struct {
	// The Apache Kafka cluster to which the connector is connected.
	ApacheKafkaCluster interface{} `field:"required" json:"apacheKafkaCluster" yaml:"apacheKafkaCluster"`
}

