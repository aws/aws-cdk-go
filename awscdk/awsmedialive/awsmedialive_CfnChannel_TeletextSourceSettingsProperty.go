package awsmedialive


// Information about the Teletext captions to extract from the input.
//
// The parent of this entity is CaptionSelectorSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   teletextSourceSettingsProperty := &teletextSourceSettingsProperty{
//   	outputRectangle: &captionRectangleProperty{
//   		height: jsii.Number(123),
//   		leftOffset: jsii.Number(123),
//   		topOffset: jsii.Number(123),
//   		width: jsii.Number(123),
//   	},
//   	pageNumber: jsii.String("pageNumber"),
//   }
//
type CfnChannel_TeletextSourceSettingsProperty struct {
	// Settings to configure the caption rectangle for an output captions that will be created using this Teletext source captions.
	OutputRectangle interface{} `field:"optional" json:"outputRectangle" yaml:"outputRectangle"`
	// Specifies the Teletext page number within the data stream from which to extract captions.
	//
	// The range is 0x100 (256) to 0x8FF (2303). This is unused for passthrough. It should be specified as a hexadecimal string with no "0x" prefix.
	PageNumber *string `field:"optional" json:"pageNumber" yaml:"pageNumber"`
}

