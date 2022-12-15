/*
 * Copyright (c) 2019 Huawei Technologies Co., Ltd.
 * A-Tune is licensed under the Mulan PSL v2.
 * You can use this software according to the terms and conditions of the Mulan PSL v2.
 * You may obtain a copy of Mulan PSL v2 at:
 *     http://license.coscl.org.cn/MulanPSL2
 * THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR
 * PURPOSE.
 * See the Mulan PSL v2 for more details.
 * Create: 2019-10-29
 */

package filters

import (
	"gitee.com/openeuler/A-Tune/common/log"
	"gitee.com/openeuler/A-Tune/common/system"
	"errors"
)

// CPUSchedule : CPU schedule filter
type CPUSchedule struct {
	Name string
}

// Tune CPU bindings according to input strategy
func (s *CPUSchedule) Tune(strategy string) error {
	switch strategy {
	case "auto":
	case "performance":
		return s.performance()
	case "powersave":
		return s.powersave()
	default:
		return errors.New("strategy don't exist")
	}
	return nil
}

func (s *CPUSchedule) performance() error {
	sys := system.GetSystem()

	for _, pid := range sys.GetPids("") {
		log.Infof("pid: %d", pid)
	}

	return nil
}

func (s *CPUSchedule) powersave() error {
	sys := system.GetSystem()

	for _, pid := range sys.GetPids("") {
		log.Infof("pid: %d", pid)
	}

	return nil
}
