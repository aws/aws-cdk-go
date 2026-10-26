package awslicensemanagerusersubscriptions


// The settings of a Remote Desktop Services Subscriber Access License server.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   rdsSalSettingsProperty := &RdsSalSettingsProperty{
//   	RdsSalCredentialsProvider: &CredentialsProviderProperty{
//   		SecretsManagerCredentialsProvider: &SecretsManagerCredentialsProviderProperty{
//   			SecretId: jsii.String("secretId"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanagerusersubscriptions-licenseserverendpoint-rdssalsettings.html
//
type CfnLicenseServerEndpointPropsMixin_RdsSalSettingsProperty struct {
	// The source of the credentials that License Manager uses to reach the license server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanagerusersubscriptions-licenseserverendpoint-rdssalsettings.html#cfn-licensemanagerusersubscriptions-licenseserverendpoint-rdssalsettings-rdssalcredentialsprovider
	//
	RdsSalCredentialsProvider interface{} `field:"optional" json:"rdsSalCredentialsProvider" yaml:"rdsSalCredentialsProvider"`
}

