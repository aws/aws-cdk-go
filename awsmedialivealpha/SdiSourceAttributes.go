package awsmedialivealpha


// Attributes for importing an existing SDI Source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   sdiSourceAttributes := &SdiSourceAttributes{
//   	SdiSourceArn: jsii.String("sdiSourceArn"),
//   	SdiSourceId: jsii.String("sdiSourceId"),
//   }
//
// Experimental.
type SdiSourceAttributes struct {
	// The SDI Source ARN.
	// Experimental.
	SdiSourceArn *string `field:"required" json:"sdiSourceArn" yaml:"sdiSourceArn"`
	// The SDI Source ID.
	// Experimental.
	SdiSourceId *string `field:"required" json:"sdiSourceId" yaml:"sdiSourceId"`
}

