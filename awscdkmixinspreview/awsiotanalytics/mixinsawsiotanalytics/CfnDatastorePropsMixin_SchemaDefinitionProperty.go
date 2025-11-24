package mixinsawsiotanalytics


// Information needed to define a schema.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   schemaDefinitionProperty := &SchemaDefinitionProperty{
//   	Columns: []interface{}{
//   		&ColumnProperty{
//   			Name: jsii.String("name"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-datastore-schemadefinition.html
//
type CfnDatastorePropsMixin_SchemaDefinitionProperty struct {
	// Specifies one or more columns that store your data.
	//
	// Each schema can have up to 100 columns. Each column can have up to 100 nested types.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-datastore-schemadefinition.html#cfn-iotanalytics-datastore-schemadefinition-columns
	//
	Columns interface{} `field:"optional" json:"columns" yaml:"columns"`
}

