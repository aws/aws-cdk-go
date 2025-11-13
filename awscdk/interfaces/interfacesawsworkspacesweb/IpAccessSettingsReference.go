package interfacesawsworkspacesweb


// A reference to a IpAccessSettings resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ipAccessSettingsReference := &IpAccessSettingsReference{
//   	IpAccessSettingsArn: jsii.String("ipAccessSettingsArn"),
//   }
//
type IpAccessSettingsReference struct {
	// The IpAccessSettingsArn of the IpAccessSettings resource.
	IpAccessSettingsArn *string `field:"required" json:"ipAccessSettingsArn" yaml:"ipAccessSettingsArn"`
}

