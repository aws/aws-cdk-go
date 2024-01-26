package awsdatazone


// Properties for defining a `CfnEnvironmentProfile`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEnvironmentProfileProps := &CfnEnvironmentProfileProps{
//   	AwsAccountId: jsii.String("awsAccountId"),
//   	AwsAccountRegion: jsii.String("awsAccountRegion"),
//   	DomainIdentifier: jsii.String("domainIdentifier"),
//   	EnvironmentBlueprintIdentifier: jsii.String("environmentBlueprintIdentifier"),
//   	Name: jsii.String("name"),
//   	ProjectIdentifier: jsii.String("projectIdentifier"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	UserParameters: []interface{}{
//   		&EnvironmentParameterProperty{
//   			Name: jsii.String("name"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-environmentprofile.html
//
type CfnEnvironmentProfileProps struct {
	// The identifier of an AWS account in which an environment profile exists.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-environmentprofile.html#cfn-datazone-environmentprofile-awsaccountid
	//
	AwsAccountId *string `field:"required" json:"awsAccountId" yaml:"awsAccountId"`
	// The AWS Region in which an environment profile exists.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-environmentprofile.html#cfn-datazone-environmentprofile-awsaccountregion
	//
	AwsAccountRegion *string `field:"required" json:"awsAccountRegion" yaml:"awsAccountRegion"`
	// The identifier of the Amazon DataZone domain in which the environment profile exists.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-environmentprofile.html#cfn-datazone-environmentprofile-domainidentifier
	//
	DomainIdentifier *string `field:"required" json:"domainIdentifier" yaml:"domainIdentifier"`
	// The identifier of a blueprint with which an environment profile is created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-environmentprofile.html#cfn-datazone-environmentprofile-environmentblueprintidentifier
	//
	EnvironmentBlueprintIdentifier *string `field:"required" json:"environmentBlueprintIdentifier" yaml:"environmentBlueprintIdentifier"`
	// The name of the environment profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-environmentprofile.html#cfn-datazone-environmentprofile-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The identifier of a project in which an environment profile exists.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-environmentprofile.html#cfn-datazone-environmentprofile-projectidentifier
	//
	ProjectIdentifier *string `field:"required" json:"projectIdentifier" yaml:"projectIdentifier"`
	// The description of the environment profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-environmentprofile.html#cfn-datazone-environmentprofile-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The user parameters of this Amazon DataZone environment profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-environmentprofile.html#cfn-datazone-environmentprofile-userparameters
	//
	UserParameters interface{} `field:"optional" json:"userParameters" yaml:"userParameters"`
}

