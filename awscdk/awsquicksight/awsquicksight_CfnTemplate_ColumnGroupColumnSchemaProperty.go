package awsquicksight


// A structure describing the name, data type, and geographic role of the columns.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   columnGroupColumnSchemaProperty := &columnGroupColumnSchemaProperty{
//   	name: jsii.String("name"),
//   }
//
type CfnTemplate_ColumnGroupColumnSchemaProperty struct {
	// The name of the column group's column schema.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

