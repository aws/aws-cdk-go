package previewawsguarddutyevents


// Type definition for Country_2.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   country2 := &Country2{
//   	CountryName: []*string{
//   		jsii.String("countryName"),
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_Country2 struct {
	// countryName property.
	//
	// Specify an array of string values to match this event if the actual value of countryName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CountryName *[]*string `field:"optional" json:"countryName" yaml:"countryName"`
}

