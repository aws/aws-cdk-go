package previewawsdatazonemixins


// Properties for CfnEnvironmentActionsPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnEnvironmentActionsMixinProps := &CfnEnvironmentActionsMixinProps{
//   	Description: jsii.String("description"),
//   	DomainIdentifier: jsii.String("domainIdentifier"),
//   	EnvironmentIdentifier: jsii.String("environmentIdentifier"),
//   	Identifier: jsii.String("identifier"),
//   	Name: jsii.String("name"),
//   	Parameters: &AwsConsoleLinkParametersProperty{
//   		Uri: jsii.String("uri"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-environmentactions.html
//
type CfnEnvironmentActionsMixinProps struct {
	// The environment action description.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-environmentactions.html#cfn-datazone-environmentactions-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The Amazon DataZone domain ID of the environment action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-environmentactions.html#cfn-datazone-environmentactions-domainidentifier
	//
	DomainIdentifier *string `field:"optional" json:"domainIdentifier" yaml:"domainIdentifier"`
	// The environment ID of the environment action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-environmentactions.html#cfn-datazone-environmentactions-environmentidentifier
	//
	EnvironmentIdentifier *string `field:"optional" json:"environmentIdentifier" yaml:"environmentIdentifier"`
	// The ID of the environment action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-environmentactions.html#cfn-datazone-environmentactions-identifier
	//
	Identifier *string `field:"optional" json:"identifier" yaml:"identifier"`
	// The name of the environment action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-environmentactions.html#cfn-datazone-environmentactions-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The parameters of the environment action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-environmentactions.html#cfn-datazone-environmentactions-parameters
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

