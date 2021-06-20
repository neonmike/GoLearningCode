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

/**@file  Test-printType.go
* @brief  学习Go中关键内容的用法
* @details  Go中 数据类型的打印
* @author      要你命三千又三千  any question please send mail to liukeshanren@gmail.com ,www.lksr.net
* @date        2018-8-17
* @version     V1.0
* @attention
* 硬件平台: windows 10 家庭版
* SDK版本：1.16.5 windows amd64
*/
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var num =16
	fmt.Printf("num Type: %T and num size %d",num,unsafe.Sizeof(num) )
}
