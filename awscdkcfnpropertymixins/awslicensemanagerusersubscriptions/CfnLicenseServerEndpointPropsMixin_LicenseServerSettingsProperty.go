package awslicensemanagerusersubscriptions


// The settings of the RDS license server to create the endpoint for.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   licenseServerSettingsProperty := &LicenseServerSettingsProperty{
//   	ServerSettings: &ServerSettingsProperty{
//   		RdsSalSettings: &RdsSalSettingsProperty{
//   			RdsSalCredentialsProvider: &CredentialsProviderProperty{
//   				SecretsManagerCredentialsProvider: &SecretsManagerCredentialsProviderProperty{
//   					SecretId: jsii.String("secretId"),
//   				},
//   			},
//   		},
//   	},
//   	ServerType: jsii.String("serverType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanagerusersubscriptions-licenseserverendpoint-licenseserversettings.html
//
type CfnLicenseServerEndpointPropsMixin_LicenseServerSettingsProperty struct {
	// The settings of a license server, by server type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanagerusersubscriptions-licenseserverendpoint-licenseserversettings.html#cfn-licensemanagerusersubscriptions-licenseserverendpoint-licenseserversettings-serversettings
	//
	ServerSettings interface{} `field:"optional" json:"serverSettings" yaml:"serverSettings"`
	// The type of license server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanagerusersubscriptions-licenseserverendpoint-licenseserversettings.html#cfn-licensemanagerusersubscriptions-licenseserverendpoint-licenseserversettings-servertype
	//
	ServerType *string `field:"optional" json:"serverType" yaml:"serverType"`
}

