package awsbcmdataexports


// The cadence for AWS to update the data export in your S3 bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   refreshCadenceProperty := &RefreshCadenceProperty{
//   	Frequency: jsii.String("frequency"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcmdataexports-export-refreshcadence.html
//
type CfnExport_RefreshCadenceProperty struct {
	// The frequency that data exports are updated.
	//
	// The export refreshes each time the source data updates, up to three times daily.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcmdataexports-export-refreshcadence.html#cfn-bcmdataexports-export-refreshcadence-frequency
	//
	Frequency *string `field:"required" json:"frequency" yaml:"frequency"`
}

