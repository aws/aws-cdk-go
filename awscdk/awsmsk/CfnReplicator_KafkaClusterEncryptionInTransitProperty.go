package awsmsk


// Details of encryption in transit to the Apache Kafka cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kafkaClusterEncryptionInTransitProperty := &KafkaClusterEncryptionInTransitProperty{
//   	EncryptionType: jsii.String("encryptionType"),
//
//   	// the properties below are optional
//   	RootCaCertificate: jsii.String("rootCaCertificate"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-kafkaclusterencryptionintransit.html
//
type CfnReplicator_KafkaClusterEncryptionInTransitProperty struct {
	// The type of encryption in transit to the Apache Kafka cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-kafkaclusterencryptionintransit.html#cfn-msk-replicator-kafkaclusterencryptionintransit-encryptiontype
	//
	EncryptionType *string `field:"required" json:"encryptionType" yaml:"encryptionType"`
	// The root CA certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-kafkaclusterencryptionintransit.html#cfn-msk-replicator-kafkaclusterencryptionintransit-rootcacertificate
	//
	RootCaCertificate *string `field:"optional" json:"rootCaCertificate" yaml:"rootCaCertificate"`
}

