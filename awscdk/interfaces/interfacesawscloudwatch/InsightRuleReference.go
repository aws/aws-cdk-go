package interfacesawscloudwatch


// A reference to a InsightRule resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   insightRuleReference := &InsightRuleReference{
//   	InsightRuleArn: jsii.String("insightRuleArn"),
//   }
//
type InsightRuleReference struct {
	// The Arn of the InsightRule resource.
	InsightRuleArn *string `field:"required" json:"insightRuleArn" yaml:"insightRuleArn"`
}

