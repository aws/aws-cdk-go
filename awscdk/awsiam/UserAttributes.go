package awsiam


// Represents a user defined outside of this stack.
//
// Example:
//   user := iam.User_FromUserAttributes(this, jsii.String("MyImportedUserByAttributes"), &UserAttributes{
//   	UserArn: jsii.String("arn:aws:iam::123456789012:user/johnsmith"),
//   })
//
type UserAttributes struct {
	// The ARN of the user.
	//
	// Format: arn:<partition>:iam::<account-id>:user/<user-name-with-path>.
	UserArn *string `field:"required" json:"userArn" yaml:"userArn"`
}

