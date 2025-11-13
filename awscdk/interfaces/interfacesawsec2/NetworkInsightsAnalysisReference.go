package interfacesawsec2


// A reference to a NetworkInsightsAnalysis resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkInsightsAnalysisReference := &NetworkInsightsAnalysisReference{
//   	NetworkInsightsAnalysisArn: jsii.String("networkInsightsAnalysisArn"),
//   	NetworkInsightsAnalysisId: jsii.String("networkInsightsAnalysisId"),
//   }
//
type NetworkInsightsAnalysisReference struct {
	// The ARN of the NetworkInsightsAnalysis resource.
	NetworkInsightsAnalysisArn *string `field:"required" json:"networkInsightsAnalysisArn" yaml:"networkInsightsAnalysisArn"`
	// The NetworkInsightsAnalysisId of the NetworkInsightsAnalysis resource.
	NetworkInsightsAnalysisId *string `field:"required" json:"networkInsightsAnalysisId" yaml:"networkInsightsAnalysisId"`
}

