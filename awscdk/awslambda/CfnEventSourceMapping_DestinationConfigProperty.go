package awslambda


// A configuration object that specifies the destination of an event after Lambda processes it.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   destinationConfigProperty := &DestinationConfigProperty{
//   	OnFailure: &OnFailureProperty{
//   		Destination: jsii.String("destination"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-destinationconfig.html
//
type CfnEventSourceMapping_DestinationConfigProperty struct {
	// The destination configuration for failed invocations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-destinationconfig.html#cfn-lambda-eventsourcemapping-destinationconfig-onfailure
	//
	OnFailure interface{} `field:"optional" json:"onFailure" yaml:"onFailure"`
}

