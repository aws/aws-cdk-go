package awsiot


// A reference to a MitigationAction resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mitigationActionReference := &MitigationActionReference{
//   	ActionName: jsii.String("actionName"),
//   	MitigationActionArn: jsii.String("mitigationActionArn"),
//   }
//
type MitigationActionReference struct {
	// The ActionName of the MitigationAction resource.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// The ARN of the MitigationAction resource.
	MitigationActionArn *string `field:"required" json:"mitigationActionArn" yaml:"mitigationActionArn"`
}

