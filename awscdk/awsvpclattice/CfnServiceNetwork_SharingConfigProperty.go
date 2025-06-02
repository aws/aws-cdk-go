package awsvpclattice


// Specify if the service network should be enabled for sharing.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sharingConfigProperty := &SharingConfigProperty{
//   	Enabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-servicenetwork-sharingconfig.html
//
type CfnServiceNetwork_SharingConfigProperty struct {
	// Specify if the service network should be enabled for sharing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-servicenetwork-sharingconfig.html#cfn-vpclattice-servicenetwork-sharingconfig-enabled
	//
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

