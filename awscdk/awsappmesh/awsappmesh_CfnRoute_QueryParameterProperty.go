package awsappmesh


// An object that represents the query parameter in the request.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   queryParameterProperty := &queryParameterProperty{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	match: &httpQueryParameterMatchProperty{
//   		exact: jsii.String("exact"),
//   	},
//   }
//
type CfnRoute_QueryParameterProperty struct {
	// A name for the query parameter that will be matched on.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The query parameter to match on.
	Match interface{} `field:"optional" json:"match" yaml:"match"`
}

