package controllers

type WatchFilter struct {
	// Value is the label value used to filter events prior to reconciliation.
	Value string
	// List of namespaces to exclude while filtering events prior to reconciliation.
	ExcludedNamespaces []string
}
