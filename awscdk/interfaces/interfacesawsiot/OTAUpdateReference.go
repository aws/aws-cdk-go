package interfacesawsiot


// A reference to a OTAUpdate resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   oTAUpdateReference := &OTAUpdateReference{
//   	OtaUpdateArn: jsii.String("otaUpdateArn"),
//   	OtaUpdateId: jsii.String("otaUpdateId"),
//   }
//
type OTAUpdateReference struct {
	// The OtaUpdateArn of the OTAUpdate resource.
	OtaUpdateArn *string `field:"required" json:"otaUpdateArn" yaml:"otaUpdateArn"`
	// The OtaUpdateId of the OTAUpdate resource.
	OtaUpdateId *string `field:"required" json:"otaUpdateId" yaml:"otaUpdateId"`
}

