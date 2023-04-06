package awsservicecatalog


// A top level `ProductViewDetail` response containing details about the product’s connection.
//
// AWS Service Catalog returns this field for the `CreateProduct` , `UpdateProduct` , `DescribeProductAsAdmin` , and `SearchProductAsAdmin` APIs. This response contains the same fields as the `ConnectionParameters` request, with the addition of the `LastSync` response.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceConnectionProperty := &SourceConnectionProperty{
//   	ConnectionParameters: &ConnectionParametersProperty{
//   		CodeStar: &CodeStarParametersProperty{
//   			ArtifactPath: jsii.String("artifactPath"),
//   			Branch: jsii.String("branch"),
//   			ConnectionArn: jsii.String("connectionArn"),
//   			Repository: jsii.String("repository"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }
//
type CfnCloudFormationProduct_SourceConnectionProperty struct {
	// The connection details based on the connection `Type` .
	ConnectionParameters interface{} `field:"required" json:"connectionParameters" yaml:"connectionParameters"`
	// The only supported `SourceConnection` type is Codestar.
	Type *string `field:"required" json:"type" yaml:"type"`
}

