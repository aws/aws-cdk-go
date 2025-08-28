package awskafkaconnect


// The client authentication information used in order to authenticate with the Apache Kafka cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kafkaClusterClientAuthenticationProperty := &KafkaClusterClientAuthenticationProperty{
//   	AuthenticationType: jsii.String("authenticationType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-kafkaclusterclientauthentication.html
//
type CfnConnector_KafkaClusterClientAuthenticationProperty struct {
	// The type of client authentication used to connect to the Apache Kafka cluster.
	//
	// Value NONE means that no client authentication is used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-kafkaclusterclientauthentication.html#cfn-kafkaconnect-connector-kafkaclusterclientauthentication-authenticationtype
	//
	AuthenticationType *string `field:"required" json:"authenticationType" yaml:"authenticationType"`
}

