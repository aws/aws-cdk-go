package awslicensemanagerusersubscriptions


// Credentials held in AWS Secrets Manager.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   secretsManagerCredentialsProviderProperty := &SecretsManagerCredentialsProviderProperty{
//   	SecretId: jsii.String("secretId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanagerusersubscriptions-licenseserverendpoint-secretsmanagercredentialsprovider.html
//
type CfnLicenseServerEndpoint_SecretsManagerCredentialsProviderProperty struct {
	// The identifier of the Secrets Manager secret that holds the license server administrator credentials.
	//
	// The name of the secret must begin with the prefix license-manager-user-.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanagerusersubscriptions-licenseserverendpoint-secretsmanagercredentialsprovider.html#cfn-licensemanagerusersubscriptions-licenseserverendpoint-secretsmanagercredentialsprovider-secretid
	//
	SecretId *string `field:"required" json:"secretId" yaml:"secretId"`
}

