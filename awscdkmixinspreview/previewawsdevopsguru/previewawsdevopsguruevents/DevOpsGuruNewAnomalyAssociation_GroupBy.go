package previewawsdevopsguruevents


// Type definition for GroupBy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   groupBy := &GroupBy{
//   	Dimensions: []*string{
//   		jsii.String("dimensions"),
//   	},
//   	Group: []*string{
//   		jsii.String("group"),
//   	},
//   }
//
// Experimental.
type DevOpsGuruNewAnomalyAssociation_GroupBy struct {
	// dimensions property.
	//
	// Specify an array of string values to match this event if the actual value of dimensions is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Dimensions *[]*string `field:"optional" json:"dimensions" yaml:"dimensions"`
	// group property.
	//
	// Specify an array of string values to match this event if the actual value of group is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Group *[]*string `field:"optional" json:"group" yaml:"group"`
}

