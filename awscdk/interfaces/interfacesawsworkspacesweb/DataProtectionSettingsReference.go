package interfacesawsworkspacesweb


// A reference to a DataProtectionSettings resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataProtectionSettingsReference := &DataProtectionSettingsReference{
//   	DataProtectionSettingsArn: jsii.String("dataProtectionSettingsArn"),
//   }
//
type DataProtectionSettingsReference struct {
	// The DataProtectionSettingsArn of the DataProtectionSettings resource.
	DataProtectionSettingsArn *string `field:"required" json:"dataProtectionSettingsArn" yaml:"dataProtectionSettingsArn"`
}

