package awsec2


// Indicates whether the instance is enabled for AWS Nitro Enclaves.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   enclaveOptionsProperty := &enclaveOptionsProperty{
//   	enabled: jsii.Boolean(false),
//   }
//
type CfnInstance_EnclaveOptionsProperty struct {
	// If this parameter is set to `true` , the instance is enabled for AWS Nitro Enclaves;
	//
	// otherwise, it is not enabled for AWS Nitro Enclaves.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

