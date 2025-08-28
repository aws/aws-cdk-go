package awsservicecatalog


// Provides connection details.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectionParametersProperty := &ConnectionParametersProperty{
//   	CodeStar: &CodeStarParametersProperty{
//   		ArtifactPath: jsii.String("artifactPath"),
//   		Branch: jsii.String("branch"),
//   		ConnectionArn: jsii.String("connectionArn"),
//   		Repository: jsii.String("repository"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicecatalog-cloudformationproduct-connectionparameters.html
//
type CfnCloudFormationProduct_ConnectionParametersProperty struct {
	// Provides `ConnectionType` details.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicecatalog-cloudformationproduct-connectionparameters.html#cfn-servicecatalog-cloudformationproduct-connectionparameters-codestar
	//
	CodeStar interface{} `field:"optional" json:"codeStar" yaml:"codeStar"`
}

