package awsdatazone


// The relational filter configurations included in the configuration details of the Amazon Redshift data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   redshiftRunConfigurationInputProperty := &RedshiftRunConfigurationInputProperty{
//   	RedshiftCredentialConfiguration: &RedshiftCredentialConfigurationProperty{
//   		SecretManagerArn: jsii.String("secretManagerArn"),
//   	},
//   	RedshiftStorage: &RedshiftStorageProperty{
//   		RedshiftClusterSource: &RedshiftClusterStorageProperty{
//   			ClusterName: jsii.String("clusterName"),
//   		},
//   		RedshiftServerlessSource: &RedshiftServerlessStorageProperty{
//   			WorkgroupName: jsii.String("workgroupName"),
//   		},
//   	},
//   	RelationalFilterConfigurations: []interface{}{
//   		&RelationalFilterConfigurationProperty{
//   			DatabaseName: jsii.String("databaseName"),
//
//   			// the properties below are optional
//   			FilterExpressions: []interface{}{
//   				&FilterExpressionProperty{
//   					Expression: jsii.String("expression"),
//   					Type: jsii.String("type"),
//   				},
//   			},
//   			SchemaName: jsii.String("schemaName"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	DataAccessRole: jsii.String("dataAccessRole"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-redshiftrunconfigurationinput.html
//
type CfnDataSource_RedshiftRunConfigurationInputProperty struct {
	// The details of the credentials required to access an Amazon Redshift cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-redshiftrunconfigurationinput.html#cfn-datazone-datasource-redshiftrunconfigurationinput-redshiftcredentialconfiguration
	//
	RedshiftCredentialConfiguration interface{} `field:"required" json:"redshiftCredentialConfiguration" yaml:"redshiftCredentialConfiguration"`
	// The details of the Amazon Redshift storage as part of the configuration of an Amazon Redshift data source run.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-redshiftrunconfigurationinput.html#cfn-datazone-datasource-redshiftrunconfigurationinput-redshiftstorage
	//
	RedshiftStorage interface{} `field:"required" json:"redshiftStorage" yaml:"redshiftStorage"`
	// The relational filter configurations included in the configuration details of the AWS Glue data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-redshiftrunconfigurationinput.html#cfn-datazone-datasource-redshiftrunconfigurationinput-relationalfilterconfigurations
	//
	RelationalFilterConfigurations interface{} `field:"required" json:"relationalFilterConfigurations" yaml:"relationalFilterConfigurations"`
	// The data access role included in the configuration details of the Amazon Redshift data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-redshiftrunconfigurationinput.html#cfn-datazone-datasource-redshiftrunconfigurationinput-dataaccessrole
	//
	DataAccessRole *string `field:"optional" json:"dataAccessRole" yaml:"dataAccessRole"`
}

