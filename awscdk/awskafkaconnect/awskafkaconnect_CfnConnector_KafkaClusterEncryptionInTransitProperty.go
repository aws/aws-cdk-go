package awskafkaconnect


// Details of encryption in transit to the Apache Kafka cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kafkaClusterEncryptionInTransitProperty := &kafkaClusterEncryptionInTransitProperty{
//   	encryptionType: jsii.String("encryptionType"),
//   }
//
type CfnConnector_KafkaClusterEncryptionInTransitProperty struct {
	// The type of encryption in transit to the Apache Kafka cluster.
	EncryptionType *string `field:"required" json:"encryptionType" yaml:"encryptionType"`
}

