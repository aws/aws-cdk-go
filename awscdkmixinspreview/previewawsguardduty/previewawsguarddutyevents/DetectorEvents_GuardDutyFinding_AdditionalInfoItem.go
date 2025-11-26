package previewawsguarddutyevents


// Type definition for AdditionalInfoItem.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   additionalInfoItem := &AdditionalInfoItem{
//   	Count: []*string{
//   		jsii.String("count"),
//   	},
//   	FirstSeen: []*string{
//   		jsii.String("firstSeen"),
//   	},
//   	LastSeen: []*string{
//   		jsii.String("lastSeen"),
//   	},
//   	Name: []*string{
//   		jsii.String("name"),
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_AdditionalInfoItem struct {
	// count property.
	//
	// Specify an array of string values to match this event if the actual value of count is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Count *[]*string `field:"optional" json:"count" yaml:"count"`
	// firstSeen property.
	//
	// Specify an array of string values to match this event if the actual value of firstSeen is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FirstSeen *[]*string `field:"optional" json:"firstSeen" yaml:"firstSeen"`
	// lastSeen property.
	//
	// Specify an array of string values to match this event if the actual value of lastSeen is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LastSeen *[]*string `field:"optional" json:"lastSeen" yaml:"lastSeen"`
	// name property.
	//
	// Specify an array of string values to match this event if the actual value of name is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Name *[]*string `field:"optional" json:"name" yaml:"name"`
}

