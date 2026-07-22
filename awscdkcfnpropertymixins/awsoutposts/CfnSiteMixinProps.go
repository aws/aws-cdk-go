package awsoutposts

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnSitePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnSiteMixinProps := &CfnSiteMixinProps{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Notes: jsii.String("notes"),
//   	OperatingAddress: &AddressProperty{
//   		AddressLine1: jsii.String("addressLine1"),
//   		AddressLine2: jsii.String("addressLine2"),
//   		AddressLine3: jsii.String("addressLine3"),
//   		City: jsii.String("city"),
//   		ContactName: jsii.String("contactName"),
//   		ContactPhoneNumber: jsii.String("contactPhoneNumber"),
//   		CountryCode: jsii.String("countryCode"),
//   		DistrictOrCounty: jsii.String("districtOrCounty"),
//   		Municipality: jsii.String("municipality"),
//   		PostalCode: jsii.String("postalCode"),
//   		StateOrRegion: jsii.String("stateOrRegion"),
//   	},
//   	RackPhysicalProperties: &RackPhysicalPropertiesProperty{
//   		FiberOpticCableType: jsii.String("fiberOpticCableType"),
//   		MaximumSupportedWeightLbs: jsii.String("maximumSupportedWeightLbs"),
//   		OpticalStandard: jsii.String("opticalStandard"),
//   		PowerConnector: jsii.String("powerConnector"),
//   		PowerDrawKva: jsii.String("powerDrawKva"),
//   		PowerFeedDrop: jsii.String("powerFeedDrop"),
//   		PowerPhase: jsii.String("powerPhase"),
//   		UplinkCount: jsii.String("uplinkCount"),
//   		UplinkGbps: jsii.String("uplinkGbps"),
//   	},
//   	ShippingAddress: &AddressProperty{
//   		AddressLine1: jsii.String("addressLine1"),
//   		AddressLine2: jsii.String("addressLine2"),
//   		AddressLine3: jsii.String("addressLine3"),
//   		City: jsii.String("city"),
//   		ContactName: jsii.String("contactName"),
//   		ContactPhoneNumber: jsii.String("contactPhoneNumber"),
//   		CountryCode: jsii.String("countryCode"),
//   		DistrictOrCounty: jsii.String("districtOrCounty"),
//   		Municipality: jsii.String("municipality"),
//   		PostalCode: jsii.String("postalCode"),
//   		StateOrRegion: jsii.String("stateOrRegion"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-outposts-site.html
//
type CfnSiteMixinProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-outposts-site.html#cfn-outposts-site-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-outposts-site.html#cfn-outposts-site-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-outposts-site.html#cfn-outposts-site-notes
	//
	Notes *string `field:"optional" json:"notes" yaml:"notes"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-outposts-site.html#cfn-outposts-site-operatingaddress
	//
	OperatingAddress interface{} `field:"optional" json:"operatingAddress" yaml:"operatingAddress"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-outposts-site.html#cfn-outposts-site-rackphysicalproperties
	//
	RackPhysicalProperties interface{} `field:"optional" json:"rackPhysicalProperties" yaml:"rackPhysicalProperties"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-outposts-site.html#cfn-outposts-site-shippingaddress
	//
	ShippingAddress interface{} `field:"optional" json:"shippingAddress" yaml:"shippingAddress"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-outposts-site.html#cfn-outposts-site-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

