package awseks


// Represents the attributes of an access entry.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessEntryAttributes := &AccessEntryAttributes{
//   	AccessEntryArn: jsii.String("accessEntryArn"),
//   	AccessEntryName: jsii.String("accessEntryName"),
//   }
//
type AccessEntryAttributes struct {
	// The Amazon Resource Name (ARN) of the access entry.
	AccessEntryArn *string `field:"required" json:"accessEntryArn" yaml:"accessEntryArn"`
	// The name of the access entry.
	AccessEntryName *string `field:"required" json:"accessEntryName" yaml:"accessEntryName"`
}

