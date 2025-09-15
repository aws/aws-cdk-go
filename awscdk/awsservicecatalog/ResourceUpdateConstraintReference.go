package awsservicecatalog


// A reference to a ResourceUpdateConstraint resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceUpdateConstraintReference := &ResourceUpdateConstraintReference{
//   	ResourceUpdateConstraintId: jsii.String("resourceUpdateConstraintId"),
//   }
//
type ResourceUpdateConstraintReference struct {
	// The Id of the ResourceUpdateConstraint resource.
	ResourceUpdateConstraintId *string `field:"required" json:"resourceUpdateConstraintId" yaml:"resourceUpdateConstraintId"`
}

