package mixinsawsdatazone


// Properties for CfnProjectPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnProjectMixinProps := &CfnProjectMixinProps{
//   	Description: jsii.String("description"),
//   	DomainIdentifier: jsii.String("domainIdentifier"),
//   	DomainUnitId: jsii.String("domainUnitId"),
//   	GlossaryTerms: []*string{
//   		jsii.String("glossaryTerms"),
//   	},
//   	Name: jsii.String("name"),
//   	ProjectProfileId: jsii.String("projectProfileId"),
//   	ProjectProfileVersion: jsii.String("projectProfileVersion"),
//   	UserParameters: []interface{}{
//   		&EnvironmentConfigurationUserParameterProperty{
//   			EnvironmentConfigurationName: jsii.String("environmentConfigurationName"),
//   			EnvironmentId: jsii.String("environmentId"),
//   			EnvironmentParameters: []interface{}{
//   				&EnvironmentParameterProperty{
//   					Name: jsii.String("name"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-project.html
//
type CfnProjectMixinProps struct {
	// The description of a project.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-project.html#cfn-datazone-project-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The identifier of a Amazon DataZone domain where the project exists.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-project.html#cfn-datazone-project-domainidentifier
	//
	DomainIdentifier *string `field:"optional" json:"domainIdentifier" yaml:"domainIdentifier"`
	// The ID of the domain unit.
	//
	// This parameter is not required and if it is not specified, then the project is created at the root domain unit level.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-project.html#cfn-datazone-project-domainunitid
	//
	DomainUnitId *string `field:"optional" json:"domainUnitId" yaml:"domainUnitId"`
	// The glossary terms that can be used in this Amazon DataZone project.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-project.html#cfn-datazone-project-glossaryterms
	//
	GlossaryTerms *[]*string `field:"optional" json:"glossaryTerms" yaml:"glossaryTerms"`
	// The name of a project.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-project.html#cfn-datazone-project-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The ID of the project profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-project.html#cfn-datazone-project-projectprofileid
	//
	ProjectProfileId *string `field:"optional" json:"projectProfileId" yaml:"projectProfileId"`
	// The project profile version to which the project should be updated.
	//
	// You can only specify the following string for this parameter: `latest` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-project.html#cfn-datazone-project-projectprofileversion
	//
	ProjectProfileVersion *string `field:"optional" json:"projectProfileVersion" yaml:"projectProfileVersion"`
	// The user parameters of the project.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-project.html#cfn-datazone-project-userparameters
	//
	UserParameters interface{} `field:"optional" json:"userParameters" yaml:"userParameters"`
}

