package awsmediatailor


// The configuration for an AWS_SERVICE_REQUEST function.
//
// Contains the target service, target Region, and request parameters that the function uses to call an AWS service API. For more information, see AWS_SERVICE_REQUEST (https://docs.aws.amazon.com/mediatailor/latest/ug/monetization-functions-types-aws-service-request.html) in the MediaTailor User Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   awsServiceRequestConfigurationProperty := &AwsServiceRequestConfigurationProperty{
//   	MethodType: jsii.String("methodType"),
//   	RequestTimeoutMilliseconds: jsii.Number(123),
//   	Runtime: jsii.String("runtime"),
//   	TargetRegion: jsii.String("targetRegion"),
//   	TargetService: jsii.String("targetService"),
//   	Url: jsii.String("url"),
//
//   	// the properties below are optional
//   	Body: jsii.String("body"),
//   	Headers: map[string]*string{
//   		"headersKey": jsii.String("headers"),
//   	},
//   	Output: map[string]*string{
//   		"outputKey": jsii.String("output"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-awsservicerequestconfiguration.html
//
type CfnFunction_AwsServiceRequestConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-awsservicerequestconfiguration.html#cfn-mediatailor-function-awsservicerequestconfiguration-methodtype
	//
	MethodType *string `field:"required" json:"methodType" yaml:"methodType"`
	// The maximum time, in milliseconds, that MediaTailor waits for a response from the AWS service.
	//
	// If the call exceeds this timeout, MediaTailor sets the response status code to null and proceeds with output expression evaluation. Valid values are 100 to 2000.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-awsservicerequestconfiguration.html#cfn-mediatailor-function-awsservicerequestconfiguration-requesttimeoutmilliseconds
	//
	RequestTimeoutMilliseconds *float64 `field:"required" json:"requestTimeoutMilliseconds" yaml:"requestTimeoutMilliseconds"`
	// The expression language used to evaluate expressions in the function configuration.
	//
	// Set this to JSONATA.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-awsservicerequestconfiguration.html#cfn-mediatailor-function-awsservicerequestconfiguration-runtime
	//
	Runtime *string `field:"required" json:"runtime" yaml:"runtime"`
	// The AWS Region for the target service.
	//
	// Specify a static Region code (for example, us-east-1) or a JSONata expression that resolves to a Region code at runtime (for example, {%inference.region%}).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-awsservicerequestconfiguration.html#cfn-mediatailor-function-awsservicerequestconfiguration-targetregion
	//
	TargetRegion *string `field:"required" json:"targetRegion" yaml:"targetRegion"`
	// The AWS service to call.
	//
	// Valid value: elemental-inference (AWS Elemental Inference).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-awsservicerequestconfiguration.html#cfn-mediatailor-function-awsservicerequestconfiguration-targetservice
	//
	TargetService *string `field:"required" json:"targetService" yaml:"targetService"`
	// An expression that evaluates to the endpoint URL for the target AWS service API operation.
	//
	// Use {%...%} delimiters for dynamic expressions. The URL must correspond to a valid endpoint for the service specified in TargetService. The maximum length after evaluation is 2,048 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-awsservicerequestconfiguration.html#cfn-mediatailor-function-awsservicerequestconfiguration-url
	//
	Url *string `field:"required" json:"url" yaml:"url"`
	// An expression that evaluates to the request body for the AWS service API call.
	//
	// The body must conform to the input format that the target service operation expects. Applies only when the target operation accepts a request body. The maximum size after evaluation is 64 KB.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-awsservicerequestconfiguration.html#cfn-mediatailor-function-awsservicerequestconfiguration-body
	//
	Body *string `field:"optional" json:"body" yaml:"body"`
	// A map of HTTP header names to expression values.
	//
	// MediaTailor evaluates each header value expression at runtime and includes the result in the outbound request to the AWS service. Use this to pass any headers required by the target service operation. You can include a maximum of 50 headers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-awsservicerequestconfiguration.html#cfn-mediatailor-function-awsservicerequestconfiguration-headers
	//
	Headers interface{} `field:"optional" json:"headers" yaml:"headers"`
	// A map of output bindings.
	//
	// Each key is a namespaced output path, such as player_params.device_type. Each value is an expression that MediaTailor evaluates at runtime and can reference the response object from the target service. For more information, see JSONata expression reference (https://docs.aws.amazon.com/mediatailor/latest/ug/monetization-functions-jsonata.html) in the MediaTailor User Guide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-awsservicerequestconfiguration.html#cfn-mediatailor-function-awsservicerequestconfiguration-output
	//
	Output interface{} `field:"optional" json:"output" yaml:"output"`
}

