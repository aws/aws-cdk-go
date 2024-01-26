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
//   	EnvironmentProfileIdentifier: jsii.String("environmentProfileIdentifier"),
//   	Name: jsii.String("name"),
//   	ProjectIdentifier: jsii.String("projectIdentifier"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
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
	// The identifier of the environment profile that is used to create this Amazon DataZone environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-environment.html#cfn-datazone-environment-environmentprofileidentifier
	//
	EnvironmentProfileIdentifier *string `field:"required" json:"environmentProfileIdentifier" yaml:"environmentProfileIdentifier"`
	// The name of the Amazon DataZone environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-environment.html#cfn-datazone-environment-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The identifier of the Amazon DataZone project in which this environment is created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-environment.html#cfn-datazone-environment-projectidentifier
	//
	ProjectIdentifier *string `field:"required" json:"projectIdentifier" yaml:"projectIdentifier"`
	// The description of the environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-environment.html#cfn-datazone-environment-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The glossary terms that can be used in this Amazon DataZone environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-environment.html#cfn-datazone-environment-glossaryterms
	//
	GlossaryTerms *[]*string `field:"optional" json:"glossaryTerms" yaml:"glossaryTerms"`
	// The user parameters of this Amazon DataZone environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-environment.html#cfn-datazone-environment-userparameters
	//
	UserParameters interface{} `field:"optional" json:"userParameters" yaml:"userParameters"`
}

