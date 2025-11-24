package mixinsawsroute53


// Properties for CfnCidrCollectionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCidrCollectionMixinProps := &CfnCidrCollectionMixinProps{
//   	Locations: []interface{}{
//   		&LocationProperty{
//   			CidrList: []*string{
//   				jsii.String("cidrList"),
//   			},
//   			LocationName: jsii.String("locationName"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-cidrcollection.html
//
type CfnCidrCollectionMixinProps struct {
	// A complex type that contains information about the list of CIDR locations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-cidrcollection.html#cfn-route53-cidrcollection-locations
	//
	Locations interface{} `field:"optional" json:"locations" yaml:"locations"`
	// The name of a CIDR collection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-cidrcollection.html#cfn-route53-cidrcollection-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

