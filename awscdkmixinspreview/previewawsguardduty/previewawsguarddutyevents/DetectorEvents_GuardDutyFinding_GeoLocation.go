package previewawsguarddutyevents


// Type definition for GeoLocation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   geoLocation := &GeoLocation{
//   	Lat: []*string{
//   		jsii.String("lat"),
//   	},
//   	Lon: []*string{
//   		jsii.String("lon"),
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_GeoLocation struct {
	// lat property.
	//
	// Specify an array of string values to match this event if the actual value of lat is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Lat *[]*string `field:"optional" json:"lat" yaml:"lat"`
	// lon property.
	//
	// Specify an array of string values to match this event if the actual value of lon is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Lon *[]*string `field:"optional" json:"lon" yaml:"lon"`
}

