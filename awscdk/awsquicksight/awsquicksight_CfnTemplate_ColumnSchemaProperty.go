package awsquicksight


// The column schema.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   columnSchemaProperty := &ColumnSchemaProperty{
//   	DataType: jsii.String("dataType"),
//   	GeographicRole: jsii.String("geographicRole"),
//   	Name: jsii.String("name"),
//   }
//
type CfnTemplate_ColumnSchemaProperty struct {
	// The data type of the column schema.
	DataType *string `field:"optional" json:"dataType" yaml:"dataType"`
	// The geographic role of the column schema.
	GeographicRole *string `field:"optional" json:"geographicRole" yaml:"geographicRole"`
	// The name of the column schema.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

