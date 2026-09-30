package interfacesawstranscribe


// A reference to a CallAnalyticsCategory resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   callAnalyticsCategoryReference := &CallAnalyticsCategoryReference{
//   	CallAnalyticsCategoryArn: jsii.String("callAnalyticsCategoryArn"),
//   }
//
type CallAnalyticsCategoryReference struct {
	// The Arn of the CallAnalyticsCategory resource.
	CallAnalyticsCategoryArn *string `field:"required" json:"callAnalyticsCategoryArn" yaml:"callAnalyticsCategoryArn"`
}

