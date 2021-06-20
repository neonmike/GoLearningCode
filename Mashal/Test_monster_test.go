
//
//Copyright [yyyy] [name of copyright owner]
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/**@file Test-Closure.go
* @brief  学习Go中关键内容的用法
* @details  Go的单元测试
* @author      要你命三千又三千  any question please send mail to liukeshanren@gmail.com ,www.lksr.net
* @date        2018-8-17
* @version     V1.0
* @attention
* 硬件平台: windows 10 家庭版
* SDK版本：1.16.5 windows amd64
*/
package Mashal

import (
	"testing"
)
func TestMonster_Store(t *testing.T){
	monster:=&Monster{
		Name:"红孩儿",
		Age:10,
		Skill: "喷火",
	}
	res:=monster.Store()
	if !res {
		t.Fatalf("monster.Store 错误，希望为=%v,实际值为 =%v",true,res)

	}
	t.Logf("monster.Store 测试成功")
}

func TestMonster_ReStore(t *testing.T) {
	var monster =&Monster{}
	res:=monster.ReStore()
	if !res {
		t.Fatalf("monster.ReStore 错误，希望为 =%v,实际为=%v",true,res)

	}
	if monster.Name!="红孩儿" {
		t.Fatalf("monster.Restore错误")
	}
	t.Logf("monster.Restore 测试成功")
}
/*
=== RUN   TestMonster_Store
    Test_monster_test.go:17: monster.Store 测试成功
--- PASS: TestMonster_Store (0.00s)
PASS
coverage: 30.0% of statements in ../../GoLearningCode/...
*/
