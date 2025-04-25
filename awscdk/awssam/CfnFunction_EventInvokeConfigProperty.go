package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventInvokeConfigProperty := &EventInvokeConfigProperty{
//   	DestinationConfig: &EventInvokeDestinationConfigProperty{
//   		OnFailure: &DestinationProperty{
//   			Destination: jsii.String("destination"),
//
//   			// the properties below are optional
//   			Type: jsii.String("type"),
//   		},
//   		OnSuccess: &DestinationProperty{
//   			Destination: jsii.String("destination"),
//
//   			// the properties below are optional
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	MaximumEventAgeInSeconds: jsii.Number(123),
//   	MaximumRetryAttempts: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-eventinvokeconfig.html
//
type CfnFunction_EventInvokeConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-eventinvokeconfig.html#cfn-serverless-function-eventinvokeconfig-destinationconfig
	//
	DestinationConfig interface{} `field:"optional" json:"destinationConfig" yaml:"destinationConfig"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-eventinvokeconfig.html#cfn-serverless-function-eventinvokeconfig-maximumeventageinseconds
	//
	MaximumEventAgeInSeconds *float64 `field:"optional" json:"maximumEventAgeInSeconds" yaml:"maximumEventAgeInSeconds"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-eventinvokeconfig.html#cfn-serverless-function-eventinvokeconfig-maximumretryattempts
	//
	MaximumRetryAttempts *float64 `field:"optional" json:"maximumRetryAttempts" yaml:"maximumRetryAttempts"`
}

