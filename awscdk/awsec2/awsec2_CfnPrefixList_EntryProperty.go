package awsec2


// An entry for a prefix list.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   entryProperty := &entryProperty{
//   	cidr: jsii.String("cidr"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   }
//
type CfnPrefixList_EntryProperty struct {
	// The CIDR block.
	Cidr *string `field:"required" json:"cidr" yaml:"cidr"`
	// A description for the entry.
	//
	// Constraints: Up to 255 characters in length.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

