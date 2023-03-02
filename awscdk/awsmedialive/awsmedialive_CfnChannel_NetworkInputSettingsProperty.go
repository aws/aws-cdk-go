package awsmedialive


// Information about how to connect to the upstream system.
//
// The parent of this entity is InputSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkInputSettingsProperty := &networkInputSettingsProperty{
//   	hlsInputSettings: &hlsInputSettingsProperty{
//   		bandwidth: jsii.Number(123),
//   		bufferSegments: jsii.Number(123),
//   		retries: jsii.Number(123),
//   		retryInterval: jsii.Number(123),
//   		scte35Source: jsii.String("scte35Source"),
//   	},
//   	serverValidation: jsii.String("serverValidation"),
//   }
//
type CfnChannel_NetworkInputSettingsProperty struct {
	// Information about how to connect to the upstream system.
	HlsInputSettings interface{} `field:"optional" json:"hlsInputSettings" yaml:"hlsInputSettings"`
	// Checks HTTPS server certificates.
	//
	// When set to checkCryptographyOnly, cryptography in the certificate is checked, but not the server's name. Certain subdomains (notably S3 buckets that use dots in the bucket name) don't strictly match the corresponding certificate's wildcard pattern and would otherwise cause the channel to error. This setting is ignored for protocols that do not use HTTPS.
	ServerValidation *string `field:"optional" json:"serverValidation" yaml:"serverValidation"`
}

