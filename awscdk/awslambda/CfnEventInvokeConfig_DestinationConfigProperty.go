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
//   	OnSuccess: &OnSuccessProperty{
//   		Destination: jsii.String("destination"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventinvokeconfig-destinationconfig.html
//
type CfnEventInvokeConfig_DestinationConfigProperty struct {
	// The destination configuration for failed invocations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventinvokeconfig-destinationconfig.html#cfn-lambda-eventinvokeconfig-destinationconfig-onfailure
	//
	OnFailure interface{} `field:"optional" json:"onFailure" yaml:"onFailure"`
	// The destination configuration for successful invocations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventinvokeconfig-destinationconfig.html#cfn-lambda-eventinvokeconfig-destinationconfig-onsuccess
	//
	OnSuccess interface{} `field:"optional" json:"onSuccess" yaml:"onSuccess"`
}

