package previewawsdataexchangeevents


// Type definition for SchemaChangeItem.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   schemaChangeItem := &SchemaChangeItem{
//   	Description: []*string{
//   		jsii.String("description"),
//   	},
//   	Name: []*string{
//   		jsii.String("name"),
//   	},
//   	Type: []*string{
//   		jsii.String("type"),
//   	},
//   }
//
// Experimental.
type SchemaChangePlannedForDataSet_SchemaChangeItem struct {
	// Description property.
	//
	// Specify an array of string values to match this event if the actual value of Description is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Description *[]*string `field:"optional" json:"description" yaml:"description"`
	// Name property.
	//
	// Specify an array of string values to match this event if the actual value of Name is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Name *[]*string `field:"optional" json:"name" yaml:"name"`
	// Type property.
	//
	// Specify an array of string values to match this event if the actual value of Type is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Type *[]*string `field:"optional" json:"type" yaml:"type"`
}

