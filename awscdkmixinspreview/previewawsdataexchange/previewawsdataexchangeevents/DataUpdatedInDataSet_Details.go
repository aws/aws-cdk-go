package previewawsdataexchangeevents


// Type definition for Details.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   details := &Details{
//   	DataUpdate: &DataUpdate{
//   		DataUpdatedAt: []*string{
//   			jsii.String("dataUpdatedAt"),
//   		},
//   	},
//   }
//
// Experimental.
type DataUpdatedInDataSet_Details struct {
	// DataUpdate property.
	//
	// Specify an array of string values to match this event if the actual value of DataUpdate is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DataUpdate *DataUpdatedInDataSet_DataUpdate `field:"optional" json:"dataUpdate" yaml:"dataUpdate"`
}

