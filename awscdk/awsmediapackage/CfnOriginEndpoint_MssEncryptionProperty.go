package awsmediapackage


// Holds encryption information so that access to the content can be controlled by a DRM solution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mssEncryptionProperty := &MssEncryptionProperty{
//   	SpekeKeyProvider: &SpekeKeyProviderProperty{
//   		ResourceId: jsii.String("resourceId"),
//   		RoleArn: jsii.String("roleArn"),
//   		SystemIds: []*string{
//   			jsii.String("systemIds"),
//   		},
//   		Url: jsii.String("url"),
//
//   		// the properties below are optional
//   		CertificateArn: jsii.String("certificateArn"),
//   		EncryptionContractConfiguration: &EncryptionContractConfigurationProperty{
//   		},
//   	},
//   }
//
type CfnOriginEndpoint_MssEncryptionProperty struct {
	// Parameters for the SPEKE key provider.
	SpekeKeyProvider interface{} `field:"required" json:"spekeKeyProvider" yaml:"spekeKeyProvider"`
}

