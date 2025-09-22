package awsquicksight


// The word cloud options for a word cloud visual.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   wordCloudOptionsProperty := &WordCloudOptionsProperty{
//   	CloudLayout: jsii.String("cloudLayout"),
//   	MaximumStringLength: jsii.Number(123),
//   	WordCasing: jsii.String("wordCasing"),
//   	WordOrientation: jsii.String("wordOrientation"),
//   	WordPadding: jsii.String("wordPadding"),
//   	WordScaling: jsii.String("wordScaling"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-wordcloudoptions.html
//
type CfnTemplate_WordCloudOptionsProperty struct {
	// The cloud layout options (fluid, normal) of a word cloud.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-wordcloudoptions.html#cfn-quicksight-template-wordcloudoptions-cloudlayout
	//
	CloudLayout *string `field:"optional" json:"cloudLayout" yaml:"cloudLayout"`
	// The length limit of each word from 1-100.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-wordcloudoptions.html#cfn-quicksight-template-wordcloudoptions-maximumstringlength
	//
	MaximumStringLength *float64 `field:"optional" json:"maximumStringLength" yaml:"maximumStringLength"`
	// The word casing options (lower_case, existing_case) for the words in a word cloud.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-wordcloudoptions.html#cfn-quicksight-template-wordcloudoptions-wordcasing
	//
	WordCasing *string `field:"optional" json:"wordCasing" yaml:"wordCasing"`
	// The word orientation options (horizontal, horizontal_and_vertical) for the words in a word cloud.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-wordcloudoptions.html#cfn-quicksight-template-wordcloudoptions-wordorientation
	//
	WordOrientation *string `field:"optional" json:"wordOrientation" yaml:"wordOrientation"`
	// The word padding options (none, small, medium, large) for the words in a word cloud.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-wordcloudoptions.html#cfn-quicksight-template-wordcloudoptions-wordpadding
	//
	WordPadding *string `field:"optional" json:"wordPadding" yaml:"wordPadding"`
	// The word scaling options (emphasize, normal) for the words in a word cloud.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-wordcloudoptions.html#cfn-quicksight-template-wordcloudoptions-wordscaling
	//
	WordScaling *string `field:"optional" json:"wordScaling" yaml:"wordScaling"`
}

