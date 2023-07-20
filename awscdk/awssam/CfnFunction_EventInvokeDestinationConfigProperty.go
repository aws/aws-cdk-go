package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventInvokeDestinationConfigProperty := &EventInvokeDestinationConfigProperty{
//   	OnFailure: &DestinationProperty{
//   		Destination: jsii.String("destination"),
//
//   		// the properties below are optional
//   		Type: jsii.String("type"),
//   	},
//   	OnSuccess: &DestinationProperty{
//   		Destination: jsii.String("destination"),
//
//   		// the properties below are optional
//   		Type: jsii.String("type"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-eventinvokedestinationconfig.html
//
type CfnFunction_EventInvokeDestinationConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-eventinvokedestinationconfig.html#cfn-serverless-function-eventinvokedestinationconfig-onfailure
	//
	OnFailure interface{} `field:"required" json:"onFailure" yaml:"onFailure"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-eventinvokedestinationconfig.html#cfn-serverless-function-eventinvokedestinationconfig-onsuccess
	//
	OnSuccess interface{} `field:"required" json:"onSuccess" yaml:"onSuccess"`
}

