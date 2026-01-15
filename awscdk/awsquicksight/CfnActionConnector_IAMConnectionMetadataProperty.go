package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iAMConnectionMetadataProperty := map[string]*string{
//   	"roleArn": jsii.String("roleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-actionconnector-iamconnectionmetadata.html
//
type CfnActionConnector_IAMConnectionMetadataProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-actionconnector-iamconnectionmetadata.html#cfn-quicksight-actionconnector-iamconnectionmetadata-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

