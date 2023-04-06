package awsvpclattice


// Describes the constraints for a header match.
//
// Matches incoming requests with rule based on request header value before applying rule action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   headerMatchProperty := &HeaderMatchProperty{
//   	Match: &HeaderMatchTypeProperty{
//   		Contains: jsii.String("contains"),
//   		Exact: jsii.String("exact"),
//   		Prefix: jsii.String("prefix"),
//   	},
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	CaseSensitive: jsii.Boolean(false),
//   }
//
type CfnRule_HeaderMatchProperty struct {
	// The header match type.
	Match interface{} `field:"required" json:"match" yaml:"match"`
	// The name of the header.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Indicates whether the match is case sensitive.
	//
	// Defaults to false.
	CaseSensitive interface{} `field:"optional" json:"caseSensitive" yaml:"caseSensitive"`
}

