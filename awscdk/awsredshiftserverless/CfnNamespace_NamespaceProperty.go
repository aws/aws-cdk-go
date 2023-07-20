package awsredshiftserverless


// A collection of database objects and users.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   namespaceProperty := &NamespaceProperty{
//   	AdminUsername: jsii.String("adminUsername"),
//   	CreationDate: jsii.String("creationDate"),
//   	DbName: jsii.String("dbName"),
//   	DefaultIamRoleArn: jsii.String("defaultIamRoleArn"),
//   	IamRoles: []*string{
//   		jsii.String("iamRoles"),
//   	},
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	LogExports: []*string{
//   		jsii.String("logExports"),
//   	},
//   	NamespaceArn: jsii.String("namespaceArn"),
//   	NamespaceId: jsii.String("namespaceId"),
//   	NamespaceName: jsii.String("namespaceName"),
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-namespace-namespace.html
//
type CfnNamespace_NamespaceProperty struct {
	// The username of the administrator for the first database created in the namespace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-namespace-namespace.html#cfn-redshiftserverless-namespace-namespace-adminusername
	//
	AdminUsername *string `field:"optional" json:"adminUsername" yaml:"adminUsername"`
	// The date of when the namespace was created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-namespace-namespace.html#cfn-redshiftserverless-namespace-namespace-creationdate
	//
	CreationDate *string `field:"optional" json:"creationDate" yaml:"creationDate"`
	// The name of the first database created in the namespace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-namespace-namespace.html#cfn-redshiftserverless-namespace-namespace-dbname
	//
	DbName *string `field:"optional" json:"dbName" yaml:"dbName"`
	// The Amazon Resource Name (ARN) of the IAM role to set as a default in the namespace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-namespace-namespace.html#cfn-redshiftserverless-namespace-namespace-defaultiamrolearn
	//
	DefaultIamRoleArn *string `field:"optional" json:"defaultIamRoleArn" yaml:"defaultIamRoleArn"`
	// A list of IAM roles to associate with the namespace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-namespace-namespace.html#cfn-redshiftserverless-namespace-namespace-iamroles
	//
	IamRoles *[]*string `field:"optional" json:"iamRoles" yaml:"iamRoles"`
	// The ID of the AWS Key Management Service key used to encrypt your data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-namespace-namespace.html#cfn-redshiftserverless-namespace-namespace-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The types of logs the namespace can export.
	//
	// Available export types are User log, Connection log, and User activity log.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-namespace-namespace.html#cfn-redshiftserverless-namespace-namespace-logexports
	//
	LogExports *[]*string `field:"optional" json:"logExports" yaml:"logExports"`
	// The Amazon Resource Name (ARN) associated with a namespace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-namespace-namespace.html#cfn-redshiftserverless-namespace-namespace-namespacearn
	//
	NamespaceArn *string `field:"optional" json:"namespaceArn" yaml:"namespaceArn"`
	// The unique identifier of a namespace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-namespace-namespace.html#cfn-redshiftserverless-namespace-namespace-namespaceid
	//
	NamespaceId *string `field:"optional" json:"namespaceId" yaml:"namespaceId"`
	// The name of the namespace.
	//
	// Must be between 3-64 alphanumeric characters in lowercase, and it cannot be a reserved word. A list of reserved words can be found in [Reserved Words](https://docs.aws.amazon.com//redshift/latest/dg/r_pg_keywords.html) in the Amazon Redshift Database Developer Guide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-namespace-namespace.html#cfn-redshiftserverless-namespace-namespace-namespacename
	//
	NamespaceName *string `field:"optional" json:"namespaceName" yaml:"namespaceName"`
	// The status of the namespace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-namespace-namespace.html#cfn-redshiftserverless-namespace-namespace-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

