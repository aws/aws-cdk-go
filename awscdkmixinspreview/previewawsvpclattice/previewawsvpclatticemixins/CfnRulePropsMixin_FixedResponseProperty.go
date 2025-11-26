package previewawsvpclatticemixins


// Describes an action that returns a custom HTTP response.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   fixedResponseProperty := &FixedResponseProperty{
//   	StatusCode: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-rule-fixedresponse.html
//
type CfnRulePropsMixin_FixedResponseProperty struct {
	// The HTTP response code.
	//
	// Only `404` and `500` status codes are supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-rule-fixedresponse.html#cfn-vpclattice-rule-fixedresponse-statuscode
	//
	StatusCode *float64 `field:"optional" json:"statusCode" yaml:"statusCode"`
}

