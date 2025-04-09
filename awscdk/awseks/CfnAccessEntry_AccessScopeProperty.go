package awseks


// The scope of an `AccessPolicy` that's associated to an `AccessEntry` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessScopeProperty := &AccessScopeProperty{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Namespaces: []*string{
//   		jsii.String("namespaces"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-accessentry-accessscope.html
//
type CfnAccessEntry_AccessScopeProperty struct {
	// The scope type of an access policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-accessentry-accessscope.html#cfn-eks-accessentry-accessscope-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// A Kubernetes `namespace` that an access policy is scoped to.
	//
	// A value is required if you specified `namespace` for `Type` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-accessentry-accessscope.html#cfn-eks-accessentry-accessscope-namespaces
	//
	Namespaces *[]*string `field:"optional" json:"namespaces" yaml:"namespaces"`
}

