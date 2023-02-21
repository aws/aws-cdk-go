package awskafkaconnect


// The details of the Apache Kafka cluster to which the connector is connected.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   apacheKafkaClusterProperty := &ApacheKafkaClusterProperty{
//   	BootstrapServers: jsii.String("bootstrapServers"),
//   	Vpc: &VpcProperty{
//   		SecurityGroups: []*string{
//   			jsii.String("securityGroups"),
//   		},
//   		Subnets: []*string{
//   			jsii.String("subnets"),
//   		},
//   	},
//   }
//
type CfnConnector_ApacheKafkaClusterProperty struct {
	// The bootstrap servers of the cluster.
	BootstrapServers *string `field:"required" json:"bootstrapServers" yaml:"bootstrapServers"`
	// Details of an Amazon VPC which has network connectivity to the Apache Kafka cluster.
	Vpc interface{} `field:"required" json:"vpc" yaml:"vpc"`
}

