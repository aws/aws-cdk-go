package awslookoutmetrics


// Contains information about the Amazon Relational Database Service (RDS) configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rDSSourceConfigProperty := &rDSSourceConfigProperty{
//   	databaseHost: jsii.String("databaseHost"),
//   	databaseName: jsii.String("databaseName"),
//   	databasePort: jsii.Number(123),
//   	dbInstanceIdentifier: jsii.String("dbInstanceIdentifier"),
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
type CfnAnomalyDetector_RDSSourceConfigProperty struct {
	// The host name of the database.
	DatabaseHost *string `field:"required" json:"databaseHost" yaml:"databaseHost"`
	// The name of the RDS database.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// The port number where the database can be accessed.
	DatabasePort *float64 `field:"required" json:"databasePort" yaml:"databasePort"`
	// A string identifying the database instance.
	DbInstanceIdentifier *string `field:"required" json:"dbInstanceIdentifier" yaml:"dbInstanceIdentifier"`
	// The Amazon Resource Name (ARN) of the role.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The Amazon Resource Name (ARN) of the AWS Secrets Manager role.
	SecretManagerArn *string `field:"required" json:"secretManagerArn" yaml:"secretManagerArn"`
	// The name of the table in the database.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// An object containing information about the Amazon Virtual Private Cloud (VPC) configuration.
	VpcConfiguration interface{} `field:"required" json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

