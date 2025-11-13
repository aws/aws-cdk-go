package interfacesawsimagebuilder


// A reference to a InfrastructureConfiguration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   infrastructureConfigurationReference := &InfrastructureConfigurationReference{
//   	InfrastructureConfigurationArn: jsii.String("infrastructureConfigurationArn"),
//   }
//
type InfrastructureConfigurationReference struct {
	// The Arn of the InfrastructureConfiguration resource.
	InfrastructureConfigurationArn *string `field:"required" json:"infrastructureConfigurationArn" yaml:"infrastructureConfigurationArn"`
}

