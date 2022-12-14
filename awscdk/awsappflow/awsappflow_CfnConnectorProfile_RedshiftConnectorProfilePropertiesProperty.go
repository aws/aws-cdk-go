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
	// The Amazon Resource Name (ARN) of the IAM role.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The object key for the destination bucket in which Amazon AppFlow places the files.
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
	// `CfnConnectorProfile.RedshiftConnectorProfilePropertiesProperty.ClusterIdentifier`.
	ClusterIdentifier *string `field:"optional" json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// `CfnConnectorProfile.RedshiftConnectorProfilePropertiesProperty.DataApiRoleArn`.
	DataApiRoleArn *string `field:"optional" json:"dataApiRoleArn" yaml:"dataApiRoleArn"`
	// `CfnConnectorProfile.RedshiftConnectorProfilePropertiesProperty.DatabaseName`.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// The JDBC URL of the Amazon Redshift cluster.
	DatabaseUrl *string `field:"optional" json:"databaseUrl" yaml:"databaseUrl"`
	// `CfnConnectorProfile.RedshiftConnectorProfilePropertiesProperty.IsRedshiftServerless`.
	IsRedshiftServerless interface{} `field:"optional" json:"isRedshiftServerless" yaml:"isRedshiftServerless"`
	// `CfnConnectorProfile.RedshiftConnectorProfilePropertiesProperty.WorkgroupName`.
	WorkgroupName *string `field:"optional" json:"workgroupName" yaml:"workgroupName"`
}

