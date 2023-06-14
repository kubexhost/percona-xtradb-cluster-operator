// Code generated by go-swagger; DO NOT EDIT.

package version_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewVersionServiceOperatorParams creates a new VersionServiceOperatorParams object
// with the default values initialized.
func NewVersionServiceOperatorParams() *VersionServiceOperatorParams {
	var ()
	return &VersionServiceOperatorParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVersionServiceOperatorParamsWithTimeout creates a new VersionServiceOperatorParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVersionServiceOperatorParamsWithTimeout(timeout time.Duration) *VersionServiceOperatorParams {
	var ()
	return &VersionServiceOperatorParams{

		timeout: timeout,
	}
}

// NewVersionServiceOperatorParamsWithContext creates a new VersionServiceOperatorParams object
// with the default values initialized, and the ability to set a context for a request
func NewVersionServiceOperatorParamsWithContext(ctx context.Context) *VersionServiceOperatorParams {
	var ()
	return &VersionServiceOperatorParams{

		Context: ctx,
	}
}

// NewVersionServiceOperatorParamsWithHTTPClient creates a new VersionServiceOperatorParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewVersionServiceOperatorParamsWithHTTPClient(client *http.Client) *VersionServiceOperatorParams {
	var ()
	return &VersionServiceOperatorParams{
		HTTPClient: client,
	}
}

