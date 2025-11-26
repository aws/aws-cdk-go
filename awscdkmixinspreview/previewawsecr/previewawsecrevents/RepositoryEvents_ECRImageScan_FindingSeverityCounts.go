package previewawsecrevents


// Type definition for FindingSeverityCounts.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   findingSeverityCounts := &FindingSeverityCounts{
//   	Critical: []*string{
//   		jsii.String("critical"),
//   	},
//   	High: []*string{
//   		jsii.String("high"),
//   	},
//   	Informational: []*string{
//   		jsii.String("informational"),
//   	},
//   	Low: []*string{
//   		jsii.String("low"),
//   	},
//   	Medium: []*string{
//   		jsii.String("medium"),
//   	},
//   	Undefined: []*string{
//   		jsii.String("undefined"),
//   	},
//   }
//
// Experimental.
type RepositoryEvents_ECRImageScan_FindingSeverityCounts struct {
	// CRITICAL property.
	//
	// Specify an array of string values to match this event if the actual value of CRITICAL is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Critical *[]*string `field:"optional" json:"critical" yaml:"critical"`
	// HIGH property.
	//
	// Specify an array of string values to match this event if the actual value of HIGH is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	High *[]*string `field:"optional" json:"high" yaml:"high"`
	// INFORMATIONAL property.
	//
	// Specify an array of string values to match this event if the actual value of INFORMATIONAL is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Informational *[]*string `field:"optional" json:"informational" yaml:"informational"`
	// LOW property.
	//
	// Specify an array of string values to match this event if the actual value of LOW is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Low *[]*string `field:"optional" json:"low" yaml:"low"`
	// MEDIUM property.
	//
	// Specify an array of string values to match this event if the actual value of MEDIUM is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Medium *[]*string `field:"optional" json:"medium" yaml:"medium"`
	// UNDEFINED property.
	//
	// Specify an array of string values to match this event if the actual value of UNDEFINED is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Undefined *[]*string `field:"optional" json:"undefined" yaml:"undefined"`
}

