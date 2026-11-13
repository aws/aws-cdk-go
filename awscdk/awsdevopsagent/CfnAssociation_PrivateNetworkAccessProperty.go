package awsdevopsagent


// Private network access to your application inside a VPC, reached through a private connection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   privateNetworkAccessProperty := &PrivateNetworkAccessProperty{
//   	PrivateConnectionName: jsii.String("privateConnectionName"),
//   	RuntimeRoleArn: jsii.String("runtimeRoleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-privatenetworkaccess.html
//
type CfnAssociation_PrivateNetworkAccessProperty struct {
	// Name of the private connection that supplies the VPC configuration for this release management environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-privatenetworkaccess.html#cfn-devopsagent-association-privatenetworkaccess-privateconnectionname
	//
	PrivateConnectionName *string `field:"required" json:"privateConnectionName" yaml:"privateConnectionName"`
	// Role ARN that AWS DevOps Agent assumes at runtime to connect to your VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-privatenetworkaccess.html#cfn-devopsagent-association-privatenetworkaccess-runtimerolearn
	//
	RuntimeRoleArn *string `field:"required" json:"runtimeRoleArn" yaml:"runtimeRoleArn"`
}

