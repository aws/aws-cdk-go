package interfacesawssecurityhub


// A reference to a Insight resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   insightReference := &InsightReference{
//   	InsightArn: jsii.String("insightArn"),
//   }
//
type InsightReference struct {
	// The InsightArn of the Insight resource.
	InsightArn *string `field:"required" json:"insightArn" yaml:"insightArn"`
}

