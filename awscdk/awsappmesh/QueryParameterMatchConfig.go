package awsappmesh


// Configuration for `QueryParameterMatch`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   queryParameterMatchConfig := &QueryParameterMatchConfig{
//   	QueryParameterMatch: &QueryParameterProperty{
//   		Name: jsii.String("name"),
//
//   		// the properties below are optional
//   		Match: &HttpQueryParameterMatchProperty{
//   			Exact: jsii.String("exact"),
//   		},
//   	},
//   }
//
type QueryParameterMatchConfig struct {
	// Route CFN configuration for route query parameter match.
	QueryParameterMatch *CfnRoute_QueryParameterProperty `field:"required" json:"queryParameterMatch" yaml:"queryParameterMatch"`
}

