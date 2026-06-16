package awselementalinference


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var cropping interface{}
//
//   getOutputProperty := &GetOutputProperty{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	OutputConfig: &OutputConfigProperty{
//   		Clipping: &ClippingConfigProperty{
//   			CallbackMetadata: jsii.String("callbackMetadata"),
//   		},
//   		Cropping: cropping,
//   		Subtitling: &SubtitlingConfigProperty{
//   			AspectRatio: &AspectRatioProperty{
//   				Height: jsii.Number(123),
//   				Width: jsii.Number(123),
//   			},
//   			Dictionary: jsii.String("dictionary"),
//   			Language: jsii.String("language"),
//   			ProfanityFilter: jsii.String("profanityFilter"),
//   		},
//   	},
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elementalinference-feed-getoutput.html
//
type CfnFeedPropsMixin_GetOutputProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elementalinference-feed-getoutput.html#cfn-elementalinference-feed-getoutput-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elementalinference-feed-getoutput.html#cfn-elementalinference-feed-getoutput-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elementalinference-feed-getoutput.html#cfn-elementalinference-feed-getoutput-outputconfig
	//
	OutputConfig interface{} `field:"optional" json:"outputConfig" yaml:"outputConfig"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elementalinference-feed-getoutput.html#cfn-elementalinference-feed-getoutput-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

