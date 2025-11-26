package previewawsvpclatticemixins


// Describes a path match type.
//
// Each rule can include only one of the following types of paths.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   pathMatchTypeProperty := &PathMatchTypeProperty{
//   	Exact: jsii.String("exact"),
//   	Prefix: jsii.String("prefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-rule-pathmatchtype.html
//
type CfnRulePropsMixin_PathMatchTypeProperty struct {
	// An exact match of the path.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-rule-pathmatchtype.html#cfn-vpclattice-rule-pathmatchtype-exact
	//
	Exact *string `field:"optional" json:"exact" yaml:"exact"`
	// A prefix match of the path.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-rule-pathmatchtype.html#cfn-vpclattice-rule-pathmatchtype-prefix
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

