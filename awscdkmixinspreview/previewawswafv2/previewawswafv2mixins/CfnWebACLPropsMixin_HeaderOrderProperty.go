package previewawswafv2mixins


// The string containing the list of a web request's header names, ordered as they appear in the web request, separated by colons.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   headerOrderProperty := &HeaderOrderProperty{
//   	OversizeHandling: jsii.String("oversizeHandling"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-headerorder.html
//
type CfnWebACLPropsMixin_HeaderOrderProperty struct {
	// Handling of requests containing oversize fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-headerorder.html#cfn-wafv2-webacl-headerorder-oversizehandling
	//
	OversizeHandling *string `field:"optional" json:"oversizeHandling" yaml:"oversizeHandling"`
}

