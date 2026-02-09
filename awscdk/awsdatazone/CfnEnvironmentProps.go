package awsdatazone


// Properties for defining a `CfnEnvironment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEnvironmentProps := &CfnEnvironmentProps{
//   	DomainIdentifier: jsii.String("domainIdentifier"),
//   	Name: jsii.String("name"),
//   	ProjectIdentifier: jsii.String("projectIdentifier"),
//
//   	// the properties below are optional
//   	DeploymentOrder: jsii.Number(123),
//   	Description: jsii.String("description"),
//   	EnvironmentAccountIdentifier: jsii.String("environmentAccountIdentifier"),
//   	EnvironmentAccountRegion: jsii.String("environmentAccountRegion"),
//   	EnvironmentBlueprintIdentifier: jsii.String("environmentBlueprintIdentifier"),
//   	EnvironmentConfigurationId: jsii.String("environmentConfigurationId"),
//   	EnvironmentProfileIdentifier: jsii.String("environmentProfileIdentifier"),
//   	EnvironmentRoleArn: jsii.String("environmentRoleArn"),
//   	GlossaryTerms: []*string{
//   		jsii.String("glossaryTerms"),
//   	},
//   	UserParameters: []interface{}{
//   		&EnvironmentParameterProperty{
//   			Name: jsii.String("name"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-environment.html
//
type CfnEnvironmentProps struct {
	// The identifier of the Amazon DataZone domain in which the environment is created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-environment.html#cfn-datazone-environment-domainidentifier
	//
	DomainIdentifier *string `field:"required" json:"domainIdentifier" yaml:"domainIdentifier"`
	// The name of the Amazon DataZone environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-environment.html#cfn-datazone-environment-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The identifier of the Amazon DataZone project in which this environment is created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-environment.html#cfn-datazone-environment-projectidentifier
	//
	ProjectIdentifier *string `field:"required" json:"projectIdentifier" yaml:"projectIdentifier"`
	// The deployment order for the environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-environment.html#cfn-datazone-environment-deploymentorder
	//
	DeploymentOrder *float64 `field:"optional" json:"deploymentOrder" yaml:"deploymentOrder"`
	// The description of the environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-environment.html#cfn-datazone-environment-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The identifier of the AWS account in which an environment exists.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-environment.html#cfn-datazone-environment-environmentaccountidentifier
	//
	EnvironmentAccountIdentifier *string `field:"optional" json:"environmentAccountIdentifier" yaml:"environmentAccountIdentifier"`
	// The AWS Region in which an environment exists.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-environment.html#cfn-datazone-environment-environmentaccountregion
	//
	EnvironmentAccountRegion *string `field:"optional" json:"environmentAccountRegion" yaml:"environmentAccountRegion"`
	// The identifier of the environment blueprint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-environment.html#cfn-datazone-environment-environmentblueprintidentifier
	//
	EnvironmentBlueprintIdentifier *string `field:"optional" json:"environmentBlueprintIdentifier" yaml:"environmentBlueprintIdentifier"`
	// The identifier of the environment configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-environment.html#cfn-datazone-environment-environmentconfigurationid
	//
	EnvironmentConfigurationId *string `field:"optional" json:"environmentConfigurationId" yaml:"environmentConfigurationId"`
	// The identifier of the environment profile that is used to create this Amazon DataZone environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-environment.html#cfn-datazone-environment-environmentprofileidentifier
	//
	EnvironmentProfileIdentifier *string `field:"optional" json:"environmentProfileIdentifier" yaml:"environmentProfileIdentifier"`
	// The ARN of the environment role.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-environment.html#cfn-datazone-environment-environmentrolearn
	//
	EnvironmentRoleArn *string `field:"optional" json:"environmentRoleArn" yaml:"environmentRoleArn"`
	// The glossary terms that can be used in this Amazon DataZone environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-environment.html#cfn-datazone-environment-glossaryterms
	//
	GlossaryTerms *[]*string `field:"optional" json:"glossaryTerms" yaml:"glossaryTerms"`
	// The user parameters of this Amazon DataZone environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-environment.html#cfn-datazone-environment-userparameters
	//
	UserParameters interface{} `field:"optional" json:"userParameters" yaml:"userParameters"`
}

