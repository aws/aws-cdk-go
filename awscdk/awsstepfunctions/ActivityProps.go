package awsstepfunctions


// Properties for defining a new Step Functions Activity.
//
// Example:
//   import kms "github.com/aws/aws-cdk-go/awscdk"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//
//   kmsKey := kms.NewKey(this, jsii.String("Key"))
//   activity := sfn.NewActivity(this, jsii.String("ActivityWithCMKEncryptionConfiguration"), &ActivityProps{
//   	ActivityName: jsii.String("ActivityWithCMKEncryptionConfiguration"),
//   	EncryptionConfiguration: sfn.NewCustomerManagedEncryptionConfiguration(kmsKey, cdk.Duration_Seconds(jsii.Number(75))),
//   })
//
type ActivityProps struct {
	// The name for this activity.
	// Default: - If not supplied, a name is generated.
	//
	ActivityName *string `field:"optional" json:"activityName" yaml:"activityName"`
	// The encryptionConfiguration object used for server-side encryption of the activity inputs.
	// Default: - data is transparently encrypted using an AWS owned key.
	//
	EncryptionConfiguration EncryptionConfiguration `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
}

