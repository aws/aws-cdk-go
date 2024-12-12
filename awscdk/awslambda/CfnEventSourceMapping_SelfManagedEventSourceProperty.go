package awslambda


// The self-managed Apache Kafka cluster for your event source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   selfManagedEventSourceProperty := &SelfManagedEventSourceProperty{
//   	Endpoints: &EndpointsProperty{
//   		KafkaBootstrapServers: []*string{
//   			jsii.String("kafkaBootstrapServers"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-selfmanagedeventsource.html
//
type CfnEventSourceMapping_SelfManagedEventSourceProperty struct {
	// The list of bootstrap servers for your Kafka brokers in the following format: `"KafkaBootstrapServers": ["abc.xyz.com:xxxx","abc2.xyz.com:xxxx"]` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-selfmanagedeventsource.html#cfn-lambda-eventsourcemapping-selfmanagedeventsource-endpoints
	//
	Endpoints interface{} `field:"optional" json:"endpoints" yaml:"endpoints"`
}

