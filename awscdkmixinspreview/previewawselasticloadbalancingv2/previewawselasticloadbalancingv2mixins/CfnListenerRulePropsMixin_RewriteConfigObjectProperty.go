package previewawselasticloadbalancingv2mixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   rewriteConfigObjectProperty := &RewriteConfigObjectProperty{
//   	Rewrites: []interface{}{
//   		&RewriteConfigProperty{
//   			Regex: jsii.String("regex"),
//   			Replace: jsii.String("replace"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-rewriteconfigobject.html
//
type CfnListenerRulePropsMixin_RewriteConfigObjectProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-rewriteconfigobject.html#cfn-elasticloadbalancingv2-listenerrule-rewriteconfigobject-rewrites
	//
	Rewrites interface{} `field:"optional" json:"rewrites" yaml:"rewrites"`
}

