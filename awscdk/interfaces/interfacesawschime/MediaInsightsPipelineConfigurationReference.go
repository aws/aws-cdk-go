package interfacesawschime


// A reference to a MediaInsightsPipelineConfiguration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mediaInsightsPipelineConfigurationReference := &MediaInsightsPipelineConfigurationReference{
//   	MediaInsightsPipelineConfigurationArn: jsii.String("mediaInsightsPipelineConfigurationArn"),
//   }
//
type MediaInsightsPipelineConfigurationReference struct {
	// The MediaInsightsPipelineConfigurationArn of the MediaInsightsPipelineConfiguration resource.
	MediaInsightsPipelineConfigurationArn *string `field:"required" json:"mediaInsightsPipelineConfigurationArn" yaml:"mediaInsightsPipelineConfigurationArn"`
}

