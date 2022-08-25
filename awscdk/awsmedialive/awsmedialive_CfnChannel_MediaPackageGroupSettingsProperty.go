package awsmedialive


// The settings for the MediaPackage group.
//
// The parent of this entity is OutputGroupSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mediaPackageGroupSettingsProperty := &mediaPackageGroupSettingsProperty{
//   	destination: &outputLocationRefProperty{
//   		destinationRefId: jsii.String("destinationRefId"),
//   	},
//   }
//
type CfnChannel_MediaPackageGroupSettingsProperty struct {
	// The MediaPackage channel destination.
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
}

