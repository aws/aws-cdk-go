package awslambda


// A destination for events that failed processing.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   onFailureProperty := &OnFailureProperty{
//   	Destination: jsii.String("destination"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventinvokeconfig-onfailure.html
//
type CfnEventInvokeConfig_OnFailureProperty struct {
	// The Amazon Resource Name (ARN) of the destination resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventinvokeconfig-onfailure.html#cfn-lambda-eventinvokeconfig-onfailure-destination
	//
	Destination *string `field:"required" json:"destination" yaml:"destination"`
}

