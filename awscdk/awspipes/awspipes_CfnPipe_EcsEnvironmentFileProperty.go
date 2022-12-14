package awspipes


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ecsEnvironmentFileProperty := &ecsEnvironmentFileProperty{
//   	type: jsii.String("type"),
//   	value: jsii.String("value"),
//   }
//
type CfnPipe_EcsEnvironmentFileProperty struct {
	// `CfnPipe.EcsEnvironmentFileProperty.Type`.
	Type *string `field:"required" json:"type" yaml:"type"`
	// `CfnPipe.EcsEnvironmentFileProperty.Value`.
	Value *string `field:"required" json:"value" yaml:"value"`
}

