package appstagingsynthesizeralpha


// Context parameters for the 'obtainStagingResources' function.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import app_staging_synthesizer_alpha "github.com/aws/aws-cdk-go/appstagingsynthesizeralpha"
//
//   obtainStagingResourcesContext := &ObtainStagingResourcesContext{
//   	EnvironmentString: jsii.String("environmentString"),
//   	Qualifier: jsii.String("qualifier"),
//
//   	// the properties below are optional
//   	DeployRoleArn: jsii.String("deployRoleArn"),
//   }
//
// Experimental.
type ObtainStagingResourcesContext struct {
	// A unique string describing the environment that is guaranteed not to have tokens in it.
	// Experimental.
	EnvironmentString *string `field:"required" json:"environmentString" yaml:"environmentString"`
	// The qualifier passed to the synthesizer.
	//
	// The staging stack shouldn't need this, but it might.
	// Experimental.
	Qualifier *string `field:"required" json:"qualifier" yaml:"qualifier"`
	// The ARN of the deploy action role, if given.
	//
	// This role will need permissions to read from to the staging resources.
	// Default: - Deploy role ARN is unknown.
	//
	// Experimental.
	DeployRoleArn *string `field:"optional" json:"deployRoleArn" yaml:"deployRoleArn"`
}

