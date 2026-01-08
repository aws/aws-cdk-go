package awsapplicationautoscaling


// Attributes for importing a scalable target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scalableTargetAttributes := &ScalableTargetAttributes{
//   	ScalableDimension: jsii.String("scalableDimension"),
//   	ScalableTargetId: jsii.String("scalableTargetId"),
//   	ServiceNamespace: jsii.String("serviceNamespace"),
//   }
//
type ScalableTargetAttributes struct {
	// The scalable dimension that's associated with the scalable target.
	ScalableDimension *string `field:"required" json:"scalableDimension" yaml:"scalableDimension"`
	// The scalable target ID.
	ScalableTargetId *string `field:"required" json:"scalableTargetId" yaml:"scalableTargetId"`
	// The namespace of the AWS service that provides the resource.
	ServiceNamespace *string `field:"required" json:"serviceNamespace" yaml:"serviceNamespace"`
}

