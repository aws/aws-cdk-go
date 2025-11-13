package interfacesawsivs


// A reference to a EncoderConfiguration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encoderConfigurationReference := &EncoderConfigurationReference{
//   	EncoderConfigurationArn: jsii.String("encoderConfigurationArn"),
//   }
//
type EncoderConfigurationReference struct {
	// The Arn of the EncoderConfiguration resource.
	EncoderConfigurationArn *string `field:"required" json:"encoderConfigurationArn" yaml:"encoderConfigurationArn"`
}

