package awslambda


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   codeArtifactProperty := &CodeArtifactProperty{
//   	Uri: jsii.String("uri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-microvmimage-codeartifact.html
//
type CfnMicrovmImage_CodeArtifactProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-microvmimage-codeartifact.html#cfn-lambda-microvmimage-codeartifact-uri
	//
	Uri *string `field:"required" json:"uri" yaml:"uri"`
}

