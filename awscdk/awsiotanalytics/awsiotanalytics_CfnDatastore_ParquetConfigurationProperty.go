package awsiotanalytics


// Contains the configuration information of the Parquet format.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parquetConfigurationProperty := &parquetConfigurationProperty{
//   	schemaDefinition: &schemaDefinitionProperty{
//   		columns: []interface{}{
//   			&columnProperty{
//   				name: jsii.String("name"),
//   				type: jsii.String("type"),
//   			},
//   		},
//   	},
//   }
//
type CfnDatastore_ParquetConfigurationProperty struct {
	// Information needed to define a schema.
	SchemaDefinition interface{} `field:"optional" json:"schemaDefinition" yaml:"schemaDefinition"`
}

