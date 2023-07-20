package awsservicecatalog


// The subtype containing details about the Codestar connection `Type` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   codeStarParametersProperty := &CodeStarParametersProperty{
//   	ArtifactPath: jsii.String("artifactPath"),
//   	Branch: jsii.String("branch"),
//   	ConnectionArn: jsii.String("connectionArn"),
//   	Repository: jsii.String("repository"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicecatalog-cloudformationproduct-codestarparameters.html
//
type CfnCloudFormationProduct_CodeStarParametersProperty struct {
	// The absolute path wehre the artifact resides within the repo and branch, formatted as "folder/file.json.".
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicecatalog-cloudformationproduct-codestarparameters.html#cfn-servicecatalog-cloudformationproduct-codestarparameters-artifactpath
	//
	ArtifactPath *string `field:"required" json:"artifactPath" yaml:"artifactPath"`
	// The specific branch where the artifact resides.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicecatalog-cloudformationproduct-codestarparameters.html#cfn-servicecatalog-cloudformationproduct-codestarparameters-branch
	//
	Branch *string `field:"required" json:"branch" yaml:"branch"`
	// The CodeStar ARN, which is the connection between AWS Service Catalog and the external repository.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicecatalog-cloudformationproduct-codestarparameters.html#cfn-servicecatalog-cloudformationproduct-codestarparameters-connectionarn
	//
	ConnectionArn *string `field:"required" json:"connectionArn" yaml:"connectionArn"`
	// The specific repository where the product’s artifact-to-be-synced resides, formatted as "Account/Repo.".
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicecatalog-cloudformationproduct-codestarparameters.html#cfn-servicecatalog-cloudformationproduct-codestarparameters-repository
	//
	Repository *string `field:"required" json:"repository" yaml:"repository"`
}

