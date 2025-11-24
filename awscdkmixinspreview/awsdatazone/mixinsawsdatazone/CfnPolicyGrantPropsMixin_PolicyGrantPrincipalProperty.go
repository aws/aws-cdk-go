package mixinsawsdatazone


// The policy grant principal.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var allDomainUnitsGrantFilter interface{}
//   var allUsersGrantFilter interface{}
//
//   policyGrantPrincipalProperty := &PolicyGrantPrincipalProperty{
//   	DomainUnit: &DomainUnitPolicyGrantPrincipalProperty{
//   		DomainUnitDesignation: jsii.String("domainUnitDesignation"),
//   		DomainUnitGrantFilter: &DomainUnitGrantFilterProperty{
//   			AllDomainUnitsGrantFilter: allDomainUnitsGrantFilter,
//   		},
//   		DomainUnitIdentifier: jsii.String("domainUnitIdentifier"),
//   	},
//   	Group: &GroupPolicyGrantPrincipalProperty{
//   		GroupIdentifier: jsii.String("groupIdentifier"),
//   	},
//   	Project: &ProjectPolicyGrantPrincipalProperty{
//   		ProjectDesignation: jsii.String("projectDesignation"),
//   		ProjectGrantFilter: &ProjectGrantFilterProperty{
//   			DomainUnitFilter: &DomainUnitFilterForProjectProperty{
//   				DomainUnit: jsii.String("domainUnit"),
//   				IncludeChildDomainUnits: jsii.Boolean(false),
//   			},
//   		},
//   		ProjectIdentifier: jsii.String("projectIdentifier"),
//   	},
//   	User: &UserPolicyGrantPrincipalProperty{
//   		AllUsersGrantFilter: allUsersGrantFilter,
//   		UserIdentifier: jsii.String("userIdentifier"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-policygrant-policygrantprincipal.html
//
type CfnPolicyGrantPropsMixin_PolicyGrantPrincipalProperty struct {
	// The domain unit of the policy grant principal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-policygrant-policygrantprincipal.html#cfn-datazone-policygrant-policygrantprincipal-domainunit
	//
	DomainUnit interface{} `field:"optional" json:"domainUnit" yaml:"domainUnit"`
	// The group of the policy grant principal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-policygrant-policygrantprincipal.html#cfn-datazone-policygrant-policygrantprincipal-group
	//
	Group interface{} `field:"optional" json:"group" yaml:"group"`
	// The project of the policy grant principal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-policygrant-policygrantprincipal.html#cfn-datazone-policygrant-policygrantprincipal-project
	//
	Project interface{} `field:"optional" json:"project" yaml:"project"`
	// The user of the policy grant principal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-policygrant-policygrantprincipal.html#cfn-datazone-policygrant-policygrantprincipal-user
	//
	User interface{} `field:"optional" json:"user" yaml:"user"`
}

