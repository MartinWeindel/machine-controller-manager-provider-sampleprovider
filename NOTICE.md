## machine-controller-manager-provider-sampleprovider
Copyright (c) 2018-2019 SAP SE or an SAP affiliate company. All rights reserved.

## Seed Source

The source code of this component was seeded based on a copy of the following files from container-storage-interface/spec. 

Container Storage Interface (CSI) Drivers  
Copyright 2017 The Kubernetes Authors.  
https://github.com/kubernetes-csi/drivers/tree/release-1.0  
Apache 2 license (https://github.com/kubernetes-csi/drivers/blob/release-1.0/LICENSE )

Release: 1.0  
Commit-ID: b776760b257e955d86d279e1bba375b06e9cbe6e.  
Commit-Message:  Merge pull request #129 from pohly/hostpath-1.0.0 -backport-2  
To the left are the list of copied files -> and to the right the current location they are at.  

    app/nfsplugin/main.go -> app/sampleprovider/cmi-plugin.go
    pkg/nfs/driver.go -> pkg/sampleprovider/plugin.go
    pkg/nfs/nodeserver.go -> pkg/sampleprovider/machine_server.go
    pkg/csi-common/driver.go -> pkg/cmicommon/plugin.go
    pkg/csi-common/identityserver-default.go -> pkg/cmicommon/identity_server_defaults.go
    pkg/csi-common/nodeserver-default.go -> pkg/cmicommon/machine_server_defaults.go
    pkg/csi-common/server.go -> pkg/cmicommon/server.go
    pkg/csi-common/utils.go -> pkg/cmicommon/utils.go


