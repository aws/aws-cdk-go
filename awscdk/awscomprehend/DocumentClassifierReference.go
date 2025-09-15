package awscomprehend


// A reference to a DocumentClassifier resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   documentClassifierReference := &DocumentClassifierReference{
//   	DocumentClassifierArn: jsii.String("documentClassifierArn"),
//   }
//
type DocumentClassifierReference struct {
	// The Arn of the DocumentClassifier resource.
	DocumentClassifierArn *string `field:"required" json:"documentClassifierArn" yaml:"documentClassifierArn"`
}

