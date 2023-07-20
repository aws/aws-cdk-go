package awslambda


// A destination for events that were processed successfully.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   onSuccessProperty := &OnSuccessProperty{
//   	Destination: jsii.String("destination"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventinvokeconfig-onsuccess.html
//
type CfnEventInvokeConfig_OnSuccessProperty struct {
	// The Amazon Resource Name (ARN) of the destination resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventinvokeconfig-onsuccess.html#cfn-lambda-eventinvokeconfig-onsuccess-destination
	//
	Destination *string `field:"required" json:"destination" yaml:"destination"`
}

