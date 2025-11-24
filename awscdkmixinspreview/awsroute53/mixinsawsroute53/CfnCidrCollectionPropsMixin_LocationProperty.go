package mixinsawsroute53


// Specifies the list of CIDR blocks for a CIDR location.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   locationProperty := &LocationProperty{
//   	CidrList: []*string{
//   		jsii.String("cidrList"),
//   	},
//   	LocationName: jsii.String("locationName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-cidrcollection-location.html
//
type CfnCidrCollectionPropsMixin_LocationProperty struct {
	// List of CIDR blocks.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-cidrcollection-location.html#cfn-route53-cidrcollection-location-cidrlist
	//
	CidrList *[]*string `field:"optional" json:"cidrList" yaml:"cidrList"`
	// The CIDR collection location name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-cidrcollection-location.html#cfn-route53-cidrcollection-location-locationname
	//
	LocationName *string `field:"optional" json:"locationName" yaml:"locationName"`
}

