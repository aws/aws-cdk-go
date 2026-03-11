package awsdatasync


// Specifies configuration information for a customer-managed secret, such as an authentication token or set of credentials that DataSync uses to access a specific transfer location, and an IAM role that DataSync can assume and access the customer-managed secret.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   customSecretConfigProperty := &CustomSecretConfigProperty{
//   	SecretAccessRoleArn: jsii.String("secretAccessRoleArn"),
//   	SecretArn: jsii.String("secretArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationfsxontap-customsecretconfig.html
//
type CfnLocationFSxONTAPPropsMixin_CustomSecretConfigProperty struct {
	// Specifies the ARN for the AWS Identity and Access Management role that DataSync uses to access the secret specified for SecretArn.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationfsxontap-customsecretconfig.html#cfn-datasync-locationfsxontap-customsecretconfig-secretaccessrolearn
	//
	SecretAccessRoleArn *string `field:"optional" json:"secretAccessRoleArn" yaml:"secretAccessRoleArn"`
	// Specifies the ARN for a customer created AWS Secrets Manager secret.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationfsxontap-customsecretconfig.html#cfn-datasync-locationfsxontap-customsecretconfig-secretarn
	//
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
}

