package awslicensemanagerusersubscriptions


// The settings of a license server, by server type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serverSettingsProperty := &ServerSettingsProperty{
//   	RdsSalSettings: &RdsSalSettingsProperty{
//   		RdsSalCredentialsProvider: &CredentialsProviderProperty{
//   			SecretsManagerCredentialsProvider: &SecretsManagerCredentialsProviderProperty{
//   				SecretId: jsii.String("secretId"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanagerusersubscriptions-licenseserverendpoint-serversettings.html
//
type CfnLicenseServerEndpoint_ServerSettingsProperty struct {
	// The settings of a Remote Desktop Services Subscriber Access License server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanagerusersubscriptions-licenseserverendpoint-serversettings.html#cfn-licensemanagerusersubscriptions-licenseserverendpoint-serversettings-rdssalsettings
	//
	RdsSalSettings interface{} `field:"required" json:"rdsSalSettings" yaml:"rdsSalSettings"`
}

