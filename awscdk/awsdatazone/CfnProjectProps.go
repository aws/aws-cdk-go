package awsdatazone


// Properties for defining a `CfnProject`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnProjectProps := &CfnProjectProps{
//   	DomainIdentifier: jsii.String("domainIdentifier"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	GlossaryTerms: []*string{
//   		jsii.String("glossaryTerms"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-project.html
//
type CfnProjectProps struct {
	// The identifier of a Amazon DataZone domain where the project exists.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-project.html#cfn-datazone-project-domainidentifier
	//
	DomainIdentifier *string `field:"required" json:"domainIdentifier" yaml:"domainIdentifier"`
	// The name of a project.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-project.html#cfn-datazone-project-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of a project.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-project.html#cfn-datazone-project-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The glossary terms that can be used in this Amazon DataZone project.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-project.html#cfn-datazone-project-glossaryterms
	//
	GlossaryTerms *[]*string `field:"optional" json:"glossaryTerms" yaml:"glossaryTerms"`
}

