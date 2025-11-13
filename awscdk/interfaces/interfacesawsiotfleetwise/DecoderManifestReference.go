package interfacesawsiotfleetwise


// A reference to a DecoderManifest resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   decoderManifestReference := &DecoderManifestReference{
//   	DecoderManifestArn: jsii.String("decoderManifestArn"),
//   	DecoderManifestName: jsii.String("decoderManifestName"),
//   }
//
type DecoderManifestReference struct {
	// The ARN of the DecoderManifest resource.
	DecoderManifestArn *string `field:"required" json:"decoderManifestArn" yaml:"decoderManifestArn"`
	// The Name of the DecoderManifest resource.
	DecoderManifestName *string `field:"required" json:"decoderManifestName" yaml:"decoderManifestName"`
}

