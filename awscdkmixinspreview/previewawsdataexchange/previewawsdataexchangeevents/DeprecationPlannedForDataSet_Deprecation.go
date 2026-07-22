package previewawsdataexchangeevents


// Type definition for Deprecation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   deprecation := &Deprecation{
//   	DeprecationAt: []*string{
//   		jsii.String("deprecationAt"),
//   	},
//   }
//
// Experimental.
type DeprecationPlannedForDataSet_Deprecation struct {
	// DeprecationAt property.
	//
	// Specify an array of string values to match this event if the actual value of DeprecationAt is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DeprecationAt *[]*string `field:"optional" json:"deprecationAt" yaml:"deprecationAt"`
}

