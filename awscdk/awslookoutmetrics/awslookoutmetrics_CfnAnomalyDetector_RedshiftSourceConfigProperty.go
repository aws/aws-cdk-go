package awslookoutmetrics


// Provides information about the Amazon Redshift database configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   redshiftSourceConfigProperty := &redshiftSourceConfigProperty{
//   	clusterIdentifier: jsii.String("clusterIdentifier"),
//   	databaseHost: jsii.String("databaseHost"),
//   	databaseName: jsii.String("databaseName"),
//   	databasePort: jsii.Number(123),
//   	roleArn: jsii.String("roleArn"),
//   	secretManagerArn: jsii.String("secretManagerArn"),
//   	tableName: jsii.String("tableName"),
//   	vpcConfiguration: &vpcConfigurationProperty{
//   		securityGroupIdList: []*string{
//   			jsii.String("securityGroupIdList"),
//   		},
//   		subnetIdList: []*string{
//   			jsii.String("subnetIdList"),
//   		},
//   	},
//   }
//
type CfnAnomalyDetector_RedshiftSourceConfigProperty struct {
	// A string identifying the Redshift cluster.
	ClusterIdentifier *string `field:"required" json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// The name of the database host.
	DatabaseHost *string `field:"required" json:"databaseHost" yaml:"databaseHost"`
	// The Redshift database name.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// The port number where the database can be accessed.
	DatabasePort *float64 `field:"required" json:"databasePort" yaml:"databasePort"`
	// The Amazon Resource Name (ARN) of the role providing access to the database.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The Amazon Resource Name (ARN) of the AWS Secrets Manager role.
	SecretManagerArn *string `field:"required" json:"secretManagerArn" yaml:"secretManagerArn"`
	// The table name of the Redshift database.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// Contains information about the Amazon Virtual Private Cloud (VPC) configuration.
	VpcConfiguration interface{} `field:"required" json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

