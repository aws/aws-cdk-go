package previewawslightsailmixins


// The AWS Region and Availability Zone where the domain was created (read-only).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   locationProperty := &LocationProperty{
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	RegionName: jsii.String("regionName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-domain-location.html
//
type CfnDomainPropsMixin_LocationProperty struct {
	// The Availability Zone.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-domain-location.html#cfn-lightsail-domain-location-availabilityzone
	//
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The AWS Region name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-domain-location.html#cfn-lightsail-domain-location-regionname
	//
	RegionName *string `field:"optional" json:"regionName" yaml:"regionName"`
}

