package awsquicksight


// The options that determine the negative value configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   negativeValueConfigurationProperty := &NegativeValueConfigurationProperty{
//   	DisplayMode: jsii.String("displayMode"),
//   }
//
type CfnTemplate_NegativeValueConfigurationProperty struct {
	// Determines the display mode of the negative value configuration.
	DisplayMode *string `field:"required" json:"displayMode" yaml:"displayMode"`
}

