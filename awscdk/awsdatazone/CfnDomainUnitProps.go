package awsdatazone


// Properties for defining a `CfnDomainUnit`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDomainUnitProps := &CfnDomainUnitProps{
//   	DomainIdentifier: jsii.String("domainIdentifier"),
//   	Name: jsii.String("name"),
//   	ParentDomainUnitIdentifier: jsii.String("parentDomainUnitIdentifier"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-domainunit.html
//
type CfnDomainUnitProps struct {
	// The ID of the domain where you want to crate a domain unit.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-domainunit.html#cfn-datazone-domainunit-domainidentifier
	//
	DomainIdentifier *string `field:"required" json:"domainIdentifier" yaml:"domainIdentifier"`
	// The name of the domain unit.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-domainunit.html#cfn-datazone-domainunit-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID of the parent domain unit.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-domainunit.html#cfn-datazone-domainunit-parentdomainunitidentifier
	//
	ParentDomainUnitIdentifier *string `field:"required" json:"parentDomainUnitIdentifier" yaml:"parentDomainUnitIdentifier"`
	// The description of the domain unit.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-domainunit.html#cfn-datazone-domainunit-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

