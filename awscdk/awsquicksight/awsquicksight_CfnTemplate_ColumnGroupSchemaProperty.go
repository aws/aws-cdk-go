package awsquicksight


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
	// `CfnTemplate.ColumnGroupSchemaProperty.ColumnGroupColumnSchemaList`.
	ColumnGroupColumnSchemaList interface{} `field:"optional" json:"columnGroupColumnSchemaList" yaml:"columnGroupColumnSchemaList"`
	// `CfnTemplate.ColumnGroupSchemaProperty.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

