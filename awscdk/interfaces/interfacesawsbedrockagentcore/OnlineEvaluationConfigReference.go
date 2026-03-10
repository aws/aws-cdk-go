package interfacesawsbedrockagentcore


// A reference to a OnlineEvaluationConfig resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   onlineEvaluationConfigReference := &OnlineEvaluationConfigReference{
//   	OnlineEvaluationConfigArn: jsii.String("onlineEvaluationConfigArn"),
//   }
//
type OnlineEvaluationConfigReference struct {
	// The OnlineEvaluationConfigArn of the OnlineEvaluationConfig resource.
	OnlineEvaluationConfigArn *string `field:"required" json:"onlineEvaluationConfigArn" yaml:"onlineEvaluationConfigArn"`
}

