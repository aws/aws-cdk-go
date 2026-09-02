package interfacesawsdevicefarm


// A reference to a Upload resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   uploadReference := &UploadReference{
//   	UploadArn: jsii.String("uploadArn"),
//   }
//
type UploadReference struct {
	// The Arn of the Upload resource.
	UploadArn *string `field:"required" json:"uploadArn" yaml:"uploadArn"`
}

