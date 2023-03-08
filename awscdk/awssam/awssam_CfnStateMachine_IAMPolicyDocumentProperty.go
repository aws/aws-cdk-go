package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var statement interface{}
//
//   iAMPolicyDocumentProperty := map[string]interface{}{
//   	"statement": statement,
//   	"version": jsii.String("version"),
//   }
//
type CfnStateMachine_IAMPolicyDocumentProperty struct {
	// `CfnStateMachine.IAMPolicyDocumentProperty.Statement`.
	Statement interface{} `field:"required" json:"statement" yaml:"statement"`
	// `CfnStateMachine.IAMPolicyDocumentProperty.Version`.
	Version *string `field:"required" json:"version" yaml:"version"`
}

