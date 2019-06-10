package report

import (
	core "k8s.io/api/core/v1"
)

func getStorageTypeFromPersistentVolume(persistentVolume core.PersistentVolume) StorageType {

	var storageType StorageType

	persistentVolumeSource := persistentVolume.Spec.PersistentVolumeSource

	if persistentVolumeSource.AWSElasticBlockStore != nil {
		storageType = AWSElasticBlockStore
	} else if persistentVolumeSource.AzureFile != nil {
		storageType = AzureFile
	} else if persistentVolumeSource.AzureDisk != nil {
		storageType = AzureDisk
	} else if persistentVolumeSource.CephFS != nil {
		storageType = CephFS
	} else if persistentVolumeSource.Cinder != nil {
		storageType = Cinder
	} else if persistentVolumeSource.CSI != nil {
		storageType = CSI
	} else if persistentVolumeSource.FC != nil {
		storageType = FC
	} else if persistentVolumeSource.Flocker != nil {
		storageType = Flocker
	} else if persistentVolumeSource.FlexVolume != nil {
		storageType = FlexVolume
	} else if persistentVolumeSource.GCEPersistentDisk != nil {
		storageType = GCEPersistentDisk
	} else if persistentVolumeSource.Glusterfs != nil {
		storageType = Glusterfs
	} else if persistentVolumeSource.HostPath != nil {
		storageType = HostPath
	} else if persistentVolumeSource.ISCSI != nil {
		storageType = ISCSI
	} else if persistentVolumeSource.Local != nil {
		storageType = Local
	} else if persistentVolumeSource.NFS != nil {
		storageType = NFS
	} else if persistentVolumeSource.Quobyte != nil {
		storageType = Quobyte
	} else if persistentVolumeSource.PortworxVolume != nil {
		storageType = PortworxVolume
	} else if persistentVolumeSource.RBD != nil {
		storageType = RBD
	} else if persistentVolumeSource.ScaleIO != nil {
		storageType = ScaleIO
	} else if persistentVolumeSource.StorageOS != nil {
		storageType = StorageOS
	} else if persistentVolumeSource.VsphereVolume != nil {
		storageType = VsphereVolume
	} else {
		storageType = Unknown
	}

	return storageType
}

func getStorageTypeFromVolume(volume core.Volume) StorageType {

	var storageType StorageType

	if volume.AWSElasticBlockStore != nil {
		storageType = AWSElasticBlockStore
	} else if volume.AzureFile != nil {
		storageType = AzureFile
	} else if volume.AzureDisk != nil {
		storageType = AzureDisk
	} else if volume.CephFS != nil {
		storageType = CephFS
	} else if volume.Cinder != nil {
		storageType = Cinder
	} else if volume.ConfigMap != nil {
		storageType = ConfigMap
	} else if volume.DownwardAPI != nil {
		storageType = DownwardAPI
	} else if volume.EmptyDir != nil {
		storageType = EmptyDir
	} else if volume.FC != nil {
		storageType = FC
	} else if volume.Flocker != nil {
		storageType = Flocker
	} else if volume.FlexVolume != nil {
		storageType = FlexVolume
	} else if volume.GCEPersistentDisk != nil {
		storageType = GCEPersistentDisk
	} else if volume.Glusterfs != nil {
		storageType = Glusterfs
	} else if volume.HostPath != nil {
		storageType = HostPath
	} else if volume.ISCSI != nil {
		storageType = ISCSI
	} else if volume.NFS != nil {
		storageType = NFS
	} else if volume.PersistentVolumeClaim != nil {
		storageType = PersistentVolumeClaim
	} else if volume.PhotonPersistentDisk != nil {
		storageType = PhotonPersistentDisk
	} else if volume.PortworxVolume != nil {
		storageType = PortworxVolume
	} else if volume.Projected != nil {
		storageType = Projected
	} else if volume.Quobyte != nil {
		storageType = Quobyte
	} else if volume.RBD != nil {
		storageType = RBD
	} else if volume.ScaleIO != nil {
		storageType = ScaleIO
	} else if volume.StorageOS != nil {
		storageType = StorageOS
	} else if volume.Secret != nil {
		storageType = Secret
	} else if volume.VsphereVolume != nil {
		storageType = VsphereVolume
	} else {
		storageType = Unknown
	}
	return storageType
}
