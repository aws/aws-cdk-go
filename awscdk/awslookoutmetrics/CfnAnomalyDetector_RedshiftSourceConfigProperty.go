package awslookoutmetrics


// Provides information about the Amazon Redshift database configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   redshiftSourceConfigProperty := &RedshiftSourceConfigProperty{
//   	ClusterIdentifier: jsii.String("clusterIdentifier"),
//   	DatabaseHost: jsii.String("databaseHost"),
//   	DatabaseName: jsii.String("databaseName"),
//   	DatabasePort: jsii.Number(123),
//   	RoleArn: jsii.String("roleArn"),
//   	SecretManagerArn: jsii.String("secretManagerArn"),
//   	TableName: jsii.String("tableName"),
//   	VpcConfiguration: &VpcConfigurationProperty{
//   		SecurityGroupIdList: []*string{
//   			jsii.String("securityGroupIdList"),
//   		},
//   		SubnetIdList: []*string{
//   			jsii.String("subnetIdList"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-anomalydetector-redshiftsourceconfig.html
//
type CfnAnomalyDetector_RedshiftSourceConfigProperty struct {
	// A string identifying the Redshift cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-anomalydetector-redshiftsourceconfig.html#cfn-lookoutmetrics-anomalydetector-redshiftsourceconfig-clusteridentifier
	//
	ClusterIdentifier *string `field:"required" json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// The name of the database host.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-anomalydetector-redshiftsourceconfig.html#cfn-lookoutmetrics-anomalydetector-redshiftsourceconfig-databasehost
	//
	DatabaseHost *string `field:"required" json:"databaseHost" yaml:"databaseHost"`
	// The Redshift database name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-anomalydetector-redshiftsourceconfig.html#cfn-lookoutmetrics-anomalydetector-redshiftsourceconfig-databasename
	//
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// The port number where the database can be accessed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-anomalydetector-redshiftsourceconfig.html#cfn-lookoutmetrics-anomalydetector-redshiftsourceconfig-databaseport
	//
	DatabasePort *float64 `field:"required" json:"databasePort" yaml:"databasePort"`
	// The Amazon Resource Name (ARN) of the role providing access to the database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-anomalydetector-redshiftsourceconfig.html#cfn-lookoutmetrics-anomalydetector-redshiftsourceconfig-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The Amazon Resource Name (ARN) of the AWS Secrets Manager role.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-anomalydetector-redshiftsourceconfig.html#cfn-lookoutmetrics-anomalydetector-redshiftsourceconfig-secretmanagerarn
	//
	SecretManagerArn *string `field:"required" json:"secretManagerArn" yaml:"secretManagerArn"`
	// The table name of the Redshift database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-anomalydetector-redshiftsourceconfig.html#cfn-lookoutmetrics-anomalydetector-redshiftsourceconfig-tablename
	//
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// Contains information about the Amazon Virtual Private Cloud (VPC) configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-anomalydetector-redshiftsourceconfig.html#cfn-lookoutmetrics-anomalydetector-redshiftsourceconfig-vpcconfiguration
	//
	VpcConfiguration interface{} `field:"required" json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

