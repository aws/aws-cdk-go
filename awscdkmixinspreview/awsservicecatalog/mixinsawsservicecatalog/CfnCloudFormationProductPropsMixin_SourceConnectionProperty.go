package mixinsawsservicecatalog


// A top level `ProductViewDetail` response containing details about the product’s connection.
//
// AWS Service Catalog returns this field for the `CreateProduct` , `UpdateProduct` , `DescribeProductAsAdmin` , and `SearchProductAsAdmin` APIs. This response contains the same fields as the `ConnectionParameters` request, with the addition of the `LastSync` response.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicecatalog-cloudformationproduct-sourceconnection.html
//
type CfnCloudFormationProductPropsMixin_SourceConnectionProperty struct {
	// The connection details based on the connection `Type` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicecatalog-cloudformationproduct-sourceconnection.html#cfn-servicecatalog-cloudformationproduct-sourceconnection-connectionparameters
	//
	ConnectionParameters interface{} `field:"optional" json:"connectionParameters" yaml:"connectionParameters"`
	// The only supported `SourceConnection` type is Codestar.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicecatalog-cloudformationproduct-sourceconnection.html#cfn-servicecatalog-cloudformationproduct-sourceconnection-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

