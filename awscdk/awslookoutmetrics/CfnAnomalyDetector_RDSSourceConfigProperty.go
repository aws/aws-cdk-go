package awslookoutmetrics


// Contains information about the Amazon Relational Database Service (RDS) configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rDSSourceConfigProperty := &RDSSourceConfigProperty{
//   	DatabaseHost: jsii.String("databaseHost"),
//   	DatabaseName: jsii.String("databaseName"),
//   	DatabasePort: jsii.Number(123),
//   	DbInstanceIdentifier: jsii.String("dbInstanceIdentifier"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-anomalydetector-rdssourceconfig.html
//
type CfnAnomalyDetector_RDSSourceConfigProperty struct {
	// The host name of the database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-anomalydetector-rdssourceconfig.html#cfn-lookoutmetrics-anomalydetector-rdssourceconfig-databasehost
	//
	DatabaseHost *string `field:"required" json:"databaseHost" yaml:"databaseHost"`
	// The name of the RDS database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-anomalydetector-rdssourceconfig.html#cfn-lookoutmetrics-anomalydetector-rdssourceconfig-databasename
	//
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// The port number where the database can be accessed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-anomalydetector-rdssourceconfig.html#cfn-lookoutmetrics-anomalydetector-rdssourceconfig-databaseport
	//
	DatabasePort *float64 `field:"required" json:"databasePort" yaml:"databasePort"`
	// A string identifying the database instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-anomalydetector-rdssourceconfig.html#cfn-lookoutmetrics-anomalydetector-rdssourceconfig-dbinstanceidentifier
	//
	DbInstanceIdentifier *string `field:"required" json:"dbInstanceIdentifier" yaml:"dbInstanceIdentifier"`
	// The Amazon Resource Name (ARN) of the role.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-anomalydetector-rdssourceconfig.html#cfn-lookoutmetrics-anomalydetector-rdssourceconfig-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The Amazon Resource Name (ARN) of the AWS Secrets Manager role.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-anomalydetector-rdssourceconfig.html#cfn-lookoutmetrics-anomalydetector-rdssourceconfig-secretmanagerarn
	//
	SecretManagerArn *string `field:"required" json:"secretManagerArn" yaml:"secretManagerArn"`
	// The name of the table in the database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-anomalydetector-rdssourceconfig.html#cfn-lookoutmetrics-anomalydetector-rdssourceconfig-tablename
	//
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// An object containing information about the Amazon Virtual Private Cloud (VPC) configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-anomalydetector-rdssourceconfig.html#cfn-lookoutmetrics-anomalydetector-rdssourceconfig-vpcconfiguration
	//
	VpcConfiguration interface{} `field:"required" json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

