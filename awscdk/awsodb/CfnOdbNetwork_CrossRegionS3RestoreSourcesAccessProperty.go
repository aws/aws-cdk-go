package awsodb


// The configuration access for the cross-Region Amazon S3 database restore source for the ODB network.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   crossRegionS3RestoreSourcesAccessProperty := &CrossRegionS3RestoreSourcesAccessProperty{
//   	Ipv4Addresses: []*string{
//   		jsii.String("ipv4Addresses"),
//   	},
//   	Region: jsii.String("region"),
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-odbnetwork-crossregions3restoresourcesaccess.html
//
type CfnOdbNetwork_CrossRegionS3RestoreSourcesAccessProperty struct {
	// The IPv4 addresses allowed for cross-Region Amazon S3 restore access.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-odbnetwork-crossregions3restoresourcesaccess.html#cfn-odb-odbnetwork-crossregions3restoresourcesaccess-ipv4addresses
	//
	Ipv4Addresses *[]*string `field:"optional" json:"ipv4Addresses" yaml:"ipv4Addresses"`
	// The AWS-Region for cross-Region Amazon S3 restore access.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-odbnetwork-crossregions3restoresourcesaccess.html#cfn-odb-odbnetwork-crossregions3restoresourcesaccess-region
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The status of the managed resource access.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-odbnetwork-crossregions3restoresourcesaccess.html#cfn-odb-odbnetwork-crossregions3restoresourcesaccess-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

