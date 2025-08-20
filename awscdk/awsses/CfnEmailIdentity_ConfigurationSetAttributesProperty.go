package awsses


// Used to associate a configuration set with an email identity.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   configurationSetAttributesProperty := &ConfigurationSetAttributesProperty{
//   	ConfigurationSetName: jsii.String("configurationSetName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-emailidentity-configurationsetattributes.html
//
type CfnEmailIdentity_ConfigurationSetAttributesProperty struct {
	// The configuration set to associate with an email identity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-emailidentity-configurationsetattributes.html#cfn-ses-emailidentity-configurationsetattributes-configurationsetname
	//
	ConfigurationSetName *string `field:"optional" json:"configurationSetName" yaml:"configurationSetName"`
}

