package awsvpclattice


// Describes the conditions that can be applied when matching a path for incoming requests.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pathMatchProperty := &PathMatchProperty{
//   	Match: &PathMatchTypeProperty{
//   		Exact: jsii.String("exact"),
//   		Prefix: jsii.String("prefix"),
//   	},
//
//   	// the properties below are optional
//   	CaseSensitive: jsii.Boolean(false),
//   }
//
type CfnRule_PathMatchProperty struct {
	// The type of path match.
	Match interface{} `field:"required" json:"match" yaml:"match"`
	// Indicates whether the match is case sensitive.
	//
	// Defaults to false.
	CaseSensitive interface{} `field:"optional" json:"caseSensitive" yaml:"caseSensitive"`
}

