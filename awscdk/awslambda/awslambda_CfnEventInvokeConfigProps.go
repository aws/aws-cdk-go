package awslambda


// Properties for defining a `CfnEventInvokeConfig`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEventInvokeConfigProps := &cfnEventInvokeConfigProps{
//   	functionName: jsii.String("functionName"),
//   	qualifier: jsii.String("qualifier"),
//
//   	// the properties below are optional
//   	destinationConfig: &destinationConfigProperty{
//   		onFailure: &onFailureProperty{
//   			destination: jsii.String("destination"),
//   		},
//   		onSuccess: &onSuccessProperty{
//   			destination: jsii.String("destination"),
//   		},
//   	},
//   	maximumEventAgeInSeconds: jsii.Number(123),
//   	maximumRetryAttempts: jsii.Number(123),
//   }
//
type CfnEventInvokeConfigProps struct {
	// The name of the Lambda function.
	//
	// *Minimum* : `1`
	//
	// *Maximum* : `64`
	//
	// *Pattern* : `([a-zA-Z0-9-_]+)`.
	FunctionName *string `field:"required" json:"functionName" yaml:"functionName"`
	// The identifier of a version or alias.
	//
	// - *Version* - A version number.
	// - *Alias* - An alias name.
	// - *Latest* - To specify the unpublished version, use `$LATEST` .
	Qualifier *string `field:"required" json:"qualifier" yaml:"qualifier"`
	// A destination for events after they have been sent to a function for processing.
	//
	// **Destinations** - *Function* - The Amazon Resource Name (ARN) of a Lambda function.
	// - *Queue* - The ARN of an SQS queue.
	// - *Topic* - The ARN of an SNS topic.
	// - *Event Bus* - The ARN of an Amazon EventBridge event bus.
	DestinationConfig interface{} `field:"optional" json:"destinationConfig" yaml:"destinationConfig"`
	// The maximum age of a request that Lambda sends to a function for processing.
	MaximumEventAgeInSeconds *float64 `field:"optional" json:"maximumEventAgeInSeconds" yaml:"maximumEventAgeInSeconds"`
	// The maximum number of times to retry when the function returns an error.
	MaximumRetryAttempts *float64 `field:"optional" json:"maximumRetryAttempts" yaml:"maximumRetryAttempts"`
}

