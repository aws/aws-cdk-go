package mixinsawsimagebuilder


// Specifies how the lifecycle policy should apply actions to selected resources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   includeResourcesProperty := &IncludeResourcesProperty{
//   	Amis: jsii.Boolean(false),
//   	Containers: jsii.Boolean(false),
//   	Snapshots: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-includeresources.html
//
type CfnLifecyclePolicyPropsMixin_IncludeResourcesProperty struct {
	// Specifies whether the lifecycle action should apply to distributed AMIs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-includeresources.html#cfn-imagebuilder-lifecyclepolicy-includeresources-amis
	//
	Amis interface{} `field:"optional" json:"amis" yaml:"amis"`
	// Specifies whether the lifecycle action should apply to distributed containers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-includeresources.html#cfn-imagebuilder-lifecyclepolicy-includeresources-containers
	//
	Containers interface{} `field:"optional" json:"containers" yaml:"containers"`
	// Specifies whether the lifecycle action should apply to snapshots associated with distributed AMIs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-includeresources.html#cfn-imagebuilder-lifecyclepolicy-includeresources-snapshots
	//
	Snapshots interface{} `field:"optional" json:"snapshots" yaml:"snapshots"`
}

