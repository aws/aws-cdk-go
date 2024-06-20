package awsquicksight


// The parameters for Amazon Redshift.
//
// The `ClusterId` field can be blank if `Host` and `Port` are both set. The `Host` and `Port` fields can be blank if the `ClusterId` field is set.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   redshiftParametersProperty := &RedshiftParametersProperty{
//   	Database: jsii.String("database"),
//
//   	// the properties below are optional
//   	ClusterId: jsii.String("clusterId"),
//   	Host: jsii.String("host"),
//   	IamParameters: &RedshiftIAMParametersProperty{
//   		RoleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		AutoCreateDatabaseUser: jsii.Boolean(false),
//   		DatabaseGroups: []*string{
//   			jsii.String("databaseGroups"),
//   		},
//   		DatabaseUser: jsii.String("databaseUser"),
//   	},
//   	IdentityCenterConfiguration: &IdentityCenterConfigurationProperty{
//   		EnableIdentityPropagation: jsii.Boolean(false),
//   	},
//   	Port: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-redshiftparameters.html
//
type CfnDataSource_RedshiftParametersProperty struct {
	// Database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-redshiftparameters.html#cfn-quicksight-datasource-redshiftparameters-database
	//
	Database *string `field:"required" json:"database" yaml:"database"`
	// Cluster ID.
	//
	// This field can be blank if the `Host` and `Port` are provided.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-redshiftparameters.html#cfn-quicksight-datasource-redshiftparameters-clusterid
	//
	ClusterId *string `field:"optional" json:"clusterId" yaml:"clusterId"`
	// Host.
	//
	// This field can be blank if `ClusterId` is provided.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-redshiftparameters.html#cfn-quicksight-datasource-redshiftparameters-host
	//
	Host *string `field:"optional" json:"host" yaml:"host"`
	// An optional parameter that uses IAM authentication to grant Amazon QuickSight access to your cluster.
	//
	// This parameter can be used instead of [DataSourceCredentials](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_DataSourceCredentials.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-redshiftparameters.html#cfn-quicksight-datasource-redshiftparameters-iamparameters
	//
	IamParameters interface{} `field:"optional" json:"iamParameters" yaml:"iamParameters"`
	// An optional parameter that configures IAM Identity Center authentication to grant Amazon QuickSight access to your cluster.
	//
	// This parameter can only be specified if your Amazon QuickSight account is configured with IAM Identity Center.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-redshiftparameters.html#cfn-quicksight-datasource-redshiftparameters-identitycenterconfiguration
	//
	IdentityCenterConfiguration interface{} `field:"optional" json:"identityCenterConfiguration" yaml:"identityCenterConfiguration"`
	// Port.
	//
	// This field can be blank if the `ClusterId` is provided.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-redshiftparameters.html#cfn-quicksight-datasource-redshiftparameters-port
	//
	// Default: - 0.
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

