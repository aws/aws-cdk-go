package awsiotsitewise


// Container associated a certain PortalType.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   portalTypeEntryProperty := &PortalTypeEntryProperty{
//   	PortalTools: []*string{
//   		jsii.String("portalTools"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-portal-portaltypeentry.html
//
type CfnPortal_PortalTypeEntryProperty struct {
	// List of enabled Tools for a certain portal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-portal-portaltypeentry.html#cfn-iotsitewise-portal-portaltypeentry-portaltools
	//
	PortalTools *[]*string `field:"required" json:"portalTools" yaml:"portalTools"`
}

