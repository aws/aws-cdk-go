package awslicensemanagerusersubscriptions


// The source of the credentials that License Manager uses to reach the license server.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   credentialsProviderProperty := &CredentialsProviderProperty{
//   	SecretsManagerCredentialsProvider: &SecretsManagerCredentialsProviderProperty{
//   		SecretId: jsii.String("secretId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanagerusersubscriptions-licenseserverendpoint-credentialsprovider.html
//
type CfnLicenseServerEndpoint_CredentialsProviderProperty struct {
	// Credentials held in AWS Secrets Manager.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanagerusersubscriptions-licenseserverendpoint-credentialsprovider.html#cfn-licensemanagerusersubscriptions-licenseserverendpoint-credentialsprovider-secretsmanagercredentialsprovider
	//
	SecretsManagerCredentialsProvider interface{} `field:"required" json:"secretsManagerCredentialsProvider" yaml:"secretsManagerCredentialsProvider"`
}

