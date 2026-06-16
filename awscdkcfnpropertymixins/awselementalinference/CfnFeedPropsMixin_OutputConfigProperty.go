package awselementalinference


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var cropping interface{}
//
//   outputConfigProperty := &OutputConfigProperty{
//   	Clipping: &ClippingConfigProperty{
//   		CallbackMetadata: jsii.String("callbackMetadata"),
//   	},
//   	Cropping: cropping,
//   	Subtitling: &SubtitlingConfigProperty{
//   		AspectRatio: &AspectRatioProperty{
//   			Height: jsii.Number(123),
//   			Width: jsii.Number(123),
//   		},
//   		Dictionary: jsii.String("dictionary"),
//   		Language: jsii.String("language"),
//   		ProfanityFilter: jsii.String("profanityFilter"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elementalinference-feed-outputconfig.html
//
type CfnFeedPropsMixin_OutputConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elementalinference-feed-outputconfig.html#cfn-elementalinference-feed-outputconfig-clipping
	//
	Clipping interface{} `field:"optional" json:"clipping" yaml:"clipping"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elementalinference-feed-outputconfig.html#cfn-elementalinference-feed-outputconfig-cropping
	//
	Cropping interface{} `field:"optional" json:"cropping" yaml:"cropping"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elementalinference-feed-outputconfig.html#cfn-elementalinference-feed-outputconfig-subtitling
	//
	Subtitling interface{} `field:"optional" json:"subtitling" yaml:"subtitling"`
}

