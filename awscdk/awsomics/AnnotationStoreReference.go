package awsomics


// A reference to a AnnotationStore resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   annotationStoreReference := &AnnotationStoreReference{
//   	AnnotationStoreName: jsii.String("annotationStoreName"),
//   }
//
type AnnotationStoreReference struct {
	// The Name of the AnnotationStore resource.
	AnnotationStoreName *string `field:"required" json:"annotationStoreName" yaml:"annotationStoreName"`
}

