package awsses


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ingressIsInAddressListProperty := &IngressIsInAddressListProperty{
//   	AddressLists: []*string{
//   		jsii.String("addressLists"),
//   	},
//   	Attribute: jsii.String("attribute"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagertrafficpolicy-ingressisinaddresslist.html
//
type CfnMailManagerTrafficPolicy_IngressIsInAddressListProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagertrafficpolicy-ingressisinaddresslist.html#cfn-ses-mailmanagertrafficpolicy-ingressisinaddresslist-addresslists
	//
	AddressLists *[]*string `field:"required" json:"addressLists" yaml:"addressLists"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagertrafficpolicy-ingressisinaddresslist.html#cfn-ses-mailmanagertrafficpolicy-ingressisinaddresslist-attribute
	//
	Attribute *string `field:"required" json:"attribute" yaml:"attribute"`
}

