package awslambda


// The endpoints used by AWS Lambda to access a self-managed event source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   endpointsProperty := &EndpointsProperty{
//   	KafkaBootstrapServers: []*string{
//   		jsii.String("kafkaBootstrapServers"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-endpoints.html
//
type CfnEventSourceMapping_EndpointsProperty struct {
	// A list of Kafka server endpoints.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-endpoints.html#cfn-lambda-eventsourcemapping-endpoints-kafkabootstrapservers
	//
	KafkaBootstrapServers *[]*string `field:"optional" json:"kafkaBootstrapServers" yaml:"kafkaBootstrapServers"`
}

