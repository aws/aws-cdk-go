package awsrds


// Properties for looking up an existing DatabaseInstance.
//
// Example:
//   var myUserRole Role
//
//
//   dbFromLookup := rds.DatabaseInstance_FromLookup(this, jsii.String("dbFromLookup"), &DatabaseInstanceLookupOptions{
//   	InstanceIdentifier: jsii.String("instanceId"),
//   })
//
//   // Grant a connection
//   dbFromLookup.GrantConnect(myUserRole, jsii.String("my-user-id"))
//
type DatabaseInstanceLookupOptions struct {
	// The instance identifier of the DatabaseInstance.
	InstanceIdentifier *string `field:"required" json:"instanceIdentifier" yaml:"instanceIdentifier"`
}

