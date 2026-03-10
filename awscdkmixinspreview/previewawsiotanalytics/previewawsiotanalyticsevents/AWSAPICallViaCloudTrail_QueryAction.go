package previewawsiotanalyticsevents


// Type definition for QueryAction.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   queryAction := &QueryAction{
//   	SqlQuery: []*string{
//   		jsii.String("sqlQuery"),
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_QueryAction struct {
	// sqlQuery property.
	//
	// Specify an array of string values to match this event if the actual value of sqlQuery is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SqlQuery *[]*string `field:"optional" json:"sqlQuery" yaml:"sqlQuery"`
}

