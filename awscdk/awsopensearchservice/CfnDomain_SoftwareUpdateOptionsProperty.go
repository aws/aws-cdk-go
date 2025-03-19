package awsopensearchservice


// Options for configuring service software updates for a domain.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   softwareUpdateOptionsProperty := &SoftwareUpdateOptionsProperty{
//   	AutoSoftwareUpdateEnabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-softwareupdateoptions.html
//
type CfnDomain_SoftwareUpdateOptionsProperty struct {
	// Specifies whether automatic service software updates are enabled for the domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-softwareupdateoptions.html#cfn-opensearchservice-domain-softwareupdateoptions-autosoftwareupdateenabled
	//
	AutoSoftwareUpdateEnabled interface{} `field:"optional" json:"autoSoftwareUpdateEnabled" yaml:"autoSoftwareUpdateEnabled"`
}

