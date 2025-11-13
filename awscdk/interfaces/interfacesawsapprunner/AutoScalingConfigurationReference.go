package interfacesawsapprunner


// A reference to a AutoScalingConfiguration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   autoScalingConfigurationReference := &AutoScalingConfigurationReference{
//   	AutoScalingConfigurationArn: jsii.String("autoScalingConfigurationArn"),
//   }
//
type AutoScalingConfigurationReference struct {
	// The AutoScalingConfigurationArn of the AutoScalingConfiguration resource.
	AutoScalingConfigurationArn *string `field:"required" json:"autoScalingConfigurationArn" yaml:"autoScalingConfigurationArn"`
}

