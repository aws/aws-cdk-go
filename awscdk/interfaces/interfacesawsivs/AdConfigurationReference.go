package interfacesawsivs


// A reference to a AdConfiguration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   adConfigurationReference := &AdConfigurationReference{
//   	AdConfigurationArn: jsii.String("adConfigurationArn"),
//   }
//
type AdConfigurationReference struct {
	// The Arn of the AdConfiguration resource.
	AdConfigurationArn *string `field:"required" json:"adConfigurationArn" yaml:"adConfigurationArn"`
}

