package awskafkaconnect


// Details of encryption in transit to the Apache Kafka cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kafkaClusterEncryptionInTransitProperty := &KafkaClusterEncryptionInTransitProperty{
//   	EncryptionType: jsii.String("encryptionType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-kafkaclusterencryptionintransit.html
//
type CfnConnector_KafkaClusterEncryptionInTransitProperty struct {
	// The type of encryption in transit to the Apache Kafka cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-kafkaclusterencryptionintransit.html#cfn-kafkaconnect-connector-kafkaclusterencryptionintransit-encryptiontype
	//
	EncryptionType *string `field:"required" json:"encryptionType" yaml:"encryptionType"`
}

