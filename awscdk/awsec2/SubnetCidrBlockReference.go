package awsec2


// A reference to a SubnetCidrBlock resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   subnetCidrBlockReference := &SubnetCidrBlockReference{
//   	SubnetCidrBlockId: jsii.String("subnetCidrBlockId"),
//   }
//
type SubnetCidrBlockReference struct {
	// The Id of the SubnetCidrBlock resource.
	SubnetCidrBlockId *string `field:"required" json:"subnetCidrBlockId" yaml:"subnetCidrBlockId"`
}

