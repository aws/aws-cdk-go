package awslex


// A card that is shown to the user by a messaging platform.
//
// You define the contents of the card, the card is displayed by the platform.
//
// When you use a response card, the response from the user is constrained to the text associated with a button on the card.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   imageResponseCardProperty := &ImageResponseCardProperty{
//   	Title: jsii.String("title"),
//
//   	// the properties below are optional
//   	Buttons: []interface{}{
//   		&ButtonProperty{
//   			Text: jsii.String("text"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ImageUrl: jsii.String("imageUrl"),
//   	Subtitle: jsii.String("subtitle"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-imageresponsecard.html
//
type CfnBot_ImageResponseCardProperty struct {
	// The title to display on the response card.
	//
	// The format of the title is determined by the platform displaying the response card.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-imageresponsecard.html#cfn-lex-bot-imageresponsecard-title
	//
	Title *string `field:"required" json:"title" yaml:"title"`
	// A list of buttons that should be displayed on the response card.
	//
	// The arrangement of the buttons is determined by the platform that displays the button.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-imageresponsecard.html#cfn-lex-bot-imageresponsecard-buttons
	//
	Buttons interface{} `field:"optional" json:"buttons" yaml:"buttons"`
	// The URL of an image to display on the response card.
	//
	// The image URL must be publicly available so that the platform displaying the response card has access to the image.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-imageresponsecard.html#cfn-lex-bot-imageresponsecard-imageurl
	//
	ImageUrl *string `field:"optional" json:"imageUrl" yaml:"imageUrl"`
	// The subtitle to display on the response card.
	//
	// The format of the subtitle is determined by the platform displaying the response card.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-imageresponsecard.html#cfn-lex-bot-imageresponsecard-subtitle
	//
	Subtitle *string `field:"optional" json:"subtitle" yaml:"subtitle"`
}

