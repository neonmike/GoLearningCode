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
* @details  Go的基准测试
* @author      要你命三千又三千  any question please send mail to liukeshanren@gmail.com ,www.lksr.net
* @date        2018-8-17
* @version     V1.0
* @attention
* 硬件平台: windows 10 家庭版
* SDK版本：1.16.5 windows amd64
*/
package Mashal

import (
	"fmt"
	"testing"
)

func Benchmarkfmt_Sprintf(b *testing.B) {
	number:=10
	b.ResetTimer()
	for i:=0;i<b.N;i++ {
		fmt.Sprintf("%d",number)
	}
}
//在这次 go test 调用里，我们给-run 选项传递了字符串"none"，来保证在运行制订的基
//准测试函数之前没有单元测试会被运行。这两个选项都可以接受正则表达式，来决定需要运行哪
//些测试。由于例子里没有单元测试函数的名字中有none，所以使用none 可以排除所有的单元
//测试
//go test -v -run="none" -bench="Benchmarkfmt_Sprintf"
//PASS
//ok      GoLearningCode/Mashal   0.235s
