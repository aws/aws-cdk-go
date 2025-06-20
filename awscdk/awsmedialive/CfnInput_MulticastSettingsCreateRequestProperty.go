package awsmedialive


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   multicastSettingsCreateRequestProperty := &MulticastSettingsCreateRequestProperty{
//   	Sources: []interface{}{
//   		&MulticastSourceCreateRequestProperty{
//   			SourceIp: jsii.String("sourceIp"),
//   			Url: jsii.String("url"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-multicastsettingscreaterequest.html
//
type CfnInput_MulticastSettingsCreateRequestProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-multicastsettingscreaterequest.html#cfn-medialive-input-multicastsettingscreaterequest-sources
	//
	Sources interface{} `field:"optional" json:"sources" yaml:"sources"`
}

