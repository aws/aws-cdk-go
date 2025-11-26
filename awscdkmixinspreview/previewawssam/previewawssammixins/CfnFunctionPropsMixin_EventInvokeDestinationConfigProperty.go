package previewawssammixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eventInvokeDestinationConfigProperty := &EventInvokeDestinationConfigProperty{
//   	OnFailure: &DestinationProperty{
//   		Destination: jsii.String("destination"),
//   		Type: jsii.String("type"),
//   	},
//   	OnSuccess: &DestinationProperty{
//   		Destination: jsii.String("destination"),
//   		Type: jsii.String("type"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-eventinvokedestinationconfig.html
//
type CfnFunctionPropsMixin_EventInvokeDestinationConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-eventinvokedestinationconfig.html#cfn-serverless-function-eventinvokedestinationconfig-onfailure
	//
	OnFailure interface{} `field:"optional" json:"onFailure" yaml:"onFailure"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-eventinvokedestinationconfig.html#cfn-serverless-function-eventinvokedestinationconfig-onsuccess
	//
	OnSuccess interface{} `field:"optional" json:"onSuccess" yaml:"onSuccess"`
}

