package mixinsawslambda


// A configuration object that specifies the destination of an event after Lambda processes it.
//
// For more information, see [Adding a destination](https://docs.aws.amazon.com/lambda/latest/dg/invocation-async-retain-records.html#invocation-async-destinations) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   destinationConfigProperty := &DestinationConfigProperty{
//   	OnFailure: &OnFailureProperty{
//   		Destination: jsii.String("destination"),
//   	},
//   	OnSuccess: &OnSuccessProperty{
//   		Destination: jsii.String("destination"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventinvokeconfig-destinationconfig.html
//
type CfnEventInvokeConfigPropsMixin_DestinationConfigProperty struct {
	// The destination configuration for failed invocations.
	//
	// > When using an Amazon SQS queue as a destination, FIFO queues cannot be used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventinvokeconfig-destinationconfig.html#cfn-lambda-eventinvokeconfig-destinationconfig-onfailure
	//
	OnFailure interface{} `field:"optional" json:"onFailure" yaml:"onFailure"`
	// The destination configuration for successful invocations.
	//
	// > When using an Amazon SQS queue as a destination, FIFO queues cannot be used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventinvokeconfig-destinationconfig.html#cfn-lambda-eventinvokeconfig-destinationconfig-onsuccess
	//
	OnSuccess interface{} `field:"optional" json:"onSuccess" yaml:"onSuccess"`
}

