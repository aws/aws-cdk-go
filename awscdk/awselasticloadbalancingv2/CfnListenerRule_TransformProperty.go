package awselasticloadbalancingv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transformProperty := &TransformProperty{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	HostHeaderRewriteConfig: &RewriteConfigObjectProperty{
//   		Rewrites: []interface{}{
//   			&RewriteConfigProperty{
//   				Regex: jsii.String("regex"),
//   				Replace: jsii.String("replace"),
//   			},
//   		},
//   	},
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
type CfnListenerRule_TransformProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-transform.html#cfn-elasticloadbalancingv2-listenerrule-transform-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-transform.html#cfn-elasticloadbalancingv2-listenerrule-transform-hostheaderrewriteconfig
	//
	HostHeaderRewriteConfig interface{} `field:"optional" json:"hostHeaderRewriteConfig" yaml:"hostHeaderRewriteConfig"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-transform.html#cfn-elasticloadbalancingv2-listenerrule-transform-urlrewriteconfig
	//
	UrlRewriteConfig interface{} `field:"optional" json:"urlRewriteConfig" yaml:"urlRewriteConfig"`
}

