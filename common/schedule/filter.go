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

package schedule

import (
	"gitee.com/openeuler/A-Tune/common/schedule/filters"
)

// Filter :schedule filters interface
type Filter interface {
	Tune(strategy string) error
}

// Factory function for schedule filters
func Factory(name string) Filter {
	switch name {
	case "irq":
		return &filters.IrqSchedule{Name: "irq"}
	case "numa":
		return &filters.NumaSchedule{Name: "numa"}
	case "cpu":
		return &filters.CPUSchedule{Name: "cpu"}
	default:
		return nil
	}
}
