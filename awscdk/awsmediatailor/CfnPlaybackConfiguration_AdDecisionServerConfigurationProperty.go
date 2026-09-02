package awsmediatailor


// The configuration for the request to the specified Ad Decision Server URL.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
//
//   	// the properties below are optional
//   	VastResponse: &VastResponseProperty{
//   		AdSequencingMode: jsii.String("adSequencingMode"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-addecisionserverconfiguration.html
//
type CfnPlaybackConfiguration_AdDecisionServerConfigurationProperty struct {
	// The configuration for the request to the Ad Decision Server URL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-addecisionserverconfiguration.html#cfn-mediatailor-playbackconfiguration-addecisionserverconfiguration-httprequest
	//
	HttpRequest interface{} `field:"required" json:"httpRequest" yaml:"httpRequest"`
	// The configuration for how MediaTailor processes the VAST response returned by the Ad Decision Server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-addecisionserverconfiguration.html#cfn-mediatailor-playbackconfiguration-addecisionserverconfiguration-vastresponse
	//
	VastResponse interface{} `field:"optional" json:"vastResponse" yaml:"vastResponse"`
}

