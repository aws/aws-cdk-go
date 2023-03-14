package awskafkaconnect


// The client authentication information used in order to authenticate with the Apache Kafka cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kafkaClusterClientAuthenticationProperty := &kafkaClusterClientAuthenticationProperty{
//   	authenticationType: jsii.String("authenticationType"),
//   }
//
type CfnConnector_KafkaClusterClientAuthenticationProperty struct {
	// The type of client authentication used to connect to the Apache Kafka cluster.
	//
	// Value NONE means that no client authentication is used.
	AuthenticationType *string `field:"required" json:"authenticationType" yaml:"authenticationType"`
}

