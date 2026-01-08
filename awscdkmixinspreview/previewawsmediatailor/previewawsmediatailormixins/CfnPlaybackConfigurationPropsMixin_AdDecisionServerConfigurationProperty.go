package previewawsmediatailormixins


// The configuration for the request to the specified Ad Decision Server URL.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   adDecisionServerConfigurationProperty := &AdDecisionServerConfigurationProperty{
//   	HttpRequest: &HttpRequestProperty{
//   		Body: jsii.String("body"),
//   		CompressRequest: jsii.String("compressRequest"),
//   		Headers: map[string]*string{
//   			"headersKey": jsii.String("headers"),
//   		},
//   		HttpMethod: jsii.String("httpMethod"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-addecisionserverconfiguration.html
//
type CfnPlaybackConfigurationPropsMixin_AdDecisionServerConfigurationProperty struct {
	// The configuration for the request to the Ad Decision Server URL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-addecisionserverconfiguration.html#cfn-mediatailor-playbackconfiguration-addecisionserverconfiguration-httprequest
	//
	HttpRequest interface{} `field:"optional" json:"httpRequest" yaml:"httpRequest"`
}

