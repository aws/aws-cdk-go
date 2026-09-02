package awsidentitystore


// The name of the user.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   nameProperty := &NameProperty{
//   	FamilyName: jsii.String("familyName"),
//   	Formatted: jsii.String("formatted"),
//   	GivenName: jsii.String("givenName"),
//   	HonorificPrefix: jsii.String("honorificPrefix"),
//   	HonorificSuffix: jsii.String("honorificSuffix"),
//   	MiddleName: jsii.String("middleName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-identitystore-user-name.html
//
type CfnUserPropsMixin_NameProperty struct {
	// The family name of the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-identitystore-user-name.html#cfn-identitystore-user-name-familyname
	//
	FamilyName *string `field:"optional" json:"familyName" yaml:"familyName"`
	// A string containing a formatted version of the name for display.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-identitystore-user-name.html#cfn-identitystore-user-name-formatted
	//
	Formatted *string `field:"optional" json:"formatted" yaml:"formatted"`
	// The given name of the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-identitystore-user-name.html#cfn-identitystore-user-name-givenname
	//
	GivenName *string `field:"optional" json:"givenName" yaml:"givenName"`
	// The honorific prefix of the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-identitystore-user-name.html#cfn-identitystore-user-name-honorificprefix
	//
	HonorificPrefix *string `field:"optional" json:"honorificPrefix" yaml:"honorificPrefix"`
	// The honorific suffix of the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-identitystore-user-name.html#cfn-identitystore-user-name-honorificsuffix
	//
	HonorificSuffix *string `field:"optional" json:"honorificSuffix" yaml:"honorificSuffix"`
	// The middle name of the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-identitystore-user-name.html#cfn-identitystore-user-name-middlename
	//
	MiddleName *string `field:"optional" json:"middleName" yaml:"middleName"`
}

