package awsmedialive


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   colorCorrectionSettingsProperty := &ColorCorrectionSettingsProperty{
//   	GlobalColorCorrections: []interface{}{
//   		&ColorCorrectionProperty{
//   			InputColorSpace: jsii.String("inputColorSpace"),
//   			OutputColorSpace: jsii.String("outputColorSpace"),
//   			Uri: jsii.String("uri"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-colorcorrectionsettings.html
//
type CfnChannel_ColorCorrectionSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-colorcorrectionsettings.html#cfn-medialive-channel-colorcorrectionsettings-globalcolorcorrections
	//
	GlobalColorCorrections interface{} `field:"optional" json:"globalColorCorrections" yaml:"globalColorCorrections"`
}

