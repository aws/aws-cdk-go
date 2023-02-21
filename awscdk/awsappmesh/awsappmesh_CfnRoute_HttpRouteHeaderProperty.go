package awsappmesh


// An object that represents the HTTP header in the request.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpRouteHeaderProperty := &httpRouteHeaderProperty{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	invert: jsii.Boolean(false),
//   	match: &headerMatchMethodProperty{
//   		exact: jsii.String("exact"),
//   		prefix: jsii.String("prefix"),
//   		range: &matchRangeProperty{
//   			end: jsii.Number(123),
//   			start: jsii.Number(123),
//   		},
//   		regex: jsii.String("regex"),
//   		suffix: jsii.String("suffix"),
//   	},
//   }
//
type CfnRoute_HttpRouteHeaderProperty struct {
	// A name for the HTTP header in the client request that will be matched on.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specify `True` to match anything except the match criteria.
	//
	// The default value is `False` .
	Invert interface{} `field:"optional" json:"invert" yaml:"invert"`
	// The `HeaderMatchMethod` object.
	Match interface{} `field:"optional" json:"match" yaml:"match"`
}

