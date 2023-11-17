package awsimagebuilder


// The included resources of the policy detail.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   includeResourcesProperty := &IncludeResourcesProperty{
//   	Amis: jsii.Boolean(false),
//   	Containers: jsii.Boolean(false),
//   	Snapshots: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-includeresources.html
//
type CfnLifecyclePolicy_IncludeResourcesProperty struct {
	// Use to configure lifecycle actions on AMIs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-includeresources.html#cfn-imagebuilder-lifecyclepolicy-includeresources-amis
	//
	Amis interface{} `field:"optional" json:"amis" yaml:"amis"`
	// Use to configure lifecycle actions on containers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-includeresources.html#cfn-imagebuilder-lifecyclepolicy-includeresources-containers
	//
	Containers interface{} `field:"optional" json:"containers" yaml:"containers"`
	// Use to configure lifecycle actions on snapshots.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-includeresources.html#cfn-imagebuilder-lifecyclepolicy-includeresources-snapshots
	//
	Snapshots interface{} `field:"optional" json:"snapshots" yaml:"snapshots"`
}

