package mixinsawswafv2


// A single label container.
//
// This is used as an element of a label array in `RuleLabels` inside a rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   labelProperty := &LabelProperty{
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-label.html
//
type CfnWebACLPropsMixin_LabelProperty struct {
	// The label string.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-label.html#cfn-wafv2-webacl-label-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

