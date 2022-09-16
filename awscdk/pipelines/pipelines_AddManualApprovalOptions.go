package pipelines


// Options for addManualApproval.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   addManualApprovalOptions := &addManualApprovalOptions{
//   	actionName: jsii.String("actionName"),
//   	runOrder: jsii.Number(123),
//   }
//
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
type AddManualApprovalOptions struct {
	// The name of the manual approval action.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ActionName *string `field:"optional" json:"actionName" yaml:"actionName"`
	// The runOrder for this action.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	RunOrder *float64 `field:"optional" json:"runOrder" yaml:"runOrder"`
}

