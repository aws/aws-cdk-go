package awsopensearchservice


// Off-peak window settings for the domain.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   offPeakWindowOptionsProperty := &OffPeakWindowOptionsProperty{
//   	Enabled: jsii.Boolean(false),
//   	OffPeakWindow: &OffPeakWindowProperty{
//   		WindowStartTime: &WindowStartTimeProperty{
//   			Hours: jsii.Number(123),
//   			Minutes: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnDomain_OffPeakWindowOptionsProperty struct {
	// Specifies whether off-peak window settings are enabled for the domain.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Off-peak window settings for the domain.
	OffPeakWindow interface{} `field:"optional" json:"offPeakWindow" yaml:"offPeakWindow"`
}

