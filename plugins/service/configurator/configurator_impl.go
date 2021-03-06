/*
 * // Copyright (c) 2018 Cisco and/or its affiliates.
 * //
 * // Licensed under the Apache License, Version 2.0 (the "License");
 * // you may not use this file except in compliance with the License.
 * // You may obtain a copy of the License at:
 * //
 * //     http://www.apache.org/licenses/LICENSE-2.0
 * //
 * // Unless required by applicable law or agreed to in writing, software
 * // distributed under the License is distributed on an "AS IS" BASIS,
 * // WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * // See the License for the specific language governing permissions and
 * // limitations under the License.
 */

package configurator

import (
	govpp "git.fd.io/govpp.git/api"
	"github.com/ligato/cn-infra/logging"

	"github.com/ligato/vpp-agent/plugins/defaultplugins"
	//epmodel "github.com/contiv/vpp/plugins/ksr/model/endpoints"
	//svcmodel "github.com/contiv/vpp/plugins/ksr/model/service"
)

// ServiceConfigurator implements ServiceConfiguratorAPI.
type ServiceConfigurator struct {
	Deps
}

// Deps lists dependencies of ServiceConfigurator.
type Deps struct {
	Log       logging.Logger
	VPP       defaultplugins.API /* interface indexes */
	GoVPPChan *govpp.Channel     /* until supported in vpp-agent, we call NAT binary APIs directly */
}

// Init initializes service configurator.
func (sc *ServiceConfigurator) Init() error {
	return nil
}

// AddService installs NAT rules for a newly added service.
func (sc *ServiceConfigurator) AddService(service *ContivService) error {
	return nil
}

// UpdateService reflects a change in the configuration of a service with
// the smallest number of VPP/NAT binary API calls necessary.
func (sc *ServiceConfigurator) UpdateService(oldService, newService *ContivService) error {
	return nil
}

// DeleteService removes NAT configuration associated with a newly undeployed
// service.
func (sc *ServiceConfigurator) DeleteService(service *ContivService) error {
	return nil
}

// UpdateFrontends updates the list of interfaces with the enabled out2in
// VPP/NAT feature.
func (sc *ServiceConfigurator) UpdateFrontends(oldIfNames []string, newIfNames []string) error {
	return nil
}

// UpdateBackends updates the list of interfaces with the enabled in2out
// VPP/NAT feature.
func (sc *ServiceConfigurator) UpdateBackends(oldIfNames []string, newIfNames []string) error {
	return nil
}

// Resync completely replaces the current NAT configuration with the provided
// full state of K8s services.
func (sc *ServiceConfigurator) Resync(services []*ContivService, frontends []string, backendIfs []string) error {
	return nil
}

// Close deallocates resource held by the configurator.
func (sc *ServiceConfigurator) Close() error {
	return nil
}
