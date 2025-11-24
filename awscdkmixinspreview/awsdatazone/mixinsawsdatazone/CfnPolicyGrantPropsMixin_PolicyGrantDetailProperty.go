package mixinsawsdatazone


// The details of the policy grant.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var createEnvironment interface{}
//   var createEnvironmentFromBlueprint interface{}
//   var delegateCreateEnvironmentProfile interface{}
//
//   policyGrantDetailProperty := &PolicyGrantDetailProperty{
//   	AddToProjectMemberPool: &AddToProjectMemberPoolPolicyGrantDetailProperty{
//   		IncludeChildDomainUnits: jsii.Boolean(false),
//   	},
//   	CreateAssetType: &CreateAssetTypePolicyGrantDetailProperty{
//   		IncludeChildDomainUnits: jsii.Boolean(false),
//   	},
//   	CreateDomainUnit: &CreateDomainUnitPolicyGrantDetailProperty{
//   		IncludeChildDomainUnits: jsii.Boolean(false),
//   	},
//   	CreateEnvironment: createEnvironment,
//   	CreateEnvironmentFromBlueprint: createEnvironmentFromBlueprint,
//   	CreateEnvironmentProfile: &CreateEnvironmentProfilePolicyGrantDetailProperty{
//   		DomainUnitId: jsii.String("domainUnitId"),
//   	},
//   	CreateFormType: &CreateFormTypePolicyGrantDetailProperty{
//   		IncludeChildDomainUnits: jsii.Boolean(false),
//   	},
//   	CreateGlossary: &CreateGlossaryPolicyGrantDetailProperty{
//   		IncludeChildDomainUnits: jsii.Boolean(false),
//   	},
//   	CreateProject: &CreateProjectPolicyGrantDetailProperty{
//   		IncludeChildDomainUnits: jsii.Boolean(false),
//   	},
//   	CreateProjectFromProjectProfile: &CreateProjectFromProjectProfilePolicyGrantDetailProperty{
//   		IncludeChildDomainUnits: jsii.Boolean(false),
//   		ProjectProfiles: []*string{
//   			jsii.String("projectProfiles"),
//   		},
//   	},
//   	DelegateCreateEnvironmentProfile: delegateCreateEnvironmentProfile,
//   	OverrideDomainUnitOwners: &OverrideDomainUnitOwnersPolicyGrantDetailProperty{
//   		IncludeChildDomainUnits: jsii.Boolean(false),
//   	},
//   	OverrideProjectOwners: &OverrideProjectOwnersPolicyGrantDetailProperty{
//   		IncludeChildDomainUnits: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-policygrant-policygrantdetail.html
//
type CfnPolicyGrantPropsMixin_PolicyGrantDetailProperty struct {
	// Specifies that the policy grant is to be added to the members of the project.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-policygrant-policygrantdetail.html#cfn-datazone-policygrant-policygrantdetail-addtoprojectmemberpool
	//
	AddToProjectMemberPool interface{} `field:"optional" json:"addToProjectMemberPool" yaml:"addToProjectMemberPool"`
	// Specifies that this is a create asset type policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-policygrant-policygrantdetail.html#cfn-datazone-policygrant-policygrantdetail-createassettype
	//
	CreateAssetType interface{} `field:"optional" json:"createAssetType" yaml:"createAssetType"`
	// Specifies that this is a create domain unit policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-policygrant-policygrantdetail.html#cfn-datazone-policygrant-policygrantdetail-createdomainunit
	//
	CreateDomainUnit interface{} `field:"optional" json:"createDomainUnit" yaml:"createDomainUnit"`
	// Specifies that this is a create environment policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-policygrant-policygrantdetail.html#cfn-datazone-policygrant-policygrantdetail-createenvironment
	//
	CreateEnvironment interface{} `field:"optional" json:"createEnvironment" yaml:"createEnvironment"`
	// The details of the policy of creating an environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-policygrant-policygrantdetail.html#cfn-datazone-policygrant-policygrantdetail-createenvironmentfromblueprint
	//
	CreateEnvironmentFromBlueprint interface{} `field:"optional" json:"createEnvironmentFromBlueprint" yaml:"createEnvironmentFromBlueprint"`
	// Specifies that this is a create environment profile policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-policygrant-policygrantdetail.html#cfn-datazone-policygrant-policygrantdetail-createenvironmentprofile
	//
	CreateEnvironmentProfile interface{} `field:"optional" json:"createEnvironmentProfile" yaml:"createEnvironmentProfile"`
	// Specifies that this is a create form type policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-policygrant-policygrantdetail.html#cfn-datazone-policygrant-policygrantdetail-createformtype
	//
	CreateFormType interface{} `field:"optional" json:"createFormType" yaml:"createFormType"`
	// Specifies that this is a create glossary policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-policygrant-policygrantdetail.html#cfn-datazone-policygrant-policygrantdetail-createglossary
	//
	CreateGlossary interface{} `field:"optional" json:"createGlossary" yaml:"createGlossary"`
	// Specifies that this is a create project policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-policygrant-policygrantdetail.html#cfn-datazone-policygrant-policygrantdetail-createproject
	//
	CreateProject interface{} `field:"optional" json:"createProject" yaml:"createProject"`
	// Specifies whether to create a project from project profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-policygrant-policygrantdetail.html#cfn-datazone-policygrant-policygrantdetail-createprojectfromprojectprofile
	//
	CreateProjectFromProjectProfile interface{} `field:"optional" json:"createProjectFromProjectProfile" yaml:"createProjectFromProjectProfile"`
	// Specifies that this is the delegation of the create environment profile policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-policygrant-policygrantdetail.html#cfn-datazone-policygrant-policygrantdetail-delegatecreateenvironmentprofile
	//
	DelegateCreateEnvironmentProfile interface{} `field:"optional" json:"delegateCreateEnvironmentProfile" yaml:"delegateCreateEnvironmentProfile"`
	// Specifies whether to override domain unit owners.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-policygrant-policygrantdetail.html#cfn-datazone-policygrant-policygrantdetail-overridedomainunitowners
	//
	OverrideDomainUnitOwners interface{} `field:"optional" json:"overrideDomainUnitOwners" yaml:"overrideDomainUnitOwners"`
	// Specifies whether to override project owners.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-policygrant-policygrantdetail.html#cfn-datazone-policygrant-policygrantdetail-overrideprojectowners
	//
	OverrideProjectOwners interface{} `field:"optional" json:"overrideProjectOwners" yaml:"overrideProjectOwners"`
}

