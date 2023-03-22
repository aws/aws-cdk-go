package awsmedialive


// The encoding information for output captions.
//
// The parent of this entity is EncoderSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   captionDescriptionProperty := &captionDescriptionProperty{
//   	captionSelectorName: jsii.String("captionSelectorName"),
//   	destinationSettings: &captionDestinationSettingsProperty{
//   		aribDestinationSettings: &aribDestinationSettingsProperty{
//   		},
//   		burnInDestinationSettings: &burnInDestinationSettingsProperty{
//   			alignment: jsii.String("alignment"),
//   			backgroundColor: jsii.String("backgroundColor"),
//   			backgroundOpacity: jsii.Number(123),
//   			font: &inputLocationProperty{
//   				passwordParam: jsii.String("passwordParam"),
//   				uri: jsii.String("uri"),
//   				username: jsii.String("username"),
//   			},
//   			fontColor: jsii.String("fontColor"),
//   			fontOpacity: jsii.Number(123),
//   			fontResolution: jsii.Number(123),
//   			fontSize: jsii.String("fontSize"),
//   			outlineColor: jsii.String("outlineColor"),
//   			outlineSize: jsii.Number(123),
//   			shadowColor: jsii.String("shadowColor"),
//   			shadowOpacity: jsii.Number(123),
//   			shadowXOffset: jsii.Number(123),
//   			shadowYOffset: jsii.Number(123),
//   			teletextGridControl: jsii.String("teletextGridControl"),
//   			xPosition: jsii.Number(123),
//   			yPosition: jsii.Number(123),
//   		},
//   		dvbSubDestinationSettings: &dvbSubDestinationSettingsProperty{
//   			alignment: jsii.String("alignment"),
//   			backgroundColor: jsii.String("backgroundColor"),
//   			backgroundOpacity: jsii.Number(123),
//   			font: &inputLocationProperty{
//   				passwordParam: jsii.String("passwordParam"),
//   				uri: jsii.String("uri"),
//   				username: jsii.String("username"),
//   			},
//   			fontColor: jsii.String("fontColor"),
//   			fontOpacity: jsii.Number(123),
//   			fontResolution: jsii.Number(123),
//   			fontSize: jsii.String("fontSize"),
//   			outlineColor: jsii.String("outlineColor"),
//   			outlineSize: jsii.Number(123),
//   			shadowColor: jsii.String("shadowColor"),
//   			shadowOpacity: jsii.Number(123),
//   			shadowXOffset: jsii.Number(123),
//   			shadowYOffset: jsii.Number(123),
//   			teletextGridControl: jsii.String("teletextGridControl"),
//   			xPosition: jsii.Number(123),
//   			yPosition: jsii.Number(123),
//   		},
//   		ebuTtDDestinationSettings: &ebuTtDDestinationSettingsProperty{
//   			copyrightHolder: jsii.String("copyrightHolder"),
//   			fillLineGap: jsii.String("fillLineGap"),
//   			fontFamily: jsii.String("fontFamily"),
//   			styleControl: jsii.String("styleControl"),
//   		},
//   		embeddedDestinationSettings: &embeddedDestinationSettingsProperty{
//   		},
//   		embeddedPlusScte20DestinationSettings: &embeddedPlusScte20DestinationSettingsProperty{
//   		},
//   		rtmpCaptionInfoDestinationSettings: &rtmpCaptionInfoDestinationSettingsProperty{
//   		},
//   		scte20PlusEmbeddedDestinationSettings: &scte20PlusEmbeddedDestinationSettingsProperty{
//   		},
//   		scte27DestinationSettings: &scte27DestinationSettingsProperty{
//   		},
//   		smpteTtDestinationSettings: &smpteTtDestinationSettingsProperty{
//   		},
//   		teletextDestinationSettings: &teletextDestinationSettingsProperty{
//   		},
//   		ttmlDestinationSettings: &ttmlDestinationSettingsProperty{
//   			styleControl: jsii.String("styleControl"),
//   		},
//   		webvttDestinationSettings: &webvttDestinationSettingsProperty{
//   			styleControl: jsii.String("styleControl"),
//   		},
//   	},
//   	languageCode: jsii.String("languageCode"),
//   	languageDescription: jsii.String("languageDescription"),
//   	name: jsii.String("name"),
//   }
//
type CfnChannel_CaptionDescriptionProperty struct {
	// Specifies which input captions selector to use as a captions source when generating output captions.
	//
	// This field should match a captionSelector name.
	CaptionSelectorName *string `field:"optional" json:"captionSelectorName" yaml:"captionSelectorName"`
	// Additional settings for a captions destination that depend on the destination type.
	DestinationSettings interface{} `field:"optional" json:"destinationSettings" yaml:"destinationSettings"`
	// An ISO 639-2 three-digit code.
	//
	// For more information, see http://www.loc.gov/standards/iso639-2/.
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// Human-readable information to indicate the captions that are available for players (for example, English or Spanish).
	LanguageDescription *string `field:"optional" json:"languageDescription" yaml:"languageDescription"`
	// The name of the captions description.
	//
	// The name is used to associate a captions description with an output. Names must be unique within a channel.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

