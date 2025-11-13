package interfacesawsappstream


// A reference to a AppBlockBuilder resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   appBlockBuilderReference := &AppBlockBuilderReference{
//   	AppBlockBuilderArn: jsii.String("appBlockBuilderArn"),
//   	AppBlockBuilderName: jsii.String("appBlockBuilderName"),
//   }
//
type AppBlockBuilderReference struct {
	// The ARN of the AppBlockBuilder resource.
	AppBlockBuilderArn *string `field:"required" json:"appBlockBuilderArn" yaml:"appBlockBuilderArn"`
	// The Name of the AppBlockBuilder resource.
	AppBlockBuilderName *string `field:"required" json:"appBlockBuilderName" yaml:"appBlockBuilderName"`
}

