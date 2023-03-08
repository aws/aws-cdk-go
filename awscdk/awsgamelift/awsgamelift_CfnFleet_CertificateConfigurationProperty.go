package awsgamelift


// Determines whether a TLS/SSL certificate is generated for a fleet.
//
// This feature must be enabled when creating the fleet. All instances in a fleet share the same certificate. The certificate can be retrieved by calling the [GameLift Server SDK](https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-serversdk.html) operation `GetInstanceCertificate` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   certificateConfigurationProperty := &certificateConfigurationProperty{
//   	certificateType: jsii.String("certificateType"),
//   }
//
type CfnFleet_CertificateConfigurationProperty struct {
	// Indicates whether a TLS/SSL certificate is generated for a fleet.
	//
	// Valid values include:
	//
	// - *GENERATED* - Generate a TLS/SSL certificate for this fleet.
	// - *DISABLED* - (default) Do not generate a TLS/SSL certificate for this fleet.
	CertificateType *string `field:"required" json:"certificateType" yaml:"certificateType"`
}

