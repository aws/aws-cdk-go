package interfacesawscleanroomsml


// A reference to a ConfiguredModelAlgorithm resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   configuredModelAlgorithmReference := &ConfiguredModelAlgorithmReference{
//   	ConfiguredModelAlgorithmArn: jsii.String("configuredModelAlgorithmArn"),
//   }
//
type ConfiguredModelAlgorithmReference struct {
	// The ConfiguredModelAlgorithmArn of the ConfiguredModelAlgorithm resource.
	ConfiguredModelAlgorithmArn *string `field:"required" json:"configuredModelAlgorithmArn" yaml:"configuredModelAlgorithmArn"`
}

