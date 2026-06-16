package awselementalinference


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cropping interface{}
//
//   getOutputProperty := &GetOutputProperty{
//   	Name: jsii.String("name"),
//   	OutputConfig: &OutputConfigProperty{
//   		Clipping: &ClippingConfigProperty{
//   			CallbackMetadata: jsii.String("callbackMetadata"),
//   		},
//   		Cropping: cropping,
//   		Subtitling: &SubtitlingConfigProperty{
//   			Language: jsii.String("language"),
//
//   			// the properties below are optional
//   			AspectRatio: &AspectRatioProperty{
//   				Height: jsii.Number(123),
//   				Width: jsii.Number(123),
//   			},
//   			Dictionary: jsii.String("dictionary"),
//   			ProfanityFilter: jsii.String("profanityFilter"),
//   		},
//   	},
//   	Status: jsii.String("status"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elementalinference-feed-getoutput.html
//
type CfnFeed_GetOutputProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elementalinference-feed-getoutput.html#cfn-elementalinference-feed-getoutput-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elementalinference-feed-getoutput.html#cfn-elementalinference-feed-getoutput-outputconfig
	//
	OutputConfig interface{} `field:"required" json:"outputConfig" yaml:"outputConfig"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elementalinference-feed-getoutput.html#cfn-elementalinference-feed-getoutput-status
	//
	Status *string `field:"required" json:"status" yaml:"status"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elementalinference-feed-getoutput.html#cfn-elementalinference-feed-getoutput-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

