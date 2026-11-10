package awscloudwatch


// Identity provider configuration for the domain.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   identityProviderConfigurationProperty := &IdentityProviderConfigurationProperty{
//   	IdentityCenterConfiguration: &IdentityCenterConfigurationProperty{
//   		IdentityCenterInstanceArn: jsii.String("identityCenterInstanceArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-domainfororganization-identityproviderconfiguration.html
//
type CfnDomainForOrganization_IdentityProviderConfigurationProperty struct {
	// Identity Center configuration.
	//
	// Required when identityProviders includes IDC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-domainfororganization-identityproviderconfiguration.html#cfn-cloudwatch-domainfororganization-identityproviderconfiguration-identitycenterconfiguration
	//
	IdentityCenterConfiguration interface{} `field:"optional" json:"identityCenterConfiguration" yaml:"identityCenterConfiguration"`
}

