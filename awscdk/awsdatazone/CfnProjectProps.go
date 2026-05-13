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
//   	DomainUnitId: jsii.String("domainUnitId"),
//   	GlossaryTerms: []*string{
//   		jsii.String("glossaryTerms"),
//   	},
//   	MembershipAssignments: []interface{}{
//   		&ProjectMembershipAssignmentProperty{
//   			Designation: jsii.String("designation"),
//   			Member: &MemberProperty{
//   				GroupIdentifier: jsii.String("groupIdentifier"),
//   				UserIdentifier: jsii.String("userIdentifier"),
//   			},
//   		},
//   	},
//   	ProjectCategory: jsii.String("projectCategory"),
//   	ProjectExecutionRole: jsii.String("projectExecutionRole"),
//   	ProjectProfileId: jsii.String("projectProfileId"),
//   	ProjectProfileVersion: jsii.String("projectProfileVersion"),
//   	ResourceTags: []interface{}{
//   		&ResourceTagProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
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
	// The project membership assignments.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-project.html#cfn-datazone-project-membershipassignments
	//
	MembershipAssignments interface{} `field:"optional" json:"membershipAssignments" yaml:"membershipAssignments"`
	// The project category.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-project.html#cfn-datazone-project-projectcategory
	//
	ProjectCategory *string `field:"optional" json:"projectCategory" yaml:"projectCategory"`
	// The project execution role ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-project.html#cfn-datazone-project-projectexecutionrole
	//
	ProjectExecutionRole *string `field:"optional" json:"projectExecutionRole" yaml:"projectExecutionRole"`
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
	// The resource tags of the project.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-project.html#cfn-datazone-project-resourcetags
	//
	ResourceTags interface{} `field:"optional" json:"resourceTags" yaml:"resourceTags"`
	// The user parameters of the project.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-project.html#cfn-datazone-project-userparameters
	//
	UserParameters interface{} `field:"optional" json:"userParameters" yaml:"userParameters"`
}

