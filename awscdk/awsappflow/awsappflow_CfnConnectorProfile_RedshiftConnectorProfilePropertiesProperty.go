package awsappflow


// The connector-specific profile properties when using Amazon Redshift.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   redshiftConnectorProfilePropertiesProperty := &redshiftConnectorProfilePropertiesProperty{
//   	bucketName: jsii.String("bucketName"),
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	bucketPrefix: jsii.String("bucketPrefix"),
//   	clusterIdentifier: jsii.String("clusterIdentifier"),
//   	dataApiRoleArn: jsii.String("dataApiRoleArn"),
//   	databaseName: jsii.String("databaseName"),
//   	databaseUrl: jsii.String("databaseUrl"),
//   	isRedshiftServerless: jsii.Boolean(false),
//   	workgroupName: jsii.String("workgroupName"),
//   }
//
type CfnConnectorProfile_RedshiftConnectorProfilePropertiesProperty struct {
	// A name for the associated Amazon S3 bucket.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// The Amazon Resource Name (ARN) of IAM role that grants Amazon Redshift read-only access to Amazon S3.
	//
	// For more information, and for the polices that you attach to this role, see [Allow Amazon Redshift to access your Amazon AppFlow data in Amazon S3](https://docs.aws.amazon.com/appflow/latest/userguide/security_iam_service-role-policies.html#redshift-access-s3) .
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The object key for the destination bucket in which Amazon AppFlow places the files.
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
	// The unique ID that's assigned to an Amazon Redshift cluster.
	ClusterIdentifier *string `field:"optional" json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// The Amazon Resource Name (ARN) of an IAM role that permits Amazon AppFlow to access your Amazon Redshift database through the Data API.
	//
	// For more information, and for the polices that you attach to this role, see [Allow Amazon AppFlow to access Amazon Redshift databases with the Data API](https://docs.aws.amazon.com/appflow/latest/userguide/security_iam_service-role-policies.html#access-redshift) .
	DataApiRoleArn *string `field:"optional" json:"dataApiRoleArn" yaml:"dataApiRoleArn"`
	// The name of an Amazon Redshift database.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// The JDBC URL of the Amazon Redshift cluster.
	DatabaseUrl *string `field:"optional" json:"databaseUrl" yaml:"databaseUrl"`
	// Indicates whether the connector profile defines a connection to an Amazon Redshift Serverless data warehouse.
	IsRedshiftServerless interface{} `field:"optional" json:"isRedshiftServerless" yaml:"isRedshiftServerless"`
	// The name of an Amazon Redshift workgroup.
	WorkgroupName *string `field:"optional" json:"workgroupName" yaml:"workgroupName"`
}

