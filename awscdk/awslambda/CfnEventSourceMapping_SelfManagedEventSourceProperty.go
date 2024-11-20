package awslambda


// The configuration used by AWS Lambda to access a self-managed event source.
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
	// The endpoints used by AWS Lambda to access a self-managed event source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-selfmanagedeventsource.html#cfn-lambda-eventsourcemapping-selfmanagedeventsource-endpoints
	//
	Endpoints interface{} `field:"optional" json:"endpoints" yaml:"endpoints"`
}

