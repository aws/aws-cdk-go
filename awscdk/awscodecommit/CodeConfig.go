package awscodecommit


// Represents the structure to pass into the underlying CfnRepository class.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   codeConfig := &CodeConfig{
//   	Code: &CodeProperty{
//   		S3: &S3Property{
//   			Bucket: jsii.String("bucket"),
//   			Key: jsii.String("key"),
//
//   			// the properties below are optional
//   			ObjectVersion: jsii.String("objectVersion"),
//   		},
//
//   		// the properties below are optional
//   		BranchName: jsii.String("branchName"),
//   	},
//   }
//
type CodeConfig struct {
	// represents the underlying code structure.
	Code *CfnRepository_CodeProperty `field:"required" json:"code" yaml:"code"`
}

