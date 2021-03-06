// Copyright (c) 2019 Tigera, Inc. All rights reserved.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package utils

import (
	"os/exec"
	"strings"

	. "github.com/onsi/gomega"
	log "github.com/sirupsen/logrus"
)

var calicoctl = "/go/src/github.com/projectcalico/calicoctl/bin/calicoctl-linux-amd64"

func Calicoctl(kdd bool, args ...string) string {
	out, err := CalicoctlMayFail(kdd, args...)
	Expect(err).NotTo(HaveOccurred())
	return out
}

func CalicoctlMayFail(kdd bool, args ...string) (string, error) {
	cmd := exec.Command(calicoctl, args...)
	cmd.Env = []string{"ETCD_ENDPOINTS=http://127.0.0.1:2379"}
	if kdd {
		cmd.Env = []string{"K8S_API_ENDPOINT=http://localhost:8080", "DATASTORE_TYPE=kubernetes"}
	}
	out, err := cmd.CombinedOutput()
	log.Infof("Run: calicoctl %v", strings.Join(args, " "))
	log.Infof("Output:\n%v", string(out))
	log.Infof("Error: %v", err)
	return string(out), err
}
