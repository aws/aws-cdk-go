package awssagemaker


// External MySQL-compatible accounting database that a Slurm cluster's slurmdbd connects to.
//
// Database credentials are supplied out-of-band through the referenced Secrets Manager secret. Supported only with Continuous node provisioning.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clusterAccountingDatabaseProperty := &ClusterAccountingDatabaseProperty{
//   	Endpoint: jsii.String("endpoint"),
//   	SecretArn: jsii.String("secretArn"),
//
//   	// the properties below are optional
//   	Name: jsii.String("name"),
//   	Port: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusteraccountingdatabase.html
//
type CfnCluster_ClusterAccountingDatabaseProperty struct {
	// Hostname or endpoint of the accounting database, such as an RDS endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusteraccountingdatabase.html#cfn-sagemaker-cluster-clusteraccountingdatabase-endpoint
	//
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// ARN of the Secrets Manager secret holding the database credentials.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusteraccountingdatabase.html#cfn-sagemaker-cluster-clusteraccountingdatabase-secretarn
	//
	SecretArn *string `field:"required" json:"secretArn" yaml:"secretArn"`
	// Name of the accounting database schema.
	//
	// Defaults to slurm_acct_db when omitted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusteraccountingdatabase.html#cfn-sagemaker-cluster-clusteraccountingdatabase-name
	//
	// Default: - "slurm_acct_db".
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// TCP port of the accounting database.
	//
	// Defaults to 3306 when omitted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusteraccountingdatabase.html#cfn-sagemaker-cluster-clusteraccountingdatabase-port
	//
	// Default: - 3306.
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

