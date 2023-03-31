package awsquicksight


// The column group schema.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   columnGroupSchemaProperty := &columnGroupSchemaProperty{
//   	columnGroupColumnSchemaList: []interface{}{
//   		&columnGroupColumnSchemaProperty{
//   			name: jsii.String("name"),
//   		},
//   	},
//   	name: jsii.String("name"),
//   }
//
type CfnTemplate_ColumnGroupSchemaProperty struct {
	// A structure containing the list of schemas for column group columns.
	ColumnGroupColumnSchemaList interface{} `field:"optional" json:"columnGroupColumnSchemaList" yaml:"columnGroupColumnSchemaList"`
	// The name of the column group schema.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

