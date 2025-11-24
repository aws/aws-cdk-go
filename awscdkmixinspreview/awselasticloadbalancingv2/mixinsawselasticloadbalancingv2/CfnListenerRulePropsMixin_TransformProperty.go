package mixinsawselasticloadbalancingv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   transformProperty := &TransformProperty{
//   	HostHeaderRewriteConfig: &RewriteConfigObjectProperty{
//   		Rewrites: []interface{}{
//   			&RewriteConfigProperty{
//   				Regex: jsii.String("regex"),
//   				Replace: jsii.String("replace"),
//   			},
//   		},
//   	},
//   	Type: jsii.String("type"),
//   	UrlRewriteConfig: &RewriteConfigObjectProperty{
//   		Rewrites: []interface{}{
//   			&RewriteConfigProperty{
//   				Regex: jsii.String("regex"),
//   				Replace: jsii.String("replace"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-transform.html
//
type CfnListenerRulePropsMixin_TransformProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-transform.html#cfn-elasticloadbalancingv2-listenerrule-transform-hostheaderrewriteconfig
	//
	HostHeaderRewriteConfig interface{} `field:"optional" json:"hostHeaderRewriteConfig" yaml:"hostHeaderRewriteConfig"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-transform.html#cfn-elasticloadbalancingv2-listenerrule-transform-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-transform.html#cfn-elasticloadbalancingv2-listenerrule-transform-urlrewriteconfig
	//
	UrlRewriteConfig interface{} `field:"optional" json:"urlRewriteConfig" yaml:"urlRewriteConfig"`
}

