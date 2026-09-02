package awsmediatailor


// Configuration for HTTP request functions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpRequestConfigurationProperty := &HttpRequestConfigurationProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-httprequestconfiguration.html
//
type CfnFunction_HttpRequestConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-httprequestconfiguration.html#cfn-mediatailor-function-httprequestconfiguration-methodtype
	//
	MethodType *string `field:"required" json:"methodType" yaml:"methodType"`
	// The timeout in milliseconds for the HTTP request.
	//
	// Maximum value is 2000.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-httprequestconfiguration.html#cfn-mediatailor-function-httprequestconfiguration-requesttimeoutmilliseconds
	//
	RequestTimeoutMilliseconds *float64 `field:"required" json:"requestTimeoutMilliseconds" yaml:"requestTimeoutMilliseconds"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-httprequestconfiguration.html#cfn-mediatailor-function-httprequestconfiguration-runtime
	//
	Runtime *string `field:"required" json:"runtime" yaml:"runtime"`
	// The URL endpoint for the HTTP request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-httprequestconfiguration.html#cfn-mediatailor-function-httprequestconfiguration-url
	//
	Url *string `field:"required" json:"url" yaml:"url"`
	// The body of the HTTP request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-httprequestconfiguration.html#cfn-mediatailor-function-httprequestconfiguration-body
	//
	Body *string `field:"optional" json:"body" yaml:"body"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-httprequestconfiguration.html#cfn-mediatailor-function-httprequestconfiguration-headers
	//
	Headers interface{} `field:"optional" json:"headers" yaml:"headers"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-httprequestconfiguration.html#cfn-mediatailor-function-httprequestconfiguration-output
	//
	Output interface{} `field:"optional" json:"output" yaml:"output"`
}

