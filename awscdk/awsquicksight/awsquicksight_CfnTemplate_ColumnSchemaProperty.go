package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   columnSchemaProperty := &columnSchemaProperty{
//   	dataType: jsii.String("dataType"),
//   	geographicRole: jsii.String("geographicRole"),
//   	name: jsii.String("name"),
//   }
//
type CfnTemplate_ColumnSchemaProperty struct {
	// `CfnTemplate.ColumnSchemaProperty.DataType`.
	DataType *string `field:"optional" json:"dataType" yaml:"dataType"`
	// `CfnTemplate.ColumnSchemaProperty.GeographicRole`.
	GeographicRole *string `field:"optional" json:"geographicRole" yaml:"geographicRole"`
	// `CfnTemplate.ColumnSchemaProperty.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

