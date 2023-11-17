package awsimagebuilder


// The action of the policy detail.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   actionProperty := &ActionProperty{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	IncludeResources: &IncludeResourcesProperty{
//   		Amis: jsii.Boolean(false),
//   		Containers: jsii.Boolean(false),
//   		Snapshots: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-action.html
//
type CfnLifecyclePolicy_ActionProperty struct {
	// The action type of the policy detail.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-action.html#cfn-imagebuilder-lifecyclepolicy-action-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The included resources of the policy detail.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-action.html#cfn-imagebuilder-lifecyclepolicy-action-includeresources
	//
	IncludeResources interface{} `field:"optional" json:"includeResources" yaml:"includeResources"`
}

