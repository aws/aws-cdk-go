package awsdatasync


// Specifies configuration information for a DataSync-managed secret, such as an authentication token or set of credentials that DataSync uses to access a specific transfer location.
//
// DataSync uses the default AWS-managed KMS key to encrypt this secret in AWS Secrets Manager.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   managedSecretConfigProperty := &ManagedSecretConfigProperty{
//   	SecretArn: jsii.String("secretArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationfsxontap-managedsecretconfig.html
//
type CfnLocationFSxONTAPPropsMixin_ManagedSecretConfigProperty struct {
	// Specifies the ARN for an AWS Secrets Manager secret.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationfsxontap-managedsecretconfig.html#cfn-datasync-locationfsxontap-managedsecretconfig-secretarn
	//
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
}

