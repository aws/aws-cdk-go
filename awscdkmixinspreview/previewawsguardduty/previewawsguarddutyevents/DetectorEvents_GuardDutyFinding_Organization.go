package previewawsguarddutyevents


// Type definition for Organization.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   organization := &Organization{
//   	Asn: []*string{
//   		jsii.String("asn"),
//   	},
//   	AsnOrg: []*string{
//   		jsii.String("asnOrg"),
//   	},
//   	Isp: []*string{
//   		jsii.String("isp"),
//   	},
//   	Org: []*string{
//   		jsii.String("org"),
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_Organization struct {
	// asn property.
	//
	// Specify an array of string values to match this event if the actual value of asn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Asn *[]*string `field:"optional" json:"asn" yaml:"asn"`
	// asnOrg property.
	//
	// Specify an array of string values to match this event if the actual value of asnOrg is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AsnOrg *[]*string `field:"optional" json:"asnOrg" yaml:"asnOrg"`
	// isp property.
	//
	// Specify an array of string values to match this event if the actual value of isp is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Isp *[]*string `field:"optional" json:"isp" yaml:"isp"`
	// org property.
	//
	// Specify an array of string values to match this event if the actual value of org is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Org *[]*string `field:"optional" json:"org" yaml:"org"`
}

