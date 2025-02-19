/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

This file was copied and modified from the kubernetes-csi/drivers project
https://github.com/kubernetes-csi/drivers/blob/release-1.0/pkg/nfs/nodeserver.go

Modifications Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved.
*/

package sampleprovider

import (
	"encoding/json"
	"fmt"

	api "github.com/gardener/machine-controller-manager-provider-sampleprovider/pkg/sampleprovider/apis"
	"github.com/gardener/machine-spec/lib/go/cmi"
	"github.com/golang/glog"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	corev1 "k8s.io/api/core/v1"
)

// NOTE
//
// The basic working of the controller will work with just implementing the CreateMachine() & DeleteMachine() methods.
// You can first implement these two methods and check the working of the controller.
// Once this works you can implement the rest of the methods.

// CreateMachine handles a machine creation request
// REQUIRED METHOD
//
// REQUEST PARAMETERS (cmi.CreateMachineRequest)
// MachineName           string             Contains the name of the machine object for whom an VM is to be created at the provider
// ProviderSpec          bytes(blob)        Template/Configuration of the machine to be created is given by at the provider
// Secrets               map<string,bytes>  (Optional) Contains a map from string to string contains any cloud specific secrets that can be used by the provider
// LastKnownState        bytes(blob)        (Optional) Last known state of VM during last operation. Could be helpful to continue operation from previous state
//
// RESPONSE PARAMETERS (cmi.CreateMachineResponse)
// ProviderID            string             Unique identification of the VM at the cloud provider. This could be the same/different from req.MachineName.
//                                          ProviderID typically matches with the node.Spec.ProviderID on the node object.
//                                          Eg: gce://project-name/region/vm-ProviderID
// NodeName              string             Returns the name of the node-object that the VM register's with Kubernetes.
//                                          This could be different from req.MachineName as well
// LastKnownState        bytes(blob)        (Optional) Last known state of VM during the current operation.
//                                          Could be helpful to continue operations in future requests.
//
// OPTIONAL IMPLEMENTATION LOGIC
// It is optionally expected by the safety controller to use an identification mechanisms to map the VM Created by a providerSpec.
// These could be done using tag(s)/resource-groups etc.
// This logic is used by safety controller to delete orphan VMs which are not backed by any machine CRD
//
func (ms *MachinePlugin) CreateMachine(ctx context.Context, req *cmi.CreateMachineRequest) (*cmi.CreateMachineResponse, error) {
	// Log messages to track start of request
	glog.V(2).Infof("Create machine request has been received for %q", req.MachineName)
	return nil, status.Error(codes.Unimplemented, "")
}

// DeleteMachine handles a machine deletion request
//
// REQUEST PARAMETERS (cmi.DeleteMachineRequest)
// MachineName          string              Contains the name of the machine object for the backing VM(s) have to be deleted
// ProviderSpec         bytes(blob)         Template/Configuration of the machine to be deleted is given by at the provider
// Secrets              map<string,bytes>   (Optional) Contains a map from string to string contains any cloud specific secrets that can be used by the provider
// LastKnownState       bytes(blob)         (Optional) Last known state of VM during last operation. Could be helpful to continue operation from previous state
//
// RESPONSE PARAMETERS (cmi.DeleteMachineResponse)
// LastKnownState       bytes(blob)        (Optional) Last known state of VM during the current operation.
//                                          Could be helpful to continue operations in future requests.
//
func (ms *MachinePlugin) DeleteMachine(ctx context.Context, req *cmi.DeleteMachineRequest) (*cmi.DeleteMachineResponse, error) {
	// Log messages to track start of request
	glog.V(2).Infof("Delete machine request has been received for %q", req.MachineName)
	return nil, status.Error(codes.Unimplemented, "")
}

// GetMachineStatus handles a machine get status request
// OPTIONAL METHOD
//
// REQUEST PARAMETERS (cmi.GetMachineStatusRequest)
// MachineName          string              Contains the name of the machine object for whose status is to be retrieved
// ProviderSpec         bytes(blob)         Template/Configuration of the machine whose status is to be retrieved
// Secrets              map<string,bytes>   (Optional) Contains a map from string to string contains any cloud specific secrets that can be used by the provider
//
// RESPONSE PARAMETERS (cmi.GetMachineStatueResponse)
// ProviderID           string              Unique identification of the VM at the cloud provider. This could be the same/different from req.MachineName.
//                                          ProviderID typically matches with the node.Spec.ProviderID on the node object.
//                                          Eg: gce://project-name/region/vm-ProviderID
// NodeName             string              Returns the name of the node-object that the VM register's with Kubernetes.
//                                          This could be different from req.MachineName as well
//
// The request should return a NOT_FOUND (5) status error code if the machine is not existing
func (ms *MachinePlugin) GetMachineStatus(ctx context.Context, req *cmi.GetMachineStatusRequest) (*cmi.GetMachineStatusResponse, error) {
	// Log messages to track start of request
	glog.V(2).Infof("Get machine request has been received for %q", req.MachineName)

	// If the machine is not existing, return a NotFound error code
	// return nil, status.Error(codes.NotFound, "")

	return nil, status.Error(codes.Unimplemented, "")
}

