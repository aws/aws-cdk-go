package awsredshiftserverless

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnNamespace`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnNamespaceProps := &CfnNamespaceProps{
//   	NamespaceName: jsii.String("namespaceName"),
//
//   	// the properties below are optional
//   	AdminUsername: jsii.String("adminUsername"),
//   	AdminUserPassword: jsii.String("adminUserPassword"),
//   	DbName: jsii.String("dbName"),
//   	DefaultIamRoleArn: jsii.String("defaultIamRoleArn"),
//   	FinalSnapshotName: jsii.String("finalSnapshotName"),
//   	FinalSnapshotRetentionPeriod: jsii.Number(123),
//   	IamRoles: []*string{
//   		jsii.String("iamRoles"),
//   	},
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	LogExports: []*string{
//   		jsii.String("logExports"),
//   	},
//   	Namespace: &NamespaceProperty{
//   		AdminUsername: jsii.String("adminUsername"),
//   		CreationDate: jsii.String("creationDate"),
//   		DbName: jsii.String("dbName"),
//   		DefaultIamRoleArn: jsii.String("defaultIamRoleArn"),
//   		IamRoles: []*string{
//   			jsii.String("iamRoles"),
//   		},
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   		LogExports: []*string{
//   			jsii.String("logExports"),
//   		},
//   		NamespaceArn: jsii.String("namespaceArn"),
//   		NamespaceId: jsii.String("namespaceId"),
//   		NamespaceName: jsii.String("namespaceName"),
//   		Status: jsii.String("status"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnNamespaceProps struct {
	// The name of the namespace.
	//
	// Must be between 3-64 alphanumeric characters in lowercase, and it cannot be a reserved word. A list of reserved words can be found in [Reserved Words](https://docs.aws.amazon.com//redshift/latest/dg/r_pg_keywords.html) in the Amazon Redshift Database Developer Guide.
	NamespaceName *string `field:"required" json:"namespaceName" yaml:"namespaceName"`
	// The username of the administrator for the primary database created in the namespace.
	AdminUsername *string `field:"optional" json:"adminUsername" yaml:"adminUsername"`
	// The password of the administrator for the primary database created in the namespace.
	AdminUserPassword *string `field:"optional" json:"adminUserPassword" yaml:"adminUserPassword"`
	// The name of the primary database created in the namespace.
	DbName *string `field:"optional" json:"dbName" yaml:"dbName"`
	// The Amazon Resource Name (ARN) of the IAM role to set as a default in the namespace.
	DefaultIamRoleArn *string `field:"optional" json:"defaultIamRoleArn" yaml:"defaultIamRoleArn"`
	// The name of the snapshot to be created before the namespace is deleted.
	FinalSnapshotName *string `field:"optional" json:"finalSnapshotName" yaml:"finalSnapshotName"`
	// How long to retain the final snapshot.
	FinalSnapshotRetentionPeriod *float64 `field:"optional" json:"finalSnapshotRetentionPeriod" yaml:"finalSnapshotRetentionPeriod"`
	// A list of IAM roles to associate with the namespace.
	IamRoles *[]*string `field:"optional" json:"iamRoles" yaml:"iamRoles"`
	// The ID of the AWS Key Management Service key used to encrypt your data.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The types of logs the namespace can export.
	//
	// Available export types are `userlog` , `connectionlog` , and `useractivitylog` .
	LogExports *[]*string `field:"optional" json:"logExports" yaml:"logExports"`
	// `AWS::RedshiftServerless::Namespace.Namespace`.
	Namespace interface{} `field:"optional" json:"namespace" yaml:"namespace"`
	// The map of the key-value pairs used to tag the namespace.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

