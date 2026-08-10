package awsredshiftserverless


// Properties for defining a `CfnRecoveryPoint`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRecoveryPointProps := &CfnRecoveryPointProps{
//   	NamespaceName: jsii.String("namespaceName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshiftserverless-recoverypoint.html
//
type CfnRecoveryPointProps struct {
	// The name of the namespace the recovery point is associated with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshiftserverless-recoverypoint.html#cfn-redshiftserverless-recoverypoint-namespacename
	//
	NamespaceName *string `field:"optional" json:"namespaceName" yaml:"namespaceName"`
}

