package awslambda


// The list of bootstrap servers for your Kafka brokers in the following format: `"KafkaBootstrapServers": ["abc.xyz.com:xxxx","abc2.xyz.com:xxxx"]` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   endpointsProperty := &endpointsProperty{
//   	kafkaBootstrapServers: []*string{
//   		jsii.String("kafkaBootstrapServers"),
//   	},
//   }
//
type CfnEventSourceMapping_EndpointsProperty struct {
	// The list of bootstrap servers for your Kafka brokers in the following format: `"KafkaBootstrapServers": ["abc.xyz.com:xxxx","abc2.xyz.com:xxxx"]` .
	KafkaBootstrapServers *[]*string `field:"optional" json:"kafkaBootstrapServers" yaml:"kafkaBootstrapServers"`
}

