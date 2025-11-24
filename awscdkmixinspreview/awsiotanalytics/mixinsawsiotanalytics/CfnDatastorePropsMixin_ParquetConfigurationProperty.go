package mixinsawsiotanalytics


// Contains the configuration information of the Parquet format.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   parquetConfigurationProperty := &ParquetConfigurationProperty{
//   	SchemaDefinition: &SchemaDefinitionProperty{
//   		Columns: []interface{}{
//   			&ColumnProperty{
//   				Name: jsii.String("name"),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-datastore-parquetconfiguration.html
//
type CfnDatastorePropsMixin_ParquetConfigurationProperty struct {
	// Information needed to define a schema.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-datastore-parquetconfiguration.html#cfn-iotanalytics-datastore-parquetconfiguration-schemadefinition
	//
	SchemaDefinition interface{} `field:"optional" json:"schemaDefinition" yaml:"schemaDefinition"`
}

