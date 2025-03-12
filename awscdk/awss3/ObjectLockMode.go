package awss3


// Modes in which S3 Object Lock retention can be configured.
// See: https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lock-overview.html#object-lock-retention-modes
//
type ObjectLockMode string

const (
	// The Governance retention mode.
	//
	// With governance mode, you protect objects against being deleted by most users, but you can
	// still grant some users permission to alter the retention settings or delete the object if
	// necessary. You can also use governance mode to test retention-period settings before
	// creating a compliance-mode retention period.
	ObjectLockMode_GOVERNANCE ObjectLockMode = "GOVERNANCE"
	// The Compliance retention mode.
	//
	// When an object is locked in compliance mode, its retention mode can't be changed, and
	// its retention period can't be shortened. Compliance mode helps ensure that an object
	// version can't be overwritten or deleted for the duration of the retention period.
	ObjectLockMode_COMPLIANCE ObjectLockMode = "COMPLIANCE"
)

