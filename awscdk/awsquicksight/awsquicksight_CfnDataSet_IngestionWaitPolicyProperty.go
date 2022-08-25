package awsquicksight


// The wait policy to use when creating or updating a Dataset.
//
// The default is to wait for SPICE ingestion to finish with timeout of 36 hours.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ingestionWaitPolicyProperty := &ingestionWaitPolicyProperty{
//   	ingestionWaitTimeInHours: jsii.Number(123),
//   	waitForSpiceIngestion: jsii.Boolean(false),
//   }
//
type CfnDataSet_IngestionWaitPolicyProperty struct {
	// The maximum time (in hours) to wait for Ingestion to complete.
	//
	// Default timeout is 36 hours. Applicable only when `DataSetImportMode` mode is set to SPICE and `WaitForSpiceIngestion` is set to true.
	IngestionWaitTimeInHours *float64 `field:"optional" json:"ingestionWaitTimeInHours" yaml:"ingestionWaitTimeInHours"`
	// Wait for SPICE ingestion to finish to mark dataset creation or update as successful.
	//
	// Default (true). Applicable only when `DataSetImportMode` mode is set to SPICE.
	WaitForSpiceIngestion interface{} `field:"optional" json:"waitForSpiceIngestion" yaml:"waitForSpiceIngestion"`
}

