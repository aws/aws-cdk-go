package awsquicksight


// Dataset schema.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataSetSchemaProperty := &dataSetSchemaProperty{
//   	columnSchemaList: []interface{}{
//   		&columnSchemaProperty{
//   			dataType: jsii.String("dataType"),
//   			geographicRole: jsii.String("geographicRole"),
//   			name: jsii.String("name"),
//   		},
//   	},
//   }
//
type CfnTemplate_DataSetSchemaProperty struct {
	// A structure containing the list of column schemas.
	ColumnSchemaList interface{} `field:"optional" json:"columnSchemaList" yaml:"columnSchemaList"`
}

