package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var statement interface{}
//
//   iAMPolicyDocumentProperty := map[string]interface{}{
//   	"statement": statement,
//
//   	// the properties below are optional
//   	"version": jsii.String("version"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-iampolicydocument.html
//
type CfnFunction_IAMPolicyDocumentProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-iampolicydocument.html#cfn-serverless-function-iampolicydocument-statement
	//
	Statement interface{} `field:"required" json:"statement" yaml:"statement"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-iampolicydocument.html#cfn-serverless-function-iampolicydocument-version
	//
	Version *string `field:"optional" json:"version" yaml:"version"`
}

