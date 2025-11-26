package previewawsgameliftmixins


// Determines whether a TLS/SSL certificate is generated for a fleet.
//
// This feature must be enabled when creating the fleet. All instances in a fleet share the same certificate. The certificate can be retrieved by calling the [GameLift Server SDK](https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-serversdk.html) operation `GetInstanceCertificate` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   certificateConfigurationProperty := &CertificateConfigurationProperty{
//   	CertificateType: jsii.String("certificateType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-certificateconfiguration.html
//
type CfnFleetPropsMixin_CertificateConfigurationProperty struct {
	// Indicates whether a TLS/SSL certificate is generated for a fleet.
	//
	// Valid values include:
	//
	// - *GENERATED* - Generate a TLS/SSL certificate for this fleet.
	// - *DISABLED* - (default) Do not generate a TLS/SSL certificate for this fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-certificateconfiguration.html#cfn-gamelift-fleet-certificateconfiguration-certificatetype
	//
	CertificateType *string `field:"optional" json:"certificateType" yaml:"certificateType"`
}

