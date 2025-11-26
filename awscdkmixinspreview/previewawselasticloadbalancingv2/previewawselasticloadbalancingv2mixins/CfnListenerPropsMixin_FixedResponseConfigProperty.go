package previewawselasticloadbalancingv2mixins


// Specifies information required when returning a custom HTTP response.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   fixedResponseConfigProperty := &FixedResponseConfigProperty{
//   	ContentType: jsii.String("contentType"),
//   	MessageBody: jsii.String("messageBody"),
//   	StatusCode: jsii.String("statusCode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-fixedresponseconfig.html
//
type CfnListenerPropsMixin_FixedResponseConfigProperty struct {
	// The content type.
	//
	// Valid Values: text/plain | text/css | text/html | application/javascript | application/json.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-fixedresponseconfig.html#cfn-elasticloadbalancingv2-listener-fixedresponseconfig-contenttype
	//
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// The message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-fixedresponseconfig.html#cfn-elasticloadbalancingv2-listener-fixedresponseconfig-messagebody
	//
	MessageBody *string `field:"optional" json:"messageBody" yaml:"messageBody"`
	// The HTTP response code (2XX, 4XX, or 5XX).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-fixedresponseconfig.html#cfn-elasticloadbalancingv2-listener-fixedresponseconfig-statuscode
	//
	StatusCode *string `field:"optional" json:"statusCode" yaml:"statusCode"`
}

