package awsredshiftserverless


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   namespaceProperty := &namespaceProperty{
//   	adminUsername: jsii.String("adminUsername"),
//   	creationDate: jsii.String("creationDate"),
//   	dbName: jsii.String("dbName"),
//   	defaultIamRoleArn: jsii.String("defaultIamRoleArn"),
//   	iamRoles: []*string{
//   		jsii.String("iamRoles"),
//   	},
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	logExports: []*string{
//   		jsii.String("logExports"),
//   	},
//   	namespaceArn: jsii.String("namespaceArn"),
//   	namespaceId: jsii.String("namespaceId"),
//   	namespaceName: jsii.String("namespaceName"),
//   	status: jsii.String("status"),
//   }
//
type CfnNamespace_NamespaceProperty struct {
	// `CfnNamespace.NamespaceProperty.AdminUsername`.
	AdminUsername *string `field:"optional" json:"adminUsername" yaml:"adminUsername"`
	// `CfnNamespace.NamespaceProperty.CreationDate`.
	CreationDate *string `field:"optional" json:"creationDate" yaml:"creationDate"`
	// `CfnNamespace.NamespaceProperty.DbName`.
	DbName *string `field:"optional" json:"dbName" yaml:"dbName"`
	// `CfnNamespace.NamespaceProperty.DefaultIamRoleArn`.
	DefaultIamRoleArn *string `field:"optional" json:"defaultIamRoleArn" yaml:"defaultIamRoleArn"`
	// `CfnNamespace.NamespaceProperty.IamRoles`.
	IamRoles *[]*string `field:"optional" json:"iamRoles" yaml:"iamRoles"`
	// `CfnNamespace.NamespaceProperty.KmsKeyId`.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// `CfnNamespace.NamespaceProperty.LogExports`.
	LogExports *[]*string `field:"optional" json:"logExports" yaml:"logExports"`
	// `CfnNamespace.NamespaceProperty.NamespaceArn`.
	NamespaceArn *string `field:"optional" json:"namespaceArn" yaml:"namespaceArn"`
	// `CfnNamespace.NamespaceProperty.NamespaceId`.
	NamespaceId *string `field:"optional" json:"namespaceId" yaml:"namespaceId"`
	// `CfnNamespace.NamespaceProperty.NamespaceName`.
	NamespaceName *string `field:"optional" json:"namespaceName" yaml:"namespaceName"`
	// `CfnNamespace.NamespaceProperty.Status`.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

