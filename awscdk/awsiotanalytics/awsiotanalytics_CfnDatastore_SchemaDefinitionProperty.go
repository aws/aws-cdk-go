package awsiotanalytics


// Information needed to define a schema.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   schemaDefinitionProperty := &schemaDefinitionProperty{
//   	columns: []interface{}{
//   		&columnProperty{
//   			name: jsii.String("name"),
//   			type: jsii.String("type"),
//   		},
//   	},
//   }
//
type CfnDatastore_SchemaDefinitionProperty struct {
	// Specifies one or more columns that store your data.
	//
	// Each schema can have up to 100 columns. Each column can have up to 100 nested types.
	Columns interface{} `field:"optional" json:"columns" yaml:"columns"`
}

