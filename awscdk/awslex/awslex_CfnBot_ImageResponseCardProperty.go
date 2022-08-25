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
//   imageResponseCardProperty := &imageResponseCardProperty{
//   	title: jsii.String("title"),
//
//   	// the properties below are optional
//   	buttons: []interface{}{
//   		&buttonProperty{
//   			text: jsii.String("text"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	imageUrl: jsii.String("imageUrl"),
//   	subtitle: jsii.String("subtitle"),
//   }
//
type CfnBot_ImageResponseCardProperty struct {
	// The title to display on the response card.
	//
	// The format of the title is determined by the platform displaying the response card.
	Title *string `field:"required" json:"title" yaml:"title"`
	// A list of buttons that should be displayed on the response card.
	//
	// The arrangement of the buttons is determined by the platform that displays the buttons.
	Buttons interface{} `field:"optional" json:"buttons" yaml:"buttons"`
	// The URL of an image to display on the response card.
	//
	// The image URL must be publicly available so that the platform displaying the response card has access to the image.
	ImageUrl *string `field:"optional" json:"imageUrl" yaml:"imageUrl"`
	// The subtitle to display on the response card.
	//
	// The format of the subtitle is determined by the platform displaying the response card.
	Subtitle *string `field:"optional" json:"subtitle" yaml:"subtitle"`
}

