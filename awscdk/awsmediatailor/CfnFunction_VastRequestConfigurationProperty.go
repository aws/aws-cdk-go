package awsmediatailor


// The configuration for a VAST_REQUEST function.
//
// Specifies the HTTP method, URL, headers, body, timeout, and output expressions for a request to a VAST endpoint. MediaTailor parses the response as VAST and resolves wrapper redirects, then makes the parsed ads available to the function's output expressions. For more information, see Function types and composition (https://docs.aws.amazon.com/mediatailor/latest/ug/monetization-functions-types.html) in the MediaTailor User Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vastRequestConfigurationProperty := &VastRequestConfigurationProperty{
//   	MethodType: jsii.String("methodType"),
//   	RequestTimeoutMilliseconds: jsii.Number(123),
//   	Runtime: jsii.String("runtime"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-vastrequestconfiguration.html
//
type CfnFunction_VastRequestConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-vastrequestconfiguration.html#cfn-mediatailor-function-vastrequestconfiguration-methodtype
	//
	MethodType *string `field:"required" json:"methodType" yaml:"methodType"`
	// The maximum time, in milliseconds, that MediaTailor waits for a response from the VAST endpoint.
	//
	// The timeout covers the entire response, including any wrapper redirects that MediaTailor follows. If the call exceeds this timeout, MediaTailor proceeds with an empty ad list and continues output expression evaluation. Valid values are 100 to 2000.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-vastrequestconfiguration.html#cfn-mediatailor-function-vastrequestconfiguration-requesttimeoutmilliseconds
	//
	RequestTimeoutMilliseconds *float64 `field:"required" json:"requestTimeoutMilliseconds" yaml:"requestTimeoutMilliseconds"`
	// The expression language used to evaluate expressions in the function configuration.
	//
	// Set this to JSONATA.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-vastrequestconfiguration.html#cfn-mediatailor-function-vastrequestconfiguration-runtime
	//
	Runtime *string `field:"required" json:"runtime" yaml:"runtime"`
	// An expression that evaluates to the VAST endpoint URL.
	//
	// Use {%...%} delimiters for dynamic expressions. A literal value must be an https:// URL. The expression can be up to 25,000 characters, and the URL after evaluation can be up to 2,048 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-vastrequestconfiguration.html#cfn-mediatailor-function-vastrequestconfiguration-url
	//
	Url *string `field:"required" json:"url" yaml:"url"`
	// An expression that evaluates to the request body, for example to send an OpenRTB bid request.
	//
	// The expression can be up to 100,000 characters, and the body after evaluation can be up to 64 KB.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-vastrequestconfiguration.html#cfn-mediatailor-function-vastrequestconfiguration-body
	//
	Body *string `field:"optional" json:"body" yaml:"body"`
	// A map of HTTP header names to expression values.
	//
	// MediaTailor evaluates each header value expression at runtime and includes the result in the outbound request. Headers beginning with X-Amz- are reserved by the service, and method override headers are not allowed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-vastrequestconfiguration.html#cfn-mediatailor-function-vastrequestconfiguration-headers
	//
	Headers interface{} `field:"optional" json:"headers" yaml:"headers"`
	// A map of output bindings.
	//
	// Each key is a namespaced output path (such as temp.wrappedAds), and each value is an expression that MediaTailor evaluates at runtime. Output expressions in a VAST_REQUEST function can reference the response object, which exposes response.parsedAds, the ads parsed from the VAST response after schema validation and wrapper resolution, and response.statusCode. For more information about expression syntax, see JSONata expression reference (https://docs.aws.amazon.com/mediatailor/latest/ug/monetization-functions-jsonata.html) in the MediaTailor User Guide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-vastrequestconfiguration.html#cfn-mediatailor-function-vastrequestconfiguration-output
	//
	Output interface{} `field:"optional" json:"output" yaml:"output"`
}

