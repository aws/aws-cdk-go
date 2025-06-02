package awsapigatewayv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   matchHeadersProperty := &MatchHeadersProperty{
//   	AnyOf: []interface{}{
//   		&MatchHeaderValueProperty{
//   			Header: jsii.String("header"),
//   			ValueGlob: jsii.String("valueGlob"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-routingrule-matchheaders.html
//
type CfnRoutingRule_MatchHeadersProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-routingrule-matchheaders.html#cfn-apigatewayv2-routingrule-matchheaders-anyof
	//
	AnyOf interface{} `field:"required" json:"anyOf" yaml:"anyOf"`
}

