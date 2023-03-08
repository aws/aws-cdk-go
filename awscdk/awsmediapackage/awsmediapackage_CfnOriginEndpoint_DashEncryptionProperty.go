package awsmediapackage


// Holds encryption information so that access to the content can be controlled by a DRM solution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dashEncryptionProperty := &dashEncryptionProperty{
//   	spekeKeyProvider: &spekeKeyProviderProperty{
//   		resourceId: jsii.String("resourceId"),
//   		roleArn: jsii.String("roleArn"),
//   		systemIds: []*string{
//   			jsii.String("systemIds"),
//   		},
//   		url: jsii.String("url"),
//
//   		// the properties below are optional
//   		certificateArn: jsii.String("certificateArn"),
//   		encryptionContractConfiguration: &encryptionContractConfigurationProperty{
//   			presetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   			presetSpeke20Video: jsii.String("presetSpeke20Video"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	keyRotationIntervalSeconds: jsii.Number(123),
//   }
//
type CfnOriginEndpoint_DashEncryptionProperty struct {
	// Parameters for the SPEKE key provider.
	SpekeKeyProvider interface{} `field:"required" json:"spekeKeyProvider" yaml:"spekeKeyProvider"`
	// Number of seconds before AWS Elemental MediaPackage rotates to a new key.
	//
	// By default, rotation is set to 60 seconds. Set to `0` to disable key rotation.
	KeyRotationIntervalSeconds *float64 `field:"optional" json:"keyRotationIntervalSeconds" yaml:"keyRotationIntervalSeconds"`
}

