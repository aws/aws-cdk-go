package interfacesawsssm


// A reference to a ServiceSetting resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceSettingReference := &ServiceSettingReference{
//   	ServiceSettingArn: jsii.String("serviceSettingArn"),
//   }
//
type ServiceSettingReference struct {
	// The Arn of the ServiceSetting resource.
	ServiceSettingArn *string `field:"required" json:"serviceSettingArn" yaml:"serviceSettingArn"`
}

