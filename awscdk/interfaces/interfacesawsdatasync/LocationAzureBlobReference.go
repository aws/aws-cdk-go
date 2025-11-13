package interfacesawsdatasync


// A reference to a LocationAzureBlob resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   locationAzureBlobReference := &LocationAzureBlobReference{
//   	LocationArn: jsii.String("locationArn"),
//   }
//
type LocationAzureBlobReference struct {
	// The LocationArn of the LocationAzureBlob resource.
	LocationArn *string `field:"required" json:"locationArn" yaml:"locationArn"`
}

