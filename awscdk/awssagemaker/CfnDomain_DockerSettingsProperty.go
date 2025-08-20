package awssagemaker


// A collection of settings that configure the domain's Docker interaction.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dockerSettingsProperty := &DockerSettingsProperty{
//   	EnableDockerAccess: jsii.String("enableDockerAccess"),
//   	VpcOnlyTrustedAccounts: []*string{
//   		jsii.String("vpcOnlyTrustedAccounts"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-dockersettings.html
//
type CfnDomain_DockerSettingsProperty struct {
	// Indicates whether the domain can access Docker.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-dockersettings.html#cfn-sagemaker-domain-dockersettings-enabledockeraccess
	//
	EnableDockerAccess *string `field:"optional" json:"enableDockerAccess" yaml:"enableDockerAccess"`
	// The list of AWS accounts that are trusted when the domain is created in VPC-only mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-dockersettings.html#cfn-sagemaker-domain-dockersettings-vpconlytrustedaccounts
	//
	VpcOnlyTrustedAccounts *[]*string `field:"optional" json:"vpcOnlyTrustedAccounts" yaml:"vpcOnlyTrustedAccounts"`
}

