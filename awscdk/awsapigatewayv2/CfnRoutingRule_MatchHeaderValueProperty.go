package awsapigatewayv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   matchHeaderValueProperty := &MatchHeaderValueProperty{
//   	Header: jsii.String("header"),
//   	ValueGlob: jsii.String("valueGlob"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-routingrule-matchheadervalue.html
//
type CfnRoutingRule_MatchHeaderValueProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-routingrule-matchheadervalue.html#cfn-apigatewayv2-routingrule-matchheadervalue-header
	//
	Header *string `field:"required" json:"header" yaml:"header"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-routingrule-matchheadervalue.html#cfn-apigatewayv2-routingrule-matchheadervalue-valueglob
	//
	ValueGlob *string `field:"required" json:"valueGlob" yaml:"valueGlob"`
}

