package awsdatazone


// The domain unit filter of the project grant filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   domainUnitFilterForProjectProperty := &DomainUnitFilterForProjectProperty{
//   	DomainUnit: jsii.String("domainUnit"),
//
//   	// the properties below are optional
//   	IncludeChildDomainUnits: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-policygrant-domainunitfilterforproject.html
//
type CfnPolicyGrant_DomainUnitFilterForProjectProperty struct {
	// The domain unit ID to use in the filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-policygrant-domainunitfilterforproject.html#cfn-datazone-policygrant-domainunitfilterforproject-domainunit
	//
	DomainUnit *string `field:"required" json:"domainUnit" yaml:"domainUnit"`
	// Specifies whether to include child domain units.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-policygrant-domainunitfilterforproject.html#cfn-datazone-policygrant-domainunitfilterforproject-includechilddomainunits
	//
	// Default: - false.
	//
	IncludeChildDomainUnits interface{} `field:"optional" json:"includeChildDomainUnits" yaml:"includeChildDomainUnits"`
}

