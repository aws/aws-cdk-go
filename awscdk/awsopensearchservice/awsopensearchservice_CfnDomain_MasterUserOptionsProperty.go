package awsopensearchservice


// Specifies information about the master user.
//
// Required if if `InternalUserDatabaseEnabled` is true in `AdvancedSecurityOptions` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   masterUserOptionsProperty := &masterUserOptionsProperty{
//   	masterUserArn: jsii.String("masterUserArn"),
//   	masterUserName: jsii.String("masterUserName"),
//   	masterUserPassword: jsii.String("masterUserPassword"),
//   }
//
type CfnDomain_MasterUserOptionsProperty struct {
	// ARN for the master user.
	//
	// Only specify if `InternalUserDatabaseEnabled` is false in `AdvancedSecurityOptions` .
	MasterUserArn *string `field:"optional" json:"masterUserArn" yaml:"masterUserArn"`
	// Username for the master user.
	//
	// Only specify if `InternalUserDatabaseEnabled` is true in `AdvancedSecurityOptions` . If you don't want to specify this value directly within the template, you can use a [dynamic reference](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/dynamic-references.html) instead.
	MasterUserName *string `field:"optional" json:"masterUserName" yaml:"masterUserName"`
	// Password for the master user.
	//
	// Only specify if `InternalUserDatabaseEnabled` is true in `AdvancedSecurityOptions` . If you don't want to specify this value directly within the template, you can use a [dynamic reference](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/dynamic-references.html) instead.
	MasterUserPassword *string `field:"optional" json:"masterUserPassword" yaml:"masterUserPassword"`
}

