package awssecretsmanager


// Properties for defining a `CfnSecretTargetAttachment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSecretTargetAttachmentProps := &CfnSecretTargetAttachmentProps{
//   	SecretId: jsii.String("secretId"),
//   	TargetId: jsii.String("targetId"),
//   	TargetType: jsii.String("targetType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-secretsmanager-secrettargetattachment.html
//
type CfnSecretTargetAttachmentProps struct {
	// The ARN or name of the secret.
	//
	// To reference a secret also created in this template, use the see [Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html) function with the secret's logical ID. This field is unique for each target attachment definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-secretsmanager-secrettargetattachment.html#cfn-secretsmanager-secrettargetattachment-secretid
	//
	SecretId *string `field:"required" json:"secretId" yaml:"secretId"`
	// The ID of the database or cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-secretsmanager-secrettargetattachment.html#cfn-secretsmanager-secrettargetattachment-targetid
	//
	TargetId *string `field:"required" json:"targetId" yaml:"targetId"`
	// A string that defines the type of service or database associated with the secret.
	//
	// This value instructs Secrets Manager how to update the secret with the details of the service or database. This value must be one of the following:
	//
	// - AWS::RDS::DBInstance
	// - AWS::RDS::DBCluster
	// - AWS::Redshift::Cluster
	// - AWS::RedshiftServerless::Namespace
	// - AWS::DocDB::DBInstance
	// - AWS::DocDB::DBCluster
	// - AWS::DocDBElastic::Cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-secretsmanager-secrettargetattachment.html#cfn-secretsmanager-secrettargetattachment-targettype
	//
	TargetType *string `field:"required" json:"targetType" yaml:"targetType"`
}

