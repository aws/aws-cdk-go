package mixinsawsdatazone


// Properties for CfnDomainUnitPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDomainUnitMixinProps := &CfnDomainUnitMixinProps{
//   	Description: jsii.String("description"),
//   	DomainIdentifier: jsii.String("domainIdentifier"),
//   	Name: jsii.String("name"),
//   	ParentDomainUnitIdentifier: jsii.String("parentDomainUnitIdentifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-domainunit.html
//
type CfnDomainUnitMixinProps struct {
	// The description of the domain unit.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-domainunit.html#cfn-datazone-domainunit-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The ID of the domain where you want to crate a domain unit.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-domainunit.html#cfn-datazone-domainunit-domainidentifier
	//
	DomainIdentifier *string `field:"optional" json:"domainIdentifier" yaml:"domainIdentifier"`
	// The name of the domain unit.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-domainunit.html#cfn-datazone-domainunit-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The ID of the parent domain unit.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-domainunit.html#cfn-datazone-domainunit-parentdomainunitidentifier
	//
	ParentDomainUnitIdentifier *string `field:"optional" json:"parentDomainUnitIdentifier" yaml:"parentDomainUnitIdentifier"`
}

