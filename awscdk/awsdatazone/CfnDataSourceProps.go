package awsdatazone


// Properties for defining a `CfnDataSource`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDataSourceProps := &CfnDataSourceProps{
//   	DomainIdentifier: jsii.String("domainIdentifier"),
//   	Name: jsii.String("name"),
//   	ProjectIdentifier: jsii.String("projectIdentifier"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	AssetFormsInput: []interface{}{
//   		&FormInputProperty{
//   			FormName: jsii.String("formName"),
//
//   			// the properties below are optional
//   			Content: jsii.String("content"),
//   			TypeIdentifier: jsii.String("typeIdentifier"),
//   			TypeRevision: jsii.String("typeRevision"),
//   		},
//   	},
//   	Configuration: &DataSourceConfigurationInputProperty{
//   		GlueRunConfiguration: &GlueRunConfigurationInputProperty{
//   			RelationalFilterConfigurations: []interface{}{
//   				&RelationalFilterConfigurationProperty{
//   					DatabaseName: jsii.String("databaseName"),
//
//   					// the properties below are optional
//   					FilterExpressions: []interface{}{
//   						&FilterExpressionProperty{
//   							Expression: jsii.String("expression"),
//   							Type: jsii.String("type"),
//   						},
//   					},
//   					SchemaName: jsii.String("schemaName"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			AutoImportDataQualityResult: jsii.Boolean(false),
//   			CatalogName: jsii.String("catalogName"),
//   			DataAccessRole: jsii.String("dataAccessRole"),
//   		},
//   		RedshiftRunConfiguration: &RedshiftRunConfigurationInputProperty{
//   			RelationalFilterConfigurations: []interface{}{
//   				&RelationalFilterConfigurationProperty{
//   					DatabaseName: jsii.String("databaseName"),
//
//   					// the properties below are optional
//   					FilterExpressions: []interface{}{
//   						&FilterExpressionProperty{
//   							Expression: jsii.String("expression"),
//   							Type: jsii.String("type"),
//   						},
//   					},
//   					SchemaName: jsii.String("schemaName"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			DataAccessRole: jsii.String("dataAccessRole"),
//   			RedshiftCredentialConfiguration: &RedshiftCredentialConfigurationProperty{
//   				SecretManagerArn: jsii.String("secretManagerArn"),
//   			},
//   			RedshiftStorage: &RedshiftStorageProperty{
//   				RedshiftClusterSource: &RedshiftClusterStorageProperty{
//   					ClusterName: jsii.String("clusterName"),
//   				},
//   				RedshiftServerlessSource: &RedshiftServerlessStorageProperty{
//   					WorkgroupName: jsii.String("workgroupName"),
//   				},
//   			},
//   		},
//   		SageMakerRunConfiguration: &SageMakerRunConfigurationInputProperty{
//   			TrackingAssets: map[string][]*string{
//   				"trackingAssetsKey": []*string{
//   					jsii.String("trackingAssets"),
//   				},
//   			},
//   		},
//   	},
//   	ConnectionIdentifier: jsii.String("connectionIdentifier"),
//   	Description: jsii.String("description"),
//   	EnableSetting: jsii.String("enableSetting"),
//   	EnvironmentIdentifier: jsii.String("environmentIdentifier"),
//   	PublishOnImport: jsii.Boolean(false),
//   	Recommendation: &RecommendationConfigurationProperty{
//   		EnableBusinessNameGeneration: jsii.Boolean(false),
//   	},
//   	Schedule: &ScheduleConfigurationProperty{
//   		Schedule: jsii.String("schedule"),
//   		Timezone: jsii.String("timezone"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-datasource.html
//
type CfnDataSourceProps struct {
	// The ID of the Amazon DataZone domain where the data source is created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-datasource.html#cfn-datazone-datasource-domainidentifier
	//
	DomainIdentifier *string `field:"required" json:"domainIdentifier" yaml:"domainIdentifier"`
	// The name of the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-datasource.html#cfn-datazone-datasource-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The identifier of the Amazon DataZone project in which you want to add this data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-datasource.html#cfn-datazone-datasource-projectidentifier
	//
	ProjectIdentifier *string `field:"required" json:"projectIdentifier" yaml:"projectIdentifier"`
	// The type of the data source.
	//
	// In Amazon DataZone, you can use data sources to import technical metadata of assets (data) from the source databases or data warehouses into Amazon DataZone. In the current release of Amazon DataZone, you can create and run data sources for AWS Glue and Amazon Redshift.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-datasource.html#cfn-datazone-datasource-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The metadata forms attached to the assets that the data source works with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-datasource.html#cfn-datazone-datasource-assetformsinput
	//
	AssetFormsInput interface{} `field:"optional" json:"assetFormsInput" yaml:"assetFormsInput"`
	// The configuration of the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-datasource.html#cfn-datazone-datasource-configuration
	//
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// The unique identifier of a connection used to fetch relevant parameters from connection during Datasource run.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-datasource.html#cfn-datazone-datasource-connectionidentifier
	//
	ConnectionIdentifier *string `field:"optional" json:"connectionIdentifier" yaml:"connectionIdentifier"`
	// The description of the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-datasource.html#cfn-datazone-datasource-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies whether the data source is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-datasource.html#cfn-datazone-datasource-enablesetting
	//
	EnableSetting *string `field:"optional" json:"enableSetting" yaml:"enableSetting"`
	// The unique identifier of the Amazon DataZone environment to which the data source publishes assets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-datasource.html#cfn-datazone-datasource-environmentidentifier
	//
	EnvironmentIdentifier *string `field:"optional" json:"environmentIdentifier" yaml:"environmentIdentifier"`
	// Specifies whether the assets that this data source creates in the inventory are to be also automatically published to the catalog.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-datasource.html#cfn-datazone-datasource-publishonimport
	//
	PublishOnImport interface{} `field:"optional" json:"publishOnImport" yaml:"publishOnImport"`
	// Specifies whether the business name generation is to be enabled for this data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-datasource.html#cfn-datazone-datasource-recommendation
	//
	Recommendation interface{} `field:"optional" json:"recommendation" yaml:"recommendation"`
	// The schedule of the data source runs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-datasource.html#cfn-datazone-datasource-schedule
	//
	Schedule interface{} `field:"optional" json:"schedule" yaml:"schedule"`
}

