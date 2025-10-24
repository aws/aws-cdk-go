package awselasticloadbalancingv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
type CfnListenerRule_RewriteConfigObjectProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-rewriteconfigobject.html#cfn-elasticloadbalancingv2-listenerrule-rewriteconfigobject-rewrites
	//
	Rewrites interface{} `field:"required" json:"rewrites" yaml:"rewrites"`
}

