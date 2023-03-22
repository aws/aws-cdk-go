package awsomics


// Server-side encryption (SSE) settings for a store.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sseConfigProperty := &SseConfigProperty{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	KeyArn: jsii.String("keyArn"),
//   }
//
type CfnVariantStore_SseConfigProperty struct {
	// The encryption type.
	Type *string `field:"required" json:"type" yaml:"type"`
	// An encryption key ARN.
	KeyArn *string `field:"optional" json:"keyArn" yaml:"keyArn"`
}

