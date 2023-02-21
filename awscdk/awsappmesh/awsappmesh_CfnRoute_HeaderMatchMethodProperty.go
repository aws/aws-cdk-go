package awsappmesh


// An object that represents the method and value to match with the header value sent in a request.
//
// Specify one match method.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   headerMatchMethodProperty := &headerMatchMethodProperty{
//   	exact: jsii.String("exact"),
//   	prefix: jsii.String("prefix"),
//   	range: &matchRangeProperty{
//   		end: jsii.Number(123),
//   		start: jsii.Number(123),
//   	},
//   	regex: jsii.String("regex"),
//   	suffix: jsii.String("suffix"),
//   }
//
type CfnRoute_HeaderMatchMethodProperty struct {
	// The value sent by the client must match the specified value exactly.
	Exact *string `field:"optional" json:"exact" yaml:"exact"`
	// The value sent by the client must begin with the specified characters.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// An object that represents the range of values to match on.
	Range interface{} `field:"optional" json:"range" yaml:"range"`
	// The value sent by the client must include the specified characters.
	Regex *string `field:"optional" json:"regex" yaml:"regex"`
	// The value sent by the client must end with the specified characters.
	Suffix *string `field:"optional" json:"suffix" yaml:"suffix"`
}

