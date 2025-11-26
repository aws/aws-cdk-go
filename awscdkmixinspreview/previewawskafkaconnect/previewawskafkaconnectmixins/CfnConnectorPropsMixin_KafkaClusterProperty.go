package previewawskafkaconnectmixins


// The details of the Apache Kafka cluster to which the connector is connected.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-kafkacluster.html
//
type CfnConnectorPropsMixin_KafkaClusterProperty struct {
	// The Apache Kafka cluster to which the connector is connected.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-kafkacluster.html#cfn-kafkaconnect-connector-kafkacluster-apachekafkacluster
	//
	ApacheKafkaCluster interface{} `field:"optional" json:"apacheKafkaCluster" yaml:"apacheKafkaCluster"`
}

