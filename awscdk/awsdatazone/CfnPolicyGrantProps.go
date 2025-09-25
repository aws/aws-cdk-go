package awsdatazone


// Properties for defining a `CfnPolicyGrant`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var allDomainUnitsGrantFilter interface{}
//   var allUsersGrantFilter interface{}
//   var createEnvironment interface{}
//   var createEnvironmentFromBlueprint interface{}
//   var delegateCreateEnvironmentProfile interface{}
//
//   cfnPolicyGrantProps := &CfnPolicyGrantProps{
//   	DomainIdentifier: jsii.String("domainIdentifier"),
//   	EntityIdentifier: jsii.String("entityIdentifier"),
//   	EntityType: jsii.String("entityType"),
//   	PolicyType: jsii.String("policyType"),
//
//   	// the properties below are optional
//   	Detail: &PolicyGrantDetailProperty{
//   		AddToProjectMemberPool: &AddToProjectMemberPoolPolicyGrantDetailProperty{
//   			IncludeChildDomainUnits: jsii.Boolean(false),
//   		},
//   		CreateAssetType: &CreateAssetTypePolicyGrantDetailProperty{
//   			IncludeChildDomainUnits: jsii.Boolean(false),
//   		},
//   		CreateDomainUnit: &CreateDomainUnitPolicyGrantDetailProperty{
//   			IncludeChildDomainUnits: jsii.Boolean(false),
//   		},
//   		CreateEnvironment: createEnvironment,
//   		CreateEnvironmentFromBlueprint: createEnvironmentFromBlueprint,
//   		CreateEnvironmentProfile: &CreateEnvironmentProfilePolicyGrantDetailProperty{
//   			DomainUnitId: jsii.String("domainUnitId"),
//   		},
//   		CreateFormType: &CreateFormTypePolicyGrantDetailProperty{
//   			IncludeChildDomainUnits: jsii.Boolean(false),
//   		},
//   		CreateGlossary: &CreateGlossaryPolicyGrantDetailProperty{
//   			IncludeChildDomainUnits: jsii.Boolean(false),
//   		},
//   		CreateProject: &CreateProjectPolicyGrantDetailProperty{
//   			IncludeChildDomainUnits: jsii.Boolean(false),
//   		},
//   		CreateProjectFromProjectProfile: &CreateProjectFromProjectProfilePolicyGrantDetailProperty{
//   			IncludeChildDomainUnits: jsii.Boolean(false),
//   			ProjectProfiles: []*string{
//   				jsii.String("projectProfiles"),
//   			},
//   		},
//   		DelegateCreateEnvironmentProfile: delegateCreateEnvironmentProfile,
//   		OverrideDomainUnitOwners: &OverrideDomainUnitOwnersPolicyGrantDetailProperty{
//   			IncludeChildDomainUnits: jsii.Boolean(false),
//   		},
//   		OverrideProjectOwners: &OverrideProjectOwnersPolicyGrantDetailProperty{
//   			IncludeChildDomainUnits: jsii.Boolean(false),
//   		},
//   	},
//   	Principal: &PolicyGrantPrincipalProperty{
//   		DomainUnit: &DomainUnitPolicyGrantPrincipalProperty{
//   			DomainUnitDesignation: jsii.String("domainUnitDesignation"),
//   			DomainUnitGrantFilter: &DomainUnitGrantFilterProperty{
//   				AllDomainUnitsGrantFilter: allDomainUnitsGrantFilter,
//   			},
//   			DomainUnitIdentifier: jsii.String("domainUnitIdentifier"),
//   		},
//   		Group: &GroupPolicyGrantPrincipalProperty{
//   			GroupIdentifier: jsii.String("groupIdentifier"),
//   		},
//   		Project: &ProjectPolicyGrantPrincipalProperty{
//   			ProjectDesignation: jsii.String("projectDesignation"),
//   			ProjectGrantFilter: &ProjectGrantFilterProperty{
//   				DomainUnitFilter: &DomainUnitFilterForProjectProperty{
//   					DomainUnit: jsii.String("domainUnit"),
//
//   					// the properties below are optional
//   					IncludeChildDomainUnits: jsii.Boolean(false),
//   				},
//   			},
//   			ProjectIdentifier: jsii.String("projectIdentifier"),
//   		},
//   		User: &UserPolicyGrantPrincipalProperty{
//   			AllUsersGrantFilter: allUsersGrantFilter,
//   			UserIdentifier: jsii.String("userIdentifier"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-policygrant.html
//
type CfnPolicyGrantProps struct {
	// The ID of the domain where you want to add a policy grant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-policygrant.html#cfn-datazone-policygrant-domainidentifier
	//
	DomainIdentifier *string `field:"required" json:"domainIdentifier" yaml:"domainIdentifier"`
	// The ID of the entity (resource) to which you want to add a policy grant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-policygrant.html#cfn-datazone-policygrant-entityidentifier
	//
	EntityIdentifier *string `field:"required" json:"entityIdentifier" yaml:"entityIdentifier"`
	// The type of entity (resource) to which the grant is added.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-policygrant.html#cfn-datazone-policygrant-entitytype
	//
	EntityType *string `field:"required" json:"entityType" yaml:"entityType"`
	// The type of policy that you want to grant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-policygrant.html#cfn-datazone-policygrant-policytype
	//
	PolicyType *string `field:"required" json:"policyType" yaml:"policyType"`
	// The details of the policy grant member.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-policygrant.html#cfn-datazone-policygrant-detail
	//
	Detail interface{} `field:"optional" json:"detail" yaml:"detail"`
	// The principal of the policy grant member.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-policygrant.html#cfn-datazone-policygrant-principal
	//
	Principal interface{} `field:"optional" json:"principal" yaml:"principal"`
}

