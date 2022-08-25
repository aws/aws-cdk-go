package awsiam


// Properties for defining an IAM access key.
//
// Example:
//   // Creates a new IAM user, access and secret keys, and stores the secret access key in a Secret.
//   user := iam.NewUser(this, jsii.String("User"))
//   accessKey := iam.NewAccessKey(this, jsii.String("AccessKey"), &accessKeyProps{
//   	user: user,
//   })
//   secret := secretsmanager.NewSecret(this, jsii.String("Secret"), &secretProps{
//   	secretStringValue: accessKey.secretAccessKey,
//   })
//
type AccessKeyProps struct {
	// The IAM user this key will belong to.
	//
	// Changing this value will result in the access key being deleted and a new
	// access key (with a different ID and secret value) being assigned to the new
	// user.
	User IUser `field:"required" json:"user" yaml:"user"`
	// A CloudFormation-specific value that signifies the access key should be replaced/rotated.
	//
	// This value can only be incremented. Incrementing this
	// value will cause CloudFormation to replace the Access Key resource.
	Serial *float64 `field:"optional" json:"serial" yaml:"serial"`
	// The status of the access key.
	//
	// An Active access key is allowed to be used
	// to make API calls; An Inactive key cannot.
	Status AccessKeyStatus `field:"optional" json:"status" yaml:"status"`
}

