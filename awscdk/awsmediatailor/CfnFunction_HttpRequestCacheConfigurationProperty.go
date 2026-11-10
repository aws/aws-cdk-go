package awsmediatailor


// The optional response-caching configuration shared by the HTTP-based function types (HTTP_REQUEST, AWS_SERVICE_REQUEST, and VAST_REQUEST).
//
// When you provide this configuration, MediaTailor caches function responses with HTTP status codes 200, 203, 204, 404, 405, 410, 414, or 501. For a cacheable response, MediaTailor uses the Cache-Control max-age directive limited to the range between TtlMinimumSeconds and TtlMaximumSeconds. If no max-age directive is present, MediaTailor uses TtlMinimumSeconds. Cached HTTP responses are scoped per playback configuration, not per function.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpRequestCacheConfigurationProperty := &HttpRequestCacheConfigurationProperty{
//   	TtlMaximumSeconds: jsii.Number(123),
//   	TtlMinimumSeconds: jsii.Number(123),
//
//   	// the properties below are optional
//   	Key: jsii.String("key"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-httprequestcacheconfiguration.html
//
type CfnFunction_HttpRequestCacheConfigurationProperty struct {
	// The upper bound, in seconds, on how long MediaTailor caches a response.
	//
	// This value must be greater than or equal to TtlMinimumSeconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-httprequestcacheconfiguration.html#cfn-mediatailor-function-httprequestcacheconfiguration-ttlmaximumseconds
	//
	TtlMaximumSeconds *float64 `field:"required" json:"ttlMaximumSeconds" yaml:"ttlMaximumSeconds"`
	// The lower bound, in seconds, on how long MediaTailor caches a response.
	//
	// MediaTailor also uses this value as the cache duration when a response has no Cache-Control max-age directive.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-httprequestcacheconfiguration.html#cfn-mediatailor-function-httprequestcacheconfiguration-ttlminimumseconds
	//
	TtlMinimumSeconds *float64 `field:"required" json:"ttlMinimumSeconds" yaml:"ttlMinimumSeconds"`
	// A JSONata expression that MediaTailor evaluates to a custom cache key.
	//
	// By default, the cache key is a hash of the HTTP URL, request body, and HTTP method; request headers are not included. The evaluated key must be smaller than 1 KB.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-httprequestcacheconfiguration.html#cfn-mediatailor-function-httprequestcacheconfiguration-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
}

