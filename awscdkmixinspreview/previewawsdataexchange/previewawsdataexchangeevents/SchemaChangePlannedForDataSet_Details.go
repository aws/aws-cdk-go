package previewawsdataexchangeevents


// Type definition for Details.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   details := &Details{
//   	SchemaChange: &SchemaChange{
//   		Changes: []SchemaChangeItem{
//   			&SchemaChangeItem{
//   				Description: []*string{
//   					jsii.String("description"),
//   				},
//   				Name: []*string{
//   					jsii.String("name"),
//   				},
//   				Type: []*string{
//   					jsii.String("type"),
//   				},
//   			},
//   		},
//   		SchemaChangeAt: []*string{
//   			jsii.String("schemaChangeAt"),
//   		},
//   	},
//   }
//
// Experimental.
type SchemaChangePlannedForDataSet_Details struct {
	// SchemaChange property.
	//
	// Specify an array of string values to match this event if the actual value of SchemaChange is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SchemaChange *SchemaChangePlannedForDataSet_SchemaChange `field:"optional" json:"schemaChange" yaml:"schemaChange"`
}

