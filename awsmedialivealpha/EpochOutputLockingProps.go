package awsmedialivealpha


// Properties for epoch output locking.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   epochOutputLockingProps := &EpochOutputLockingProps{
//   	CustomEpoch: jsii.String("customEpoch"),
//   	JamSyncTime: jsii.String("jamSyncTime"),
//   }
//
// Experimental.
type EpochOutputLockingProps struct {
	// A custom epoch (ISO-8601 timestamp) to lock outputs to.
	// Default: - service default.
	//
	// Experimental.
	CustomEpoch *string `field:"optional" json:"customEpoch" yaml:"customEpoch"`
	// A jam-sync time (ISO-8601 timestamp).
	// Default: - service default.
	//
	// Experimental.
	JamSyncTime *string `field:"optional" json:"jamSyncTime" yaml:"jamSyncTime"`
}

