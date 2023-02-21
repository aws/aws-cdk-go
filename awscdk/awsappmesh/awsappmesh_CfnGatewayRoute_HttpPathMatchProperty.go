package awsappmesh


// An object representing the path to match in the request.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpPathMatchProperty := &httpPathMatchProperty{
//   	exact: jsii.String("exact"),
//   	regex: jsii.String("regex"),
//   }
//
type CfnGatewayRoute_HttpPathMatchProperty struct {
	// The exact path to match on.
	Exact *string `field:"optional" json:"exact" yaml:"exact"`
	// The regex used to match the path.
	Regex *string `field:"optional" json:"regex" yaml:"regex"`
}

