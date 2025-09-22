package awsapplicationautoscaling


// A reference to a ScalableTarget resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scalableTargetReference := &ScalableTargetReference{
//   	ResourceId: jsii.String("resourceId"),
//   	ScalableDimension: jsii.String("scalableDimension"),
//   	ServiceNamespace: jsii.String("serviceNamespace"),
//   }
//
type ScalableTargetReference struct {
	// The ResourceId of the ScalableTarget resource.
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// The ScalableDimension of the ScalableTarget resource.
	ScalableDimension *string `field:"required" json:"scalableDimension" yaml:"scalableDimension"`
	// The ServiceNamespace of the ScalableTarget resource.
	ServiceNamespace *string `field:"required" json:"serviceNamespace" yaml:"serviceNamespace"`
}

