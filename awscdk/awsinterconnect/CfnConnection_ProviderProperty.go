package awsinterconnect


// The partner cloud service provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   providerProperty := &ProviderProperty{
//   	CloudServiceProvider: jsii.String("cloudServiceProvider"),
//   	LastMileProvider: jsii.String("lastMileProvider"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-interconnect-connection-provider.html
//
type CfnConnection_ProviderProperty struct {
	// The name of the cloud service provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-interconnect-connection-provider.html#cfn-interconnect-connection-provider-cloudserviceprovider
	//
	CloudServiceProvider *string `field:"optional" json:"cloudServiceProvider" yaml:"cloudServiceProvider"`
	// The name of the last mile provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-interconnect-connection-provider.html#cfn-interconnect-connection-provider-lastmileprovider
	//
	LastMileProvider *string `field:"optional" json:"lastMileProvider" yaml:"lastMileProvider"`
}

