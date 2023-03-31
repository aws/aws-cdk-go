package awsservicecatalog


// The subtype containing details about the Codestar connection `Type` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   codeStarParametersProperty := &codeStarParametersProperty{
//   	artifactPath: jsii.String("artifactPath"),
//   	branch: jsii.String("branch"),
//   	connectionArn: jsii.String("connectionArn"),
//   	repository: jsii.String("repository"),
//   }
//
type CfnCloudFormationProduct_CodeStarParametersProperty struct {
	// The absolute path wehre the artifact resides within the repo and branch, formatted as "folder/file.json.".
	ArtifactPath *string `field:"required" json:"artifactPath" yaml:"artifactPath"`
	// The specific branch where the artifact resides.
	Branch *string `field:"required" json:"branch" yaml:"branch"`
	// The CodeStar ARN, which is the connection between AWS Service Catalog and the external repository.
	ConnectionArn *string `field:"required" json:"connectionArn" yaml:"connectionArn"`
	// The specific repository where the product’s artifact-to-be-synced resides, formatted as "Account/Repo.".
	Repository *string `field:"required" json:"repository" yaml:"repository"`
}

