package interfacesawshealthlake


// A reference to a DataTransformationProfile resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataTransformationProfileReference := &DataTransformationProfileReference{
//   	DataTransformationProfileArn: jsii.String("dataTransformationProfileArn"),
//   }
//
type DataTransformationProfileReference struct {
	// The Arn of the DataTransformationProfile resource.
	DataTransformationProfileArn *string `field:"required" json:"dataTransformationProfileArn" yaml:"dataTransformationProfileArn"`
}

