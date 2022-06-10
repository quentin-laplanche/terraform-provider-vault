package consts

const (
	// common field names
	FieldPath           = "path"
	FieldParameters     = "parameters"
	FieldMethod         = "method"
	FieldNamespace      = "namespace"
	FieldData           = "data"
	FieldMount          = "mount"
	FieldName           = "name"
	FieldVersion        = "version"
	FieldMetadata       = "metadata"
	FieldNames          = "names"
	FieldLeaseID        = "lease_id"
	FieldLeaseDuration  = "lease_duration"
	FieldLeaseRenewable = "lease_renewable"

	// env vars
	EnvVarVaultNamespaceImport = "TERRAFORM_VAULT_NAMESPACE_IMPORT"
	EnvVarSkipChildToken       = "TERRAFORM_VAULT_SKIP_CHILD_TOKEN"
)
