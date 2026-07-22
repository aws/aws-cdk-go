package awscontrolcatalog


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   regionConfigurationProperty := &RegionConfigurationProperty{
//   	Scope: jsii.String("scope"),
//
//   	// the properties below are optional
//   	DeployableRegions: []*string{
//   		jsii.String("deployableRegions"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-controlcatalog-control-regionconfiguration.html
//
type CfnControl_RegionConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-controlcatalog-control-regionconfiguration.html#cfn-controlcatalog-control-regionconfiguration-scope
	//
	Scope *string `field:"required" json:"scope" yaml:"scope"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-controlcatalog-control-regionconfiguration.html#cfn-controlcatalog-control-regionconfiguration-deployableregions
	//
	DeployableRegions *[]*string `field:"optional" json:"deployableRegions" yaml:"deployableRegions"`
}

