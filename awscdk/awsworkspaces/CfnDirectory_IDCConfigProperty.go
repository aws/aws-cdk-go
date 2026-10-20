package awsworkspaces


// Specifies the configurations of the identity center.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iDCConfigProperty := map[string]*string{
//   	"applicationArn": jsii.String("applicationArn"),
//   	"instanceArn": jsii.String("instanceArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-idcconfig.html
//
type CfnDirectory_IDCConfigProperty struct {
	// The Amazon Resource Name (ARN) of the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-idcconfig.html#cfn-workspaces-directory-idcconfig-applicationarn
	//
	ApplicationArn *string `field:"optional" json:"applicationArn" yaml:"applicationArn"`
	// The Amazon Resource Name (ARN) of the identity center instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-idcconfig.html#cfn-workspaces-directory-idcconfig-instancearn
	//
	InstanceArn *string `field:"optional" json:"instanceArn" yaml:"instanceArn"`
}

