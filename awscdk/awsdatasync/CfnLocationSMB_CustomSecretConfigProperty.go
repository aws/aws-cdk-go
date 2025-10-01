package awsdatasync


// Specifies configuration information for a customer-managed Secrets Manager secret where a storage location authentication token or secret key is stored in plain text.
//
// This configuration includes the secret ARN, and the ARN for an IAM role that provides access to the secret.
//
// > You can use either `CmkSecretConfig` or `CustomSecretConfig` to provide credentials for a `CreateLocation` request. Do not provide both parameters for the same request.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customSecretConfigProperty := &CustomSecretConfigProperty{
//   	SecretAccessRoleArn: jsii.String("secretAccessRoleArn"),
//   	SecretArn: jsii.String("secretArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationsmb-customsecretconfig.html
//
type CfnLocationSMB_CustomSecretConfigProperty struct {
	// Specifies the ARN for the AWS Identity and Access Management role that DataSync uses to access the secret specified for `SecretArn` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationsmb-customsecretconfig.html#cfn-datasync-locationsmb-customsecretconfig-secretaccessrolearn
	//
	SecretAccessRoleArn *string `field:"required" json:"secretAccessRoleArn" yaml:"secretAccessRoleArn"`
	// Specifies the ARN for an AWS Secrets Manager secret.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationsmb-customsecretconfig.html#cfn-datasync-locationsmb-customsecretconfig-secretarn
	//
	SecretArn *string `field:"required" json:"secretArn" yaml:"secretArn"`
}

