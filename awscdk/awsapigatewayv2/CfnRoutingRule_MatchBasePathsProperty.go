package awsapigatewayv2


// Represents a `MatchBasePaths` condition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   matchBasePathsProperty := &MatchBasePathsProperty{
//   	AnyOf: []*string{
//   		jsii.String("anyOf"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-routingrule-matchbasepaths.html
//
type CfnRoutingRule_MatchBasePathsProperty struct {
	// The string of the case sensitive base path to be matched.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-routingrule-matchbasepaths.html#cfn-apigatewayv2-routingrule-matchbasepaths-anyof
	//
	AnyOf *[]*string `field:"required" json:"anyOf" yaml:"anyOf"`
}

