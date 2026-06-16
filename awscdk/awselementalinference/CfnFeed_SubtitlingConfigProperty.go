package awselementalinference


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   subtitlingConfigProperty := &SubtitlingConfigProperty{
//   	Language: jsii.String("language"),
//
//   	// the properties below are optional
//   	AspectRatio: &AspectRatioProperty{
//   		Height: jsii.Number(123),
//   		Width: jsii.Number(123),
//   	},
//   	Dictionary: jsii.String("dictionary"),
//   	ProfanityFilter: jsii.String("profanityFilter"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elementalinference-feed-subtitlingconfig.html
//
type CfnFeed_SubtitlingConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elementalinference-feed-subtitlingconfig.html#cfn-elementalinference-feed-subtitlingconfig-language
	//
	Language *string `field:"required" json:"language" yaml:"language"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elementalinference-feed-subtitlingconfig.html#cfn-elementalinference-feed-subtitlingconfig-aspectratio
	//
	AspectRatio interface{} `field:"optional" json:"aspectRatio" yaml:"aspectRatio"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elementalinference-feed-subtitlingconfig.html#cfn-elementalinference-feed-subtitlingconfig-dictionary
	//
	Dictionary *string `field:"optional" json:"dictionary" yaml:"dictionary"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elementalinference-feed-subtitlingconfig.html#cfn-elementalinference-feed-subtitlingconfig-profanityfilter
	//
	ProfanityFilter *string `field:"optional" json:"profanityFilter" yaml:"profanityFilter"`
}

