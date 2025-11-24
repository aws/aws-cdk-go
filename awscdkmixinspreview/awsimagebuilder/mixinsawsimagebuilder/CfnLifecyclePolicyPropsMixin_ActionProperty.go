package mixinsawsimagebuilder


// Contains selection criteria for the lifecycle policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   actionProperty := &ActionProperty{
//   	IncludeResources: &IncludeResourcesProperty{
//   		Amis: jsii.Boolean(false),
//   		Containers: jsii.Boolean(false),
//   		Snapshots: jsii.Boolean(false),
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-action.html
//
type CfnLifecyclePolicyPropsMixin_ActionProperty struct {
	// Specifies the resources that the lifecycle policy applies to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-action.html#cfn-imagebuilder-lifecyclepolicy-action-includeresources
	//
	IncludeResources interface{} `field:"optional" json:"includeResources" yaml:"includeResources"`
	// Specifies the lifecycle action to take.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-action.html#cfn-imagebuilder-lifecyclepolicy-action-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

