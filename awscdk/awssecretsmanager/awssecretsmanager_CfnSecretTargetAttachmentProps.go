package awssecretsmanager


// Properties for defining a `CfnSecretTargetAttachment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSecretTargetAttachmentProps := &cfnSecretTargetAttachmentProps{
//   	secretId: jsii.String("secretId"),
//   	targetId: jsii.String("targetId"),
//   	targetType: jsii.String("targetType"),
//   }
//
type CfnSecretTargetAttachmentProps struct {
	// The ARN or name of the secret.
	//
	// To reference a secret also created in this template, use the see [Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html) function with the secret's logical ID.
	SecretId *string `field:"required" json:"secretId" yaml:"secretId"`
	// The ID of the database or cluster.
	TargetId *string `field:"required" json:"targetId" yaml:"targetId"`
	// A string that defines the type of service or database associated with the secret.
	//
	// This value instructs Secrets Manager how to update the secret with the details of the service or database. This value must be one of the following:
	//
	// - AWS::RDS::DBInstance
	// - AWS::RDS::DBCluster
	// - AWS::Redshift::Cluster
	// - AWS::DocDB::DBInstance
	// - AWS::DocDB::DBCluster.
	TargetType *string `field:"required" json:"targetType" yaml:"targetType"`
}

