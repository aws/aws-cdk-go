package awspipes


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ecsResourceRequirementProperty := &ecsResourceRequirementProperty{
//   	type: jsii.String("type"),
//   	value: jsii.String("value"),
//   }
//
type CfnPipe_EcsResourceRequirementProperty struct {
	// `CfnPipe.EcsResourceRequirementProperty.Type`.
	Type *string `field:"required" json:"type" yaml:"type"`
	// `CfnPipe.EcsResourceRequirementProperty.Value`.
	Value *string `field:"required" json:"value" yaml:"value"`
}

