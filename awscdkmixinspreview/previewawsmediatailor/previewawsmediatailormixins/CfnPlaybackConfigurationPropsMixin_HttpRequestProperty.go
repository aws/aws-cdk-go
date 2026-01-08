package previewawsmediatailormixins


// The configuration for the request to the Ad Decision Server URL.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   httpRequestProperty := &HttpRequestProperty{
//   	Body: jsii.String("body"),
//   	CompressRequest: jsii.String("compressRequest"),
//   	Headers: map[string]*string{
//   		"headersKey": jsii.String("headers"),
//   	},
//   	HttpMethod: jsii.String("httpMethod"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-httprequest.html
//
type CfnPlaybackConfigurationPropsMixin_HttpRequestProperty struct {
	// The body of the request to the Ad Decision Server URL.
	//
	// The maximum length is 100,000 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-httprequest.html#cfn-mediatailor-playbackconfiguration-httprequest-body
	//
	Body *string `field:"optional" json:"body" yaml:"body"`
	// The compression type of the request sent to the Ad Decision Server URL.
	//
	// Only the POST HTTP Method permits compression other than NONE.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-httprequest.html#cfn-mediatailor-playbackconfiguration-httprequest-compressrequest
	//
	CompressRequest *string `field:"optional" json:"compressRequest" yaml:"compressRequest"`
	// The headers in the request sent to the Ad Decision Server URL.
	//
	// The max length is 10,000 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-httprequest.html#cfn-mediatailor-playbackconfiguration-httprequest-headers
	//
	Headers interface{} `field:"optional" json:"headers" yaml:"headers"`
	// Supported HTTP Methods for the request to the Ad Decision Server URL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-httprequest.html#cfn-mediatailor-playbackconfiguration-httprequest-httpmethod
	//
	HttpMethod *string `field:"optional" json:"httpMethod" yaml:"httpMethod"`
}

