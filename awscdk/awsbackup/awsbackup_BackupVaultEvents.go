package awsbackup


// Backup vault events.
// Experimental.
type BackupVaultEvents string

const (
	// BACKUP_JOB_STARTED.
	// Experimental.
	BackupVaultEvents_BACKUP_JOB_STARTED BackupVaultEvents = "BACKUP_JOB_STARTED"
	// BACKUP_JOB_COMPLETED.
	// Experimental.
	BackupVaultEvents_BACKUP_JOB_COMPLETED BackupVaultEvents = "BACKUP_JOB_COMPLETED"
	// BACKUP_JOB_SUCCESSFUL.
	// Experimental.
	BackupVaultEvents_BACKUP_JOB_SUCCESSFUL BackupVaultEvents = "BACKUP_JOB_SUCCESSFUL"
	// BACKUP_JOB_FAILED.
	// Experimental.
	BackupVaultEvents_BACKUP_JOB_FAILED BackupVaultEvents = "BACKUP_JOB_FAILED"
	// BACKUP_JOB_EXPIRED.
	// Experimental.
	BackupVaultEvents_BACKUP_JOB_EXPIRED BackupVaultEvents = "BACKUP_JOB_EXPIRED"
	// RESTORE_JOB_STARTED.
	// Experimental.
	BackupVaultEvents_RESTORE_JOB_STARTED BackupVaultEvents = "RESTORE_JOB_STARTED"
	// RESTORE_JOB_COMPLETED.
	// Experimental.
	BackupVaultEvents_RESTORE_JOB_COMPLETED BackupVaultEvents = "RESTORE_JOB_COMPLETED"
	// RESTORE_JOB_SUCCESSFUL.
	// Experimental.
	BackupVaultEvents_RESTORE_JOB_SUCCESSFUL BackupVaultEvents = "RESTORE_JOB_SUCCESSFUL"
	// RESTORE_JOB_FAILED.
	// Experimental.
	BackupVaultEvents_RESTORE_JOB_FAILED BackupVaultEvents = "RESTORE_JOB_FAILED"
	// COPY_JOB_STARTED.
	// Experimental.
	BackupVaultEvents_COPY_JOB_STARTED BackupVaultEvents = "COPY_JOB_STARTED"
	// COPY_JOB_SUCCESSFUL.
	// Experimental.
	BackupVaultEvents_COPY_JOB_SUCCESSFUL BackupVaultEvents = "COPY_JOB_SUCCESSFUL"
	// COPY_JOB_FAILED.
	// Experimental.
	BackupVaultEvents_COPY_JOB_FAILED BackupVaultEvents = "COPY_JOB_FAILED"
	// RECOVERY_POINT_MODIFIED.
	// Experimental.
	BackupVaultEvents_RECOVERY_POINT_MODIFIED BackupVaultEvents = "RECOVERY_POINT_MODIFIED"
	// BACKUP_PLAN_CREATED.
	// Experimental.
	BackupVaultEvents_BACKUP_PLAN_CREATED BackupVaultEvents = "BACKUP_PLAN_CREATED"
	// BACKUP_PLAN_MODIFIED.
	// Experimental.
	BackupVaultEvents_BACKUP_PLAN_MODIFIED BackupVaultEvents = "BACKUP_PLAN_MODIFIED"
)

