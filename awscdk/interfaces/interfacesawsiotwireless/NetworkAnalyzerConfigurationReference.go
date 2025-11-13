package interfacesawsiotwireless


// A reference to a NetworkAnalyzerConfiguration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkAnalyzerConfigurationReference := &NetworkAnalyzerConfigurationReference{
//   	NetworkAnalyzerConfigurationArn: jsii.String("networkAnalyzerConfigurationArn"),
//   	NetworkAnalyzerConfigurationName: jsii.String("networkAnalyzerConfigurationName"),
//   }
//
type NetworkAnalyzerConfigurationReference struct {
	// The ARN of the NetworkAnalyzerConfiguration resource.
	NetworkAnalyzerConfigurationArn *string `field:"required" json:"networkAnalyzerConfigurationArn" yaml:"networkAnalyzerConfigurationArn"`
	// The Name of the NetworkAnalyzerConfiguration resource.
	NetworkAnalyzerConfigurationName *string `field:"required" json:"networkAnalyzerConfigurationName" yaml:"networkAnalyzerConfigurationName"`
}

