package awsgameliftstreams


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   locationConfigurationProperty := &LocationConfigurationProperty{
//   	LocationName: jsii.String("locationName"),
//
//   	// the properties below are optional
//   	AlwaysOnCapacity: jsii.Number(123),
//   	OnDemandCapacity: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gameliftstreams-streamgroup-locationconfiguration.html
//
type CfnStreamGroup_LocationConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gameliftstreams-streamgroup-locationconfiguration.html#cfn-gameliftstreams-streamgroup-locationconfiguration-locationname
	//
	LocationName *string `field:"required" json:"locationName" yaml:"locationName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gameliftstreams-streamgroup-locationconfiguration.html#cfn-gameliftstreams-streamgroup-locationconfiguration-alwaysoncapacity
	//
	AlwaysOnCapacity *float64 `field:"optional" json:"alwaysOnCapacity" yaml:"alwaysOnCapacity"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gameliftstreams-streamgroup-locationconfiguration.html#cfn-gameliftstreams-streamgroup-locationconfiguration-ondemandcapacity
	//
	OnDemandCapacity *float64 `field:"optional" json:"onDemandCapacity" yaml:"onDemandCapacity"`
}

