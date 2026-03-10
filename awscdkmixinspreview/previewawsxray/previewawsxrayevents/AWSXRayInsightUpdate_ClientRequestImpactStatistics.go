package previewawsxrayevents


// Type definition for ClientRequestImpactStatistics.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   clientRequestImpactStatistics := &ClientRequestImpactStatistics{
//   	FaultCount: []*string{
//   		jsii.String("faultCount"),
//   	},
//   	OkCount: []*string{
//   		jsii.String("okCount"),
//   	},
//   	TotalCount: []*string{
//   		jsii.String("totalCount"),
//   	},
//   }
//
// Experimental.
type AWSXRayInsightUpdate_ClientRequestImpactStatistics struct {
	// FaultCount property.
	//
	// Specify an array of string values to match this event if the actual value of FaultCount is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FaultCount *[]*string `field:"optional" json:"faultCount" yaml:"faultCount"`
	// OkCount property.
	//
	// Specify an array of string values to match this event if the actual value of OkCount is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	OkCount *[]*string `field:"optional" json:"okCount" yaml:"okCount"`
	// TotalCount property.
	//
	// Specify an array of string values to match this event if the actual value of TotalCount is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TotalCount *[]*string `field:"optional" json:"totalCount" yaml:"totalCount"`
}

