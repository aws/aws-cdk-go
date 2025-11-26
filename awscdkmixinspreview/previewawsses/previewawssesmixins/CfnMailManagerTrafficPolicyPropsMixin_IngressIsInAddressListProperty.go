package previewawssesmixins


// The address lists and the address list attribute value that is evaluated in a policy statement's conditional expression to either deny or block the incoming email.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
type CfnMailManagerTrafficPolicyPropsMixin_IngressIsInAddressListProperty struct {
	// The address lists that will be used for evaluation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagertrafficpolicy-ingressisinaddresslist.html#cfn-ses-mailmanagertrafficpolicy-ingressisinaddresslist-addresslists
	//
	AddressLists *[]*string `field:"optional" json:"addressLists" yaml:"addressLists"`
	// The email attribute that needs to be evaluated against the address list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagertrafficpolicy-ingressisinaddresslist.html#cfn-ses-mailmanagertrafficpolicy-ingressisinaddresslist-attribute
	//
	Attribute *string `field:"optional" json:"attribute" yaml:"attribute"`
}

