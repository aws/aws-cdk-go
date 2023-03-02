package awsec2


// The operating Regions for an IPAM.
//
// Operating Regions are AWS Regions where the IPAM is allowed to manage IP address CIDRs. IPAM only discovers and monitors resources in the AWS Regions you select as operating Regions.
//
// For more information about operating Regions, see [Create an IPAM](https://docs.aws.amazon.com//vpc/latest/ipam/create-ipam.html) in the *Amazon VPC IPAM User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ipamOperatingRegionProperty := &ipamOperatingRegionProperty{
//   	regionName: jsii.String("regionName"),
//   }
//
type CfnIPAM_IpamOperatingRegionProperty struct {
	// The name of the operating Region.
	RegionName *string `field:"required" json:"regionName" yaml:"regionName"`
}

