package mixinsawsdatazone


// The configuration details of the AWS Glue data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   glueRunConfigurationInputProperty := &GlueRunConfigurationInputProperty{
//   	AutoImportDataQualityResult: jsii.Boolean(false),
//   	CatalogName: jsii.String("catalogName"),
//   	DataAccessRole: jsii.String("dataAccessRole"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-gluerunconfigurationinput.html
//
type CfnDataSourcePropsMixin_GlueRunConfigurationInputProperty struct {
	// Specifies whether to automatically import data quality metrics as part of the data source run.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-gluerunconfigurationinput.html#cfn-datazone-datasource-gluerunconfigurationinput-autoimportdataqualityresult
	//
	AutoImportDataQualityResult interface{} `field:"optional" json:"autoImportDataQualityResult" yaml:"autoImportDataQualityResult"`
	// The catalog name in the AWS Glue run configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-gluerunconfigurationinput.html#cfn-datazone-datasource-gluerunconfigurationinput-catalogname
	//
	CatalogName *string `field:"optional" json:"catalogName" yaml:"catalogName"`
	// The data access role included in the configuration details of the AWS Glue data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-gluerunconfigurationinput.html#cfn-datazone-datasource-gluerunconfigurationinput-dataaccessrole
	//
	DataAccessRole *string `field:"optional" json:"dataAccessRole" yaml:"dataAccessRole"`
	// The relational filter configurations included in the configuration details of the AWS Glue data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-gluerunconfigurationinput.html#cfn-datazone-datasource-gluerunconfigurationinput-relationalfilterconfigurations
	//
	RelationalFilterConfigurations interface{} `field:"optional" json:"relationalFilterConfigurations" yaml:"relationalFilterConfigurations"`
}

