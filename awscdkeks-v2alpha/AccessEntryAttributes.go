package awscdkeks-v2alpha


// Represents the attributes of an access entry.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import eks_v2_alpha "github.com/aws/aws-cdk-go/awscdkeks-v2alpha"
//
//   accessEntryAttributes := &AccessEntryAttributes{
//   	AccessEntryArn: jsii.String("accessEntryArn"),
//   	AccessEntryName: jsii.String("accessEntryName"),
//   }
//
// Experimental.
type AccessEntryAttributes struct {
	// The Amazon Resource Name (ARN) of the access entry.
	// Experimental.
	AccessEntryArn *string `field:"required" json:"accessEntryArn" yaml:"accessEntryArn"`
	// The name of the access entry.
	// Experimental.
	AccessEntryName *string `field:"required" json:"accessEntryName" yaml:"accessEntryName"`
}

