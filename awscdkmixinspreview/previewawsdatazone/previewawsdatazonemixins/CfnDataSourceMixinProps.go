package previewawsdatazonemixins


// Properties for CfnDataSourcePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDataSourceMixinProps := &CfnDataSourceMixinProps{
//   	AssetFormsInput: []interface{}{
//   		&FormInputProperty{
//   			Content: jsii.String("content"),
//   			FormName: jsii.String("formName"),
//   			TypeIdentifier: jsii.String("typeIdentifier"),
//   			TypeRevision: jsii.String("typeRevision"),
//   		},
//   	},
//   	Configuration: &DataSourceConfigurationInputProperty{
//   		GlueRunConfiguration: &GlueRunConfigurationInputProperty{
//   			AutoImportDataQualityResult: jsii.Boolean(false),
//   			CatalogName: jsii.String("catalogName"),
//   			DataAccessRole: jsii.String("dataAccessRole"),
//   			RelationalFilterConfigurations: []interface{}{
//   				&RelationalFilterConfigurationProperty{
//   					DatabaseName: jsii.String("databaseName"),
//   					FilterExpressions: []interface{}{
//   						&FilterExpressionProperty{
//   							Expression: jsii.String("expression"),
//   							Type: jsii.String("type"),
//   						},
//   					},
//   					SchemaName: jsii.String("schemaName"),
//   				},
//   			},
//   		},
//   		RedshiftRunConfiguration: &RedshiftRunConfigurationInputProperty{
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
//   			RelationalFilterConfigurations: []interface{}{
//   				&RelationalFilterConfigurationProperty{
//   					DatabaseName: jsii.String("databaseName"),
//   					FilterExpressions: []interface{}{
//   						&FilterExpressionProperty{
//   							Expression: jsii.String("expression"),
//   							Type: jsii.String("type"),
//   						},
//   					},
//   					SchemaName: jsii.String("schemaName"),
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
//   	DomainIdentifier: jsii.String("domainIdentifier"),
//   	EnableSetting: jsii.String("enableSetting"),
//   	EnvironmentIdentifier: jsii.String("environmentIdentifier"),
//   	Name: jsii.String("name"),
//   	ProjectIdentifier: jsii.String("projectIdentifier"),
//   	PublishOnImport: jsii.Boolean(false),
//   	Recommendation: &RecommendationConfigurationProperty{
//   		EnableBusinessNameGeneration: jsii.Boolean(false),
//   	},
//   	Schedule: &ScheduleConfigurationProperty{
//   		Schedule: jsii.String("schedule"),
//   		Timezone: jsii.String("timezone"),
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-datasource.html
//
type CfnDataSourceMixinProps struct {
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
	// The ID of the Amazon DataZone domain where the data source is created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-datasource.html#cfn-datazone-datasource-domainidentifier
	//
	DomainIdentifier *string `field:"optional" json:"domainIdentifier" yaml:"domainIdentifier"`
	// Specifies whether the data source is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-datasource.html#cfn-datazone-datasource-enablesetting
	//
	EnableSetting *string `field:"optional" json:"enableSetting" yaml:"enableSetting"`
	// The unique identifier of the Amazon DataZone environment to which the data source publishes assets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-datasource.html#cfn-datazone-datasource-environmentidentifier
	//
	EnvironmentIdentifier *string `field:"optional" json:"environmentIdentifier" yaml:"environmentIdentifier"`
	// The name of the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-datasource.html#cfn-datazone-datasource-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The identifier of the Amazon DataZone project in which you want to add this data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-datasource.html#cfn-datazone-datasource-projectidentifier
	//
	ProjectIdentifier *string `field:"optional" json:"projectIdentifier" yaml:"projectIdentifier"`
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
	// The type of the data source.
	//
	// In Amazon DataZone, you can use data sources to import technical metadata of assets (data) from the source databases or data warehouses into Amazon DataZone. In the current release of Amazon DataZone, you can create and run data sources for AWS Glue and Amazon Redshift.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-datasource.html#cfn-datazone-datasource-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

