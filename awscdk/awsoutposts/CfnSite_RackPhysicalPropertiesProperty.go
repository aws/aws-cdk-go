package awsoutposts


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rackPhysicalPropertiesProperty := &RackPhysicalPropertiesProperty{
//   	FiberOpticCableType: jsii.String("fiberOpticCableType"),
//   	MaximumSupportedWeightLbs: jsii.String("maximumSupportedWeightLbs"),
//   	OpticalStandard: jsii.String("opticalStandard"),
//   	PowerConnector: jsii.String("powerConnector"),
//   	PowerDrawKva: jsii.String("powerDrawKva"),
//   	PowerFeedDrop: jsii.String("powerFeedDrop"),
//   	PowerPhase: jsii.String("powerPhase"),
//   	UplinkCount: jsii.String("uplinkCount"),
//   	UplinkGbps: jsii.String("uplinkGbps"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-outposts-site-rackphysicalproperties.html
//
type CfnSite_RackPhysicalPropertiesProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-outposts-site-rackphysicalproperties.html#cfn-outposts-site-rackphysicalproperties-fiberopticcabletype
	//
	FiberOpticCableType *string `field:"optional" json:"fiberOpticCableType" yaml:"fiberOpticCableType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-outposts-site-rackphysicalproperties.html#cfn-outposts-site-rackphysicalproperties-maximumsupportedweightlbs
	//
	MaximumSupportedWeightLbs *string `field:"optional" json:"maximumSupportedWeightLbs" yaml:"maximumSupportedWeightLbs"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-outposts-site-rackphysicalproperties.html#cfn-outposts-site-rackphysicalproperties-opticalstandard
	//
	OpticalStandard *string `field:"optional" json:"opticalStandard" yaml:"opticalStandard"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-outposts-site-rackphysicalproperties.html#cfn-outposts-site-rackphysicalproperties-powerconnector
	//
	PowerConnector *string `field:"optional" json:"powerConnector" yaml:"powerConnector"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-outposts-site-rackphysicalproperties.html#cfn-outposts-site-rackphysicalproperties-powerdrawkva
	//
	PowerDrawKva *string `field:"optional" json:"powerDrawKva" yaml:"powerDrawKva"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-outposts-site-rackphysicalproperties.html#cfn-outposts-site-rackphysicalproperties-powerfeeddrop
	//
	PowerFeedDrop *string `field:"optional" json:"powerFeedDrop" yaml:"powerFeedDrop"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-outposts-site-rackphysicalproperties.html#cfn-outposts-site-rackphysicalproperties-powerphase
	//
	PowerPhase *string `field:"optional" json:"powerPhase" yaml:"powerPhase"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-outposts-site-rackphysicalproperties.html#cfn-outposts-site-rackphysicalproperties-uplinkcount
	//
	UplinkCount *string `field:"optional" json:"uplinkCount" yaml:"uplinkCount"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-outposts-site-rackphysicalproperties.html#cfn-outposts-site-rackphysicalproperties-uplinkgbps
	//
	UplinkGbps *string `field:"optional" json:"uplinkGbps" yaml:"uplinkGbps"`
}

