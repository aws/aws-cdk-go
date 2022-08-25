package awsmedialive


// The configuration of channel behavior when the input is lost.
//
// The parent of this entity is GlobalConfiguration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputLossBehaviorProperty := &inputLossBehaviorProperty{
//   	blackFrameMsec: jsii.Number(123),
//   	inputLossImageColor: jsii.String("inputLossImageColor"),
//   	inputLossImageSlate: &inputLocationProperty{
//   		passwordParam: jsii.String("passwordParam"),
//   		uri: jsii.String("uri"),
//   		username: jsii.String("username"),
//   	},
//   	inputLossImageType: jsii.String("inputLossImageType"),
//   	repeatFrameMsec: jsii.Number(123),
//   }
//
type CfnChannel_InputLossBehaviorProperty struct {
	// On input loss, the number of milliseconds to substitute black into the output before switching to the frame specified by inputLossImageType.
	//
	// A value x, where 0 <= x <= 1,000,000 and a value of 1,000,000, is interpreted as infinite.
	BlackFrameMsec *float64 `field:"optional" json:"blackFrameMsec" yaml:"blackFrameMsec"`
	// When the input loss image type is "color," this field specifies the color to use.
	//
	// Value: 6 hex characters that represent the values of RGB.
	InputLossImageColor *string `field:"optional" json:"inputLossImageColor" yaml:"inputLossImageColor"`
	// When the input loss image type is "slate," these fields specify the parameters for accessing the slate.
	InputLossImageSlate interface{} `field:"optional" json:"inputLossImageSlate" yaml:"inputLossImageSlate"`
	// Indicates whether to substitute a solid color or a slate into the output after the input loss exceeds blackFrameMsec.
	InputLossImageType *string `field:"optional" json:"inputLossImageType" yaml:"inputLossImageType"`
	// On input loss, the number of milliseconds to repeat the previous picture before substituting black into the output.
	//
	// A value x, where 0 <= x <= 1,000,000 and a value of 1,000,000, is interpreted as infinite.
	RepeatFrameMsec *float64 `field:"optional" json:"repeatFrameMsec" yaml:"repeatFrameMsec"`
}

