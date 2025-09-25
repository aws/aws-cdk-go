package awsworkspacesinstances


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   licenseConfigurationRequestProperty := &LicenseConfigurationRequestProperty{
//   	LicenseConfigurationArn: jsii.String("licenseConfigurationArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-licenseconfigurationrequest.html
//
type CfnWorkspaceInstance_LicenseConfigurationRequestProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-licenseconfigurationrequest.html#cfn-workspacesinstances-workspaceinstance-licenseconfigurationrequest-licenseconfigurationarn
	//
	LicenseConfigurationArn *string `field:"optional" json:"licenseConfigurationArn" yaml:"licenseConfigurationArn"`
}

