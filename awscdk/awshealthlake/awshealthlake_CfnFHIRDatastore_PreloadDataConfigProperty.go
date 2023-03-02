package awshealthlake


// Optional parameter to preload data upon creation of the Data Store.
//
// Currently, the only supported preloaded data is synthetic data generated from Synthea.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   preloadDataConfigProperty := &preloadDataConfigProperty{
//   	preloadDataType: jsii.String("preloadDataType"),
//   }
//
type CfnFHIRDatastore_PreloadDataConfigProperty struct {
	// The type of preloaded data.
	//
	// Only Synthea preloaded data is supported.
	PreloadDataType *string `field:"required" json:"preloadDataType" yaml:"preloadDataType"`
}

