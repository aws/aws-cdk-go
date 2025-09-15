package awsec2


// A reference to a NetworkInsightsAccessScopeAnalysis resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkInsightsAccessScopeAnalysisReference := &NetworkInsightsAccessScopeAnalysisReference{
//   	NetworkInsightsAccessScopeAnalysisArn: jsii.String("networkInsightsAccessScopeAnalysisArn"),
//   	NetworkInsightsAccessScopeAnalysisId: jsii.String("networkInsightsAccessScopeAnalysisId"),
//   }
//
type NetworkInsightsAccessScopeAnalysisReference struct {
	// The ARN of the NetworkInsightsAccessScopeAnalysis resource.
	NetworkInsightsAccessScopeAnalysisArn *string `field:"required" json:"networkInsightsAccessScopeAnalysisArn" yaml:"networkInsightsAccessScopeAnalysisArn"`
	// The NetworkInsightsAccessScopeAnalysisId of the NetworkInsightsAccessScopeAnalysis resource.
	NetworkInsightsAccessScopeAnalysisId *string `field:"required" json:"networkInsightsAccessScopeAnalysisId" yaml:"networkInsightsAccessScopeAnalysisId"`
}

