package awsredshiftserverless


// A collection of database objects and users.
//
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
	// The username of the administrator for the first database created in the namespace.
	AdminUsername *string `field:"optional" json:"adminUsername" yaml:"adminUsername"`
	// The date of when the namespace was created.
	CreationDate *string `field:"optional" json:"creationDate" yaml:"creationDate"`
	// The name of the first database created in the namespace.
	DbName *string `field:"optional" json:"dbName" yaml:"dbName"`
	// The Amazon Resource Name (ARN) of the IAM role to set as a default in the namespace.
	DefaultIamRoleArn *string `field:"optional" json:"defaultIamRoleArn" yaml:"defaultIamRoleArn"`
	// A list of IAM roles to associate with the namespace.
	IamRoles *[]*string `field:"optional" json:"iamRoles" yaml:"iamRoles"`
	// The ID of the AWS Key Management Service key used to encrypt your data.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The types of logs the namespace can export.
	//
	// Available export types are User log, Connection log, and User activity log.
	LogExports *[]*string `field:"optional" json:"logExports" yaml:"logExports"`
	// The Amazon Resource Name (ARN) associated with a namespace.
	NamespaceArn *string `field:"optional" json:"namespaceArn" yaml:"namespaceArn"`
	// The unique identifier of a namespace.
	NamespaceId *string `field:"optional" json:"namespaceId" yaml:"namespaceId"`
	// The name of the namespace.
	//
	// Must be between 3-64 alphanumeric characters in lowercase, and it cannot be a reserved word. A list of reserved words can be found in [Reserved Words](https://docs.aws.amazon.com//redshift/latest/dg/r_pg_keywords.html) in the Amazon Redshift Database Developer Guide.
	NamespaceName *string `field:"optional" json:"namespaceName" yaml:"namespaceName"`
	// The status of the namespace.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

