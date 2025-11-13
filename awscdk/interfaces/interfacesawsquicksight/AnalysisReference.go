package interfacesawsquicksight


// A reference to a Analysis resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   analysisReference := &AnalysisReference{
//   	AnalysisArn: jsii.String("analysisArn"),
//   	AnalysisId: jsii.String("analysisId"),
//   	AwsAccountId: jsii.String("awsAccountId"),
//   }
//
type AnalysisReference struct {
	// The ARN of the Analysis resource.
	AnalysisArn *string `field:"required" json:"analysisArn" yaml:"analysisArn"`
	// The AnalysisId of the Analysis resource.
	AnalysisId *string `field:"required" json:"analysisId" yaml:"analysisId"`
	// The AwsAccountId of the Analysis resource.
	AwsAccountId *string `field:"required" json:"awsAccountId" yaml:"awsAccountId"`
}

