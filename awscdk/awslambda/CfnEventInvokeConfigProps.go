package awslambda


// Properties for defining a `CfnEventInvokeConfig`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEventInvokeConfigProps := &CfnEventInvokeConfigProps{
//   	FunctionName: jsii.String("functionName"),
//   	Qualifier: jsii.String("qualifier"),
//
//   	// the properties below are optional
//   	DestinationConfig: &DestinationConfigProperty{
//   		OnFailure: &OnFailureProperty{
//   			Destination: jsii.String("destination"),
//   		},
//   		OnSuccess: &OnSuccessProperty{
//   			Destination: jsii.String("destination"),
//   		},
//   	},
//   	MaximumEventAgeInSeconds: jsii.Number(123),
//   	MaximumRetryAttempts: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventinvokeconfig.html
//
type CfnEventInvokeConfigProps struct {
	// The name of the Lambda function.
	//
	// *Minimum* : `1`
	//
	// *Maximum* : `64`
	//
	// *Pattern* : `([a-zA-Z0-9-_]+)`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventinvokeconfig.html#cfn-lambda-eventinvokeconfig-functionname
	//
	FunctionName *string `field:"required" json:"functionName" yaml:"functionName"`
	// The identifier of a version or alias.
	//
	// - *Version* - A version number.
	// - *Alias* - An alias name.
	// - *Latest* - To specify the unpublished version, use `$LATEST` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventinvokeconfig.html#cfn-lambda-eventinvokeconfig-qualifier
	//
	Qualifier *string `field:"required" json:"qualifier" yaml:"qualifier"`
	// A destination for events after they have been sent to a function for processing.
	//
	// **Destinations** - *Function* - The Amazon Resource Name (ARN) of a Lambda function.
	// - *Queue* - The ARN of a standard SQS queue.
	// - *Topic* - The ARN of a standard SNS topic.
	// - *Event Bus* - The ARN of an Amazon EventBridge event bus.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventinvokeconfig.html#cfn-lambda-eventinvokeconfig-destinationconfig
	//
	DestinationConfig interface{} `field:"optional" json:"destinationConfig" yaml:"destinationConfig"`
	// The maximum age of a request that Lambda sends to a function for processing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventinvokeconfig.html#cfn-lambda-eventinvokeconfig-maximumeventageinseconds
	//
	MaximumEventAgeInSeconds *float64 `field:"optional" json:"maximumEventAgeInSeconds" yaml:"maximumEventAgeInSeconds"`
	// The maximum number of times to retry when the function returns an error.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventinvokeconfig.html#cfn-lambda-eventinvokeconfig-maximumretryattempts
	//
	MaximumRetryAttempts *float64 `field:"optional" json:"maximumRetryAttempts" yaml:"maximumRetryAttempts"`
}

