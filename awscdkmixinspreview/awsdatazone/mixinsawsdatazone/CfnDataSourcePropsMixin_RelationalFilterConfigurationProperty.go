package mixinsawsdatazone


// The relational filter configuration for the data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   relationalFilterConfigurationProperty := &RelationalFilterConfigurationProperty{
//   	DatabaseName: jsii.String("databaseName"),
//   	FilterExpressions: []interface{}{
//   		&FilterExpressionProperty{
//   			Expression: jsii.String("expression"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	SchemaName: jsii.String("schemaName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-relationalfilterconfiguration.html
//
type CfnDataSourcePropsMixin_RelationalFilterConfigurationProperty struct {
	// The database name specified in the relational filter configuration for the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-relationalfilterconfiguration.html#cfn-datazone-datasource-relationalfilterconfiguration-databasename
	//
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// The filter expressions specified in the relational filter configuration for the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-relationalfilterconfiguration.html#cfn-datazone-datasource-relationalfilterconfiguration-filterexpressions
	//
	FilterExpressions interface{} `field:"optional" json:"filterExpressions" yaml:"filterExpressions"`
	// The schema name specified in the relational filter configuration for the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-relationalfilterconfiguration.html#cfn-datazone-datasource-relationalfilterconfiguration-schemaname
	//
	SchemaName *string `field:"optional" json:"schemaName" yaml:"schemaName"`
}

