package awslicensemanagerusersubscriptions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnLicenseServerEndpoint`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLicenseServerEndpointProps := &CfnLicenseServerEndpointProps{
//   	IdentityProviderArn: jsii.String("identityProviderArn"),
//
//   	// the properties below are optional
//   	LicenseServerSettings: &LicenseServerSettingsProperty{
//   		ServerSettings: &ServerSettingsProperty{
//   			RdsSalSettings: &RdsSalSettingsProperty{
//   				RdsSalCredentialsProvider: &CredentialsProviderProperty{
//   					SecretsManagerCredentialsProvider: &SecretsManagerCredentialsProviderProperty{
//   						SecretId: jsii.String("secretId"),
//   					},
//   				},
//   			},
//   		},
//   		ServerType: jsii.String("serverType"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanagerusersubscriptions-licenseserverendpoint.html
//
type CfnLicenseServerEndpointProps struct {
	// The Amazon Resource Name (ARN) that identifies the registered identity provider the license server endpoint is created for.
	//
	// For Active Directory, that can be a self-managed Active Directory or an AWS Managed Microsoft AD.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanagerusersubscriptions-licenseserverendpoint.html#cfn-licensemanagerusersubscriptions-licenseserverendpoint-identityproviderarn
	//
	IdentityProviderArn *string `field:"required" json:"identityProviderArn" yaml:"identityProviderArn"`
	// The settings of the RDS license server to create the endpoint for.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanagerusersubscriptions-licenseserverendpoint.html#cfn-licensemanagerusersubscriptions-licenseserverendpoint-licenseserversettings
	//
	LicenseServerSettings interface{} `field:"optional" json:"licenseServerSettings" yaml:"licenseServerSettings"`
	// The tags that apply for the license server endpoint.
	//
	// A maximum of 50 tags.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanagerusersubscriptions-licenseserverendpoint.html#cfn-licensemanagerusersubscriptions-licenseserverendpoint-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

