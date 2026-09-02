package previewawsdataexchangeevents


// Type definition for DataUpdate.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataUpdate := &DataUpdate{
//   	DataUpdatedAt: []*string{
//   		jsii.String("dataUpdatedAt"),
//   	},
//   }
//
// Experimental.
type DataUpdatedInDataSet_DataUpdate struct {
	// DataUpdatedAt property.
	//
	// Specify an array of string values to match this event if the actual value of DataUpdatedAt is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DataUpdatedAt *[]*string `field:"optional" json:"dataUpdatedAt" yaml:"dataUpdatedAt"`
}

