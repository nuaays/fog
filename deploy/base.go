package deploy

import (
	"fmt"

	"github.com/cheyang/fog/types"
)

type Deployer interface {
	SetHosts(hosts []types.Host)
	Run(h) error
}

type ansibleDeployer struct {
	hosts   []types.Host
	roleMap map[string][]types.Host
}

func (this ansibleDeployer) Run() error {

	for k, hosts := range this.roleMap {
		fmt.Printf("[%s]\n", k)

		for _, h := range hosts {
			fmt.Printf("%s %s %s", h.MachineName, h.SSHHostname, h.SSHUserName)
		}

		fmt.Println("")
	}
}

func (this *ansibleDeployer) SetHost(hosts []types.Host) {

	this.hosts = hosts
	this.roleMap = make(map[string]types.Host)

	for _, host := range hosts {

		for _, role := range host.Roles {

			if v, found := this.roleMap[role]; !found {
				this.roleMap[role] = make([]types.Host)
			}

			this.roleMap[role] = append(this.roleMap[role], host)
		}

	}
}