// ListMachines lists all the machines possibilly created by a providerSpec
// Identifying machines created by a given providerSpec depends on the OPTIONAL IMPLEMENTATION LOGIC
// you have used to identify machines created by a providerSpec. It could be tags/resource-groups etc
// OPTIONAL METHOD
//
// REQUEST PARAMETERS (cmi.ListMachinesRequest)
// ProviderSpec          bytes(blob)         Template/Configuration of the machine that wouldn've been created by this ProviderSpec (Machine Class)
// Secrets               map<string,bytes>   (Optional) Contains a map from string to string contains any cloud specific secrets that can be used by the provider
//
// RESPONSE PARAMETERS (cmi.ListMachinesResponse)
// MachineList           map<string,string>  A map containing the keys as the MachineID and value as the MachineName
//                                           for all machine's who where possibilly created by this ProviderSpec
//
func (ms *MachinePlugin) ListMachines(ctx context.Context, req *cmi.ListMachinesRequest) (*cmi.ListMachinesResponse, error) {
	// Log messages to track start of request
	glog.V(2).Infof("List machines request has been received")
	glog.V(4).Infof("ProviderSpec = %q", req.ProviderSpec)
	return nil, status.Error(codes.Unimplemented, "")
}

// GetVolumeIDs returns a list of Volume IDs for all PV Specs for whom an provider volume was found
//
// REQUEST PARAMETERS (cmi.GetVolumeIDsRequest)
// PVSpecList            bytes(blob)         PVSpecsList is a list PV specs for whom volume-IDs are required. Plugin should parse this raw data into pre-defined list of PVSpecs.
//
// RESPONSE PARAMETERS (cmi.GetVolumeIDsResponse)
// VolumeIDs             repeated string     VolumeIDs is a repeated list of VolumeIDs.
//
func (ms *MachinePlugin) GetVolumeIDs(ctx context.Context, req *cmi.GetVolumeIDsRequest) (*cmi.GetVolumeIDsResponse, error) {
	// Log messages to track start of request
	glog.V(2).Infof("GetVolumeIDs request has been received")
	glog.V(4).Infof("PVSpecList = %q", req.PVSpecList)

	var volumeSpecs []*corev1.PersistentVolumeSpec
	err := json.Unmarshal(req.PVSpecList, &volumeSpecs)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return nil, status.Error(codes.Unimplemented, "")
}

// ShutDownMachine handles a machine shutdown/power-off/stop request
// OPTIONAL METHOD
//
// REQUEST PARAMETERS (cmi.ShutDownMachineRequest)
// ProviderSpec          bytes(blob)         Template/Configuration of the machine that wouldn've been created by this ProviderSpec (Machine Class)
// Secrets               map<string,bytes>   (Optional) Contains a map from string to string contains any cloud specific secrets that can be used by the provider
// LastKnownState        bytes(blob)        (Optional) Last known state of VM during last operation. Could be helpful to continue operation from previous state
//
// RESPONSE PARAMETERS (cmi.DeleteMachineResponse)
// LastKnownState        bytes(blob)        (Optional) Last known state of VM during the current operation.
//                                          Could be helpful to continue operations in future requests.
//
func (ms *MachinePlugin) ShutDownMachine(ctx context.Context, req *cmi.ShutDownMachineRequest) (*cmi.ShutDownMachineResponse, error) {
	// Log messages to track start of request
	glog.V(2).Infof("ShutDown machine request has been received for %q", req.MachineName)
	return nil, status.Error(codes.Unimplemented, "")
}

// Remove the method below in the final implementation
func dummyMethod() {
	// Sample code to access provider spec
	// Delete the 4 following line in the controller implementation
	dummyObject := api.SampleProviderProviderSpec{
		APIVersion: "v1alpha1",
	}
	fmt.Println("APIVersion of object ", dummyObject.APIVersion)
}