/*VersionServiceOperatorParams contains all the parameters to send to the API endpoint
for the version service operator operation typically these are written to a http.Request
*/
type VersionServiceOperatorParams struct {

	/*BackupVersion*/
	BackupVersion *string
	/*BackupsEnabled*/
	BackupsEnabled *bool
	/*ClusterSize*/
	ClusterSize *int32
	/*ClusterWideEnabled*/
	ClusterWideEnabled *bool
	/*CustomResourceUID*/
	CustomResourceUID *string
	/*DatabaseVersion*/
	DatabaseVersion *string
	/*HaproxyVersion*/
	HaproxyVersion *string
	/*HashicorpVaultEnabled*/
	HashicorpVaultEnabled *bool
	/*HelmDeployCr*/
	HelmDeployCr *bool
	/*HelmDeployOperator*/
	HelmDeployOperator *bool
	/*KubeVersion*/
	KubeVersion *string
	/*LogCollectorVersion*/
	LogCollectorVersion *string
	/*NamespaceUID*/
	NamespaceUID *string
	/*OperatorVersion*/
	OperatorVersion string
	/*PhysicalBackupScheduled*/
	PhysicalBackupScheduled *bool
	/*PitrEnabled*/
	PitrEnabled *bool
	/*Platform*/
	Platform *string
	/*PmmEnabled*/
	PmmEnabled *bool
	/*PmmVersion*/
	PmmVersion *string
	/*Product*/
	Product string
	/*ProxysqlScheduler*/
	ProxysqlScheduler *string
	/*ProxysqlVersion*/
	ProxysqlVersion *string
	/*ShardingEnabled*/
	ShardingEnabled *bool
	/*SidecarsUsed*/
	SidecarsUsed *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the version service operator params
func (o *VersionServiceOperatorParams) WithTimeout(timeout time.Duration) *VersionServiceOperatorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the version service operator params
func (o *VersionServiceOperatorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the version service operator params
func (o *VersionServiceOperatorParams) WithContext(ctx context.Context) *VersionServiceOperatorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the version service operator params
func (o *VersionServiceOperatorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the version service operator params
func (o *VersionServiceOperatorParams) WithHTTPClient(client *http.Client) *VersionServiceOperatorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the version service operator params
func (o *VersionServiceOperatorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBackupVersion adds the backupVersion to the version service operator params
func (o *VersionServiceOperatorParams) WithBackupVersion(backupVersion *string) *VersionServiceOperatorParams {
	o.SetBackupVersion(backupVersion)
	return o
}

// SetBackupVersion adds the backupVersion to the version service operator params
func (o *VersionServiceOperatorParams) SetBackupVersion(backupVersion *string) {
	o.BackupVersion = backupVersion
}

// WithBackupsEnabled adds the backupsEnabled to the version service operator params
func (o *VersionServiceOperatorParams) WithBackupsEnabled(backupsEnabled *bool) *VersionServiceOperatorParams {
	o.SetBackupsEnabled(backupsEnabled)
	return o
}

// SetBackupsEnabled adds the backupsEnabled to the version service operator params
func (o *VersionServiceOperatorParams) SetBackupsEnabled(backupsEnabled *bool) {
	o.BackupsEnabled = backupsEnabled
}

// WithClusterSize adds the clusterSize to the version service operator params
func (o *VersionServiceOperatorParams) WithClusterSize(clusterSize *int32) *VersionServiceOperatorParams {
	o.SetClusterSize(clusterSize)
	return o
}

// SetClusterSize adds the clusterSize to the version service operator params
func (o *VersionServiceOperatorParams) SetClusterSize(clusterSize *int32) {
	o.ClusterSize = clusterSize
}

// WithClusterWideEnabled adds the clusterWideEnabled to the version service operator params
func (o *VersionServiceOperatorParams) WithClusterWideEnabled(clusterWideEnabled *bool) *VersionServiceOperatorParams {
	o.SetClusterWideEnabled(clusterWideEnabled)
	return o
}

// SetClusterWideEnabled adds the clusterWideEnabled to the version service operator params
func (o *VersionServiceOperatorParams) SetClusterWideEnabled(clusterWideEnabled *bool) {
	o.ClusterWideEnabled = clusterWideEnabled
}

// WithCustomResourceUID adds the customResourceUID to the version service operator params
func (o *VersionServiceOperatorParams) WithCustomResourceUID(customResourceUID *string) *VersionServiceOperatorParams {
	o.SetCustomResourceUID(customResourceUID)
	return o
}

// SetCustomResourceUID adds the customResourceUid to the version service operator params
func (o *VersionServiceOperatorParams) SetCustomResourceUID(customResourceUID *string) {
	o.CustomResourceUID = customResourceUID
}

// WithDatabaseVersion adds the databaseVersion to the version service operator params
func (o *VersionServiceOperatorParams) WithDatabaseVersion(databaseVersion *string) *VersionServiceOperatorParams {
	o.SetDatabaseVersion(databaseVersion)
	return o
}

// SetDatabaseVersion adds the databaseVersion to the version service operator params
func (o *VersionServiceOperatorParams) SetDatabaseVersion(databaseVersion *string) {
	o.DatabaseVersion = databaseVersion
}

// WithHaproxyVersion adds the haproxyVersion to the version service operator params
func (o *VersionServiceOperatorParams) WithHaproxyVersion(haproxyVersion *string) *VersionServiceOperatorParams {
	o.SetHaproxyVersion(haproxyVersion)
	return o
}

// SetHaproxyVersion adds the haproxyVersion to the version service operator params
func (o *VersionServiceOperatorParams) SetHaproxyVersion(haproxyVersion *string) {
	o.HaproxyVersion = haproxyVersion
}

// WithHashicorpVaultEnabled adds the hashicorpVaultEnabled to the version service operator params
func (o *VersionServiceOperatorParams) WithHashicorpVaultEnabled(hashicorpVaultEnabled *bool) *VersionServiceOperatorParams {
	o.SetHashicorpVaultEnabled(hashicorpVaultEnabled)
	return o
}

// SetHashicorpVaultEnabled adds the hashicorpVaultEnabled to the version service operator params
func (o *VersionServiceOperatorParams) SetHashicorpVaultEnabled(hashicorpVaultEnabled *bool) {
	o.HashicorpVaultEnabled = hashicorpVaultEnabled
}

// WithHelmDeployCr adds the helmDeployCr to the version service operator params
func (o *VersionServiceOperatorParams) WithHelmDeployCr(helmDeployCr *bool) *VersionServiceOperatorParams {
	o.SetHelmDeployCr(helmDeployCr)
	return o
}

// SetHelmDeployCr adds the helmDeployCr to the version service operator params
func (o *VersionServiceOperatorParams) SetHelmDeployCr(helmDeployCr *bool) {
	o.HelmDeployCr = helmDeployCr
}

// WithHelmDeployOperator adds the helmDeployOperator to the version service operator params
func (o *VersionServiceOperatorParams) WithHelmDeployOperator(helmDeployOperator *bool) *VersionServiceOperatorParams {
	o.SetHelmDeployOperator(helmDeployOperator)
	return o
}

// SetHelmDeployOperator adds the helmDeployOperator to the version service operator params
func (o *VersionServiceOperatorParams) SetHelmDeployOperator(helmDeployOperator *bool) {
	o.HelmDeployOperator = helmDeployOperator
}

// WithKubeVersion adds the kubeVersion to the version service operator params
func (o *VersionServiceOperatorParams) WithKubeVersion(kubeVersion *string) *VersionServiceOperatorParams {
	o.SetKubeVersion(kubeVersion)
	return o
}

// SetKubeVersion adds the kubeVersion to the version service operator params
func (o *VersionServiceOperatorParams) SetKubeVersion(kubeVersion *string) {
	o.KubeVersion = kubeVersion
}

// WithLogCollectorVersion adds the logCollectorVersion to the version service operator params
func (o *VersionServiceOperatorParams) WithLogCollectorVersion(logCollectorVersion *string) *VersionServiceOperatorParams {
	o.SetLogCollectorVersion(logCollectorVersion)
	return o
}

// SetLogCollectorVersion adds the logCollectorVersion to the version service operator params
func (o *VersionServiceOperatorParams) SetLogCollectorVersion(logCollectorVersion *string) {
	o.LogCollectorVersion = logCollectorVersion
}

// WithNamespaceUID adds the namespaceUID to the version service operator params
func (o *VersionServiceOperatorParams) WithNamespaceUID(namespaceUID *string) *VersionServiceOperatorParams {
	o.SetNamespaceUID(namespaceUID)
	return o
}

// SetNamespaceUID adds the namespaceUid to the version service operator params
func (o *VersionServiceOperatorParams) SetNamespaceUID(namespaceUID *string) {
	o.NamespaceUID = namespaceUID
}

// WithOperatorVersion adds the operatorVersion to the version service operator params
func (o *VersionServiceOperatorParams) WithOperatorVersion(operatorVersion string) *VersionServiceOperatorParams {
	o.SetOperatorVersion(operatorVersion)
	return o
}

// SetOperatorVersion adds the operatorVersion to the version service operator params
func (o *VersionServiceOperatorParams) SetOperatorVersion(operatorVersion string) {
	o.OperatorVersion = operatorVersion
}

// WithPhysicalBackupScheduled adds the physicalBackupScheduled to the version service operator params
func (o *VersionServiceOperatorParams) WithPhysicalBackupScheduled(physicalBackupScheduled *bool) *VersionServiceOperatorParams {
	o.SetPhysicalBackupScheduled(physicalBackupScheduled)
	return o
}

// SetPhysicalBackupScheduled adds the physicalBackupScheduled to the version service operator params
func (o *VersionServiceOperatorParams) SetPhysicalBackupScheduled(physicalBackupScheduled *bool) {
	o.PhysicalBackupScheduled = physicalBackupScheduled
}

// WithPitrEnabled adds the pitrEnabled to the version service operator params
func (o *VersionServiceOperatorParams) WithPitrEnabled(pitrEnabled *bool) *VersionServiceOperatorParams {
	o.SetPitrEnabled(pitrEnabled)
	return o
}

// SetPitrEnabled adds the pitrEnabled to the version service operator params
func (o *VersionServiceOperatorParams) SetPitrEnabled(pitrEnabled *bool) {
	o.PitrEnabled = pitrEnabled
}

// WithPlatform adds the platform to the version service operator params
func (o *VersionServiceOperatorParams) WithPlatform(platform *string) *VersionServiceOperatorParams {
	o.SetPlatform(platform)
	return o
}

// SetPlatform adds the platform to the version service operator params
func (o *VersionServiceOperatorParams) SetPlatform(platform *string) {
	o.Platform = platform
}

// WithPmmEnabled adds the pmmEnabled to the version service operator params
func (o *VersionServiceOperatorParams) WithPmmEnabled(pmmEnabled *bool) *VersionServiceOperatorParams {
	o.SetPmmEnabled(pmmEnabled)
	return o
}

// SetPmmEnabled adds the pmmEnabled to the version service operator params
func (o *VersionServiceOperatorParams) SetPmmEnabled(pmmEnabled *bool) {
	o.PmmEnabled = pmmEnabled
}

// WithPmmVersion adds the pmmVersion to the version service operator params
func (o *VersionServiceOperatorParams) WithPmmVersion(pmmVersion *string) *VersionServiceOperatorParams {
	o.SetPmmVersion(pmmVersion)
	return o
}

// SetPmmVersion adds the pmmVersion to the version service operator params
func (o *VersionServiceOperatorParams) SetPmmVersion(pmmVersion *string) {
	o.PmmVersion = pmmVersion
}

// WithProduct adds the product to the version service operator params
func (o *VersionServiceOperatorParams) WithProduct(product string) *VersionServiceOperatorParams {
	o.SetProduct(product)
	return o
}

// SetProduct adds the product to the version service operator params
func (o *VersionServiceOperatorParams) SetProduct(product string) {
	o.Product = product
}

// WithProxysqlScheduler adds the proxysqlScheduler to the version service operator params
func (o *VersionServiceOperatorParams) WithProxysqlScheduler(proxysqlScheduler *string) *VersionServiceOperatorParams {
	o.SetProxysqlScheduler(proxysqlScheduler)
	return o
}

// SetProxysqlScheduler adds the proxysqlScheduler to the version service operator params
func (o *VersionServiceOperatorParams) SetProxysqlScheduler(proxysqlScheduler *string) {
	o.ProxysqlScheduler = proxysqlScheduler
}

// WithProxysqlVersion adds the proxysqlVersion to the version service operator params
func (o *VersionServiceOperatorParams) WithProxysqlVersion(proxysqlVersion *string) *VersionServiceOperatorParams {
	o.SetProxysqlVersion(proxysqlVersion)
	return o
}

// SetProxysqlVersion adds the proxysqlVersion to the version service operator params
func (o *VersionServiceOperatorParams) SetProxysqlVersion(proxysqlVersion *string) {
	o.ProxysqlVersion = proxysqlVersion
}

// WithShardingEnabled adds the shardingEnabled to the version service operator params
func (o *VersionServiceOperatorParams) WithShardingEnabled(shardingEnabled *bool) *VersionServiceOperatorParams {
	o.SetShardingEnabled(shardingEnabled)
	return o
}

// SetShardingEnabled adds the shardingEnabled to the version service operator params
func (o *VersionServiceOperatorParams) SetShardingEnabled(shardingEnabled *bool) {
	o.ShardingEnabled = shardingEnabled
}

// WithSidecarsUsed adds the sidecarsUsed to the version service operator params
func (o *VersionServiceOperatorParams) WithSidecarsUsed(sidecarsUsed *bool) *VersionServiceOperatorParams {
	o.SetSidecarsUsed(sidecarsUsed)
	return o
}

// SetSidecarsUsed adds the sidecarsUsed to the version service operator params
func (o *VersionServiceOperatorParams) SetSidecarsUsed(sidecarsUsed *bool) {
	o.SidecarsUsed = sidecarsUsed
}

// WriteToRequest writes these params to a swagger request
func (o *VersionServiceOperatorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.BackupVersion != nil {

		// query param backupVersion
		var qrBackupVersion string
		if o.BackupVersion != nil {
			qrBackupVersion = *o.BackupVersion
		}
		qBackupVersion := qrBackupVersion
		if qBackupVersion != "" {
			if err := r.SetQueryParam("backupVersion", qBackupVersion); err != nil {
				return err
			}
		}

	}

	if o.BackupsEnabled != nil {

		// query param backupsEnabled
		var qrBackupsEnabled bool
		if o.BackupsEnabled != nil {
			qrBackupsEnabled = *o.BackupsEnabled
		}
		qBackupsEnabled := swag.FormatBool(qrBackupsEnabled)
		if qBackupsEnabled != "" {
			if err := r.SetQueryParam("backupsEnabled", qBackupsEnabled); err != nil {
				return err
			}
		}

	}

	if o.ClusterSize != nil {

		// query param clusterSize
		var qrClusterSize int32
		if o.ClusterSize != nil {
			qrClusterSize = *o.ClusterSize
		}
		qClusterSize := swag.FormatInt32(qrClusterSize)
		if qClusterSize != "" {
			if err := r.SetQueryParam("clusterSize", qClusterSize); err != nil {
				return err
			}
		}

	}

	if o.ClusterWideEnabled != nil {

		// query param clusterWideEnabled
		var qrClusterWideEnabled bool
		if o.ClusterWideEnabled != nil {
			qrClusterWideEnabled = *o.ClusterWideEnabled
		}
		qClusterWideEnabled := swag.FormatBool(qrClusterWideEnabled)
		if qClusterWideEnabled != "" {
			if err := r.SetQueryParam("clusterWideEnabled", qClusterWideEnabled); err != nil {
				return err
			}
		}

	}

	if o.CustomResourceUID != nil {

		// query param customResourceUid
		var qrCustomResourceUID string
		if o.CustomResourceUID != nil {
			qrCustomResourceUID = *o.CustomResourceUID
		}
		qCustomResourceUID := qrCustomResourceUID
		if qCustomResourceUID != "" {
			if err := r.SetQueryParam("customResourceUid", qCustomResourceUID); err != nil {
				return err
			}
		}

	}

	if o.DatabaseVersion != nil {

		// query param databaseVersion
		var qrDatabaseVersion string
		if o.DatabaseVersion != nil {
			qrDatabaseVersion = *o.DatabaseVersion
		}
		qDatabaseVersion := qrDatabaseVersion
		if qDatabaseVersion != "" {
			if err := r.SetQueryParam("databaseVersion", qDatabaseVersion); err != nil {
				return err
			}
		}

	}

	if o.HaproxyVersion != nil {

		// query param haproxyVersion
		var qrHaproxyVersion string
		if o.HaproxyVersion != nil {
			qrHaproxyVersion = *o.HaproxyVersion
		}
		qHaproxyVersion := qrHaproxyVersion
		if qHaproxyVersion != "" {
			if err := r.SetQueryParam("haproxyVersion", qHaproxyVersion); err != nil {
				return err
			}
		}

	}

	if o.HashicorpVaultEnabled != nil {

		// query param hashicorpVaultEnabled
		var qrHashicorpVaultEnabled bool
		if o.HashicorpVaultEnabled != nil {
			qrHashicorpVaultEnabled = *o.HashicorpVaultEnabled
		}
		qHashicorpVaultEnabled := swag.FormatBool(qrHashicorpVaultEnabled)
		if qHashicorpVaultEnabled != "" {
			if err := r.SetQueryParam("hashicorpVaultEnabled", qHashicorpVaultEnabled); err != nil {
				return err
			}
		}

	}

	if o.HelmDeployCr != nil {

		// query param helmDeployCr
		var qrHelmDeployCr bool
		if o.HelmDeployCr != nil {
			qrHelmDeployCr = *o.HelmDeployCr
		}
		qHelmDeployCr := swag.FormatBool(qrHelmDeployCr)
		if qHelmDeployCr != "" {
			if err := r.SetQueryParam("helmDeployCr", qHelmDeployCr); err != nil {
				return err
			}
		}

	}

	if o.HelmDeployOperator != nil {

		// query param helmDeployOperator
		var qrHelmDeployOperator bool
		if o.HelmDeployOperator != nil {
			qrHelmDeployOperator = *o.HelmDeployOperator
		}
		qHelmDeployOperator := swag.FormatBool(qrHelmDeployOperator)
		if qHelmDeployOperator != "" {
			if err := r.SetQueryParam("helmDeployOperator", qHelmDeployOperator); err != nil {
				return err
			}
		}

	}

	if o.KubeVersion != nil {

		// query param kubeVersion
		var qrKubeVersion string
		if o.KubeVersion != nil {
			qrKubeVersion = *o.KubeVersion
		}
		qKubeVersion := qrKubeVersion
		if qKubeVersion != "" {
			if err := r.SetQueryParam("kubeVersion", qKubeVersion); err != nil {
				return err
			}
		}

	}

	if o.LogCollectorVersion != nil {

		// query param logCollectorVersion
		var qrLogCollectorVersion string
		if o.LogCollectorVersion != nil {
			qrLogCollectorVersion = *o.LogCollectorVersion
		}
		qLogCollectorVersion := qrLogCollectorVersion
		if qLogCollectorVersion != "" {
			if err := r.SetQueryParam("logCollectorVersion", qLogCollectorVersion); err != nil {
				return err
			}
		}

	}

	if o.NamespaceUID != nil {

		// query param namespaceUid
		var qrNamespaceUID string
		if o.NamespaceUID != nil {
			qrNamespaceUID = *o.NamespaceUID
		}
		qNamespaceUID := qrNamespaceUID
		if qNamespaceUID != "" {
			if err := r.SetQueryParam("namespaceUid", qNamespaceUID); err != nil {
				return err
			}
		}

	}

	// path param operatorVersion
	if err := r.SetPathParam("operatorVersion", o.OperatorVersion); err != nil {
		return err
	}

	if o.PhysicalBackupScheduled != nil {

		// query param physicalBackupScheduled
		var qrPhysicalBackupScheduled bool
		if o.PhysicalBackupScheduled != nil {
			qrPhysicalBackupScheduled = *o.PhysicalBackupScheduled
		}
		qPhysicalBackupScheduled := swag.FormatBool(qrPhysicalBackupScheduled)
		if qPhysicalBackupScheduled != "" {
			if err := r.SetQueryParam("physicalBackupScheduled", qPhysicalBackupScheduled); err != nil {
				return err
			}
		}

	}

	if o.PitrEnabled != nil {

		// query param pitrEnabled
		var qrPitrEnabled bool
		if o.PitrEnabled != nil {
			qrPitrEnabled = *o.PitrEnabled
		}
		qPitrEnabled := swag.FormatBool(qrPitrEnabled)
		if qPitrEnabled != "" {
			if err := r.SetQueryParam("pitrEnabled", qPitrEnabled); err != nil {
				return err
			}
		}

	}

	if o.Platform != nil {

		// query param platform
		var qrPlatform string
		if o.Platform != nil {
			qrPlatform = *o.Platform
		}
		qPlatform := qrPlatform
		if qPlatform != "" {
			if err := r.SetQueryParam("platform", qPlatform); err != nil {
				return err
			}
		}

	}

	if o.PmmEnabled != nil {

		// query param pmmEnabled
		var qrPmmEnabled bool
		if o.PmmEnabled != nil {
			qrPmmEnabled = *o.PmmEnabled
		}
		qPmmEnabled := swag.FormatBool(qrPmmEnabled)
		if qPmmEnabled != "" {
			if err := r.SetQueryParam("pmmEnabled", qPmmEnabled); err != nil {
				return err
			}
		}

	}

	if o.PmmVersion != nil {

		// query param pmmVersion
		var qrPmmVersion string
		if o.PmmVersion != nil {
			qrPmmVersion = *o.PmmVersion
		}
		qPmmVersion := qrPmmVersion
		if qPmmVersion != "" {
			if err := r.SetQueryParam("pmmVersion", qPmmVersion); err != nil {
				return err
			}
		}

	}

	// path param product
	if err := r.SetPathParam("product", o.Product); err != nil {
		return err
	}

	if o.ProxysqlScheduler != nil {

		// query param proxysqlScheduler
		var qrProxysqlScheduler string
		if o.ProxysqlScheduler != nil {
			qrProxysqlScheduler = *o.ProxysqlScheduler
		}
		qProxysqlScheduler := qrProxysqlScheduler
		if qProxysqlScheduler != "" {
			if err := r.SetQueryParam("proxysqlScheduler", qProxysqlScheduler); err != nil {
				return err
			}
		}

	}

	if o.ProxysqlVersion != nil {

		// query param proxysqlVersion
		var qrProxysqlVersion string
		if o.ProxysqlVersion != nil {
			qrProxysqlVersion = *o.ProxysqlVersion
		}
		qProxysqlVersion := qrProxysqlVersion
		if qProxysqlVersion != "" {
			if err := r.SetQueryParam("proxysqlVersion", qProxysqlVersion); err != nil {
				return err
			}
		}

	}

	if o.ShardingEnabled != nil {

		// query param shardingEnabled
		var qrShardingEnabled bool
		if o.ShardingEnabled != nil {
			qrShardingEnabled = *o.ShardingEnabled
		}
		qShardingEnabled := swag.FormatBool(qrShardingEnabled)
		if qShardingEnabled != "" {
			if err := r.SetQueryParam("shardingEnabled", qShardingEnabled); err != nil {
				return err
			}
		}

	}

	if o.SidecarsUsed != nil {

		// query param sidecarsUsed
		var qrSidecarsUsed bool
		if o.SidecarsUsed != nil {
			qrSidecarsUsed = *o.SidecarsUsed
		}
		qSidecarsUsed := swag.FormatBool(qrSidecarsUsed)
		if qSidecarsUsed != "" {
			if err := r.SetQueryParam("sidecarsUsed", qSidecarsUsed); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
