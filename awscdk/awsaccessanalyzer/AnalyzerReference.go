package awsaccessanalyzer


// A reference to a Analyzer resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   analyzerReference := &AnalyzerReference{
//   	AnalyzerArn: jsii.String("analyzerArn"),
//   }
//
type AnalyzerReference struct {
	// The Arn of the Analyzer resource.
	AnalyzerArn *string `field:"required" json:"analyzerArn" yaml:"analyzerArn"`
}

