package mixinsawsdatazone


// The grant filter for the domain unit.
//
// In the current release of Amazon DataZone, the only supported filter is the `allDomainUnitsGrantFilter` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var allDomainUnitsGrantFilter interface{}
//
//   domainUnitGrantFilterProperty := &DomainUnitGrantFilterProperty{
//   	AllDomainUnitsGrantFilter: allDomainUnitsGrantFilter,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-policygrant-domainunitgrantfilter.html
//
type CfnPolicyGrantPropsMixin_DomainUnitGrantFilterProperty struct {
	// Specifies a grant filter containing all domain units.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-policygrant-domainunitgrantfilter.html#cfn-datazone-policygrant-domainunitgrantfilter-alldomainunitsgrantfilter
	//
	AllDomainUnitsGrantFilter interface{} `field:"optional" json:"allDomainUnitsGrantFilter" yaml:"allDomainUnitsGrantFilter"`
}

