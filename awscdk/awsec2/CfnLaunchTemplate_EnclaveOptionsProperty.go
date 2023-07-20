package awsec2


// Indicates whether the instance is enabled for AWS Nitro Enclaves.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   enclaveOptionsProperty := &EnclaveOptionsProperty{
//   	Enabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-enclaveoptions.html
//
type CfnLaunchTemplate_EnclaveOptionsProperty struct {
	// If this parameter is set to `true` , the instance is enabled for AWS Nitro Enclaves;
	//
	// otherwise, it is not enabled for AWS Nitro Enclaves.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-enclaveoptions.html#cfn-ec2-launchtemplate-enclaveoptions-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

