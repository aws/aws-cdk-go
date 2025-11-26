package previewawsdatazonemixins


// The relational filter configurations included in the configuration details of the Amazon Redshift data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   redshiftRunConfigurationInputProperty := &RedshiftRunConfigurationInputProperty{
//   	DataAccessRole: jsii.String("dataAccessRole"),
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
//   			FilterExpressions: []interface{}{
//   				&FilterExpressionProperty{
//   					Expression: jsii.String("expression"),
//   					Type: jsii.String("type"),
//   				},
//   			},
//   			SchemaName: jsii.String("schemaName"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-redshiftrunconfigurationinput.html
//
type CfnDataSourcePropsMixin_RedshiftRunConfigurationInputProperty struct {
	// The data access role included in the configuration details of the Amazon Redshift data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-redshiftrunconfigurationinput.html#cfn-datazone-datasource-redshiftrunconfigurationinput-dataaccessrole
	//
	DataAccessRole *string `field:"optional" json:"dataAccessRole" yaml:"dataAccessRole"`
	// The details of the credentials required to access an Amazon Redshift cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-redshiftrunconfigurationinput.html#cfn-datazone-datasource-redshiftrunconfigurationinput-redshiftcredentialconfiguration
	//
	RedshiftCredentialConfiguration interface{} `field:"optional" json:"redshiftCredentialConfiguration" yaml:"redshiftCredentialConfiguration"`
	// The details of the Amazon Redshift storage as part of the configuration of an Amazon Redshift data source run.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-redshiftrunconfigurationinput.html#cfn-datazone-datasource-redshiftrunconfigurationinput-redshiftstorage
	//
	RedshiftStorage interface{} `field:"optional" json:"redshiftStorage" yaml:"redshiftStorage"`
	// The relational filter configurations included in the configuration details of the AWS Glue data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-redshiftrunconfigurationinput.html#cfn-datazone-datasource-redshiftrunconfigurationinput-relationalfilterconfigurations
	//
	RelationalFilterConfigurations interface{} `field:"optional" json:"relationalFilterConfigurations" yaml:"relationalFilterConfigurations"`
}

