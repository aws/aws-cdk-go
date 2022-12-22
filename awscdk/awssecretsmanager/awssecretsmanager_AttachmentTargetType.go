package awssecretsmanager


// The type of service or database that's being associated with the secret.
// Experimental.
type AttachmentTargetType string

const (
	// A database instance.
	// Deprecated: use RDS_DB_INSTANCE instead.
	AttachmentTargetType_INSTANCE AttachmentTargetType = "INSTANCE"
	// A database cluster.
	// Deprecated: use RDS_DB_CLUSTER instead.
	AttachmentTargetType_CLUSTER AttachmentTargetType = "CLUSTER"
	// AWS::RDS::DBInstance.
	// Experimental.
	AttachmentTargetType_RDS_DB_INSTANCE AttachmentTargetType = "RDS_DB_INSTANCE"
	// AWS::RDS::DBCluster.
	// Experimental.
	AttachmentTargetType_RDS_DB_CLUSTER AttachmentTargetType = "RDS_DB_CLUSTER"
	// AWS::RDS::DBProxy.
	// Experimental.
	AttachmentTargetType_RDS_DB_PROXY AttachmentTargetType = "RDS_DB_PROXY"
	// AWS::Redshift::Cluster.
	// Experimental.
	AttachmentTargetType_REDSHIFT_CLUSTER AttachmentTargetType = "REDSHIFT_CLUSTER"
	// AWS::DocDB::DBInstance.
	// Experimental.
	AttachmentTargetType_DOCDB_DB_INSTANCE AttachmentTargetType = "DOCDB_DB_INSTANCE"
	// AWS::DocDB::DBCluster.
	// Experimental.
	AttachmentTargetType_DOCDB_DB_CLUSTER AttachmentTargetType = "DOCDB_DB_CLUSTER"
)

