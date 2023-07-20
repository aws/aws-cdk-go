package awsec2


// An entry for a prefix list.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   entryProperty := &EntryProperty{
//   	Cidr: jsii.String("cidr"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-prefixlist-entry.html
//
type CfnPrefixList_EntryProperty struct {
	// The CIDR block.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-prefixlist-entry.html#cfn-ec2-prefixlist-entry-cidr
	//
	Cidr *string `field:"required" json:"cidr" yaml:"cidr"`
	// A description for the entry.
	//
	// Constraints: Up to 255 characters in length.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-prefixlist-entry.html#cfn-ec2-prefixlist-entry-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

