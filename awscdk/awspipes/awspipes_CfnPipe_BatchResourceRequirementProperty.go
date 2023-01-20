package awspipes


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   batchResourceRequirementProperty := &batchResourceRequirementProperty{
//   	type: jsii.String("type"),
//   	value: jsii.String("value"),
//   }
//
type CfnPipe_BatchResourceRequirementProperty struct {
	// `CfnPipe.BatchResourceRequirementProperty.Type`.
	Type *string `field:"required" json:"type" yaml:"type"`
	// `CfnPipe.BatchResourceRequirementProperty.Value`.
	Value *string `field:"required" json:"value" yaml:"value"`
}

