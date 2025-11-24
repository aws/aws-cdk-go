package mixinsawsgreengrassv2


// Contains information about a component dependency for a Lambda function component.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   componentDependencyRequirementProperty := &ComponentDependencyRequirementProperty{
//   	DependencyType: jsii.String("dependencyType"),
//   	VersionRequirement: jsii.String("versionRequirement"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-componentversion-componentdependencyrequirement.html
//
type CfnComponentVersionPropsMixin_ComponentDependencyRequirementProperty struct {
	// The type of this dependency. Choose from the following options:.
	//
	// - `SOFT` – The component doesn't restart if the dependency changes state.
	// - `HARD` – The component restarts if the dependency changes state.
	//
	// Default: `HARD`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-componentversion-componentdependencyrequirement.html#cfn-greengrassv2-componentversion-componentdependencyrequirement-dependencytype
	//
	DependencyType *string `field:"optional" json:"dependencyType" yaml:"dependencyType"`
	// The component version requirement for the component dependency.
	//
	// AWS IoT Greengrass uses semantic version constraints. For more information, see [Semantic Versioning](https://docs.aws.amazon.com/https://semver.org/) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-componentversion-componentdependencyrequirement.html#cfn-greengrassv2-componentversion-componentdependencyrequirement-versionrequirement
	//
	VersionRequirement *string `field:"optional" json:"versionRequirement" yaml:"versionRequirement"`
}

