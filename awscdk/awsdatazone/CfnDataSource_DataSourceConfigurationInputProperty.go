package awsdatazone


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataSourceConfigurationInputProperty := &DataSourceConfigurationInputProperty{
//   	GlueRunConfiguration: &GlueRunConfigurationInputProperty{
//   		RelationalFilterConfigurations: []interface{}{
//   			&RelationalFilterConfigurationProperty{
//   				DatabaseName: jsii.String("databaseName"),
//
//   				// the properties below are optional
//   				FilterExpressions: []interface{}{
//   					&FilterExpressionProperty{
//   						Expression: jsii.String("expression"),
//   						Type: jsii.String("type"),
//   					},
//   				},
//   				SchemaName: jsii.String("schemaName"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		DataAccessRole: jsii.String("dataAccessRole"),
//   	},
//   	RedshiftRunConfiguration: &RedshiftRunConfigurationInputProperty{
//   		RedshiftCredentialConfiguration: &RedshiftCredentialConfigurationProperty{
//   			SecretManagerArn: jsii.String("secretManagerArn"),
//   		},
//   		RedshiftStorage: &RedshiftStorageProperty{
//   			RedshiftClusterSource: &RedshiftClusterStorageProperty{
//   				ClusterName: jsii.String("clusterName"),
//   			},
//   			RedshiftServerlessSource: &RedshiftServerlessStorageProperty{
//   				WorkgroupName: jsii.String("workgroupName"),
//   			},
//   		},
//   		RelationalFilterConfigurations: []interface{}{
//   			&RelationalFilterConfigurationProperty{
//   				DatabaseName: jsii.String("databaseName"),
//
//   				// the properties below are optional
//   				FilterExpressions: []interface{}{
//   					&FilterExpressionProperty{
//   						Expression: jsii.String("expression"),
//   						Type: jsii.String("type"),
//   					},
//   				},
//   				SchemaName: jsii.String("schemaName"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		DataAccessRole: jsii.String("dataAccessRole"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-datasourceconfigurationinput.html
//
type CfnDataSource_DataSourceConfigurationInputProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-datasourceconfigurationinput.html#cfn-datazone-datasource-datasourceconfigurationinput-gluerunconfiguration
	//
	GlueRunConfiguration interface{} `field:"optional" json:"glueRunConfiguration" yaml:"glueRunConfiguration"`
	// The configuration details of the Amazon Redshift data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-datasourceconfigurationinput.html#cfn-datazone-datasource-datasourceconfigurationinput-redshiftrunconfiguration
	//
	RedshiftRunConfiguration interface{} `field:"optional" json:"redshiftRunConfiguration" yaml:"redshiftRunConfiguration"`
}

