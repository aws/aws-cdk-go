package mixinsawsrekognition


// Connected home settings to use on a streaming video.
//
// Defining the settings is required in the request parameter for `CreateStreamProcessor` . Including this setting in the CreateStreamProcessor request lets you use the stream processor for connected home features. You can then select what you want the stream processor to detect, such as people or pets.
//
// When the stream processor has started, one notification is sent for each object class specified. For example, if packages and pets are selected, one SNS notification is published the first time a package is detected and one SNS notification is published the first time a pet is detected. An end-of-session summary is also published. For more information, see the ConnectedHome section of [StreamProcessorSettings](https://docs.aws.amazon.com/rekognition/latest/APIReference/API_StreamProcessorSettings) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   connectedHomeSettingsProperty := &ConnectedHomeSettingsProperty{
//   	Labels: []*string{
//   		jsii.String("labels"),
//   	},
//   	MinConfidence: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rekognition-streamprocessor-connectedhomesettings.html
//
type CfnStreamProcessorPropsMixin_ConnectedHomeSettingsProperty struct {
	// Specifies what you want to detect in the video, such as people, packages, or pets.
	//
	// The current valid labels you can include in this list are: "PERSON", "PET", "PACKAGE", and "ALL".
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rekognition-streamprocessor-connectedhomesettings.html#cfn-rekognition-streamprocessor-connectedhomesettings-labels
	//
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
	// The minimum confidence required to label an object in the video.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rekognition-streamprocessor-connectedhomesettings.html#cfn-rekognition-streamprocessor-connectedhomesettings-minconfidence
	//
	MinConfidence *float64 `field:"optional" json:"minConfidence" yaml:"minConfidence"`
}

