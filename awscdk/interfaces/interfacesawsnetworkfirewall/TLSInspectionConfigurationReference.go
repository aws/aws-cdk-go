package interfacesawsnetworkfirewall


// A reference to a TLSInspectionConfiguration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tLSInspectionConfigurationReference := &TLSInspectionConfigurationReference{
//   	TlsInspectionConfigurationArn: jsii.String("tlsInspectionConfigurationArn"),
//   }
//
type TLSInspectionConfigurationReference struct {
	// The TLSInspectionConfigurationArn of the TLSInspectionConfiguration resource.
	TlsInspectionConfigurationArn *string `field:"required" json:"tlsInspectionConfigurationArn" yaml:"tlsInspectionConfigurationArn"`
}

