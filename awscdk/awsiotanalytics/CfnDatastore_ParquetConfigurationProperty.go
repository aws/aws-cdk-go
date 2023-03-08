package awsiotanalytics


// Contains the configuration information of the Parquet format.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
type CfnDatastore_ParquetConfigurationProperty struct {
	// Information needed to define a schema.
	SchemaDefinition interface{} `field:"optional" json:"schemaDefinition" yaml:"schemaDefinition"`
}

