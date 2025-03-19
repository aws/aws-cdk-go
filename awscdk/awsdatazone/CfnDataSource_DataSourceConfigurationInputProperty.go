package awsdatazone


// The configuration of the data source.
//
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
//   		AutoImportDataQualityResult: jsii.Boolean(false),
//   		CatalogName: jsii.String("catalogName"),
//   		DataAccessRole: jsii.String("dataAccessRole"),
//   	},
//   	RedshiftRunConfiguration: &RedshiftRunConfigurationInputProperty{
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
//   	},
//   	SageMakerRunConfiguration: &SageMakerRunConfigurationInputProperty{
//   		TrackingAssets: map[string][]*string{
//   			"trackingAssetsKey": []*string{
//   				jsii.String("trackingAssets"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-datasourceconfigurationinput.html
//
type CfnDataSource_DataSourceConfigurationInputProperty struct {
	// The configuration of the AWS Glue data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-datasourceconfigurationinput.html#cfn-datazone-datasource-datasourceconfigurationinput-gluerunconfiguration
	//
	GlueRunConfiguration interface{} `field:"optional" json:"glueRunConfiguration" yaml:"glueRunConfiguration"`
	// The configuration of the Amazon Redshift data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-datasourceconfigurationinput.html#cfn-datazone-datasource-datasourceconfigurationinput-redshiftrunconfiguration
	//
	RedshiftRunConfiguration interface{} `field:"optional" json:"redshiftRunConfiguration" yaml:"redshiftRunConfiguration"`
	// The configuration details of the Amazon SageMaker data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-datasourceconfigurationinput.html#cfn-datazone-datasource-datasourceconfigurationinput-sagemakerrunconfiguration
	//
	SageMakerRunConfiguration interface{} `field:"optional" json:"sageMakerRunConfiguration" yaml:"sageMakerRunConfiguration"`
}

