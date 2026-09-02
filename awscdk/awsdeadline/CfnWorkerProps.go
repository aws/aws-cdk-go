package awsdeadline


// Properties for defining a `CfnWorker`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnWorkerProps := &CfnWorkerProps{
//   	FarmId: jsii.String("farmId"),
//   	FleetId: jsii.String("fleetId"),
//
//   	// the properties below are optional
//   	HostProperties: &HostPropertiesRequestProperty{
//   		HostName: jsii.String("hostName"),
//   		IpAddresses: &IpAddressesProperty{
//   			IpV4Addresses: []*string{
//   				jsii.String("ipV4Addresses"),
//   			},
//   			IpV6Addresses: []*string{
//   				jsii.String("ipV6Addresses"),
//   			},
//   		},
//   	},
//   	Tags: []TagsItemsProperty{
//   		&TagsItemsProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-worker.html
//
type CfnWorkerProps struct {
	// The farm ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-worker.html#cfn-deadline-worker-farmid
	//
	FarmId *string `field:"required" json:"farmId" yaml:"farmId"`
	// The fleet ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-worker.html#cfn-deadline-worker-fleetid
	//
	FleetId *string `field:"required" json:"fleetId" yaml:"fleetId"`
	// The host property details.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-worker.html#cfn-deadline-worker-hostproperties
	//
	HostProperties interface{} `field:"optional" json:"hostProperties" yaml:"hostProperties"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-worker.html#cfn-deadline-worker-tags
	//
	Tags *[]*CfnWorker_TagsItemsProperty `field:"optional" json:"tags" yaml:"tags"`
}

