package awsquicksight


// Analysis error.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   analysisErrorProperty := &analysisErrorProperty{
//   	message: jsii.String("message"),
//   	type: jsii.String("type"),
//   }
//
type CfnAnalysis_AnalysisErrorProperty struct {
	// The message associated with the analysis error.
	Message *string `field:"optional" json:"message" yaml:"message"`
	// The type of the analysis error.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

