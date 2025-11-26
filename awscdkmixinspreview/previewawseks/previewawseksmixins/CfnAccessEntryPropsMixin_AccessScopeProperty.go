package previewawseksmixins


// The scope of an `AccessPolicy` that's associated to an `AccessEntry` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   accessScopeProperty := &AccessScopeProperty{
//   	Namespaces: []*string{
//   		jsii.String("namespaces"),
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-accessentry-accessscope.html
//
type CfnAccessEntryPropsMixin_AccessScopeProperty struct {
	// A Kubernetes `namespace` that an access policy is scoped to.
	//
	// A value is required if you specified `namespace` for `Type` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-accessentry-accessscope.html#cfn-eks-accessentry-accessscope-namespaces
	//
	Namespaces *[]*string `field:"optional" json:"namespaces" yaml:"namespaces"`
	// The scope type of an access policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-accessentry-accessscope.html#cfn-eks-accessentry-accessscope-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

