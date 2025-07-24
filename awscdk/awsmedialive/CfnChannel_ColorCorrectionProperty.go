package awsmedialive


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   colorCorrectionProperty := &ColorCorrectionProperty{
//   	InputColorSpace: jsii.String("inputColorSpace"),
//   	OutputColorSpace: jsii.String("outputColorSpace"),
//   	Uri: jsii.String("uri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-colorcorrection.html
//
type CfnChannel_ColorCorrectionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-colorcorrection.html#cfn-medialive-channel-colorcorrection-inputcolorspace
	//
	InputColorSpace *string `field:"optional" json:"inputColorSpace" yaml:"inputColorSpace"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-colorcorrection.html#cfn-medialive-channel-colorcorrection-outputcolorspace
	//
	OutputColorSpace *string `field:"optional" json:"outputColorSpace" yaml:"outputColorSpace"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-colorcorrection.html#cfn-medialive-channel-colorcorrection-uri
	//
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

