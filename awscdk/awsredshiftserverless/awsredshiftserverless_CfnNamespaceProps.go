package awsredshiftserverless

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnNamespace`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnNamespaceProps := &cfnNamespaceProps{
//   	namespaceName: jsii.String("namespaceName"),
//
//   	// the properties below are optional
//   	adminUsername: jsii.String("adminUsername"),
//   	adminUserPassword: jsii.String("adminUserPassword"),
//   	dbName: jsii.String("dbName"),
//   	defaultIamRoleArn: jsii.String("defaultIamRoleArn"),
//   	finalSnapshotName: jsii.String("finalSnapshotName"),
//   	finalSnapshotRetentionPeriod: jsii.Number(123),
//   	iamRoles: []*string{
//   		jsii.String("iamRoles"),
//   	},
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	logExports: []*string{
//   		jsii.String("logExports"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnNamespaceProps struct {
	// `AWS::RedshiftServerless::Namespace.NamespaceName`.
	NamespaceName *string `field:"required" json:"namespaceName" yaml:"namespaceName"`
	// `AWS::RedshiftServerless::Namespace.AdminUsername`.
	AdminUsername *string `field:"optional" json:"adminUsername" yaml:"adminUsername"`
	// `AWS::RedshiftServerless::Namespace.AdminUserPassword`.
	AdminUserPassword *string `field:"optional" json:"adminUserPassword" yaml:"adminUserPassword"`
	// `AWS::RedshiftServerless::Namespace.DbName`.
	DbName *string `field:"optional" json:"dbName" yaml:"dbName"`
	// `AWS::RedshiftServerless::Namespace.DefaultIamRoleArn`.
	DefaultIamRoleArn *string `field:"optional" json:"defaultIamRoleArn" yaml:"defaultIamRoleArn"`
	// `AWS::RedshiftServerless::Namespace.FinalSnapshotName`.
	FinalSnapshotName *string `field:"optional" json:"finalSnapshotName" yaml:"finalSnapshotName"`
	// `AWS::RedshiftServerless::Namespace.FinalSnapshotRetentionPeriod`.
	FinalSnapshotRetentionPeriod *float64 `field:"optional" json:"finalSnapshotRetentionPeriod" yaml:"finalSnapshotRetentionPeriod"`
	// `AWS::RedshiftServerless::Namespace.IamRoles`.
	IamRoles *[]*string `field:"optional" json:"iamRoles" yaml:"iamRoles"`
	// `AWS::RedshiftServerless::Namespace.KmsKeyId`.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// `AWS::RedshiftServerless::Namespace.LogExports`.
	LogExports *[]*string `field:"optional" json:"logExports" yaml:"logExports"`
	// `AWS::RedshiftServerless::Namespace.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

