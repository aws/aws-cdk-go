package awsappmesh


// An object representing the path to match in the request.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpPathMatchProperty := &HttpPathMatchProperty{
//   	Exact: jsii.String("exact"),
//   	Regex: jsii.String("regex"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-httppathmatch.html
//
type CfnRoute_HttpPathMatchProperty struct {
	// The exact path to match on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-httppathmatch.html#cfn-appmesh-route-httppathmatch-exact
	//
	Exact *string `field:"optional" json:"exact" yaml:"exact"`
	// The regex used to match the path.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-httppathmatch.html#cfn-appmesh-route-httppathmatch-regex
	//
	Regex *string `field:"optional" json:"regex" yaml:"regex"`
}

