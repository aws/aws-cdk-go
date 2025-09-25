package awsdevopsguru


// A reference to a ResourceCollection resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceCollectionReference := &ResourceCollectionReference{
//   	ResourceCollectionType: jsii.String("resourceCollectionType"),
//   }
//
type ResourceCollectionReference struct {
	// The ResourceCollectionType of the ResourceCollection resource.
	ResourceCollectionType *string `field:"required" json:"resourceCollectionType" yaml:"resourceCollectionType"`
}

