package previewawsguarddutyevents


// Type definition for City.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   city := &City{
//   	CityName: []*string{
//   		jsii.String("cityName"),
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_City struct {
	// cityName property.
	//
	// Specify an array of string values to match this event if the actual value of cityName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CityName *[]*string `field:"optional" json:"cityName" yaml:"cityName"`
}

