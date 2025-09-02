package awsdatazone


// The domain unit principal to whom the policy is granted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var allDomainUnitsGrantFilter interface{}
//
//   domainUnitPolicyGrantPrincipalProperty := &DomainUnitPolicyGrantPrincipalProperty{
//   	DomainUnitDesignation: jsii.String("domainUnitDesignation"),
//   	DomainUnitGrantFilter: &DomainUnitGrantFilterProperty{
//   		AllDomainUnitsGrantFilter: allDomainUnitsGrantFilter,
//   	},
//   	DomainUnitIdentifier: jsii.String("domainUnitIdentifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-policygrant-domainunitpolicygrantprincipal.html
//
type CfnPolicyGrant_DomainUnitPolicyGrantPrincipalProperty struct {
	// Specifes the designation of the domain unit users.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-policygrant-domainunitpolicygrantprincipal.html#cfn-datazone-policygrant-domainunitpolicygrantprincipal-domainunitdesignation
	//
	DomainUnitDesignation *string `field:"optional" json:"domainUnitDesignation" yaml:"domainUnitDesignation"`
	// The grant filter for the domain unit.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-policygrant-domainunitpolicygrantprincipal.html#cfn-datazone-policygrant-domainunitpolicygrantprincipal-domainunitgrantfilter
	//
	DomainUnitGrantFilter interface{} `field:"optional" json:"domainUnitGrantFilter" yaml:"domainUnitGrantFilter"`
	// The ID of the domain unit.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-policygrant-domainunitpolicygrantprincipal.html#cfn-datazone-policygrant-domainunitpolicygrantprincipal-domainunitidentifier
	//
	DomainUnitIdentifier *string `field:"optional" json:"domainUnitIdentifier" yaml:"domainUnitIdentifier"`
}

