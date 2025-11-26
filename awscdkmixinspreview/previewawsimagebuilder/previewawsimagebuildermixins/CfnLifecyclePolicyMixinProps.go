package previewawsimagebuildermixins


// Properties for CfnLifecyclePolicyPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLifecyclePolicyMixinProps := &CfnLifecyclePolicyMixinProps{
//   	Description: jsii.String("description"),
//   	ExecutionRole: jsii.String("executionRole"),
//   	Name: jsii.String("name"),
//   	PolicyDetails: []interface{}{
//   		&PolicyDetailProperty{
//   			Action: &ActionProperty{
//   				IncludeResources: &IncludeResourcesProperty{
//   					Amis: jsii.Boolean(false),
//   					Containers: jsii.Boolean(false),
//   					Snapshots: jsii.Boolean(false),
//   				},
//   				Type: jsii.String("type"),
//   			},
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
//   			Filter: &FilterProperty{
//   				RetainAtLeast: jsii.Number(123),
//   				Type: jsii.String("type"),
//   				Unit: jsii.String("unit"),
//   				Value: jsii.Number(123),
//   			},
//   		},
//   	},
//   	ResourceSelection: &ResourceSelectionProperty{
//   		Recipes: []interface{}{
//   			&RecipeSelectionProperty{
//   				Name: jsii.String("name"),
//   				SemanticVersion: jsii.String("semanticVersion"),
//   			},
//   		},
//   		TagMap: map[string]*string{
//   			"tagMapKey": jsii.String("tagMap"),
//   		},
//   	},
//   	ResourceType: jsii.String("resourceType"),
//   	Status: jsii.String("status"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-lifecyclepolicy.html
//
type CfnLifecyclePolicyMixinProps struct {
	// Optional description for the lifecycle policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-lifecyclepolicy.html#cfn-imagebuilder-lifecyclepolicy-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name or Amazon Resource Name (ARN) for the IAM role you create that grants Image Builder access to run lifecycle actions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-lifecyclepolicy.html#cfn-imagebuilder-lifecyclepolicy-executionrole
	//
	ExecutionRole *string `field:"optional" json:"executionRole" yaml:"executionRole"`
	// The name of the lifecycle policy to create.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-lifecyclepolicy.html#cfn-imagebuilder-lifecyclepolicy-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Configuration details for the lifecycle policy rules.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-lifecyclepolicy.html#cfn-imagebuilder-lifecyclepolicy-policydetails
	//
	PolicyDetails interface{} `field:"optional" json:"policyDetails" yaml:"policyDetails"`
	// Selection criteria for the resources that the lifecycle policy applies to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-lifecyclepolicy.html#cfn-imagebuilder-lifecyclepolicy-resourceselection
	//
	ResourceSelection interface{} `field:"optional" json:"resourceSelection" yaml:"resourceSelection"`
	// The type of Image Builder resource that the lifecycle policy applies to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-lifecyclepolicy.html#cfn-imagebuilder-lifecyclepolicy-resourcetype
	//
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// Indicates whether the lifecycle policy resource is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-lifecyclepolicy.html#cfn-imagebuilder-lifecyclepolicy-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Tags to apply to the lifecycle policy resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-lifecyclepolicy.html#cfn-imagebuilder-lifecyclepolicy-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

