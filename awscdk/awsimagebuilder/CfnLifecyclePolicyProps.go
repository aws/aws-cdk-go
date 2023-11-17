package awsimagebuilder


// Properties for defining a `CfnLifecyclePolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLifecyclePolicyProps := &CfnLifecyclePolicyProps{
//   	ExecutionRole: jsii.String("executionRole"),
//   	Name: jsii.String("name"),
//   	PolicyDetails: []interface{}{
//   		&PolicyDetailProperty{
//   			Action: &ActionProperty{
//   				Type: jsii.String("type"),
//
//   				// the properties below are optional
//   				IncludeResources: &IncludeResourcesProperty{
//   					Amis: jsii.Boolean(false),
//   					Containers: jsii.Boolean(false),
//   					Snapshots: jsii.Boolean(false),
//   				},
//   			},
//   			Filter: &FilterProperty{
//   				Type: jsii.String("type"),
//   				Value: jsii.Number(123),
//
//   				// the properties below are optional
//   				RetainAtLeast: jsii.Number(123),
//   				Unit: jsii.String("unit"),
//   			},
//
//   			// the properties below are optional
//   			ExclusionRules: &ExclusionRulesProperty{
//   				Amis: &AmiExclusionRulesProperty{
//   					IsPublic: jsii.Boolean(false),
//   					LastLaunched: &LastLaunchedProperty{
//   						Unit: jsii.String("unit"),
//   						Value: jsii.Number(123),
//   					},
//   					Regions: []*string{
//   						jsii.String("regions"),
//   					},
//   					SharedAccounts: []*string{
//   						jsii.String("sharedAccounts"),
//   					},
//   					TagMap: map[string]*string{
//   						"tagMapKey": jsii.String("tagMap"),
//   					},
//   				},
//   				TagMap: map[string]*string{
//   					"tagMapKey": jsii.String("tagMap"),
//   				},
//   			},
//   		},
//   	},
//   	ResourceSelection: &ResourceSelectionProperty{
//   		Recipes: []interface{}{
//   			&RecipeSelectionProperty{
//   				Name: jsii.String("name"),
//
//   				// the properties below are optional
//   				SemanticVersion: jsii.String("semanticVersion"),
//   			},
//   		},
//   		TagMap: map[string]*string{
//   			"tagMapKey": jsii.String("tagMap"),
//   		},
//   	},
//   	ResourceType: jsii.String("resourceType"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Status: jsii.String("status"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-lifecyclepolicy.html
//
type CfnLifecyclePolicyProps struct {
	// The execution role of the lifecycle policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-lifecyclepolicy.html#cfn-imagebuilder-lifecyclepolicy-executionrole
	//
	ExecutionRole *string `field:"required" json:"executionRole" yaml:"executionRole"`
	// The name of the lifecycle policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-lifecyclepolicy.html#cfn-imagebuilder-lifecyclepolicy-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The policy details of the lifecycle policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-lifecyclepolicy.html#cfn-imagebuilder-lifecyclepolicy-policydetails
	//
	PolicyDetails interface{} `field:"required" json:"policyDetails" yaml:"policyDetails"`
	// The resource selection for the lifecycle policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-lifecyclepolicy.html#cfn-imagebuilder-lifecyclepolicy-resourceselection
	//
	ResourceSelection interface{} `field:"required" json:"resourceSelection" yaml:"resourceSelection"`
	// The resource type of the lifecycle policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-lifecyclepolicy.html#cfn-imagebuilder-lifecyclepolicy-resourcetype
	//
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
	// The description of the lifecycle policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-lifecyclepolicy.html#cfn-imagebuilder-lifecyclepolicy-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The status of the lifecycle policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-lifecyclepolicy.html#cfn-imagebuilder-lifecyclepolicy-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The tags associated with the lifecycle policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-lifecyclepolicy.html#cfn-imagebuilder-lifecyclepolicy-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

