package assertions


// Match failure details.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var matcher matcher
//
//   matchFailure := &MatchFailure{
//   	Matcher: matcher,
//   	Message: jsii.String("message"),
//   	Path: []*string{
//   		jsii.String("path"),
//   	},
//
//   	// the properties below are optional
//   	Cost: jsii.Number(123),
//   }
//
type MatchFailure struct {
	// The matcher that had the failure.
	Matcher Matcher `field:"required" json:"matcher" yaml:"matcher"`
	// Failure message.
	Message *string `field:"required" json:"message" yaml:"message"`
	// The relative path in the target where the failure occurred.
	//
	// If the failure occurred at root of the match tree, set the path to an empty list.
	// If it occurs in the 5th index of an array nested within the 'foo' key of an object,
	// set the path as `['/foo', '[5]']`.
	Path *[]*string `field:"required" json:"path" yaml:"path"`
	// The cost of this particular mismatch.
	// Default: 1.
	//
	Cost *float64 `field:"optional" json:"cost" yaml:"cost"`
}

